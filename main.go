package main

import (
	"flash/flashcore"
	"flash/gitfile"
	"fmt"
	"os"
	"time"
)

func main() {
	// ❗ Argumentlar tekshiruvi
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	// 🔀 Variant: flash -b dev  → default branch o‘zgartirish
	if len(os.Args) == 3 && os.Args[1] == "-b" {
		newBranch := os.Args[2]
		err := flashcore.UpdateDefaultBranch(newBranch)
		if err != nil {
			fmt.Println("❌ Branch yangilab bo‘lmadi:", err)
			return
		}
		fmt.Println("✅ Default branch yangilandi:", newBranch)
		return
	}

	// 🔄 Komanda: init | start | help
	command := os.Args[1]

	switch command {
	case "init":
		flashcore.InitProject()
		time.Sleep(time.Millisecond * 1) // kichik delay

	case "start":
		if len(os.Args) < 3 {
			fmt.Println("❗ start uchun commit xabarini kiriting. Masalan: flash start \"Initial commit\"")
			return
		}

		msg := os.Args[2]
		gitfile.Start(msg)

	case "help":
		printHelp()

	default:
		fmt.Println("❓ Noma’lum komanda:", command)
		printHelp()
	}
}

// 📘 CLI yordam funksiyasi
func printHelp() {
	fmt.Println("🛠️ Yordam:")
	fmt.Println("  init               → Loyihani boshlash (.flash papka va config)")
	fmt.Println("  start \"msg\"        → Git add, commit, push bajarish")
	fmt.Println("  -b 'branch'        → Default branchni o‘zgartirish (config.json)")
	fmt.Println("  help               → Yordam oynasi")
}
