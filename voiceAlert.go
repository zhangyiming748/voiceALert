package voiceAlert

import (
	"log"
	"os/exec"
	"runtime"
)

const (
	success  = iota + 1 // 单次转码成功
	failed              // 转码失败,程序退出
	complete            // 转码进程完成
)

func Voice(msg int) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("发生命令出现错误", err)
		}
	}()
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		// 查询发音人 `say -v ?`
		voice := "Serena"

		switch msg {
		case success:
			cmd = exec.Command("say", "-v", voice, "Rocket was launched successfully")
			if err := cmd.Start(); err != nil {
				log.Printf("执行发声命令发生错误:%v\n", err)
			}
		case failed:
			cmd = exec.Command("say", "-v", voice, "Rocket launch failed")
			if err := cmd.Start(); err != nil {
				log.Printf("执行发声命令发生错误:%v\n", err)
			}
		case complete:
			cmd = exec.Command("say", "-v", voice, "mission complete!")
			if err := cmd.Start(); err != nil {
				log.Printf("执行发声命令发生错误:%v\n", err)
			}
		}
	case "linux":
		switch msg {
		case success:
			cmd = exec.Command("espeak", "Rocket was launched successfully")
			if err := cmd.Start(); err != nil {
				log.Printf("执行发声命令发生错误:%v\n", err)
			}
		case failed:
			cmd = exec.Command("espeak", "Rocket launch failed")
			if err := cmd.Start(); err != nil {
				log.Printf("执行发声命令发生错误:%v\n", err)
			}
		case complete:
			//espeak "enter the text that you want to listen to"
			cmd = exec.Command("espeak", "mission complete!")
			if err := cmd.Start(); err != nil {
				log.Printf("执行发声命令发生错误:%v\n", err)
			}
		}
	}
	if err := cmd.Wait(); err != nil {
		log.Println("命令执行中有错误产生", err)
	}
}
