package box

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)


func CreateBox(name string){
 
	homeDir, err := os.UserHomeDir()
		if err != nil {
		fmt.Println(" Home directory topilmadi:", err)
		return
	}

	boxDir := filepath.Join(homeDir, ".flash", "boxes")
	boxFilePath := filepath.Join(boxDir, name+".box")

	err = os.MkdirAll(boxDir, 0755)
	if err != nil {
		fmt.Println(" Papka yaratib bo‘lmadi:", err)
		return
	}

	if _, err := os.Stat(boxFilePath); os.IsNotExist(err){
		file, err := os.Create(boxFilePath)
		if err != nil {
			fmt.Println(" Fayl yaratishda xatolik:", err)
			return
		}
		file.Close()
		fmt.Println(" Yangi box fayl yaratildi:", boxFilePath)
	} else {
		fmt.Println(" Box fayl allaqachon mavjud:", boxFilePath)
	}

		cmd := exec.Command("nano", boxFilePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr


	err = cmd.Run()

		if err != nil {
		fmt.Println(" Editorni ochishda xatolik:", err)
		return
	}


	fmt.Println(" Box tayyor:", name)


}