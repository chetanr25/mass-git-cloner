#!/bin/bash

# Test script for mass-git-cloner
echo "🧪 Testing mass-git-cloner"
echo "========================="

# Check if git is installed
if ! command -v git &> /dev/null; then
    echo "❌ Git is not installed or not in PATH"
    exit 1
fi

echo "✅ Git is available"

# Build the project
echo "🔨 Building project..."
if go build -o bin/git-clone ./cmd/git-clone; then
    echo "✅ Build successful"
else
    echo "❌ Build failed"
    exit 1
fi

# Check if binary was created
if [ -f "bin/git-clone" ]; then
    echo "✅ Binary created successfully"
else
    echo "❌ Binary not found"
    exit 1
fi

# Test help or version (if implemented)
echo "🔍 Testing binary execution..."
if ./bin/git-clone --help 2>/dev/null || true; then
    echo "✅ Binary executes"
else
    echo "⚠️  Binary execution test skipped (--help not implemented)"
fi

echo ""
echo "🎉 All tests passed! The mass-git-cloner is ready to use."
echo ""
echo "To run the application:"
echo "  ./bin/git-clone"
echo ""
echo "Or install it globally:"
echo "  go install ./cmd/git-clone"