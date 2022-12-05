package voiceAlert

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

const (
	success  = iota + 1 // 单次转码成功
	failed              // 转码失败,程序退出
	complete            // 转码进程完成
)

func Voice(msg int) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		// 查询发音人 `say -v ?`
		voice := "Serena"
		if !HasSpoker(voice) {
			voice = "Victoria"
		}
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

func HasSpoker(key string) bool {
	cmd := exec.Command("say", "-v", "?")
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		log.Printf("cmd.StdoutPipe产生的错误:%v\n", err)
	}
	if err = cmd.Start(); err != nil {
		log.Printf("cmd.Run产生的错误:%v\n", err)
	}
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		//写成输出日志
		t := string(tmp)
		t = strings.Replace(t, "\u0000", "", -1)
		if strings.Contains(t, key) {
			return true
		}
		//log.Println(t)
		if err != nil {
			break
		}
	}
	if err = cmd.Wait(); err != nil {
		log.Println("命令执行中有错误产生", err)
	}
	return false
}
