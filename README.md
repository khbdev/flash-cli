
# ⚡ Flash CLI

Git komandalarini va developerlar ishini soddalashtiruvchi CLI vosita.

---

## 🔧 O‘rnatish

### Linux / macOS:

```bash
wget https://github.com/khbdev/flash-cli/releases/download/v1.0.0/flash
chmod +x flash
sudo mv flash /usr/local/bin/
````

### Windows (PowerShell):

```powershell
Invoke-WebRequest -Uri "https://github.com/khbdev/flash-cli/releases/download/v1.0.0/flash.exe" -OutFile "$env:USERPROFILE\flash.exe"
Move-Item "$env:USERPROFILE\flash.exe" "C:\Windows\System32\flash.exe"
```

✅ Endi `flash` komandasi terminalda ishlaydi.

---

## 🚀 Tez foydalanish

```bash
flash init               # .flash/ yaratadi
flash start "msg"       # Git add, commit va push
flash -b main           # Default branch o‘zgartiradi
```

---

## 📦 Boxlar

```bash
flash box -c laravel       # laravel.box yaratadi
flash box laravel          # komandalarni bajaradi
flash box edit laravel     # tahrirlash
flash box status           # ro‘yxatini ko‘rsatadi
flash box remove laravel   # o‘chirish
```

Box fayllar: `~/.flash/boxes/`

---

## 👤 Muallif

Made by [Azizbek Xasanov](https://github.com/khbdev)

