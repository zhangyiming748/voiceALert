package voiceAlert

import (
	"fmt"
	"golang.org/x/exp/slog"
	"io"
	"os"
	"os/exec"
	"runtime"
)

const (
	Default  = Serena
	Allison  = "Allison"  // 深沉美式女声
	Ava      = "Ava"      // 深沉美式女声
	Daniel   = "Daniel"   // 正式英式男声
	Lanlan   = "Lanlan"   // 童声中文女声
	Meijia   = "Meijia"   // 正式中文女声
	Lilian   = "Lilian"   // 柔和中文女声
	Samantha = "Samantha" // 正经美式女声
	Serena   = "Serena"   // 沉稳英式女声
	Shanshan = "Shanshan" // 浑厚中文女声
	Shasha   = "Shasha"   // 成熟中文女声
	Sinji    = "Sinji"    // 粤语中文女声
	Tingting = "Tingting" // 机械中文女声
	Victoria = "Victoria" // 性感美式女声
)

var mylog *slog.Logger

func SetLog(level string) {
	var opt slog.HandlerOptions
	switch level {
	case "Debug":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}
	case "Info":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelInfo, // slog 默认日志级别是 info
		}
	case "Warn":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelWarn, // slog 默认日志级别是 info
		}
	case "Err":
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelError, // slog 默认日志级别是 info
		}
	default:
		slog.Warn("需要正确设置环境变量 Debug,Info,Warn or Err")
		slog.Info("默认使用Debug等级")
		opt = slog.HandlerOptions{ // 自定义option
			AddSource: true,
			Level:     slog.LevelDebug, // slog 默认日志级别是 info
		}

	}
	file := "processVideo.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	//defer logf.Close() //如果不关闭可能造成内存泄露
	mylog = slog.New(opt.NewJSONHandler(io.MultiWriter(logf, os.Stdout)))
}
func init() {
	l := os.Getenv("LEVEL")
	SetLog(l)
}

/*
运行在mac上的发声命令
*/
func customizedOnMac(spoker, content string) {
	defer func() {
		if err := recover(); err != nil {
			mylog.Warn("执行发声命令出现错误", slog.Any("错误信息", err))
		}
	}()
	cmd := exec.Command("say", "-v", spoker, content)
	mylog.Debug("成功执行命令", slog.String("命令", fmt.Sprint(cmd)))
	cmd.Run()
}

/*
运行在linux上的发声命令
*/
func customizedOnLinux(content string) {
	defer func() {
		if err := recover(); err != nil {
			mylog.Warn("执行发声命令出现错误", slog.Any("是否安装espeak？", err))
		}
	}()
	//espeak "Testing espeak from the Ubuntu 18.04 terminal"
	cmd := exec.Command("espeak", "-v", "zh", content)
	cmd.Run()
}
func Customize(content, teller string) {
	if os.Getenv("QUIET") == "True" {
		return
	}
	switch runtime.GOOS {
	case "darwin":
		customizedOnMac(teller, content)
	case "linux":
		customizedOnLinux(content)
	default:
		mylog.Warn("系统问题")
	}
}
