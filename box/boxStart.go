package box

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunBox(name string) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Home directory topilmadi:", err)
		return
	}


	boxPath := filepath.Join(homeDir, ".flash", "boxes", name+".box")

	
	if _, err := os.Stat(boxPath); os.IsNotExist(err) {
		fmt.Println(" Box topilmadi:", name)
		return
	}

	
	content, err := os.ReadFile(boxPath)
	if err != nil {
		fmt.Println("❌ Faylni o‘qishda xatolik:", err)
		return
	}

	fmt.Printf("🚀 Box ishga tushyapti: %s\n\n", name)

	
	cmd := exec.Command("bash", "-c", string(content))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("❌ Komandalarni bajarishda xatolik:", err)
		return
	}

	fmt.Println("\n✅ Box bajarildi:", name)
}
