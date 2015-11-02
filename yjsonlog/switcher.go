package yjsonlog

import (
	"time"
)

var (
	// DaySwitcher 按天切换文件
	DaySwitcher = daySwitch{}
	// HoursSwitcher 按小时切换文件
	HoursSwitcher = hoursSwitch{}
)

// Switcher 文件切换器
type Switcher interface {
	FirstSwitchTime() time.Duration
	NextSwitchTime() time.Duration
	DirAndFileName(base string) (dir, file string)
}

type daySwitch struct{}

// FirstSwitchTime 到明天凌晨间隔多长时间
func (s daySwitch) FirstSwitchTime() time.Duration {
	now := time.Now()
	return time.Date(
		now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0, now.Location(),
	).Add(24 * time.Hour).Sub(now)
}

// NextSwitchTime 下次切换时间
func (s daySwitch) NextSwitchTime() time.Duration {
	return 24 * time.Hour
}

// DirAndFileName 文件夹和文件名
func (s daySwitch) DirAndFileName(base string) (dir, file string) {
	now := time.Now()
	dir = base + "/" + now.Format("2006-01/")
	file = dir + now.Format("2006-01-02")
	return
}

type hoursSwitch struct{}

// FirstSwitchTime 到下一个整点间隔多长时间
func (s hoursSwitch) FirstSwitchTime() time.Duration {
	now := time.Now()
	return time.Date(
		now.Year(), now.Month(), now.Day(),
		now.Hour(), 0, 0, 0, now.Location(),
	).Add(time.Hour).Sub(now)
}

// NextSwitchTime 下次切换时间
func (s hoursSwitch) NextSwitchTime() time.Duration {
	return time.Hour
}

// DirAndFileName 文件夹和文件名
func (s hoursSwitch) DirAndFileName(base string) (dir, file string) {
	now := time.Now()
	dir = base + "/" + now.Format("2006-01/2006-01-02/")
	file = dir + now.Format("2006-01-02_03")
	return
}
