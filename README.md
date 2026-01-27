

# ⚡ Flash CLI

A CLI tool that simplifies Git commands and automates developers' workflows.

---

## 🔧 Installation

### Linux / macOS:


#### Version 2.0.0:

```bash
wget https://github.com/khbdev/flash-cli/releases/download/v2.00/flash
chmod +x flash
sudo mv flash /usr/local/bin/
```

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
flash box -e laravel     # Edit the box file (with nano)
flash box -s           # Show list of existing boxes
flash box -r laravel   # Delete a box file
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

