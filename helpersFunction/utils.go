package helpersFunction

import (
	"fmt"
	"os"
	"os/exec"
)

// 📦 Terminal komandani bajaradi
func RunCommand(command string) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		LogRed(fmt.Sprintf("❌ Xatolik: %s", err))
	}
}

// 📗 Yashil rangdagi log
func LogGreen(text string) {
	fmt.Printf("\033[32m%s\033[0m\n", text)
}

// 📕 Qizil rangdagi log
func LogRed(text string) {
	fmt.Printf("\033[31m%s\033[0m\n", text)
}

// 📝 Faylga yozish (agar mavjud bo‘lmasa, yaratadi)
func WriteFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// 📖 Fayldan o‘qish
func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}
