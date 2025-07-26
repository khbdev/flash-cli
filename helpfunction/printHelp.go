package helpfunction

import "fmt"

func PrintHelp() {
	fmt.Println("🛠️ Flash CLI Yordam:")
	fmt.Println()
	
	fmt.Println("📁 Loyihaviy komandalar:")
	fmt.Println("  init               → Loyihani boshlash (.flash papka va config)")
	fmt.Println("  start \"msg\"        → Git add, commit, push bajarish")
	fmt.Println("  -b 'branch'        → Default branchni o‘zgartirish (config.json)")
	fmt.Println()

	fmt.Println("📦 Box komandalar:")
	fmt.Println("  box -c <name>      → Yangi box yaratish (masalan: flash box -c laravel)")
	fmt.Println("  box <name>         → Box ichidagi komandalarni ishga tushurish")
	fmt.Println("  box status         → Mavjud box'lar ro‘yxatini ko‘rsatish")
	fmt.Println("  box edit <name>    → Box faylni tahrirlash (nano editor bilan)")
	fmt.Println("  box remove <name>  → Box faylni o‘chirish")
	fmt.Println()

	fmt.Println("🔑 Token va repo komandalar:")
	fmt.Println("  token set <t>      → GitHub Personal Access Token saqlash")
	fmt.Println("  token show         → Tokenni ko‘rsatish")
	fmt.Println("  token clear        → Tokenni o‘chirish")
	fmt.Println()
	fmt.Println("  repo create <n>    → GitHub'da yangi repository yaratish")
	fmt.Println("  repo clone <url>   → Mavjud repository'ni clone qilish")
	fmt.Println("  repo delete <n>    → GitHub'dagi repository'ni o‘chirish")
	fmt.Println()
}
