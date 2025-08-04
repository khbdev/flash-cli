

# ⚡ Flash CLI

A CLI tool that simplifies Git commands and automates developers' workflows.

---

## 🔧 Installation

### Linux / macOS:

#### Version 1.0.0:

```bash
wget https://github.com/khbdev/flash-cli/releases/download/v1.0.0/flash
chmod +x flash
sudo mv flash /usr/local/bin/
```

#### Version 2.0.0:

```bash
wget https://github.com/khbdev/flash-cli/releases/download/v2.00/flash
chmod +x flash
sudo mv flash /usr/local/bin/
```

---

## 🪟 Install on Windows via PowerShell (v2.0.0)

1. **Open PowerShell as Administrator**

2. Run the following commands step-by-step:

```powershell
Invoke-WebRequest -Uri "https://github.com/khbdev/flash-cli/releases/download/v2.00/flash.exe" -OutFile "$env:USERPROFILE\flash.exe"
New-Item -ItemType Directory -Path "$env:USERPROFILE\bin" -Force
Move-Item "$env:USERPROFILE\flash.exe" "$env:USERPROFILE\bin\flash.exe" -Force
$oldPath = [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::User)
$newPath = $oldPath + ";$env:USERPROFILE\bin"
[Environment]::SetEnvironmentVariable("Path", $newPath, [EnvironmentVariableTarget]::User)
```

3. Close and reopen PowerShell. Now type `flash` to check if it's working.

---

If you face any issues with PATH, add it manually from Windows Environment Variables.

✅ Now you can run the `flash` command from anywhere in the terminal.

---

## 📁 Project Commands

```bash
flash init                 # Start a project (.flash folder and config.json)
flash start "msg"          # Run git add, commit, and push
flash -b main              # Set default branch (saved in config.json)
```

---

## 📦 Box Commands

```bash
flash box -c laravel       # Create a new box (laravel.box)
flash box laravel          # Execute commands inside the box
flash box edit laravel     # Edit the box file (with nano)
flash box status           # Show list of existing boxes
flash box remove laravel   # Delete a box file
```

📁 Box files location: `~/.flash/boxes/`

---

## 🔑 Token & Repo Commands

```bash
flash token -c             # Save GitHub Personal Access Token
flash token                # Show the current token
flash token -d             # Delete the token

flash repo                 # Create a new public GitHub repository
flash repo -p              # Create a new private GitHub repository
```

📁 Token is stored at: `~/.flash/config.json`

---

## 👤 Author

Made by [Azizbek Xasanov](https://github.com/khbdev)

