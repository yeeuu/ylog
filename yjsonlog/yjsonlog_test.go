package yjsonlog

import (
	"testing"
	"time"

	"github.com/funny/unitest"
)

func Test_SwitchByDay(t *testing.T) {
	log, err := New(Config{
		Dir:      ".",
		Switcher: DaySwitcher,
		FileType: ".log",
	})
	unitest.NotError(t, err)
	log.Log(M{"Time": time.Now()})
	log.Log(M{"Time": time.Now()})
	log.Log(M{"Time": time.Now()})
	log.Close()
}

func Test_SwitchByHours(t *testing.T) {
	log, err := New(Config{
		Dir:      ".",
		Switcher: HoursSwitcher,
		FileType: ".log",
	})
	unitest.NotError(t, err)
	log.Log(M{"Time": time.Now()})
	log.Log(M{"Time": time.Now()})
	log.Log(M{"Time": time.Now()})
	log.Close()
}
