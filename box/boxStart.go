package box

import (
	"fmt"
	"os"
	"os/exec"

)

func RunBox(name string) {

boxFilePath, err := getBoxPath(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	if _, err := os.Stat(boxFilePath); os.IsNotExist(err) {
		fmt.Println("Box topilmadi:", name)
		return
	}

	
	content, err := os.ReadFile(boxFilePath)
	if err != nil {
		fmt.Println("Faylni o‘qishda xatolik:", err)
		return
	}

	fmt.Printf("Box ishga tushyapti: %s\n\n", name)

	
	cmd := exec.Command("bash", "-c", string(content))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println(" Komandalarni bajarishda xatolik:", err)
		return
	}

	fmt.Println("\n Box bajarildi:", name)
}
