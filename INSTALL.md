# 📦 Installation Guide

## Method 1: Go Install (Recommended)
If you have Go installed:
```bash
go install github.com/chetanr25/mass-git-cloner/cmd/git-clone@latest
```

## Method 2: Download Binary
Download the appropriate binary for your system from the releases page:

### macOS
```bash
# For Intel Macs
curl -L -o git-clone https://github.com/chetanr25/mass-git-cloner/releases/latest/download/git-clone-darwin-amd64
chmod +x git-clone
sudo mv git-clone /usr/local/bin/

# For Apple Silicon Macs
curl -L -o git-clone https://github.com/chetanr25/mass-git-cloner/releases/latest/download/git-clone-darwin-arm64
chmod +x git-clone
sudo mv git-clone /usr/local/bin/
```

### Linux
```bash
curl -L -o git-clone https://github.com/chetanr25/mass-git-cloner/releases/latest/download/git-clone-linux-amd64
chmod +x git-clone
sudo mv git-clone /usr/local/bin/
```

### Windows
Download `git-clone-windows-amd64.exe` from the releases page and add to your PATH.

## Usage
After installation, simply run:
```bash
git-clone
```

## Requirements
- Git must be installed on your system
- Internet connection
- Optional: GitHub token for higher API limits
