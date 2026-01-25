package box

import (
	"fmt"
	"os"

	"strings"
)

func BoxStatus() {
	
		boxFilePath, err := getBoxDir()
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := os.ReadDir(boxFilePath)
	if err != nil {
		fmt.Println("Box papkasi topilmadi:", err)
		return
	}


	if len(files) == 0 {
		fmt.Println(" Hech qanday box topilmadi.")
		return
	}

	fmt.Println(" Topilgan box'lar:")
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".box") {
			boxName := strings.TrimSuffix(file.Name(), ".box")
			fmt.Println(" -", boxName)
		}
	}
}
