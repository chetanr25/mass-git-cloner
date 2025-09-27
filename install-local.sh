#!/bin/bash

# Installation Script for gclone
# Can be used locally or via curl from GitHub

set -e

BINARY_NAME="gclone"
INSTALL_DIR="/usr/local/bin"
REPO_URL="https://github.com/chetanr25/mass-git-cloner"
TEMP_DIR=$(mktemp -d)

echo "🚀 Installing gclone..."
echo "======================="

# Function to build from source
build_from_source() {
    echo "📦 Building $BINARY_NAME from source..."
    
    # Check if we have Go installed
    if ! command -v go &> /dev/null; then
        echo "❌ Go is not installed. Please install Go first."
        echo "💡 Visit: https://golang.org/doc/install"
        exit 1
    fi
    
    # Clone repository to temp directory
    echo "📥 Downloading source code..."
    cd "$TEMP_DIR"
    git clone "$REPO_URL" .
    
    # Build the binary
    echo "🔨 Building..."
    go build -o "$BINARY_NAME" ./cmd/git-clone
    
    echo "✅ Build complete!"
}

# Check if we're in the project directory (local install)
if [ -f "go.mod" ] && [ -d "cmd/git-clone" ]; then
    echo "📍 Detected local repository"
    # Check if binary exists
    if [ ! -f "bin/$BINARY_NAME" ]; then
        echo "📦 Building $BINARY_NAME first..."
        if [ -f "Makefile" ]; then
            make build
        else
            mkdir -p bin
            go build -o "bin/$BINARY_NAME" ./cmd/git-clone
        fi
    fi
    BINARY_PATH="bin/$BINARY_NAME"
else
    echo "📍 Installing from remote repository"
    build_from_source
    BINARY_PATH="$TEMP_DIR/$BINARY_NAME"
fi

if [ ! -f "$BINARY_PATH" ]; then
    echo "❌ Could not build $BINARY_NAME"
    echo "💡 Please check the error messages above"
    exit 1
fi

echo "📍 Installing $BINARY_NAME to $INSTALL_DIR"

# Install binary
if [ -w "$INSTALL_DIR" ]; then
    cp "$BINARY_PATH" "$INSTALL_DIR/"
else
    sudo cp "$BINARY_PATH" "$INSTALL_DIR/"
fi

# Make sure it's executable
if [ -w "$INSTALL_DIR/$BINARY_NAME" ]; then
    chmod +x "$INSTALL_DIR/$BINARY_NAME"
else
    sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
fi

# Cleanup temp directory if used
if [ "$BINARY_PATH" = "$TEMP_DIR/$BINARY_NAME" ]; then
    echo "🧹 Cleaning up..."
    rm -rf "$TEMP_DIR"
fi

echo "✅ Successfully installed $BINARY_NAME!"
echo ""
echo "🎉 You can now run: $BINARY_NAME"
echo ""
echo "📝 To test it:"
echo "   $BINARY_NAME"
echo ""
echo "🚀 To uninstall later:"
echo "   sudo rm $INSTALL_DIR/$BINARY_NAME"