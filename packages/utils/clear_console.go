package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func clearConsole() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "linux", "darwin":
		cmd = exec.Command("clear")
	default:
		panic("Your platform is unsupported!")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
