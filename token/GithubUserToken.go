package token

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var tokenFile = os.Getenv("HOME") + "/.flash_token"

func Token(args []string) {
	if len(args) > 0 {
		switch args[0] {
		case "-c":
			createToken()
			return
		case "-d":
			deleteToken()
			return
		default:
			fmt.Println("Noma'lum flag:", args[0])
			return
		}
	}

	ShowToken()
}

func ShowToken() {
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		fmt.Println("Token topilmadi.")
		return
	}
	fmt.Println("Token:", strings.TrimSpace(string(data)))
	
}

func createToken() {
	cmd := exec.Command("sudo", "nano", tokenFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Token yozishda xatolik:", err)
		return
	}

	fmt.Println("Token saqlandi:", tokenFile)
}

func deleteToken() {
	err := os.Remove(tokenFile)
	if err != nil {
		fmt.Println(" Token o‘chirishda xatolik:", err)
		return
	}
	fmt.Println("Token o‘chirildi.")
}


func GetToken() string {
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		fmt.Println("Token topilmadi.")
		return ""
	}
	return strings.TrimSpace(string(data))
}