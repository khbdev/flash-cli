package flashcore

import (
	"fmt"
	"os"
)

func InitProject() {
	err := os.MkdirAll(".flash", 0755)
	if err != nil {
		fmt.Println("Papka yaratishda xatolik:", err)
		return
	}

	cfg := Config{
		DefaultBranch: "main",
	}

	err = SaveConfig(cfg)
	if err != nil {
		fmt.Println("Configni saqlashda xatolik:", err)
		return
	}

	fmt.Println("Init tugadi. Default branch: main")
}
