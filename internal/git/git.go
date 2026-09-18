// Package git — git ustidan yupqa qatlam. Har bir funksiya `git` ni
// argumentlar ro'yxati bilan chaqiradi (shell orqali emas), shuning uchun
// commit xabaridagi tirnoq yoki `$(...)` kod sifatida bajarilmaydi.
package git

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ErrNotRepo joriy papka git repozitoriya emasligini bildiradi.
var ErrNotRepo = errors.New("bu papka git repozitoriya emas — `git init` qiling")

// Run git'ni ishga tushiradi va chiqishini to'g'ridan-to'g'ri terminalga uzatadi.
func Run(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git %s: %w", strings.Join(args, " "), err)
	}
	return nil
}

// Output git'ni ishga tushiradi va stdout'ini qaytaradi.
func Output(args ...string) (string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("git", args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		if msg := strings.TrimSpace(stderr.String()); msg != "" {
			return "", fmt.Errorf("git %s: %s", strings.Join(args, " "), msg)
		}
		return "", fmt.Errorf("git %s: %w", strings.Join(args, " "), err)
	}
	return strings.TrimSpace(stdout.String()), nil
}

// IsRepo joriy papka git repozitoriya ichidami-yo'qmi.
func IsRepo() bool {
	_, err := Output("rev-parse", "--git-dir")
	return err == nil
}

// CurrentBranch joriy branch nomini qaytaradi.
func CurrentBranch() (string, error) {
	return Output("rev-parse", "--abbrev-ref", "HEAD")
}

// HasStagedChanges commit qilish uchun staged o'zgarish bor-yo'qligini aytadi.
// `git diff --cached --quiet` farq bo'lsa nolga teng bo'lmagan kod qaytaradi.
func HasStagedChanges() bool {
	return exec.Command("git", "diff", "--cached", "--quiet").Run() != nil
}

// HasRemote berilgan nomli remote mavjudligini tekshiradi.
func HasRemote(name string) bool {
	out, err := Output("remote")
	if err != nil {
		return false
	}
	for _, line := range strings.Fields(out) {
		if line == name {
			return true
		}
	}
	return false
}

// Add barcha o'zgarishlarni staging'ga qo'shadi.
func Add() error { return Run("add", "-A") }

// Commit staged o'zgarishlarni commit qiladi.
func Commit(message string) error { return Run("commit", "-m", message) }

// Push branchni origin'ga yuboradi va upstream sifatida belgilaydi.
func Push(branch string) error { return Run("push", "-u", "origin", branch) }

// AddRemote yangi remote qo'shadi.
func AddRemote(name, url string) error { return Run("remote", "add", name, url) }
