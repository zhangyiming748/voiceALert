package voiceAlert

import (
	"os/exec"
	"runtime"
)

const (
	SUCCESS  = iota + 1 // 单次转码成功
	FAILED              // 转码失败,程序退出
	COMPLETE            // 转码进程完成
)

func Voice(msg int) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		// 查询发音人 `say -v ?`
		voice := "Serena"

		switch msg {
		case SUCCESS:
			cmd = exec.Command("say", "-v", voice, "Rocket was launched successfully")
			cmd.Start()
		case FAILED:
			cmd = exec.Command("say", "-v", voice, "Rocket launch FAILED")
			cmd.Start()
		case COMPLETE:
			cmd = exec.Command("say", "-v", voice, "mission COMPLETE!")
			cmd.Start()
		}
	case "linux":
		cmd = exec.Command("echo", "-e", "\\a")
		switch msg {
		case SUCCESS:
			for i := 0; i < 2; i++ {
				cmd.Start()
			}
		case FAILED:
			for i := 0; i < 50; i++ {
				cmd.Start()
			}
		case COMPLETE:
			for i := 0; i < 100; i++ {
				cmd.Start()
			}
		}
	}
	cmd.Wait()
}
