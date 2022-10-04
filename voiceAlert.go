package voiceAlert

import (
	"github.com/zhangyiming748/voiceAlert/util"
	"os/exec"
	"runtime"
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
		if !util.HasSpoker("Kate") {
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
