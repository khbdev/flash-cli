
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

Albatta! Mana qisqaroq ko‘rsatma:

---

## Windows PowerShell orqali `flash.exe` ni yuklab olish va o‘rnatish (v2.0.0)

1. **PowerShell-ni administrator sifatida oching**

2. Quyidagi buyruqlarni ketma-ket bajaring:

```powershell
Invoke-WebRequest -Uri "https://github.com/khbdev/flash-cli/releases/download/v2.00/flash.exe" -OutFile "$env:USERPROFILE\flash.exe"
New-Item -ItemType Directory -Path "$env:USERPROFILE\bin" -Force
Move-Item "$env:USERPROFILE\flash.exe" "$env:USERPROFILE\bin\flash.exe" -Force
$oldPath = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::User)
$newPath = $oldPath + ";$env:USERPROFILE\bin"
[Environment]::SetEnvironmentVariable("Path", $newPath, [EnvironmentVariableTarget]::User)
```

3. PowerShell-ni yoping va qaytadan oching, keyin `flash` deb yozing, ishlashi kerak.

---

Agar PATH’ga qo‘shishda muammo bo‘lsa, uni Windows sozlamalaridan qo‘lda qo‘shing.

---

]
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


