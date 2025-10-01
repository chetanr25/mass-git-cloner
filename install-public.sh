#!/bin/bash

# Universal installer for gclone
# Usage: curl -sSL https://raw.githubusercontent.com/chetanr25/mass-git-cloner/main/install-public.sh | bash

set -e

REPO="chetanr25/mass-git-cloner"
BINARY_NAME="gclone"
INSTALL_DIR="/usr/local/bin"

echo "🚀 Installing gclone - Mass Git Cloner"
echo "======================================"

# Detect OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case $ARCH in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    armv7l) ARCH="arm" ;;
    *) echo "❌ Unsupported architecture: $ARCH" && exit 1 ;;
esac

case $OS in
    linux) PLATFORM="linux-$ARCH" ;;
    darwin) PLATFORM="darwin-$ARCH" ;;
    *) echo "❌ Unsupported OS: $OS" && exit 1 ;;
esac

echo "📍 Detected platform: $PLATFORM"

# Check if we have write access to /usr/local/bin
if [ ! -w "$INSTALL_DIR" ]; then
    echo "🔐 Need sudo access to install to $INSTALL_DIR"
    USE_SUDO="sudo"
else
    USE_SUDO=""
fi

# Get latest release info from GitHub API
echo "🔍 Checking for latest release..."
LATEST_RELEASE=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep -o '"tag_name": "[^"]*' | cut -d'"' -f4)

if [ -z "$LATEST_RELEASE" ]; then
    echo "❌ Could not fetch latest release information"
    exit 1
fi

echo "📦 Latest version: $LATEST_RELEASE"

# Download URL
if [ "$OS" = "windows" ]; then
    DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/gclone-$PLATFORM.exe"
    BINARY_NAME="gclone.exe"
else
    DOWNLOAD_URL="https://github.com/$REPO/releases/download/$LATEST_RELEASE/gclone-$PLATFORM"
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
    $USE_SUDO mv "$BINARY_NAME" "$INSTALL_DIR/"
    
    # Cleanup
    cd - > /dev/null
    rm -rf "$TMP_DIR"
    
    echo ""
    echo "✅ Installation complete!"
    echo ""
    echo "🎯 You can now run:"
    echo "   gclone"
    echo ""
    echo "🔍 To verify installation:"
    echo "   which gclone"
    echo "   gclone --help"
    echo ""
    echo "📚 For usage instructions and documentation:"
    echo "   https://github.com/$REPO"
    echo ""
    echo "🚀 To uninstall:"
    echo "   $USE_SUDO rm $INSTALL_DIR/$BINARY_NAME"
    
else
    echo "❌ Download failed"
    cd - > /dev/null
    rm -rf "$TMP_DIR"
    exit 1
fi