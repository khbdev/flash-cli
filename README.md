# ⚡ Flash CLI

Git ish oqimini tezlashtiruvchi CLI: bitta komanda bilan `add + commit + push`,
qayta ishlatiladigan bash skriptlar ("box") va GitHub repozitoriyasini terminaldan
yaratish. Tashqi kutubxonalarsiz, faqat Go standart kutubxonasida yozilgan.

## O'rnatish

```bash
go install github.com/khbdev/flash-cli/cmd/flash@latest
```

Yoki manbadan:

```bash
git clone https://github.com/khbdev/flash-cli && cd flash-cli
make install          # /usr/local/bin/flash
```

Tayyor binar:

```bash
wget https://github.com/khbdev/flash-cli/releases/latest/download/flash-linux-amd64 -O flash
chmod +x flash && sudo mv flash /usr/local/bin/
```

## Loyiha komandalari

```bash
flash init            # .flash/config.json yaratadi (default: joriy branch)
flash init develop    # default branchni ko'rsatib
flash start "xabar"   # git add -A + commit + push -u origin <default-branch>
flash -b main         # default branchni o'zgartirish
flash -b              # joriy default branchni ko'rsatish
```

`flash start` commit xabarini git'ga argument sifatida uzatadi, shuning uchun
xabardagi tirnoq yoki `$(...)` shell tomonidan bajarilmaydi. Har bir qadam
muvaffaqiyatsiz bo'lsa, keyingisi bajarilmaydi va exit kodi `1` bo'ladi.

## Box — qayta ishlatiladigan skriptlar

```bash
flash box -s                   # ro'yxat
flash box -c laravel           # yaratish + muharrirda ochish
flash box -e laravel           # tahrirlash
flash box -r laravel           # o'chirish
flash box laravel myapp        # ishga tushirish ($1 = myapp)
```

Box fayllari `~/.flash/boxes/*.box` da saqlanadi va oddiy bash skript bo'ladi:

```bash
#!/usr/bin/env bash
set -euo pipefail

composer create-project laravel/laravel "$1"
cd "$1" && php artisan key:generate
```

Muharrir `$VISUAL` → `$EDITOR` → `nano`/`vim`/`vi` tartibida tanlanadi.

## GitHub

```bash
flash token -c        # tokenni ko'rinmas kiritish orqali saqlash
flash token           # yashirilgan token + qaysi akkaunt ekanini tekshirish
flash token -d        # o'chirish

flash repo            # joriy papka nomi bilan public repo
flash repo my-app     # nom ko'rsatib
flash repo my-app -p  # private
```

Repo yaratilgandan so'ng, joriy papka git repozitoriyasi bo'lsa va `origin`
mavjud bo'lmasa, remote avtomatik ulanadi.

## Sozlamalar

| Joy | Nima saqlanadi |
|---|---|
| `./.flash/config.json` | loyiha default branchi (git'ga commit qilinadi) |
| `~/.flash/config.json` | GitHub token, `0600` huquq bilan |
| `~/.flash/boxes/` | box fayllari |

Muhit o'zgaruvchilari:

| O'zgaruvchi | Vazifasi |
|---|---|
| `FLASH_GITHUB_TOKEN` | tokenni fayl o'rniga muhitdan olish (CI uchun) |
| `EDITOR` / `VISUAL` | box tahrirlash uchun muharrir |
| `NO_COLOR` | ranglarni o'chirish |

## Arxitektura

```
cmd/flash/            main() — faqat xatoni chop etadi va exit kodini beradi
internal/cli/         komandalarni tahlil qilish, chiqarish (yagona I/O qatlami)
internal/config/      .flash/config.json va ~/.flash/config.json
internal/git/         git ustidan yupqa qatlam (shell'siz, argumentlar bilan)
internal/github/      GitHub REST API klienti
internal/box/         box fayllarini saqlash va bajarish
internal/ui/          rangli chiqish, maxfiy kiritish
```

Qoida: domen paketlari (`config`, `git`, `github`, `box`) hech narsa chop
etmaydi — ular `error` qaytaradi. Faqat `cli` va `ui` terminalga yozadi. Shu
sababli domen mantiqini test qilish oson va exit kodlari to'g'ri ishlaydi.

## Ishlab chiqish

```bash
make test     # go test -race ./...
make lint     # go vet + gofmt tekshiruvi
make build    # bin/flash
make release  # linux/macOS, amd64/arm64
```

## Muallif

[Azizbek Xasanov](https://github.com/khbdev)
