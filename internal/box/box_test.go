package box

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestPathRejectsTraversal(t *testing.T) {
	store := &Store{dir: t.TempDir()}

	for _, name := range []string{"", "../etc/passwd", "a/b", `a\b`, ".hidden", "."} {
		if _, err := store.Path(name); err == nil {
			t.Errorf("Path(%q) xato qaytarishi kerak edi", name)
		}
	}
}

func TestPathAcceptsPlainNames(t *testing.T) {
	dir := t.TempDir()
	store := &Store{dir: dir}

	got, err := store.Path("laravel-9_x")
	if err != nil {
		t.Fatalf("kutilmagan xato: %v", err)
	}
	if want := filepath.Join(dir, "laravel-9_x.box"); got != want {
		t.Errorf("Path() = %q, kutilgan %q", got, want)
	}
}

func TestCreateIsIdempotent(t *testing.T) {
	store := &Store{dir: filepath.Join(t.TempDir(), "boxes")}

	path, created, err := store.Create("deploy")
	if err != nil || !created {
		t.Fatalf("Create() = %q, %v, %v; yangi fayl kutilgan edi", path, created, err)
	}

	_, created, err = store.Create("deploy")
	if err != nil {
		t.Fatalf("kutilmagan xato: %v", err)
	}
	if created {
		t.Error("ikkinchi Create() mavjud faylni qayta yaratmasligi kerak")
	}
}

func TestListSkipsNonBoxFiles(t *testing.T) {
	store := &Store{dir: filepath.Join(t.TempDir(), "boxes")}
	for _, name := range []string{"zeta", "alpha"} {
		if _, _, err := store.Create(name); err != nil {
			t.Fatal(err)
		}
	}

	names, err := store.List()
	if err != nil {
		t.Fatal(err)
	}
	if got := strings.Join(names, ","); got != "alpha,zeta" {
		t.Errorf("List() = %q, kutilgan \"alpha,zeta\"", got)
	}
}

func TestListOnMissingDir(t *testing.T) {
	store := &Store{dir: filepath.Join(t.TempDir(), "yoq")}

	names, err := store.List()
	if err != nil {
		t.Fatalf("mavjud bo'lmagan papka xato bermasligi kerak: %v", err)
	}
	if len(names) != 0 {
		t.Errorf("List() = %v, bo'sh kutilgan", names)
	}
}
