package voiceAlert

import (
	"github.com/zhangyiming748/log"
	"os/exec"
	"runtime"
)

const (
	SUCCESS  = iota + 1 // 单次转码成功
	FAILED              // 转码失败,程序退出
	COMPLETE            // 转码进程完成
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

var (
	cmd *exec.Cmd
)

/*
运行在mac上的发声命令
*/
func CustomizedOnMac(spoker, content string) {
	defer func() {
		if err := recover(); err != nil {
			log.Warn.Printf("执行发生命令出现错误:%v", err)
		}
	}()
	cmd = exec.Command("say", "-v", spoker, content)
	cmd.Run()
}

/*
运行在linux上的发声命令
*/
func CustomizedOnLinux(content string) {
	defer func() {
		if err := recover(); err != nil {
			log.Warn.Printf("执行发生命令出现错误:%v", err)
		}
	}()
	//espeak "Testing espeak from the Ubuntu 18.04 terminal"
	cmd = exec.Command("espeak", content)
	cmd.Run()
}
func Voice(msg int) {
	defer func() {
		log.Warn.Printf("这是一个即将废弃的函数\n请及时切换到CustomizedOnMac或CustomizedOnLinux函数\n")
		if err := recover(); err != nil {
			log.Debug.Printf("执行发生命令出现错误:%v", err)
		}
	}()
	switch runtime.GOOS {
	case "darwin":
		switch msg {
		case SUCCESS:
			CustomizedOnMac(Victoria, "Rocket was launched successfully")
		case FAILED:
			CustomizedOnMac(Victoria, "Rocket launch FAILED")
		case COMPLETE:
			CustomizedOnMac(Victoria, "Mission COMPLETE!")
		}
	case "linux":
		switch msg {
		case SUCCESS:
			CustomizedOnLinux("Rocket was launched successfully")
		case FAILED:
			CustomizedOnLinux("Rocket launch FAILED")
		case COMPLETE:
			CustomizedOnLinux("Rocket launch FAILED")
		}
	}
}
