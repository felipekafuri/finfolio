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

```bash
go build -o finfolio
```

## Usage

```bash
# Add a new investment
finfolio add

# List all investments
finfolio list

# View specific investment details
finfolio show <id>
```

## Learning Go

This project is built as a learning exercise, exploring:
- CLI development with Cobra
- Database operations with SQLite
- Date/time calculations
- Percentage and compound interest formulas