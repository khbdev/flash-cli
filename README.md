# ⚡ Flash CLI

Git komandalarini avtomatlashtiruvchi oddiy CLI vosita.

---

## 🛠 O‘rnatish

### 🔹 Linux / macOS

```bash
git clone https://github.com/azizbekxasanov/flashcli.git
cd flashcli
go build -o flash
sudo mv flash /usr/local/bin/
```

> Endi `flash` terminalda istalgan joyda ishlaydi.

---

### 🔹 Windows

```powershell
git clone https://github.com/azizbekxasanov/flashcli.git
cd flashcli
go build -o flash.exe
```

✅ `flash.exe` faylini `C:\Windows\System32` yoki PATH listga qo‘shilgan boshqa folderga nusxa ko‘chiring.

---

## 💻 Ishlatish

### 🔹 1. Loyiha papkasiga kir:

```bash
cd my-project/
```

### 🔹 2. Flash’ni boshlash (bir martalik):

```bash
flash init
```

Bu `.flash/` papkasini yaratadi.

### 🔹 3. Git komandalarini avtomatlashtirish:

```bash
flash start "commit xabari"
```

Bu quyidagilarni bajaradi:

```bash
git add .
git commit -m "commit xabari"
git push
```

---

## 📋 Komandalar

| Komanda               | Tavsifi                                 |
|-----------------------|------------------------------------------|
| `flash init`          | Loyiha ichida `.flash/` papka yaratadi  |
| `flash start "xabar"` | Git add, commit va push qiladi           |
| `flash -b new-branch` | branch ozgartiradi            |
| `flash help` | nimadir chunmsangiz help commandasi            |

---

