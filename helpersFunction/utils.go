package helpersFunction

import (
	"fmt"
	"os"
	"os/exec"
)


func RunCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		LogRed(fmt.Sprintf("Xatolik: %s", err))
	}
}


func LogGreen(text string) {
	fmt.Printf("\033[32m%s\033[0m\n", text)
}


func LogRed(text string) {
	fmt.Printf("\033[31m%s\033[0m\n", text)
}


func WriteFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}


func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}
