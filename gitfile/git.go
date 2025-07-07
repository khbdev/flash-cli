package gitfile

import (
	"encoding/json"
	"flash/helpersFunction"
	"fmt"
	"os"
)

type Config struct {
    DefaultBranch string `json:"default_branch"`
}


func Start(msg string) {

    // 📦 1. config.json dan default branchni o‘qish////
    configData, err := os.ReadFile(".flash/config.json")
    if err != nil {
        fmt.Println("❌ config.json topilmadi:", err)
        return
    }

    var config Config
    if err := json.Unmarshal(configData, &config); err != nil {
        fmt.Println("❌ config.json noto‘g‘ri formatda:", err)
        return
    }

    // ✅ 2. Git komandalarni bajarish
    helpersFunction.RunCommand("git add .")
    helpersFunction.RunCommand(fmt.Sprintf(`git commit -m "%s"`, msg))
    helpersFunction.RunCommand(fmt.Sprintf("git push origin %s", config.DefaultBranch))
}
