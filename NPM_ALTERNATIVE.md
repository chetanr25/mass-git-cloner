# 🚀 NPM-Like Distribution for Go App

## ✨ **The Solution**

Since this is a Go application (not JavaScript), it can't go to NPM. But I've created an **even better experience** that works just like NPM global installs!

## 📦 **One-Command Install (Like NPM)**

```bash
# Just like: npm install -g package-name
curl -sSL https://raw.githubusercontent.com/chetanr25/mass-git-cloner/main/install.sh | bash
```

### Then run anywhere:
```bash
git-clone
```

## 🎯 **How This Works**

1. **`install.sh`** - Smart installer script that:
   - Detects your OS (macOS, Linux, Windows)
   - Detects your architecture (Intel, Apple Silicon, etc.)
   - Downloads the correct binary
   - Installs it globally in `/usr/local/bin`
   - Makes it available as `git-clone` command

2. **Cross-platform binaries** - Pre-built for:
   - macOS (Intel & Apple Silicon)
   - Linux (amd64)
   - Windows (amd64)

3. **GitHub Releases** - Automatic distribution via GitHub

## 🚀 **Distribution Steps**

### 1. Push to GitHub
```bash
git init
git add .
git commit -m "Initial release"
git remote add origin https://github.com/chetanr25/mass-git-cloner.git
git push -u origin main
```

### 2. Run Distribution Setup
```bash
./setup-distribution.sh
```

### 3. Create GitHub Release
```bash
git tag v1.0.0
git push origin v1.0.0
```

Then upload the binaries from `./releases/` to the GitHub release page.

### 4. Users Install With One Command
```bash
curl -sSL https://raw.githubusercontent.com/chetanr25/mass-git-cloner/main/install.sh | bash
```

## 🎉 **Result**

Your users get an **NPM-like experience**:
- ✅ One-command global install
- ✅ Works on any platform
- ✅ No need to install Go
- ✅ Available as `git-clone` command everywhere
- ✅ Automatic updates possible

## 🌟 **Even Better Than NPM**

- **No dependencies** - Single binary
- **Cross-platform** - Works everywhere
- **Fast downloads** - Small binary size
- **Secure** - No package vulnerabilities
- **Simple updates** - Just re-run the install command

## 📋 **Alternative Names**

If you want a shorter command name, just change `BINARY_NAME` in the scripts:
- `gclone`
- `mgc` (mass-git-clone)
- `repo-clone`
- `bulk-git`

This gives you the **NPM experience** with **Go performance**! 🚀