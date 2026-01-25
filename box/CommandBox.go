package box

import (
	"fmt"

)

func CommandBox(args []string) {
	if len(args) == 0 {
		fmt.Println("❗ Box komandasi kiritilmadi.")
		fmt.Println("  Misollar:")
		fmt.Println("    flash box -c laravel")
		fmt.Println("    flash box laravel")
		fmt.Println("    flash box status")
		return
	}

	switch args[0] {

	case "-c":
		if len(args) < 2 {
			fmt.Println("Qanday box yaratish kerakligini yozing: flash box -c golang")
			return
		}
	CreateBox(args[1])

	case "-s":
	BoxStatus()

	case "-e":
		if len(args) < 2 {
			fmt.Println("Qaysi box'ni tahrirlash kerak: flash box edit golang")
			return
		}
	EditBox(args[1])

	case "-r":
		if len(args) < 2 {
			fmt.Println("Qaysi box'ni o‘chirish kerak: flash box remove golang")
			return
		}
	RemoveBox(args[1])

	default:
	RunBox(args[0])
	}
}
