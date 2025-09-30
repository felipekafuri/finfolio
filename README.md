# Finfolio

A CLI tool for managing your fixed-income investment portfolio.

## Overview

Finfolio helps you track and analyze your investments locally. It stores information about your applications, redemptions, returns, and taxes, with automatic calculations for:

- Investment period (days)
- Net returns (after taxes)
- Period return percentage
- Annualized return percentage

## Features

- **Add investments**: Record application date, value, bank, title, and expected redemption
- **Track returns**: Calculate gross/net returns and tax deductions
- **View portfolio**: List all investments with status indicators
- **Automatic calculations**: Period, percentage returns, and annualized rates
- **Status indicators**: Identify redeemed, upcoming redemptions, and active investments

## Data Model

Each investment tracks:
- Application date
- Initial value
- Redemption date
- Bank and title
- Gross return, tax (IR), and net return
- Calculated: period (days), % return, annualized %

## Storage

Data is stored locally using SQLite, keeping your financial information private and portable.

## Installation

### From Source

```bash
git clone https://github.com/felipekafuri/finfolio.git
cd finfolio
go build -o finfolio
```

### From Release

Download the latest release from the [releases page](https://github.com/felipekafuri/finfolio/releases).

## Usage

```bash
# Add a new investment (beautiful interactive form)
finfolio add

# List all investments (coming soon)
finfolio list

# View specific investment details (coming soon)
finfolio show <id>
```

## Development

### Prerequisites

- Go 1.21+
- SQLite3
- [golang-migrate](https://github.com/golang-migrate/migrate) CLI (for migrations)

### Setup

```bash
# Install dependencies
go mod download

# Install migration CLI
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Create a new migration
make migrate-create

# Run migrations
make migrate-up

# Run in development
make dev CMD="add"
```

### Makefile Commands

```bash
make help            # Show all available commands
make build           # Build the binary
make dev             # Run in development mode
make test            # Run tests
make migrate-create  # Create a new migration
make migrate-up      # Run pending migrations
make migrate-down    # Rollback last migration
make migrate-status  # Show migration status
```

## Tech Stack

- **CLI Framework**: [Cobra](https://github.com/spf13/cobra)
- **TUI**: [Bubble Tea](https://github.com/charmbracelet/bubbletea) + [Lipgloss](https://github.com/charmbracelet/lipgloss)
- **Database**: SQLite with [go-sqlite3](https://github.com/mattn/go-sqlite3)
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)

## Learning Go

This project is built as a learning exercise, exploring:
- CLI development with Cobra
- Beautiful TUIs with Bubble Tea and Lipgloss
- Database operations with SQLite and migrations
- Date/time calculations
- Percentage and compound interest formulas
- Cross-platform builds with GoReleaser