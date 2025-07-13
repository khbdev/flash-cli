
# ⚡ Flash CLI

Git komandalarini va  developerlar uchun paketlarni avtomatlashtiruvchi CLI vosita.

---

## 🔧 O‘rnatish (Linux / macOS)

```bash
git clone https://github.com/khbdev/flash-cli.git
cd flash-cli
go build -o flash
sudo mv flash /usr/local/bin/
````

✅ Endi `flash` terminalda istalgan joyda ishlaydi.

---

## 🚀 Tezkor foydalanish

```bash
cd my-project/
flash init                    # .flash/ yaratadi
flash start "msg"            # git add . + commit + push
flash -b main                # default branch o‘zgartirish
```

---

## 📦 Box — doimiy komandalar uchun qutilar

Laravel, Go, React, Django kabi texnologiyalar uchun **o‘z komandalaringni saqlab** ishlat!

```bash
flash box -c laravel         # laravel.box fayl yaratadi
flash box laravel            # ichidagi komandalarni bajaradi
flash box edit laravel       # faylni tahrirlash
flash box status             # mavjud box’lar ro‘yxati
flash box remove laravel     # box’ni o‘chirish
```

Box fayllar manzili: `~/.flash/boxes/`

---

## 🔍 Komandalar jadvali

| Komanda                  | Tavsifi                                   |
| ------------------------ | ----------------------------------------- |
| `flash init`             | Loyiha ichida `.flash/` yaratadi          |
| `flash start "xabar"`    | Git add, commit va push qiladi            |
| `flash -b branch`        | Default branch’ni o‘zgartiradi            |
| `flash box -c <nom>`     | Yangi box yaratadi                        |
| `flash box <nom>`        | Box ichidagi komandalarni ishga tushuradi |
| `flash box status`       | Barcha box’lar ro‘yxatini ko‘rsatadi      |
| `flash box edit <nom>`   | Box faylni tahrirlash                     |
| `flash box remove <nom>` | Box faylni o‘chirish                      |
| `flash help`             | Yordam oynasini ko‘rsatadi                |

---

## 👤 Muallif

Made with ❤️ by [Azizbek Xasanov](https://github.com/khbdev)

```


