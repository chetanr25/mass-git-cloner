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
    i686|i386) ARCH="386" ;;
    *) echo "❌ Unsupported architecture: $ARCH" && exit 1 ;;
esac

# Handle Windows environments (including Git Bash, MSYS2, Cygwin)
case $OS in
    linux) PLATFORM="linux-$ARCH" ;;
    darwin) PLATFORM="darwin-$ARCH" ;;
    mingw64*|msys*|cygwin*) 
        OS="windows"
        PLATFORM="windows-$ARCH" 
        BINARY_NAME="gclone.exe"
        # For Windows environments, try to install in a directory that's in PATH
        if [ -d "/usr/local/bin" ]; then
            INSTALL_DIR="/usr/local/bin"
        elif [ -d "/usr/bin" ]; then
            INSTALL_DIR="/usr/bin"
        else
            INSTALL_DIR="$HOME/bin"
            mkdir -p "$INSTALL_DIR"
            echo "📍 Created directory: $INSTALL_DIR"
            echo "⚠️  Make sure $INSTALL_DIR is in your PATH"
        fi
        ;;
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

# Check download tools availability for Windows
if [ "$OS" = "windows" ]; then
    if ! command -v curl >/dev/null 2>&1 && ! command -v wget >/dev/null 2>&1; then
        echo "❌ Neither curl nor wget found"
        echo "💡 Please install Git for Windows which includes curl, or install wget"
        exit 1
    fi
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
    
    # Handle Windows executable extension
    GO_BINARY_NAME="git-clone"
    if [ "$OS" = "windows" ]; then
        GO_BINARY_NAME="git-clone.exe"
    fi
    
    echo "📍 Go installed binary to: $GOBIN/$GO_BINARY_NAME"
    
    # Create symlink with shorter name
    if [ -f "$GOBIN/$GO_BINARY_NAME" ]; then
        echo "📍 Creating symlink: $INSTALL_DIR/$BINARY_NAME -> $GOBIN/$GO_BINARY_NAME"
        if [ -w "$INSTALL_DIR" ]; then
            ln -sf "$GOBIN/$GO_BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
        else
            sudo ln -sf "$GOBIN/$GO_BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
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
    
    # Create temporary directory - handle Windows path issues
    if [ "$OS" = "windows" ]; then
        TMP_DIR="/tmp/gclone-install-$$"
        mkdir -p "$TMP_DIR"
    else
        TMP_DIR=$(mktemp -d)
    fi
    cd "$TMP_DIR"
    
    # Download binary with Windows-specific curl options
    if [ "$OS" = "windows" ]; then
        # Use different curl options for Windows to avoid "Failed writing body" error
        if curl -L --fail --show-error --output "$BINARY_NAME" "$DOWNLOAD_URL" 2>/dev/null; then
            echo "✅ Download successful"
        elif command -v wget >/dev/null 2>&1; then
            echo "⚠️  curl failed, trying wget..."
            if wget -q -O "$BINARY_NAME" "$DOWNLOAD_URL"; then
                echo "✅ Download successful with wget"
            else
                echo "❌ wget also failed"
                return 1
            fi
        elif command -v powershell.exe >/dev/null 2>&1; then
            echo "⚠️  curl/wget failed, trying PowerShell..."
            if powershell.exe -Command "Invoke-WebRequest -Uri '$DOWNLOAD_URL' -OutFile '$BINARY_NAME'" 2>/dev/null; then
                echo "✅ Download successful with PowerShell"
            else
                echo "❌ All download methods failed"
                return 1
            fi
        else
            echo "❌ Download failed and no alternative tools available"
            return 1
        fi
    else
        # Use standard curl options for Linux/macOS
        if curl -sSL -o "$BINARY_NAME" "$DOWNLOAD_URL"; then
            echo "✅ Download successful"
        else
            return 1
        fi
    fi
    
    chmod +x "$BINARY_NAME"
    
    # Install binary
    echo "📦 Installing to $INSTALL_DIR..."
    if [ -w "$INSTALL_DIR" ]; then
        mv "$BINARY_NAME" "$INSTALL_DIR/"
    else
        if [ "$OS" = "windows" ]; then
            # On Windows, avoid sudo which might not work in Git Bash
            mv "$BINARY_NAME" "$INSTALL_DIR/" 2>/dev/null || {
                echo "❌ Cannot write to $INSTALL_DIR"
                echo "💡 Try running as administrator or choose a different directory"
                return 1
            }
        else
            sudo mv "$BINARY_NAME" "$INSTALL_DIR/"
        fi
    fi
    
    # Cleanup
    cd - > /dev/null
    rm -rf "$TMP_DIR"
    return 0
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
        if [ "$OS" = "windows" ]; then
            echo ""
            echo "🪟 Windows Troubleshooting:"
            echo "   • Ensure you're running in Git Bash, PowerShell, or WSL"
            echo "   • Try running as Administrator"
            echo "   • Check if antivirus is blocking the download"
            echo "   • Manual installation:"
            echo "     1. Download from: https://github.com/$REPO/releases"
            echo "     2. Extract to a folder in your PATH"
            echo ""
        fi
        exit 1
    fi
else
    # Method 2: Try pre-built binary if Go is not available
    if install_prebuilt; then
        echo "✅ Successfully installed via pre-built binary!"
    else
        echo "❌ Installation failed!"
        if [ "$OS" = "windows" ]; then
            echo ""
            echo "🪟 Windows Troubleshooting:"
            echo "   • Ensure you're running in Git Bash, PowerShell, or WSL"
            echo "   • Try running as Administrator"
            echo "   • Check if antivirus is blocking the download"
            echo "   • Manual installation:"
            echo "     1. Download from: https://github.com/$REPO/releases"
            echo "     2. Extract to a folder in your PATH"
            echo ""
        fi
        echo "💡 Alternative: Install Go and try:"
        echo "   go install github.com/$REPO/cmd/git-clone@latest"
        exit 1
    fi
fi

echo ""
echo "🎉 Installation complete!"
echo ""
echo "🛠️ You can now run:"
echo "   $BINARY_NAME"
echo ""
echo "🔍 To verify installation:"
echo "   which $BINARY_NAME"
echo ""

# Windows-specific guidance
if [ "$OS" = "windows" ]; then
    echo "🪟 Windows Note:"
    echo "   If command not found, ensure $INSTALL_DIR is in your PATH"
    echo "   You may need to restart your terminal/Git Bash"
    echo ""
fi

echo "📚 For usage instructions, visit:"
echo "   https://github.com/$REPO"
echo ""
echo "🚀 To uninstall later:"
if [ "$OS" = "windows" ]; then
    echo "   rm $INSTALL_DIR/$BINARY_NAME"
else
    echo "   sudo rm $INSTALL_DIR/$BINARY_NAME"
fi
