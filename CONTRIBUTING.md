# Contributing to Finfolio

Thank you for your interest in contributing to Finfolio!

## Development Setup

### Prerequisites

- Go 1.24.5 or later
- Git
- GoReleaser (for releases)
- GitHub CLI (optional, for easier repo management)

### Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/felipekafuri/finfolio.git
   cd finfolio
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the application**
   ```bash
   make run
   # or
   go run main.go add
   ```

## Development Workflow

### Using Make Commands

We provide a Makefile with common tasks:

```bash
make help          # Show all available commands
make build         # Build the binary
make install       # Install to $GOPATH/bin
make run           # Run the application
make dev CMD="add" # Run with specific command
make test          # Run tests
make clean         # Remove build artifacts
make fmt           # Format code
make tidy          # Tidy go modules
```

### Building

```bash
# Build locally
make build

# Test the binary
./finfolio add
```

### Running Tests

```bash
make test
```

### Code Style

- Run `make fmt` before committing
- Follow standard Go conventions
- Use meaningful variable and function names

## Release Process

### Creating a New Release

#### 1. Prerequisites

- Ensure all tests pass: `make test`
- Ensure code is formatted: `make fmt`
- Commit all changes

#### 2. Get GitHub Token

Create a Personal Access Token:

**Option A: Fine-grained token (recommended)**
1. Go to https://github.com/settings/personal-access-tokens/new
2. Token name: `finfolio-releases`
3. Expiration: 90 days (set calendar reminder)
4. Repository access: Only select repositories → `finfolio`
5. Permissions:
   - Contents: Read and write
   - Metadata: Read-only
6. Generate token and save it securely

**Option B: Classic token**
1. Go to https://github.com/settings/tokens/new
2. Select `public_repo` scope
3. Set expiration (90 days recommended)
4. Generate token

#### 3. Export Token

```bash
export GITHUB_TOKEN="your_github_token_here"
```

**Tip:** Add to your `.zshrc` or `.bashrc` (but never commit it!):
```bash
export GITHUB_TOKEN="your_token_here"
```

#### 4. Create Release

**Manual approach:**
```bash
# Create and push tag
git tag -a v0.2.0 -m "Release v0.2.0: Description of changes"
git push origin v0.2.0

# Build and publish release
goreleaser release --clean
```

**Using Make:**
```bash
make release
# This will prompt you for the version number and handle everything
```

#### 5. Test Local Build (Optional)

Before creating a real release, test the build locally:

```bash
make snapshot
# Check dist/ folder for binaries
./dist/finfolio_darwin_arm64_v8.0/finfolio add
```

### Release Checklist

- [ ] All changes committed
- [ ] Tests pass (`make test`)
- [ ] Code formatted (`make fmt`)
- [ ] GitHub token exported
- [ ] Version number decided (semantic versioning)
- [ ] Changelog updated (if maintaining one)
- [ ] Tag created and pushed
- [ ] Release created with goreleaser
- [ ] GitHub release verified with binaries
- [ ] Release notes updated on GitHub

### Versioning

We follow [Semantic Versioning](https://semver.org/):

- `v0.1.0` → `v0.2.0`: New features (minor)
- `v0.2.0` → `v0.2.1`: Bug fixes (patch)
- `v0.9.0` → `v1.0.0`: Breaking changes (major)

### What GoReleaser Does

When you run `goreleaser release --clean`:

1. Validates git state and tag
2. Runs `go mod tidy`
3. Builds binaries for:
   - Linux (amd64, arm64)
   - macOS (amd64, arm64)
   - Windows (amd64, arm64)
4. Creates archives (tar.gz for Unix, zip for Windows)
5. Generates checksums
6. Creates GitHub release
7. Uploads all binaries to the release
8. Generates changelog from commits

## Project Structure

```
finfolio/
├── cmd/                    # CLI commands
│   ├── root.go            # Root command
│   └── add.go             # Add investment command
├── internal/              # Private application code
│   ├── investment/        # Investment domain
│   │   ├── model.go       # Investment struct
│   │   └── service.go     # Business logic
│   └── ui/                # User interface
│       └── add_form.go    # Bubbletea form
├── main.go                # Entry point
├── go.mod                 # Go modules
├── Makefile               # Build commands
├── .goreleaser.yaml       # Release configuration
└── README.md              # User documentation

```

## Adding New Commands

1. Create new file in `cmd/` (e.g., `cmd/list.go`)
2. Define command using Cobra
3. Register command in `init()` function
4. Add any UI components to `internal/ui/`
5. Add business logic to appropriate service

Example:
```go
// cmd/list.go
package cmd

import (
    "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all investments",
    Run: func(cmd *cobra.Command, args []string) {
        // Implementation
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
```

## Questions?

Feel free to open an issue for any questions about contributing or the release process!