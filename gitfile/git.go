package gitfile

import (
	"encoding/json"

	"flash/config"
	"flash/helpersFunction"
	"fmt"
	"os"
)



func Start(msg string) {

    
    configData, err := os.ReadFile(".flash/config.json")
    if err != nil {
        fmt.Println(" config.json topilmadi:", err)
        return
    }

    var config config.Config
    if err := json.Unmarshal(configData, &config); err != nil {
        fmt.Println(" config.json noto‘g‘ri formatda:", err)
        return
    }

 
    helpersFunction.RunCommand("git add .")
    helpersFunction.RunCommand(fmt.Sprintf(`git commit -m "%s"`, msg))
    helpersFunction.RunCommand(fmt.Sprintf("git push origin %s", config.DefaultBranch))
}
