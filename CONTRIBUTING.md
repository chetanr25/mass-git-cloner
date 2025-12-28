# Contributing to Mass Git Cloner

Thank you for your interest in contributing to Mass Git Cloner (gclone)! This document provides guidelines and instructions for contributing to the project.

## How to Contribute

Contributions come in many forms:
- 🐛 Reporting bugs
- 💡 Suggesting new features
- 📝 Improving documentation
- 🔧 Submitting code changes
- 🧪 Writing tests

## 🚀 Getting Started

### Prerequisites

- **Go 1.25.1 or later** - [Install Go](https://go.dev/doc/install)
- **Git** - For version control
- **GitHub account** - For forking and creating pull requests

### Setting Up Your Development Environment

1. **Fork the repository**
   ```bash
   # Navigate to: https://github.com/chetanr25/mass-git-cloner
   # Click the "Fork" button
   ```

2. **Clone your fork**
   ```bash
   git clone https://github.com/YOUR_USERNAME/mass-git-cloner.git
   cd mass-git-cloner
   ```

3. **Add the upstream repository**
   ```bash
   git remote add upstream https://github.com/chetanr25/mass-git-cloner.git
   ```

4. **Install dependencies**
   ```bash
   go mod download
   ```

5. **Build the project**
   ```bash
   go build -o gclone ./cmd/git-clone
   ```

6. **Test the installation**
   ```bash
   ./gclone --help  # or ./gclone (if help flag exists)
   ```

## 📁 Project Structure

Understanding the project structure will help you navigate the codebase:

```
mass-git-cloner/
├── cmd/
│   └── git-clone/          # Main application entry point
│       └── main.go
├── internal/               # Internal packages (not for external use)
│   ├── cloner/            # Git cloning logic
│   ├── config/            # Configuration management
│   ├── github/            # GitHub API client
│   └── ui/                # User interface (Bubbletea)
├── pkg/
│   └── models/            # Public models/types
├── bin/                   # Binary output directory
├── releases/              # Pre-built binaries
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
└── README.md             # Project documentation
```

## 🔨 Development Workflow

### Making Changes

1. **Create a new branch**
   ```bash
   git checkout -b feature/your-feature-name
   # or
   git checkout -b fix/your-bug-fix
   ```

2. **Make your changes**
   - Write clean, readable code
   - Follow Go conventions and best practices
   - Add comments for complex logic
   - Update documentation if needed

3. **Test your changes**
   ```bash
   # Run tests (if available)
   go test ./...
   
   # Build and test manually
   go build -o gclone ./cmd/git-clone
   ./gclone
   ```

4. **Commit your changes**
   ```bash
   git add .
   git commit -m "feat: add new feature description"
   ```

5. **Keep your branch updated**
   ```bash
   git fetch upstream
   git rebase upstream/main
   ```

6. **Push to your fork**
   ```bash
   git push origin feature/your-feature-name
   ```

### Commit Message Guidelines

We follow [Conventional Commits](https://www.conventionalcommits.org/) format:

- **feat**: New feature
- **fix**: Bug fix
- **docs**: Documentation changes
- **style**: Code style changes (formatting, etc.)
- **refactor**: Code refactoring
- **test**: Adding or updating tests
- **chore**: Maintenance tasks

Examples:
```
feat: add support for cloning private repositories
fix: resolve issue with repository filtering
docs: update installation instructions
refactor: improve error handling in cloner package
```

## Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test ./internal/cloner

# Run tests with coverage
go test -cover ./...
```

### Writing Tests

- Write tests for new features and bug fixes
- Aim for good test coverage
- Use descriptive test names
- Test both success and error cases

## Code Style Guidelines

### Go Code Style

- Follow [Effective Go](https://go.dev/doc/effective_go) guidelines
- Use `gofmt` to format your code:
  ```bash
  gofmt -w .
  ```
- Use `golint` or `golangci-lint` for linting (if configured)
- Keep functions small and focused
- Use meaningful variable and function names
- Add comments for exported functions and types

### Code Organization

- Keep related code together
- Separate concerns (UI, business logic, API clients)
- Use interfaces where appropriate for testability
- Avoid deep nesting

## Reporting Bugs

When reporting bugs, please include:

1. **Description**: Clear description of the bug
2. **Steps to Reproduce**: Detailed steps to reproduce the issue
3. **Expected Behavior**: What should happen
4. **Actual Behavior**: What actually happens
5. **Environment**:
   - OS and version
   - Go version
   - gclone version (if applicable)
6. **Error Messages**: Full error messages or logs
7. **Screenshots**: If applicable

Create an issue using the [GitHub issue template](https://github.com/chetanr25/mass-git-cloner/issues/new).

## Suggesting Features

When suggesting features:

1. **Check existing issues** to avoid duplicates
2. **Describe the use case** - Why is this feature needed?
3. **Provide examples** - How would it be used?
4. **Consider alternatives** - Are there workarounds?

## Code Review Process

1. **Submit a Pull Request** from your fork to the main repository
2. **Ensure your PR**:
   - Has a clear title and description
   - References related issues (if any)
   - Includes tests (if applicable)
   - Updates documentation (if needed)
   - Builds successfully
3. **Respond to feedback** from maintainers
4. **Make requested changes** and update your PR
5. **Wait for approval** before merging

### Pull Request Checklist

- [ ] Code follows the project's style guidelines
- [ ] Code is tested and working
- [ ] Documentation is updated (if needed)
- [ ] Commit messages follow the conventional format
- [ ] Branch is up to date with upstream/main
- [ ] No merge conflicts

## 🏗️ Building for Different Platforms

The project supports multiple platforms. To build for a specific platform:

```bash
# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o gclone-darwin-amd64 ./cmd/git-clone

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o gclone-darwin-arm64 ./cmd/git-clone

# Linux (x86_64)
GOOS=linux GOARCH=amd64 go build -o gclone-linux-amd64 ./cmd/git-clone

# Linux (ARM64)
GOOS=linux GOARCH=arm64 go build -o gclone-linux-arm64 ./cmd/git-clone

# Windows (x86_64)
GOOS=windows GOARCH=amd64 go build -o gclone-windows-amd64.exe ./cmd/git-clone

# Windows (32-bit)
GOOS=windows GOARCH=386 go build -o gclone-windows-386.exe ./cmd/git-clone
```

## Additional Resources

- [Go Documentation](https://go.dev/doc/)
- [Bubbletea Documentation](https://github.com/charmbracelet/bubbletea)
- [GitHub API Documentation](https://docs.github.com/en/rest)

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to Mass Git Cloner! 🎉

