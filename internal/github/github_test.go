package github

import (
	"strings"
	"testing"
)

func TestAPIErrorUsesGitHubMessage(t *testing.T) {
	body := []byte(`{"message":"Repository creation failed.","errors":[{"field":"name","code":"custom","message":"name already exists on this account"}]}`)

	err := APIError(422, body)
	for _, want := range []string{"422", "Repository creation failed.", "already exists"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("APIError() = %q, %q ni o'z ichiga olishi kerak", err, want)
		}
	}
}

func TestAPIErrorFallsBackToStatusText(t *testing.T) {
	err := APIError(401, []byte("html emas json"))
	if !strings.Contains(err.Error(), "Unauthorized") {
		t.Errorf("APIError() = %q, \"Unauthorized\" kutilgan", err)
	}
}

func TestNewRejectsEmptyToken(t *testing.T) {
	if _, err := New(""); err != ErrNoToken {
		t.Errorf("New(\"\") = %v, ErrNoToken kutilgan", err)
	}
}
