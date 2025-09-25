# 🚀 Mass Git Cloner

A beautiful, interactive command-line tool for cloning multiple repositories from a GitHub user or organization. Built with Go and featuring a stunning terminal UI powered by Bubbletea.

## ✨ Features

### 🎨 Beautiful Terminal UI
- **Interactive Navigation**: Use arrow keys or vim keys (j/k) to navigate
- **Visual Feedback**: Color-coded selections and status indicators  
- **Smooth Animations**: Elegant transitions and progress indicators
- **Responsive Design**: Adapts to different terminal sizes

### 📊 Smart Repository Management
- **Repository Statistics**: Visual breakdown of forks, originals, public/private repos
- **Intelligent Filtering**: Choose between all repos, non-forks only, or forks only
- **Batch Selection**: Select individual repos, ranges, or use select all/none
- **Confirmation Dialog**: Review your selection before cloning

### ⚡ High-Performance Cloning
- **Concurrent Cloning**: Clone multiple repositories simultaneously
- **Progress Tracking**: Real-time progress bars with success/failure indicators
- **Error Handling**: Retry failed clones automatically with graceful error recovery
- **HTTPS Cloning**: Secure HTTPS-based repository cloning

### 🛠️ Advanced Configuration
- **Environment Variables**: Customize behavior via config
- **GitHub Token Support**: Higher API limits with personal access tokens
- **Flexible Directory Structure**: Organized output in `username/` folders
- **Signal Handling**: Graceful shutdown with Ctrl+C

## 🎮 Controls

### Repository Selection
- `↑/k`: Move up
- `↓/j`: Move down  
- `Space`: Toggle repository selection
- `a`: Select all repositories
- `n`: Select none (clear all selections)
- `Enter`: Confirm selection and proceed
- `q`: Quit application

### Filter Selection
- `↑/↓` or `j/k`: Navigate filter options
- `Enter` or `Space`: Select filter type
- `q`: Quit

## 🚀 Quick Start

### One-Command Installation

```bash
# Install with one command (like NPM global install)
curl -sSL https://raw.githubusercontent.com/chetanr25/mass-git-cloner/main/install.sh | bash
```

### Alternative Installation Methods

#### Via Go (if you have Go installed)
```bash
go install github.com/chetanr25/mass-git-cloner/cmd/git-clone@latest
```

#### Manual Installation
```bash
# Clone and build
git clone https://github.com/chetanr25/mass-git-cloner.git
cd mass-git-cloner
make build

# Or use the setup script
./setup-distribution.sh
```

### Usage

```bash
# Run the application locally
./bin/gclone

# Or if installed globally
gclone
```

### First Run Experience

1. **Enter Username**: Type a GitHub username or organization name
2. **View Statistics**: See a beautiful breakdown of repository statistics
3. **Choose Filter**: Select what type of repositories to show
4. **Select Repositories**: Use keyboard navigation to select repos
5. **Confirm Selection**: Review and confirm your choices
6. **Watch Progress**: Monitor real-time cloning progress

## 🎨 UI Screenshots

### Welcome & Statistics Display
```
🚀 Mass Git Cloner - username
┌─────────────────────────────────────────────┐
│ ✅ Found 24 repositories                    │
└─────────────────────────────────────────────┘

┌─────────────────────────────────────────────┐
│ 📊 Repository Statistics                    │
│                                            │
│ Total repositories:      ████████████ 24   │
│ Original repositories:   █████████ 17      │
│ Forked repositories:     ███ 7             │
│ Public repositories:     ████████████ 24   │
│ Private repositories:    ████████████ 0    │
│                                            │
│ 📈 Breakdown:                              │
│    Original: 70.8%                        │
│    Forks: 29.2%                          │
└─────────────────────────────────────────────┘
```

### Interactive Repository Selection
```
🚀 Mass Git Cloner - Repository Selection

Non-fork repositories only - 3 selected
════════════════════════════════════════════════════════════════

❯ ☑ quick-docs              Dart      ★ 5    ⑂ 0  Document management app
  ☐ chetanr25               N/A       ★ 0    ⑂ 0  No description  
  ☐ locadora                Dart      ★ 0    ⑂ 0  No description
  ☑ aventus-website         JavaScript ★ 0    ⑂ 0  No description
  ☑ qr_code                 Dart      ★ 5    ⑂ 0  QR Code scanner app
  ☐ BugVengers              Dart      ★ 0    ⑂ 0  No description

Controls:
  ↑/k: Move up    ↓/j: Move down    Space: Toggle selection
  a: Select all   n: Select none    Enter: Confirm selection
  q: Quit
```

### Cloning Progress
```
🚀 Mass Git Cloner - Cloning Progress

[██████████████████████░░░░░░░░░░] 75.0% (3/4)

┌─────────────────────────────────┐
│ ⏱️  Elapsed: 1m23s              │
│ ✅ Completed: 3                 │
│ ❌ Failed: 0                    │
│ 📈 Success Rate: 100.0%         │
└─────────────────────────────────┘

⏳ Cloning: my-awesome-project

📋 Recent completions:
  ✅ quick-docs (2.3s)
  ✅ qr_code (5.1s)
  ✅ aventus-website (1.8s)
```

## ⚙️ Configuration

### Environment Variables

Create a `.env` file or set environment variables:

```bash
# GitHub personal access token (recommended for higher API limits)
GITHUB_TOKEN=ghp_your_token_here

# Maximum concurrent clones (default: 5)  
MAX_CONCURRENCY=3

# Timeout for each clone operation (default: 10m)
CLONE_TIMEOUT=5m

# Base directory for cloned repositories (default: current directory)
BASE_DIR=/path/to/clone/directory
```

### GitHub Token Setup

1. Go to GitHub Settings > Developer Settings > Personal Access Tokens
2. Generate a new token with `public_repo` scope (and `repo` for private repos)
3. Set the `GITHUB_TOKEN` environment variable

## 🔧 Development

### Project Structure

```
mass-git-cloner/
├── cmd/git-clone/           # Main application entry point
├── internal/
│   ├── github/             # GitHub API client and operations
│   ├── ui/                 # Bubbletea UI components
│   ├── cloner/             # Git cloning logic and worker pools
│   └── config/             # Configuration management
├── pkg/models/             # Shared data models
└── Makefile               # Build automation
```

### Building from Source

```bash
# Download dependencies
make deps

# Format code
make fmt

# Build application
make build

# Run tests
make test

# Clean build artifacts  
make clean
```

### Available Make Commands

```bash
make build     # Build the application
make clean     # Clean build artifacts  
make test      # Run tests
make run       # Build and run
make install   # Install globally
make fmt       # Format code
make deps      # Download dependencies
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [Bubbletea](https://github.com/charmbracelet/bubbletea) - The amazing TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling library
- GitHub API for providing excellent repository data

## 🚀 Roadmap

- [ ] Support for GitLab and other Git hosting services
- [ ] Repository search and filtering by language, stars, etc.
- [ ] Clone progress persistence and resume functionality  
- [ ] Integration with Git LFS for large files
- [ ] Custom clone templates and post-clone scripts
- [ ] Repository dependency analysis and smart ordering

---

**Made with ❤️ and lots of ☕ by [chetanr25](https://github.com/chetanr25)**