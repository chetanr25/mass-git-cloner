# Go Module Distribution Setup

## 📦 Publishing to Go Modules

### 1. Push to GitHub
```bash
# Initialize git repository
git init
git add .
git commit -m "Initial commit: Mass Git Cloner with Bubbletea UI"

# Add GitHub remote (replace with your repo)
git remote add origin https://github.com/chetanr25/mass-git-cloner.git
git branch -M main
git push -u origin main
```

### 2. Tag a Release
```bash
# Create and push a version tag
git tag v1.0.0
git push origin v1.0.0
```

### 3. Users Install With
```bash
# Install directly from GitHub
go install github.com/chetanr25/mass-git-cloner/cmd/git-clone@latest

# Or install specific version
go install github.com/chetanr25/mass-git-cloner/cmd/git-clone@v1.0.0
```

### 4. Update go.mod for Distribution
```go
module github.com/chetanr25/mass-git-cloner

go 1.21

require (
    github.com/charmbracelet/bubbletea v1.3.10
    github.com/charmbracelet/lipgloss v1.1.0
)
```

## 🔧 Alternative Distribution Methods

### Option 1: GitHub Releases with Binaries
Create pre-built binaries for different platforms:

```bash
# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o releases/git-clone-linux-amd64 ./cmd/git-clone
GOOS=darwin GOARCH=amd64 go build -o releases/git-clone-darwin-amd64 ./cmd/git-clone
GOOS=darwin GOARCH=arm64 go build -o releases/git-clone-darwin-arm64 ./cmd/git-clone
GOOS=windows GOARCH=amd64 go build -o releases/git-clone-windows-amd64.exe ./cmd/git-clone
```

### Option 2: Homebrew (macOS)
Create a Homebrew formula for easy installation on macOS.

### Option 3: Package Managers
- **Arch Linux**: AUR package
- **Ubuntu/Debian**: .deb package
- **Red Hat/CentOS**: .rpm package

## 🎯 Recommended Approach

**Use Go Modules** - this is the standard way to distribute Go applications:

1. Push code to GitHub
2. Tag releases (v1.0.0, v1.1.0, etc.)
3. Users install with: `go install github.com/chetanr25/mass-git-cloner/cmd/git-clone@latest`

Would you like me to help you set this up?