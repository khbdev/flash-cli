package configbranchedit
import (
	"flash/flashcore"
	"fmt"
	"os"
)



func Configbranchedit() bool {
		// 🔀 Variant: flash -b dev  → default branch o‘zgartirish
	if len(os.Args) == 3 && os.Args[1] == "-b" {
		newBranch := os.Args[2]
		err := flashcore.UpdateDefaultBranch(newBranch)
		if err != nil {
			fmt.Println("❌ Branch yangilab bo‘lmadi:", err)
			return true
		}
		fmt.Println("✅ Default branch yangilandi:", newBranch)
		return true
	}
return 	false
}