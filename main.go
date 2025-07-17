package main

import (
	"flash/box"
	"flash/configbranchedit"

	"flash/flashcore"
	"flash/gitfile"
	"flash/helpfunction"
	"fmt"
	"os"
	"time"
)

func main() {
	// ❗ Argumentlar tekshiruvi
	if len(os.Args) < 2 {
		helpfunction.PrintHelp()
		return
	}

	// 🔀 Variant: flash -b dev  → default branch o‘zgartirish
    if configbranchedit.Configbranchedit(){
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
		
	case "box":
	box.CommandBox(os.Args[2:])

	case "help":
		  helpfunction.PrintHelp()

	default:
		fmt.Println("❓ Noma’lum komanda:", command)
		helpfunction.PrintHelp()
	}
}



