package cli

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/khbdev/flash-cli/internal/config"
	"github.com/khbdev/flash-cli/internal/git"
	"github.com/khbdev/flash-cli/internal/github"
	"github.com/khbdev/flash-cli/internal/ui"
)

const apiTimeout = 30 * time.Second

// tokenCmd: flash token [-c | -d]
func tokenCmd(args []string) error {
	if len(args) == 0 {
		return showToken()
	}

	switch args[0] {
	case "-c", "set":
		return saveToken()
	case "-d", "delete":
		if err := config.SetGitHubToken(""); err != nil {
			return err
		}
		ui.Success("Token o'chirildi.")
		return nil
	default:
		return fmt.Errorf("noma'lum flag: %s (mavjud: -c, -d)", args[0])
	}
}

func saveToken() error {
	token, err := ui.Secret("GitHub Personal Access Token (ko'rinmaydi): ")
	if err != nil {
		return err
	}
	if token == "" {
		return errors.New("token bo'sh — saqlanmadi")
	}
	if err := config.SetGitHubToken(token); err != nil {
		return err
	}

	ui.Success("Token saqlandi (~/.flash/config.json, faqat siz o'qiy olasiz).")
	return verify(token)
}

func showToken() error {
	token, err := config.GitHubToken()
	if err != nil {
		return err
	}
	if token == "" {
		return github.ErrNoToken
	}

	ui.Plain("Token: %s", ui.Mask(token))
	return verify(token)
}

// verify token haqiqiy ishlashini GitHub'dan so'rab tekshiradi.
func verify(token string) error {
	client, err := github.New(token)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), apiTimeout)
	defer cancel()

	login, err := client.Login(ctx)
	if err != nil {
		ui.Warn("Tokenni tekshirib bo'lmadi: %v", err)
		return nil
	}
	ui.Success("GitHub: %s", login)
	return nil
}

// repoCmd: flash repo [nom] [-p]
func repoCmd(args []string) error {
	name, private, err := parseRepoArgs(args)
	if err != nil {
		return err
	}

	token, err := config.GitHubToken()
	if err != nil {
		return err
	}
	client, err := github.New(token)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), apiTimeout)
	defer cancel()

	ui.Step("GitHub repozitoriyasi yaratilmoqda: %s", name)
	repo, err := client.CreateRepo(ctx, name, private)
	if err != nil {
		return err
	}

	visibility := "public"
	if repo.Private {
		visibility = "private"
	}
	ui.Success("Yaratildi (%s): %s", visibility, repo.HTMLURL)
	return linkRemote(repo)
}

func parseRepoArgs(args []string) (name string, private bool, err error) {
	for _, arg := range args {
		switch arg {
		case "-p", "--private":
			private = true
		case "--public":
			private = false
		default:
			if strings.HasPrefix(arg, "-") {
				return "", false, fmt.Errorf("noma'lum flag: %s (mavjud: -p)", arg)
			}
			if name != "" {
				return "", false, errors.New("faqat bitta repo nomini kiriting")
			}
			name = arg
		}
	}

	if name == "" {
		// Nom berilmasa joriy papka nomini ishlatamiz.
		wd, wdErr := os.Getwd()
		if wdErr != nil {
			return "", false, errors.New("repo nomini kiriting: flash repo my-repo")
		}
		name = filepath.Base(wd)
	}
	return name, private, nil
}

// linkRemote joriy git repoda origin bo'lmasa, uni avtomatik ulaydi.
func linkRemote(repo *github.Repository) error {
	if !git.IsRepo() {
		ui.Plain("\nKeyingi qadamlar:")
		ui.Plain("  git init")
		ui.Plain("  git remote add origin %s", repo.CloneURL)
		ui.Plain("  flash init && flash start \"birinchi commit\"")
		return nil
	}
	if git.HasRemote("origin") {
		ui.Warn("origin allaqachon mavjud — qo'lda ulang: git remote add <nom> %s", repo.CloneURL)
		return nil
	}

	if err := git.AddRemote("origin", repo.CloneURL); err != nil {
		return err
	}
	ui.Success("origin ulandi: %s", repo.CloneURL)
	ui.Plain("Endi: flash start \"birinchi commit\"")
	return nil
}
