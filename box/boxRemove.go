package box

import (
	"fmt"
	"os"
)

func RemoveBox(name string) {

		boxFilePath, err := getBoxPath(name)
	if err != nil {
		fmt.Println(err)
		return
	}
   
	if _, err := os.Stat(boxFilePath); os.IsNotExist(err) {
		fmt.Println("Box topilmadi:", name)
		return
	}

	
	err = os.Remove(boxFilePath)
	if err != nil {
		fmt.Println("O‘chirishda xatolik:", err)
		return
	}

	fmt.Println("Box o‘chirildi:", name)
}
