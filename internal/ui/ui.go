// Package ui terminalga chiqarish va foydalanuvchidan ma'lumot so'rash uchun
// yagona nuqta. Domen paketlari hech narsa chop etmaydi — faqat shu paket va
// cli qatlami chop etadi.
package ui

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var color = colorEnabled()

func colorEnabled() bool {
	if os.Getenv("NO_COLOR") != "" || os.Getenv("TERM") == "dumb" {
		return false
	}
	info, err := os.Stdout.Stat()
	return err == nil && info.Mode()&os.ModeCharDevice != 0
}

func paint(code, text string) string {
	if !color {
		return text
	}
	return "\033[" + code + "m" + text + "\033[0m"
}

// Success yashil ✓ bilan muvaffaqiyatli natijani chiqaradi.
func Success(format string, a ...any) { line("32", "✓", format, a...) }

// Step ko'k → bilan bajarilayotgan qadamni chiqaradi.
func Step(format string, a ...any) { line("36", "→", format, a...) }

// Warn sariq ! bilan ogohlantirish chiqaradi.
func Warn(format string, a ...any) { line("33", "!", format, a...) }

// Plain hech qanday belgisiz matn chiqaradi.
func Plain(format string, a ...any) { fmt.Printf(format+"\n", a...) }

func line(code, mark, format string, a ...any) {
	fmt.Println(paint(code, mark) + " " + fmt.Sprintf(format, a...))
}

// Secret terminalda ko'rinmaydigan qilib bir qator o'qiydi (parol/token uchun).
// stty mavjud bo'lmasa oddiy o'qishga qaytadi.
func Secret(label string) (string, error) {
	fmt.Print(label)
	restore := disableEcho()
	defer restore()

	value, err := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println()
	if err != nil && value == "" {
		return "", fmt.Errorf("kiritishni o'qib bo'lmadi: %w", err)
	}
	return strings.TrimSpace(value), nil
}

func disableEcho() func() {
	stty := func(arg string) error {
		cmd := exec.Command("stty", arg)
		cmd.Stdin = os.Stdin
		return cmd.Run()
	}
	if stty("-echo") != nil {
		return func() {}
	}
	return func() { _ = stty("echo") }
}

// Mask tokenni faqat oxirgi 4 belgisini ko'rsatib yashiradi.
func Mask(secret string) string {
	if len(secret) <= 4 {
		return strings.Repeat("*", len(secret))
	}
	return strings.Repeat("*", 8) + secret[len(secret)-4:]
}
