package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/khbdev/flash-cli/internal/config"
	"github.com/khbdev/flash-cli/internal/git"
	"github.com/khbdev/flash-cli/internal/ui"
)

// initCmd: flash init [branch]
func initCmd(args []string) error {
	branch := "main"
	switch {
	case len(args) > 0:
		branch = args[0]
	case git.IsRepo():
		// Joriy branchni default sifatida olish — eng kutilgan xatti-harakat.
		if current, err := git.CurrentBranch(); err == nil && current != "HEAD" {
			branch = current
		}
	}

	if err := config.SaveProject(config.Project{DefaultBranch: branch}); err != nil {
		return err
	}
	ui.Success("%s yaratildi (default branch: %s)", config.ProjectPath(), branch)
	return nil
}

// startCmd: flash start "xabar"
func startCmd(args []string) error {
	message := strings.TrimSpace(strings.Join(args, " "))
	if message == "" {
		return errors.New(`commit xabarini kiriting, masalan: flash start "login tuzatildi"`)
	}
	if !git.IsRepo() {
		return git.ErrNotRepo
	}

	branch, err := targetBranch()
	if err != nil {
		return err
	}

	if err := git.Add(); err != nil {
		return err
	}
	if !git.HasStagedChanges() {
		ui.Warn("Commit qilinadigan o'zgarish yo'q — faqat push qilinadi.")
	} else if err := git.Commit(message); err != nil {
		return err
	}
	if err := git.Push(branch); err != nil {
		return err
	}

	ui.Success("Push tayyor: origin/%s", branch)
	return nil
}

// targetBranch push qilinadigan branchni aniqlaydi: config'dagi default,
// config bo'lmasa — joriy branch.
func targetBranch() (string, error) {
	cfg, err := config.LoadProject()
	switch {
	case err == nil && cfg.DefaultBranch != "":
		return cfg.DefaultBranch, nil
	case err != nil && !errors.Is(err, config.ErrNoProject):
		return "", err
	}

	current, cerr := git.CurrentBranch()
	if cerr != nil {
		return "", config.ErrNoProject
	}
	ui.Warn("default branch sozlanmagan — joriy branch ishlatilmoqda: %s", current)
	return current, nil
}

// branchCmd: flash -b [branch]
func branchCmd(args []string) error {
	if len(args) == 0 {
		cfg, err := config.LoadProject()
		if err != nil {
			return err
		}
		fmt.Println(cfg.DefaultBranch)
		return nil
	}

	branch := args[0]
	cfg, err := config.LoadProject()
	if err != nil && !errors.Is(err, config.ErrNoProject) {
		return err
	}
	cfg.DefaultBranch = branch

	if err := config.SaveProject(cfg); err != nil {
		return err
	}
	ui.Success("Default branch: %s", branch)
	return nil
}
