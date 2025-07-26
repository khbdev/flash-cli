package helpfunction

import "fmt"

func PrintHelp() {
	fmt.Println("🛠️ Flash CLI Yordam:")
	fmt.Println()
	
	fmt.Println("📁 Loyihaviy komandalar:")
	fmt.Println(" flash init               → Loyihani boshlash (.flash papka va config)")
	fmt.Println(" flash start \"msg\"        → Git add, commit, push bajarish")
	fmt.Println(" flash -b 'branch'        → Default branchni o‘zgartirish (config.json)")
	fmt.Println()

	fmt.Println("📦 Box komandalar:")
	fmt.Println(" flash box -c <name>      → Yangi box yaratish (masalan: flash box -c laravel)")
	fmt.Println(" flash box <name>         → Box ichidagi komandalarni ishga tushurish")
	fmt.Println(" flash box status         → Mavjud box'lar ro‘yxatini ko‘rsatish")
	fmt.Println(" flash box edit <name>    → Box faylni tahrirlash (nano editor bilan)")
	fmt.Println(" flash box remove <name>  → Box faylni o‘chirish")
	fmt.Println()

	fmt.Println("🔑 Token va repo komandalar:")
	fmt.Println(" flash token -c     → GitHub Personal Access Token saqlash")
	fmt.Println(" flash token         → Tokenni ko‘rsatish")
	fmt.Println(" flash  token -d      → Tokenni o‘chirish")
	fmt.Println()
	fmt.Println("  flash repo    → GitHub'da yangi repository yaratish")
	fmt.Println("  flash repo  -p  → GitHub'da yangi repository privative qilib yarartish")
	fmt.Println()
}
