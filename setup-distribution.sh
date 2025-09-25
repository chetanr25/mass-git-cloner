#!/bin/bash

echo "🚀 Mass Git Cloner - Distribution Setup"
echo "======================================"
echo ""

# Check if we're in a git repository
if [ ! -d ".git" ]; then
    echo "📁 Initializing Git repository..."
    git init
    echo "✅ Git repository initialized"
    echo ""
fi

# Check if go.mod needs updating
echo "📦 Checking Go module configuration..."
if grep -q "github.com/chetanr25/mass-git-cloner" go.mod; then
    echo "✅ Go module already configured correctly"
else
    echo "⚠️  Go module needs to be updated for distribution"
    echo "   Current module path might not match your GitHub repository"
fi
echo ""

# Build for multiple platforms
echo "🔨 Building for multiple platforms..."
mkdir -p releases

echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o releases/gclone-linux-amd64 ./cmd/git-clone

echo "Building for macOS (amd64)..."
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o releases/gclone-darwin-amd64 ./cmd/git-clone

echo "Building for macOS (arm64)..."
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o releases/gclone-darwin-arm64 ./cmd/git-clone

echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o releases/gclone-windows-amd64.exe ./cmd/git-clone

echo "✅ Built binaries for multiple platforms in ./releases/"
echo ""

# Show file sizes
echo "📊 Binary sizes:"
ls -lh releases/
echo ""

# Create distribution README
echo "📝 Creating distribution instructions..."
cat > INSTALL.md << 'EOF'
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
EOF

echo "✅ Created INSTALL.md with installation instructions"
echo ""

echo "🎯 Next Steps for Distribution:"
echo "1. Push your code to GitHub:"
echo "   git add ."
echo "   git commit -m 'Ready for distribution'"
echo "   git remote add origin https://github.com/YOUR_USERNAME/mass-git-cloner.git"
echo "   git push -u origin main"
echo ""
echo "2. Create a release tag:"
echo "   git tag v1.0.0"
echo "   git push origin v1.0.0"
echo ""
echo "3. Upload the binaries from ./releases/ to GitHub releases"
echo ""
echo "4. Users can then install with:"
echo "   go install github.com/YOUR_USERNAME/mass-git-cloner/cmd/git-clone@latest"
echo ""
echo "🎉 Your Go application will be distributable worldwide!"