package repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"flash/token"
	"io"
	"net/http"
	"os"
	"strings"
)

func Repo() {
	if len(os.Args) < 3 {
		fmt.Println("❌ Iltimos, repo nomini bering: `flash repo my-repo-name [-pri|-pub]`")
		return
	}

	repoName := os.Args[2]

	// Default: public
	isPrivate := false
	if len(os.Args) > 3 {
		arg := strings.ToLower(os.Args[3])
		if arg == "-p" {
			isPrivate = true
		}  else {
			fmt.Println("⚠️ Belgilanmagan parametr:", arg)
			fmt.Println("👉 Iltimos, `-pri` (private) yoki `-pub` (public) yozing.")
			return
		}
	}

	githubToken := token.GetToken()
	fmt.Println("⏳ GitHub repo yaratilmoqda...")

	data := map[string]interface{}{
		"name":    repoName,
		"private": isPrivate,
	}

	body, err := json.Marshal(data)
	if err != nil {
		fmt.Println("❌ JSON marshal xatosi:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://api.github.com/user/repos", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("❌ So‘rov yaratishda xato:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+githubToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ So‘rov yuborishda xato:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		bodyBytes, _ := io.ReadAll(resp.Body)

		var result map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &result); err != nil {
			fmt.Println("✅ Repo yaratildi:", repoName)
			fmt.Println("⚠️ Ammo linkni o‘qib bo‘lmadi.")
			return
		}

		htmlURL := result["html_url"].(string)

		fmt.Println("✅ Repozitoriya yaratildi:", repoName)
		fmt.Println("🔗 Link:", htmlURL)
		fmt.Println("\n💡 Git komandalar:")
		fmt.Printf("git remote add origin %s.git\n", htmlURL)
		fmt.Println("git branch -M main")
		fmt.Println("git push -u origin main")

	} else {
		fmt.Println("❌ Repozitoriya yaratilmadi. Status:", resp.Status)
	}
}
