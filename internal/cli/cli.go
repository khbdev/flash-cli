// Package cli komandalarni tahlil qiladi va domen paketlarini chaqiradi.
// Faqat shu qatlam chop etadi va xatoni main'ga qaytaradi — main esa exit
// kodini belgilaydi.
package cli

import "fmt"

// Version — build vaqtida -ldflags orqali almashtiriladi.
var Version = "dev"

// Run bitta komandani bajaradi. args — os.Args[1:].
func Run(args []string) error {
	if len(args) == 0 {
		PrintHelp()
		return nil
	}

	command, rest := args[0], args[1:]
	switch command {
	case "init":
		return initCmd(rest)
	case "start":
		return startCmd(rest)
	case "-b", "branch":
		return branchCmd(rest)
	case "box":
		return boxCmd(rest)
	case "repo":
		return repoCmd(rest)
	case "token":
		return tokenCmd(rest)
	case "help", "-h", "--help":
		PrintHelp()
		return nil
	case "version", "-v", "--version":
		fmt.Println("flash", Version)
		return nil
	default:
		PrintHelp()
		return fmt.Errorf("noma'lum komanda: %s", command)
	}
}
