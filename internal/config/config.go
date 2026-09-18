// Package config ikkita konfiguratsiyani boshqaradi:
//   - loyihaviy: ./.flash/config.json  (default branch)
//   - global:    ~/.flash/config.json  (GitHub token, 0600)
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	projectDir  = ".flash"
	projectFile = "config.json"
	globalFile  = "config.json"

	// TokenEnv o'rnatilgan bo'lsa, faylga yozilgan tokendan ustun turadi.
	TokenEnv = "FLASH_GITHUB_TOKEN"
)

// ErrNoProject .flash/config.json mavjud emasligini bildiradi.
var ErrNoProject = errors.New("loyiha sozlanmagan — `flash init` ni ishga tushiring")

// Project — loyihaga xos sozlamalar (repo ichida, git'ga commit qilinadi).
type Project struct {
	DefaultBranch string `json:"default_branch"`
}

// Global — foydalanuvchiga xos sozlamalar (home papkada, maxfiy).
type Global struct {
	GitHubToken string `json:"github_token,omitempty"`
}

// ProjectPath loyihaviy config fayl yo'lini qaytaradi.
func ProjectPath() string { return filepath.Join(projectDir, projectFile) }

// LoadProject loyihaviy configni o'qiydi. Fayl bo'lmasa ErrNoProject qaytadi.
func LoadProject() (Project, error) {
	var p Project
	data, err := os.ReadFile(ProjectPath())
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return p, ErrNoProject
	case err != nil:
		return p, fmt.Errorf("%s o'qib bo'lmadi: %w", ProjectPath(), err)
	}
	if err := json.Unmarshal(data, &p); err != nil {
		return p, fmt.Errorf("%s noto'g'ri JSON: %w", ProjectPath(), err)
	}
	return p, nil
}

// SaveProject loyihaviy configni yozadi, kerak bo'lsa .flash papkasini yaratadi.
func SaveProject(p Project) error {
	if err := os.MkdirAll(projectDir, 0o755); err != nil {
		return fmt.Errorf("%s papkasini yaratib bo'lmadi: %w", projectDir, err)
	}
	return writeJSON(ProjectPath(), p, 0o644)
}

// HomeDir ~/.flash papkasini qaytaradi va kerak bo'lsa yaratadi.
func HomeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home papka topilmadi: %w", err)
	}
	dir := filepath.Join(home, ".flash")
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return "", fmt.Errorf("%s papkasini yaratib bo'lmadi: %w", dir, err)
	}
	return dir, nil
}

func globalPath() (string, error) {
	dir, err := HomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, globalFile), nil
}

// LoadGlobal global configni o'qiydi. Fayl bo'lmasa bo'sh struct qaytadi.
func LoadGlobal() (Global, error) {
	var g Global
	path, err := globalPath()
	if err != nil {
		return g, err
	}
	data, err := os.ReadFile(path)
	switch {
	case errors.Is(err, fs.ErrNotExist):
		return g, nil
	case err != nil:
		return g, fmt.Errorf("%s o'qib bo'lmadi: %w", path, err)
	}
	if err := json.Unmarshal(data, &g); err != nil {
		return g, fmt.Errorf("%s noto'g'ri JSON: %w", path, err)
	}
	return g, nil
}

// SaveGlobal global configni faqat egasi o'qiy oladigan qilib (0600) yozadi.
func SaveGlobal(g Global) error {
	path, err := globalPath()
	if err != nil {
		return err
	}
	return writeJSON(path, g, 0o600)
}

// GitHubToken tokenni shu tartibda qidiradi: muhit o'zgaruvchisi →
// ~/.flash/config.json → eski ~/.flash_token (orqaga moslik uchun).
func GitHubToken() (string, error) {
	if token := strings.TrimSpace(os.Getenv(TokenEnv)); token != "" {
		return token, nil
	}
	g, err := LoadGlobal()
	if err != nil {
		return "", err
	}
	if g.GitHubToken != "" {
		return g.GitHubToken, nil
	}
	return legacyToken(), nil
}

// SetGitHubToken tokenni global configga saqlaydi. Bo'sh qiymat uni o'chiradi.
func SetGitHubToken(token string) error {
	g, err := LoadGlobal()
	if err != nil {
		return err
	}
	g.GitHubToken = strings.TrimSpace(token)
	return SaveGlobal(g)
}

// legacyToken v2 da ishlatilgan ~/.flash_token faylini o'qiydi.
func legacyToken() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	data, err := os.ReadFile(filepath.Join(home, ".flash_token"))
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}

// writeJSON faylni atomik yozadi: avval vaqtinchalik faylga, keyin rename.
func writeJSON(path string, value any, perm os.FileMode) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("JSONga o'girib bo'lmadi: %w", err)
	}
	data = append(data, '\n')

	tmp, err := os.CreateTemp(filepath.Dir(path), ".flash-*.tmp")
	if err != nil {
		return fmt.Errorf("%s yozib bo'lmadi: %w", path, err)
	}
	defer os.Remove(tmp.Name())

	if err := tmp.Chmod(perm); err != nil {
		tmp.Close()
		return fmt.Errorf("%s huquqlarini o'rnatib bo'lmadi: %w", path, err)
	}
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return fmt.Errorf("%s yozib bo'lmadi: %w", path, err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("%s yopib bo'lmadi: %w", path, err)
	}
	if err := os.Rename(tmp.Name(), path); err != nil {
		return fmt.Errorf("%s ni saqlab bo'lmadi: %w", path, err)
	}
	return nil
}
