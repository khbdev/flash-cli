
# ⚡ Flash CLI

Git komandalarini va developerlar ishini soddalashtiruvchi CLI vosita.

---

## 🔧 O‘rnatish

### Linux / macOS:

#### Versiya 1.0.0:
```bash
wget https://github.com/khbdev/flash-cli/releases/download/v1.0.0/flash
chmod +x flash
sudo mv flash /usr/local/bin/
````

#### Versiya 2.0.0:

```bash
wget https://github.com/khbdev/flash-cli/releases/download/v2.00/flash
chmod +x flash
sudo mv flash /usr/local/bin/

```

### Windows (PowerShell):

#### Versiya 1.0.0:

```powershell
Invoke-WebRequest -Uri "https://github.com/khbdev/flash-cli/releases/download/v1.0.0/flash.exe" -OutFile "$env:USERPROFILE\flash.exe"
Move-Item "$env:USERPROFILE\flash.exe" "C:\Windows\System32\flash.exe"
```

#### Versiya 2.0.0:

```powershell
Invoke-WebRequest -Uri "https://github.com/khbdev/flash-cli/releases/download/v2.0.0/flash.exe" -OutFile "$env:USERPROFILE\flash.exe"
Move-Item "$env:USERPROFILE\flash.exe" "C:\Windows\System32\flash.exe"
```
⚠️ Eslatma
🚫 Flash CLI hozircha Windows operatsion tizimida ishlamaydi.
Linux va macOS foydalanuvchilari uchun to‘liq qo‘llab-quvvatlanad

✅ Endi `flash` komandasi terminalda ishlaydi.

---

Albatta! Quyidagicha **toza va bir xil uslubda** tuzib berdim. Formatlash, imlo, tushunarlilik jihatidan yaxshilandi:

---

## 📁 Loyihaviy komandalar

```bash
flash init                 # Loyihani boshlash (.flash papka va config.json)
flash start "msg"          # Git add, commit, push bajarish
flash -b main              # Default branchni o‘zgartirish (config.json)
```

---

## 📦 Box komandalar

```bash
flash box -c laravel       # Yangi box yaratish (laravel.box)
flash box laravel          # Box ichidagi komandalarni ishga tushurish
flash box edit laravel     # Box faylni tahrirlash (nano bilan)
flash box status           # Mavjud box'lar ro‘yxatini ko‘rsatish
flash box remove laravel   # Box faylni o‘chirish
```

📁 Box fayllar joylashuvi: `~/.flash/boxes/`

---

## 🔑 Token va repo komandalar

```bash
flash token -c             # GitHub Personal Access Token saqlash
flash token                # Tokenni ko‘rsatish
flash token -d             # Tokenni o‘chirish

flash repo                 # GitHub'da yangi repository yaratish (public)
flash repo -p              # GitHub'da yangi repository yaratish (private)
```

📁 Token saqlanish joyi: `~/.flash/config.json`

---

## 👤 Muallif

Made by [Azizbek Xasanov](https://github.com/khbdev)


