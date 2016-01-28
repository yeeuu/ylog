package ylog

import (
	"runtime"
	"time"

	"github.com/yeeuu/ylog/yjsonlog"
)

var gLogger *L

func init() {
	gLogger = nil
}

// Init 初始化全局日志
func Init(dir string) {
	l, err := New(dir)
	if err != nil {
		panic(err)
	}
	gLogger = l
}

// Close 关闭全局日志系统
func Close() {
	if gLogger != nil {
		gLogger.Close()
		gLogger = nil
	}
}

// Info 在全局日志中输出信息
func Info(msg string, data ...interface{}) {
	if gLogger != nil {
		gLogger.Info(msg, data...)
	}
}

// Warning 在全局日志中输出警告信息
func Warning(msg string, data ...interface{}) {
	if gLogger != nil {
		gLogger.Warning(msg, data...)
	}
}

// Error 在全局日志中输出错误信息
func Error(msg string, data ...interface{}) {
	if gLogger != nil {
		gLogger.Error(msg, data...)
	}
}

// Debug 在全局日志中输出调试信息
func Debug(msg string, data ...interface{}) {
	if gLogger != nil {
		gLogger.Debug(msg, data...)
	}
}

// SetDebug 全局日志开启或关闭调试信息的输出
func SetDebug(debug bool) {
	if gLogger != nil {
		gLogger.SetDebug(debug)
	}
}

// M 日志数据
type M map[string]interface{}

// L 日志记录器
type L struct {
	l     *yjsonlog.L
	debug bool
}

// New 新建一个日志记录器
func New(dir string) (*L, error) {
	l, err := yjsonlog.New(yjsonlog.Config{
		Dir:      dir,
		Switcher: yjsonlog.DaySwitcher,
		FileType: ".log",
	})
	if err != nil {
		return nil, err
	}
	return &L{l, true}, nil
}

// Close 关闭日志系统
func (logger *L) Close() {
	logger.l.Close()
}

// Log 生成jsonlog
func (logger *L) Log(msg string, typ string, data ...interface{}) {
	_, file, line, _ := runtime.Caller(3)
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	file = short
	m := yjsonlog.M{
		"Time":    time.Now().Format("2006-01-02 15:04:05"),
		"Type":    typ,
		"File":    file,
		"Line":    line,
		"Message": msg,
	}
	if data != nil {
		m["Data"] = data
	}
	logger.l.Log(m)
}

// Info 在日志文件中输出信息
func (logger *L) Info(msg string, data ...interface{}) {
	logger.Log(msg, "info", data...)
}

// Warning 在日志文件中输出警告信息
func (logger *L) Warning(msg string, data ...interface{}) {
	logger.Log(msg, "warning", data...)
}

// Error 在日志文件中输出错误信息
func (logger *L) Error(msg string, data ...interface{}) {
	logger.Log(msg, "error", data...)
}

// Debug 在日志文件中输出调试信息
func (logger *L) Debug(msg string, data ...interface{}) {
	if logger.debug {
		logger.Log(msg, "debug", data...)
	}
}

// SetDebug 开启或关闭调试信息的输出
func (logger *L) SetDebug(debug bool) {
	logger.debug = debug
}
