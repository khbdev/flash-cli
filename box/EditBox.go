package box

import (
	"fmt"
	"os"
	"os/exec"
)

func EditBox(name string) {
	boxFilePath, err := getBoxPath(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	if _, err := os.Stat(boxFilePath); os.IsNotExist(err) {
		fmt.Println(" Bu nomli box topilmadi:", name)
		return
	}


	cmd := exec.Command("nano", boxFilePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(" Faylni tahrirlab bo‘lmadi:", err)
		return
	}

	fmt.Println(" Box tahrirlandi:", name)
}
