package box

import (
	"fmt"
	"os"
	"path/filepath"
)

func RemoveBox(name string) {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Home directory topilmadi:", err)
		return
	}


	boxPath := filepath.Join(homeDir, ".flash", "boxes", name+".box")


	if _, err := os.Stat(boxPath); os.IsNotExist(err) {
		fmt.Println("❌ Box topilmadi:", name)
		return
	}

	
	err = os.Remove(boxPath)
	if err != nil {
		fmt.Println("❌ O‘chirishda xatolik:", err)
		return
	}

	fmt.Println("🗑️ Box o‘chirildi:", name)
}
