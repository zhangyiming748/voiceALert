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
		voice := "Kate"
		if !HasSpoker(voice) {
			voice = "Victoria"
		}
		switch msg {
		case success:
			cmd = exec.Command("say", "-v", voice, "Rocket was launched successfully")
			cmd.Start()
		case failed:
			cmd = exec.Command("say", "-v", voice, "Rocket launch failed")
			cmd.Start()
		case complete:
			cmd = exec.Command("say", "-v", voice, "mission complete!")
			cmd.Start()
		}
	case "linux":
		cmd = exec.Command("echo", "-e", "\\a")
		switch msg {
		case success:
			for i := 0; i < 2; i++ {
				cmd.Start()
			}
		case failed:
			for i := 0; i < 50; i++ {
				cmd.Start()
			}
		case complete:
			for i := 0; i < 100; i++ {
				cmd.Start()
			}
		}
	}
	cmd.Wait()
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
