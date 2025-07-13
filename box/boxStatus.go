package box

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func BoxStatus() {
	
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Home directory topilmadi:", err)
		return
	}


	boxDir := filepath.Join(homeDir, ".flash", "boxes")


	files, err := os.ReadDir(boxDir)
	if err != nil {
		fmt.Println("❌ Box papkasi topilmadi:", err)
		return
	}


	if len(files) == 0 {
		fmt.Println("📭 Hech qanday box topilmadi.")
		return
	}

	fmt.Println("📦 Topilgan box'lar:")
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".box") {
			boxName := strings.TrimSuffix(file.Name(), ".box")
			fmt.Println(" -", boxName)
		}
	}
}
