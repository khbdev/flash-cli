package cli

import "fmt"

// PrintHelp barcha komandalar ro'yxatini chiqaradi.
func PrintHelp() {
	fmt.Print(`⚡ Flash CLI — git ishlarini tezlashtiruvchi vosita

Loyiha:
  flash init [branch]        .flash/config.json yaratish (default: joriy branch)
  flash start "xabar"        git add + commit + push
  flash -b <branch>          default branchni o'zgartirish
  flash -b                   joriy default branchni ko'rsatish

Box (qayta ishlatiladigan bash skriptlar):
  flash box -s               mavjud box'lar ro'yxati
  flash box -c <nom>         yangi box yaratish va muharrirda ochish
  flash box -e <nom>         box'ni tahrirlash
  flash box -r <nom>         box'ni o'chirish
  flash box <nom> [args...]  box ichidagi komandalarni bajarish

GitHub:
  flash token                saqlangan tokenni tekshirish
  flash token -c             tokenni saqlash (ko'rinmas kiritish)
  flash token -d             tokenni o'chirish
  flash repo [nom] [-p]      yangi repozitoriya yaratish (-p = private)

Boshqa:
  flash help                 shu yordam
  flash version              versiya

Muhit o'zgaruvchilari:
  FLASH_GITHUB_TOKEN         tokenni fayl o'rniga muhitdan olish
  EDITOR / VISUAL            box'ni tahrirlash uchun muharrir
  NO_COLOR                   ranglarni o'chirish

Fayllar:
  ./.flash/config.json       loyiha sozlamalari
  ~/.flash/config.json       token (0600)
  ~/.flash/boxes/            box fayllari
`)
}
