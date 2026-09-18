package cli

import (
	"errors"
	"fmt"

	"github.com/khbdev/flash-cli/internal/box"
	"github.com/khbdev/flash-cli/internal/ui"
)

// boxCmd: flash box [-s | -c <nom> | -e <nom> | -r <nom> | <nom> [args...]]
func boxCmd(args []string) error {
	store, err := box.NewStore()
	if err != nil {
		return err
	}
	if len(args) == 0 {
		return errors.New("box komandasi kerak: flash box -s | -c <nom> | <nom>")
	}

	switch args[0] {
	case "-s", "-l", "list":
		return listBoxes(store)
	case "-c", "create":
		name, err := boxName(args, "yaratish", "flash box -c laravel")
		if err != nil {
			return err
		}
		return createBox(store, name)
	case "-e", "edit":
		name, err := boxName(args, "tahrirlash", "flash box -e laravel")
		if err != nil {
			return err
		}
		if err := store.Edit(name); err != nil {
			return err
		}
		ui.Success("Box saqlandi: %s", name)
		return nil
	case "-r", "remove":
		name, err := boxName(args, "o'chirish", "flash box -r laravel")
		if err != nil {
			return err
		}
		if err := store.Remove(name); err != nil {
			return err
		}
		ui.Success("Box o'chirildi: %s", name)
		return nil
	default:
		return store.Run(args[0], args[1:])
	}
}

func boxName(args []string, action, example string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("qaysi box'ni %s kerak? Masalan: %s", action, example)
	}
	return args[1], nil
}

func listBoxes(store *box.Store) error {
	names, err := store.List()
	if err != nil {
		return err
	}
	if len(names) == 0 {
		ui.Warn("Box topilmadi. Yangisini yarating: flash box -c <nom>")
		return nil
	}

	ui.Plain("Box'lar (%s):", store.Dir())
	for _, name := range names {
		ui.Plain("  • %s", name)
	}
	return nil
}

func createBox(store *box.Store, name string) error {
	path, created, err := store.Create(name)
	if err != nil {
		return err
	}
	if !created {
		ui.Warn("Box allaqachon mavjud — tahrirlash uchun ochilmoqda: %s", name)
	}
	if err := box.OpenEditor(path); err != nil {
		return err
	}
	ui.Success("Box tayyor: %s → flash box %s", name, name)
	return nil
}
