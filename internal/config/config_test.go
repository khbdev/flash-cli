package config

import (
	"errors"
	"os"
	"testing"
)

// chdirTemp testni vaqtinchalik papkaga ko'chiradi (loyihaviy config uchun).
func chdirTemp(t *testing.T) {
	t.Helper()
	t.Chdir(t.TempDir())
}

func TestLoadProjectMissing(t *testing.T) {
	chdirTemp(t)

	if _, err := LoadProject(); !errors.Is(err, ErrNoProject) {
		t.Errorf("LoadProject() = %v, ErrNoProject kutilgan", err)
	}
}

func TestSaveLoadProjectRoundTrip(t *testing.T) {
	chdirTemp(t)

	if err := SaveProject(Project{DefaultBranch: "develop"}); err != nil {
		t.Fatal(err)
	}
	got, err := LoadProject()
	if err != nil {
		t.Fatal(err)
	}
	if got.DefaultBranch != "develop" {
		t.Errorf("DefaultBranch = %q, \"develop\" kutilgan", got.DefaultBranch)
	}
}

func TestLoadProjectInvalidJSON(t *testing.T) {
	chdirTemp(t)

	if err := os.MkdirAll(projectDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(ProjectPath(), []byte("{bu json emas"), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := LoadProject(); err == nil {
		t.Error("buzilgan JSON uchun xato kutilgan edi")
	}
}

func TestGlobalTokenIsPrivate(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv(TokenEnv, "")

	if err := SetGitHubToken("ghp_secret"); err != nil {
		t.Fatal(err)
	}

	path, err := globalPath()
	if err != nil {
		t.Fatal(err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if perm := info.Mode().Perm(); perm != 0o600 {
		t.Errorf("token fayli huquqi = %o, 600 kutilgan", perm)
	}

	token, err := GitHubToken()
	if err != nil {
		t.Fatal(err)
	}
	if token != "ghp_secret" {
		t.Errorf("GitHubToken() = %q", token)
	}
}

func TestEnvTokenWins(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	if err := SetGitHubToken("fayldagi"); err != nil {
		t.Fatal(err)
	}
	t.Setenv(TokenEnv, "muhitdagi")

	token, err := GitHubToken()
	if err != nil {
		t.Fatal(err)
	}
	if token != "muhitdagi" {
		t.Errorf("GitHubToken() = %q, muhit o'zgaruvchisi ustun turishi kerak", token)
	}
}
