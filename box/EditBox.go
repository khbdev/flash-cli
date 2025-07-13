package box

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func EditBox(name string) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Home directory topilmadi:", err)
		return
	}


	boxPath := filepath.Join(homeDir, ".flash", "boxes", name+".box")

	
	if _, err := os.Stat(boxPath); os.IsNotExist(err) {
		fmt.Println("❌ Bu nomli box topilmadi:", name)
		return
	}


	cmd := exec.Command("nano", boxPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("❌ Faylni tahrirlab bo‘lmadi:", err)
		return
	}

	fmt.Println("✅ Box tahrirlandi:", name)
}
