// Package box qayta ishlatiladigan bash skriptlarini ("box") ~/.flash/boxes
// ichida saqlaydi va ishga tushiradi.
package box

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/khbdev/flash-cli/internal/config"
)

const ext = ".box"

// ErrNotFound shu nomli box mavjud emasligini bildiradi.
var ErrNotFound = errors.New("box topilmadi")

const template = `#!/usr/bin/env bash
# flash box: %s
# Komandalarni shu yerga yozing. Ishga tushirish: flash box %s [argumentlar]
# Qo'shimcha argumentlar $1, $2 ... orqali keladi.
set -euo pipefail

`

// Store — box fayllari saqlanadigan papka ustidagi qatlam.
type Store struct {
	dir string
}

// NewStore ~/.flash/boxes papkasiga bog'langan Store qaytaradi.
func NewStore() (*Store, error) {
	home, err := config.HomeDir()
	if err != nil {
		return nil, err
	}
	return &Store{dir: filepath.Join(home, "boxes")}, nil
}

// Dir box fayllari papkasini qaytaradi.
func (s *Store) Dir() string { return s.dir }

// Path nomni tekshirib, box fayl yo'lini qaytaradi.
func (s *Store) Path(name string) (string, error) {
	if err := validateName(name); err != nil {
		return "", err
	}
	return filepath.Join(s.dir, name+ext), nil
}

// validateName papkadan chiqib ketadigan nomlarni ("../etc") rad etadi.
func validateName(name string) error {
	if name == "" {
		return errors.New("box nomi bo'sh")
	}
	if strings.HasPrefix(name, ".") || name != filepath.Base(name) ||
		strings.ContainsAny(name, `/\`) {
		return fmt.Errorf("box nomi noto'g'ri: %q", name)
	}
	return nil
}

// Exists shu nomli box borligini tekshiradi.
func (s *Store) Exists(name string) (bool, error) {
	path, err := s.Path(name)
	if err != nil {
		return false, err
	}
	_, err = os.Stat(path)
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, fs.ErrNotExist):
		return false, nil
	default:
		return false, err
	}
}

// List mavjud box nomlarini alifbo tartibida qaytaradi.
func (s *Store) List() ([]string, error) {
	entries, err := os.ReadDir(s.dir)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("%s o'qib bo'lmadi: %w", s.dir, err)
	}

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ext) {
			names = append(names, strings.TrimSuffix(entry.Name(), ext))
		}
	}
	sort.Strings(names)
	return names, nil
}

// Create yangi box faylini shablon bilan yaratadi. Fayl allaqachon mavjud
// bo'lsa, uni o'zgartirmaydi va created=false qaytaradi.
func (s *Store) Create(name string) (path string, created bool, err error) {
	path, err = s.Path(name)
	if err != nil {
		return "", false, err
	}
	if err := os.MkdirAll(s.dir, 0o700); err != nil {
		return "", false, fmt.Errorf("%s papkasini yaratib bo'lmadi: %w", s.dir, err)
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o700)
	if errors.Is(err, fs.ErrExist) {
		return path, false, nil
	}
	if err != nil {
		return "", false, fmt.Errorf("%s yaratib bo'lmadi: %w", path, err)
	}
	defer file.Close()

	if _, err := fmt.Fprintf(file, template, name, name); err != nil {
		return "", false, fmt.Errorf("%s yozib bo'lmadi: %w", path, err)
	}
	return path, true, nil
}

// Remove box faylini o'chiradi.
func (s *Store) Remove(name string) error {
	path, err := s.Path(name)
	if err != nil {
		return err
	}
	err = os.Remove(path)
	if errors.Is(err, fs.ErrNotExist) {
		return fmt.Errorf("%w: %s", ErrNotFound, name)
	}
	if err != nil {
		return fmt.Errorf("%s o'chirib bo'lmadi: %w", path, err)
	}
	return nil
}

// Run box ichidagi komandalarni bajaradi. args skript ichida $1, $2 ... bo'ladi.
func (s *Store) Run(name string, args []string) error {
	path, err := s.Path(name)
	if err != nil {
		return err
	}
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("%w: %s", ErrNotFound, name)
	}

	cmd := exec.Command("bash", append([]string{path}, args...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("box %q bajarilmadi: %w", name, err)
	}
	return nil
}

// Edit box faylini foydalanuvchi tanlagan muharrirda ochadi.
func (s *Store) Edit(name string) error {
	path, err := s.Path(name)
	if err != nil {
		return err
	}
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("%w: %s", ErrNotFound, name)
	}
	return OpenEditor(path)
}

// OpenEditor $VISUAL / $EDITOR ni hurmat qiladi, topilmasa nano yoki vi'ga o'tadi.
func OpenEditor(path string) error {
	name := pickEditor()
	if name == "" {
		return fmt.Errorf("muharrir topilmadi — $EDITOR ni o'rnating yoki faylni qo'lda oching: %s", path)
	}

	cmd := exec.Command(name, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("muharrirni ochib bo'lmadi (%s): %w", name, err)
	}
	return nil
}

func pickEditor() string {
	for _, candidate := range []string{os.Getenv("VISUAL"), os.Getenv("EDITOR"), "nano", "vim", "vi"} {
		if candidate == "" {
			continue
		}
		if _, err := exec.LookPath(candidate); err == nil {
			return candidate
		}
	}
	return ""
}
