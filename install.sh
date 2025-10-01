#!/bin/bash

# Mass Git Cloner - Global Installer Script
# This provides an NPM-like installation experience for the Go binary

set -e

REPO="chetanr25/mass-git-cloner"
BINARY_NAME="gclone"
INSTALL_DIR="/usr/local/bin"

echo "🚀 Installing Mass Git Cloner..."
echo "================================"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) echo "❌ Unsupported architecture: $ARCH" && exit 1 ;;
esac

case $OS in
    linux) PLATFORM="linux-$ARCH" ;;
    darwin) PLATFORM="darwin-$ARCH" ;;
    *) echo "❌ Unsupported OS: $OS" && exit 1 ;;
esac

echo "📍 Detected platform: $PLATFORM"

# Check if Go is installed for fallback
GO_INSTALLED=false
if command -v go &> /dev/null; then
    GO_INSTALLED=true
    echo "✅ Go found: $(go version | awk '{print $3}' | sed 's/go//')"
else
    echo "⚠️  Go not found - will try to download pre-built binary"
fi

# Function to install via Go (most reliable method)
install_via_go() {
    echo "🔨 Installing via Go..."
    
    # Check if we're in the project directory with local source
    if [ -f "cmd/git-clone/main.go" ] && [ -f "go.mod" ]; then
        echo "📍 Found local source code, installing from local directory..."
        go install ./cmd/git-clone
    else
        echo "📍 Installing from GitHub repository..."
        go install "github.com/$REPO/cmd/git-clone@latest"
    fi
    
    # Get GOPATH/GOBIN
    GOBIN=$(go env GOBIN)
    if [ -z "$GOBIN" ]; then
        GOPATH=$(go env GOPATH)
        if [ -z "$GOPATH" ]; then
            GOPATH="$HOME/go"
        fi
        GOBIN="$GOPATH/bin"
    fi
    
    echo "📍 Go installed binary to: $GOBIN/git-clone"
    
    # Create symlink with shorter name
    if [ -f "$GOBIN/git-clone" ]; then
        echo "📍 Creating symlink: $INSTALL_DIR/$BINARY_NAME -> $GOBIN/git-clone"
        if [ -w "$INSTALL_DIR" ]; then
            ln -sf "$GOBIN/git-clone" "$INSTALL_DIR/$BINARY_NAME"
        else
            sudo ln -sf "$GOBIN/git-clone" "$INSTALL_DIR/$BINARY_NAME"
        fi
        return 0
    else
        return 1
    fi
}

# Function to install from pre-built binary
install_prebuilt() {
    echo "📦 Downloading pre-built binary..."
    
    # Try to download from releases directory in repo
    DOWNLOAD_URL="https://github.com/$REPO/raw/main/releases/gclone-$PLATFORM"
    if [ "$OS" = "windows" ]; then
        DOWNLOAD_URL="${DOWNLOAD_URL}.exe"
    fi
    
    echo "⬇️  Downloading from: $DOWNLOAD_URL"
    
    # Create temporary directory
    TMP_DIR=$(mktemp -d)
    cd "$TMP_DIR"
    
    # Download binary
    if curl -sSL -o "$BINARY_NAME" "$DOWNLOAD_URL"; then
        chmod +x "$BINARY_NAME"
        
        # Install binary
        echo "📦 Installing to $INSTALL_DIR..."
        if [ -w "$INSTALL_DIR" ]; then
            mv "$BINARY_NAME" "$INSTALL_DIR/"
        else
            sudo mv "$BINARY_NAME" "$INSTALL_DIR/"
        fi
        
        # Cleanup
        cd - > /dev/null
        rm -rf "$TMP_DIR"
        return 0
    else
        cd - > /dev/null
        rm -rf "$TMP_DIR"
        return 1
    fi
}

# Try installation methods
echo "🎯 Attempting installation..."

if [ "$GO_INSTALLED" = true ]; then
    # Method 1: Try Go installation first (most reliable)
    if install_via_go; then
        echo "✅ Successfully installed via Go!"
    elif install_prebuilt; then
        echo "✅ Successfully installed via pre-built binary!"
    else
        echo "❌ Both installation methods failed"
        exit 1
    fi
else
    # Method 2: Try pre-built binary if Go is not available
    if install_prebuilt; then
        echo "✅ Successfully installed via pre-built binary!"
    else
        echo "❌ Installation failed!"
        echo "💡 Please install Go and try again:"
        echo "   go install github.com/$REPO/cmd/git-clone@latest"
        exit 1
    fi
fi

echo ""
echo "🎉 Installation complete!"
echo ""
echo "� You can now run:"
echo "   $BINARY_NAME"
echo "   # or"
echo "   git-clone"
echo ""
echo "🔍 To verify installation:"
echo "   which $BINARY_NAME"
echo ""
echo "📚 For usage instructions, visit:"
echo "   https://github.com/$REPO"
echo ""
echo "🚀 To uninstall later:"
echo "   sudo rm $INSTALL_DIR/$BINARY_NAME"