package yjsonlog

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"log"
	"os"
	"sync"
)

type M map[string]interface{}

type File struct {
	mutex   sync.Mutex
	f       *os.File
	bufio   *bufio.Writer
	gzip    *gzip.Writer
	json    *json.Encoder
	changed bool
}

func NewFile(fileName, fileType string) (*File, error) {
	fullName := fileName + fileType

	f, err := os.OpenFile(fullName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}

	log := &File{f: f}
	log.bufio = bufio.NewWriter(log.f)
	log.json = json.NewEncoder(log.bufio)

	return log, nil
}

func (file *File) Write(r M) {
	file.mutex.Lock()
	defer file.mutex.Unlock()

	if err := file.json.Encode(r); err != nil {
		log.Println("jsonlog encode failed:", err.Error())
	}
	file.changed = true
}

func (file *File) flush(isClose bool) error {
	file.mutex.Lock()
	defer file.mutex.Unlock()

	if !file.changed {
		return nil
	}

	if err := file.bufio.Flush(); err != nil {
		return err
	}

	if err := file.f.Sync(); err != nil {
		return err
	}

	file.changed = false
	return nil
}

func (file *File) Flush() error {
	return file.flush(false)
}

func (file *File) Close() error {
	if err := file.flush(true); err != nil {
		return err
	}
	return file.f.Close()
}
