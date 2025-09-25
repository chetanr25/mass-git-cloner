#!/bin/bash

# Mass Git Cloner - Global Installer Script
# This provides an NPM-like installation experience for the Go binary

set -e

REPO="chetanr25/mass-git-cloner"
BINARY_NAME="gclone"  # Shorter, more memorable name
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

# Get latest release info
RELEASE_URL="https://api.github.com/repos/$REPO/releases/latest"
DOWNLOAD_URL=$(curl -s $RELEASE_URL | grep "browser_download_url.*$BINARY_NAME-$PLATFORM" | cut -d '"' -f 4)

if [ -z "$DOWNLOAD_URL" ]; then
    echo "❌ Could not find binary for platform $PLATFORM"
    echo "💡 Try installing with Go: go install github.com/$REPO/cmd/git-clone@latest"
    exit 1
fi

echo "⬇️  Downloading from: $DOWNLOAD_URL"

# Create temporary directory
TMP_DIR=$(mktemp -d)
cd "$TMP_DIR"

# Download binary
curl -L -o "$BINARY_NAME" "$DOWNLOAD_URL"
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

echo "✅ Successfully installed $BINARY_NAME!"
echo ""
echo "🎉 You can now run: $BINARY_NAME"
echo ""
echo "📚 For usage instructions, visit:"
echo "   https://github.com/$REPO"