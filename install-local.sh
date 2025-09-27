#!/bin/bash

# Local Installation Script for gclone
# Use this before pushing to GitHub

set -e

BINARY_NAME="gclone"
INSTALL_DIR="/usr/local/bin"

echo "🚀 Installing gclone locally..."
echo "==============================="

# Check if binary exists
if [ ! -f "bin/$BINARY_NAME" ]; then
    echo "📦 Building $BINARY_NAME first..."
    make build
fi

if [ ! -f "bin/$BINARY_NAME" ]; then
    echo "❌ Could not find bin/$BINARY_NAME"
    echo "💡 Run 'make build' first"
    exit 1
fi

echo "📍 Installing $BINARY_NAME to $INSTALL_DIR"

# Install binary
if [ -w "$INSTALL_DIR" ]; then
    cp "bin/$BINARY_NAME" "$INSTALL_DIR/"
else
    sudo cp "bin/$BINARY_NAME" "$INSTALL_DIR/"
fi

# Make sure it's executable
if [ -w "$INSTALL_DIR/$BINARY_NAME" ]; then
    chmod +x "$INSTALL_DIR/$BINARY_NAME"
else
    sudo chmod +x "$INSTALL_DIR/$BINARY_NAME"
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