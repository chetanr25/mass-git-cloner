# Mass Git Cloner (gclone)

A powerful command-line tool that allows you to clone multiple Git repositories simultaneously with ease. Perfect for developers who work with multiple repositories, need to set up development environments quickly, or want to batch clone repositories from organizations or users.

## 🎯 What Problem Does It Solve?

**Before gclone:**
- Manual cloning of multiple repositories one by one
- Repetitive `git clone` commands
- Difficulty managing multiple repository downloads
- No easy way to clone all repositories from a GitHub user/organization

**After gclone:**
- Clone multiple repositories with a single command
- Batch download from GitHub users, organizations, or custom lists 
- Simple CLI interface 

<!--
## Table of Contents

- [Installation](#-installation)
- [Usage](#-usage)
- [Contributing](#-contributing)
- [License](#-license)

-=>

## Installation

### Quick Install (Recommended)

<!-- **macOS and Linux:** -->
```bash
curl -sSL https://raw.githubusercontent.com/chetanr25/mass-git-cloner/main/install.sh | bash
```

<!-- **Windows (PowerShell):**
```powershell
# Coming soon - for now use Go install method below
``` -->

### Alternative Installation Methods

#### 📦 Using Go (if you have Go installed)
```bash
go install github.com/chetanr25/mass-git-cloner/cmd/git-clone@latest
```


#### 📥 Manual Download
1. Go to [Releases](https://github.com/chetanr25/mass-git-cloner/releases)
2. Download the binary for your platform
3. Make it executable and move to your PATH:

<!-- **macOS/Linux:**
```bash
chmod +x gclone-*
sudo mv gclone-* /usr/local/bin/gclone
``` -->
<!-- 
**Windows:**
```cmd
# Move the .exe file to a directory in your PATH
``` -->

<!-- 
#### 🍺 Homebrew (macOS/Linux)
```bash
# Coming soon
``` -->

### Platform Support

| Platform | Architecture | Status |
|----------|--------------|--------|
| 🍎 macOS | Intel (x86_64) | ✅ Supported |
| 🍎 macOS | Apple Silicon (ARM64) | ✅ Supported |
| 🐧 Linux | x86_64 | ✅ Supported |
| 🐧 Linux | ARM64 | ✅ Supported |
| 🪟 Windows | x86_64 | ✅ Supported |
| 🪟 Windows | x86 (32-bit) | ✅ Supported |

### Project Structure

```
mass-git-cloner/
├── cmd/git-clone/          # Main application entry point
├── internal/               # Internal packages
├── releases/               # Pre-built binaries
```


## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
