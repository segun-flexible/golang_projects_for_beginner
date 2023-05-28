package helpers

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "linux", "darwin":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		// Unsupported platform, do nothing
	}
}

func ScanInput(name *string) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	*name = scanner.Text()
}
