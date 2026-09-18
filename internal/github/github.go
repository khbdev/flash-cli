// Package github GitHub REST API bilan ishlaydi (tashqi kutubxonasiz).
package github

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	baseURL      = "https://api.github.com"
	maxBodyBytes = 1 << 20 // 1 MiB — javobni o'qishda yuqori chegara
)

// ErrNoToken token sozlanmaganini bildiradi.
var ErrNoToken = errors.New("GitHub token topilmadi — `flash token -c` bilan saqlang")

// Client — autentifikatsiya qilingan GitHub API klienti.
type Client struct {
	token string
	http  *http.Client
}

// New berilgan token bilan klient yaratadi.
func New(token string) (*Client, error) {
	if token == "" {
		return nil, ErrNoToken
	}
	return &Client{token: token, http: &http.Client{Timeout: 20 * time.Second}}, nil
}

// Repository — API qaytaradigan repozitoriya ma'lumotlari.
type Repository struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	HTMLURL  string `json:"html_url"`
	SSHURL   string `json:"ssh_url"`
	CloneURL string `json:"clone_url"`
	Private  bool   `json:"private"`
}

// CreateRepo foydalanuvchi hisobida yangi repozitoriya yaratadi.
func (c *Client) CreateRepo(ctx context.Context, name string, private bool) (*Repository, error) {
	payload := map[string]any{"name": name, "private": private}

	var repo Repository
	if err := c.do(ctx, http.MethodPost, "/user/repos", payload, http.StatusCreated, &repo); err != nil {
		return nil, err
	}
	return &repo, nil
}

// Login token egasining GitHub username'ini qaytaradi (tokenni tekshirish uchun).
func (c *Client) Login(ctx context.Context) (string, error) {
	var user struct {
		Login string `json:"login"`
	}
	if err := c.do(ctx, http.MethodGet, "/user", nil, http.StatusOK, &user); err != nil {
		return "", err
	}
	return user.Login, nil
}

func (c *Client) do(ctx context.Context, method, path string, payload any, wantStatus int, out any) error {
	var body io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("so'rovni JSONga o'girib bo'lmadi: %w", err)
		}
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, baseURL+path, body)
	if err != nil {
		return fmt.Errorf("so'rov yaratib bo'lmadi: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("GitHub bilan bog'lanib bo'lmadi: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(io.LimitReader(resp.Body, maxBodyBytes))
	if err != nil {
		return fmt.Errorf("javobni o'qib bo'lmadi: %w", err)
	}
	if resp.StatusCode != wantStatus {
		return APIError(resp.StatusCode, data)
	}
	if out == nil {
		return nil
	}
	if err := json.Unmarshal(data, out); err != nil {
		return fmt.Errorf("javobni tahlil qilib bo'lmadi: %w", err)
	}
	return nil
}

// APIError GitHub xato javobidan o'qiladigan xabar yasaydi.
func APIError(status int, body []byte) error {
	var payload struct {
		Message string `json:"message"`
		Errors  []struct {
			Message string `json:"message"`
			Field   string `json:"field"`
			Code    string `json:"code"`
		} `json:"errors"`
	}
	_ = json.Unmarshal(body, &payload)

	message := payload.Message
	if len(payload.Errors) > 0 {
		detail := payload.Errors[0].Message
		if detail == "" {
			detail = payload.Errors[0].Field + " " + payload.Errors[0].Code
		}
		message += " (" + detail + ")"
	}
	if message == "" {
		message = http.StatusText(status)
	}
	return fmt.Errorf("GitHub API %d: %s", status, message)
}
