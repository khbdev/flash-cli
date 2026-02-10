package box

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CreateBox(name string) {

	boxFilePath, err := getBoxPath(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	boxDir := filepath.Dir(boxFilePath)

	err = os.MkdirAll(boxDir, 0755)
	if err != nil {
		fmt.Println(" Papka yaratib bo‘lmadi:", err)
		return
	}

	if _, err := os.Stat(boxFilePath); os.IsNotExist(err) {
		file, err := os.Create(boxFilePath)
		if err != nil {
			fmt.Println(" Fayl yaratishd:", err)
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
