# ylog

ylog是一个简单的日志记录模块，以json格式记录程序日志，日志文件每天切换。根据[funny/log](https://github.com/funny/log)修改而来。

# 例子

```
package main

import (
	"github.com/yeeuu/ylog"
)

func main(t *testing.T) {
	ylog.Init(".")
	ylog.SetDebug(true)
	ylog.Debug("debug")
	ylog.Error("error")
	ylog.Warning("warning")
	ylog.Info("info")
	ylog.Close()
}
```
