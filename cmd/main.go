package main

import (
	"fmt"

	"os"
	"flash/box"
	"flash/configbranchedit"
	"flash/flashcore"
	"flash/gitfile"
	"flash/helpfunction"
	"flash/repo"
	"flash/token"
)

func main() {

	if len(os.Args) < 2 {
		helpfunction.PrintHelp()
		return
	}

	command := os.Args[1]


	if command == "-b" {
		if configbranchedit.Configbranchedit() {
			return
		}
	}


	switch command {
	case "init":
		flashcore.InitProject()
	case "start":
		if len(os.Args) < 3 {
			fmt.Println("start uchun commit xabarini kiriting. Masalan: flash start flash")
			return
		}
		msg := os.Args[2]
		gitfile.Start(msg)

	case "box":
		box.CommandBox(os.Args[2:])

	case "repo":
		repo.Repo()

	case "token":
		token.Token(os.Args[2:])

	case "help":
		helpfunction.PrintHelp()

	default:
		fmt.Println("Noma’lum komanda:", command)
		helpfunction.PrintHelp()
	}
}
