#!/bin/bash

# Build script for creating releases
set -e

VERSION=${1:-"v1.0.0"}
BUILD_DIR="releases"
CMD_DIR="cmd/git-clone"

echo "🚀 Building releases for version $VERSION"
echo "========================================"

# Clean and create build directory
rm -rf $BUILD_DIR
mkdir -p $BUILD_DIR

# Build for different platforms
platforms=(
    "linux/amd64"
    "linux/arm64" 
    "darwin/amd64"
    "darwin/arm64"
    "windows/amd64"
    "windows/386"
)

for platform in "${platforms[@]}"; do
    GOOS=${platform%/*}
    GOARCH=${platform#*/}
    
    if [ "$GOOS" = "windows" ]; then
        output_name="gclone-${GOOS}-${GOARCH}.exe"
    else
        output_name="gclone-${GOOS}-${GOARCH}"
    fi
    
    echo "📦 Building for $GOOS/$GOARCH..."
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "$BUILD_DIR/$output_name" ./$CMD_DIR
done

echo ""
echo "✅ Build complete! Files created:"
ls -la $BUILD_DIR/

echo ""
echo "📋 Next steps to create a GitHub release:"
echo "1. git tag $VERSION"
echo "2. git push origin $VERSION"
echo "3. Go to GitHub > Releases > Create new release"
echo "4. Upload the files from $BUILD_DIR/"
echo ""
echo "🌐 Or use GitHub CLI:"
echo "gh release create $VERSION $BUILD_DIR/* --title \"Release $VERSION\" --notes \"Release notes here\""