# Finfolio Roadmap

This document outlines the planned features and improvements for Finfolio.

## Current Status (v0.1.0)

✅ Basic CLI structure with Cobra
✅ Interactive `add` command with Bubbletea UI
✅ Investment model and service layer
✅ Date and percentage calculations
✅ Cross-platform builds with GoReleaser

## Short-term Goals (v0.2.0 - v0.4.0)

### Database Integration (v0.2.0) - IN PROGRESS

**Goal:** Persist investment data locally using SQLite

- ✅ Add SQLite dependency (`mattn/go-sqlite3`)
- ✅ Create database schema
  - ✅ `investments` table (id, application_date, value, redemption_date, bank, title, period_days, created_at)
  - [ ] `selic_rates` table (id, date, rate, created_at)
- ✅ Create database initialization and migration system
  - ✅ Set up `golang-migrate` for migrations
  - ✅ Create migration files (up/down)
  - ✅ Makefile commands for migration management
- [ ] Implement repository layer (`internal/investment/repository.go`)
  - [ ] `Create(investment)` - Save new investment
  - [ ] `FindAll()` - Get all investments
  - [ ] `FindByID(id)` - Get specific investment
  - [ ] `Update(investment)` - Update investment
  - [ ] `Delete(id)` - Remove investment
- [ ] Update `add` command to save to database
- ✅ Database file path configured (`./finfolio.db`)

**Files created:**
```
internal/
  database/
    db.go             # ✅ Connection, setup, and migrations
db/
  migrations/
    000001_create_investments_table.up.sql   # ✅
    000001_create_investments_table.down.sql # ✅
```

### List Command (v0.2.0)

**Goal:** View all investments in a formatted table

- [ ] Create `cmd/list.go`
- [ ] Build table UI with Bubbletea/Lipgloss
- [ ] Display key information (ID, Bank, Title, Value, Period, Status)
- [ ] Add status indicators:
  - 🟢 Green: Already redeemed or redeems today
  - 🟡 Yellow: Redeems within 10 days
  - ⚪ White: Active investment (future redemption)
- [ ] Add sorting options (by date, value, bank)
- [ ] Add filtering options (by bank, status)

### Update Command (v0.3.0)

**Goal:** Update investment with returns and tax information

- [ ] Create `cmd/update.go`
- [ ] Interactive form to:
  - [ ] Select investment (from list)
  - [ ] Input gross return
  - [ ] Input tax (IR)
  - [ ] Auto-calculate net return, percentages
- [ ] Update database record
- [ ] Display calculated results

### Show/Detail Command (v0.3.0)

**Goal:** View detailed information about a specific investment

- [ ] Create `cmd/show.go`
- [ ] Accept investment ID as argument
- [ ] Display all fields in a formatted view
- [ ] Show calculated metrics prominently
- [ ] Add comparison with SELIC rate (if available)

### Delete Command (v0.3.0)

**Goal:** Remove investments from portfolio

- [ ] Create `cmd/delete.go`
- [ ] Accept investment ID as argument
- [ ] Show confirmation prompt
- [ ] Remove from database

### SELIC Rate Integration (v0.4.0)

**Goal:** Track Brazilian SELIC rate for benchmarking

**API Options:**
1. **Banco Central do Brasil API** (Official, recommended)
   - Endpoint: `https://api.bcb.gov.br/dados/serie/bcdata.sgs.11/dados?formato=json`
   - Series 11: SELIC daily rate
   - Series 4189: SELIC monthly target
   - Free, no authentication required

2. **Brasil API** (Aggregator)
   - Endpoint: `https://brasilapi.com.br/api/taxas/v1/selic`
   - Simpler interface
   - Community maintained

**Implementation:**
- [ ] Create `internal/selic/` package
  - [ ] `client.go` - API client
  - [ ] `service.go` - Business logic
  - [ ] `repository.go` - Database operations
- [ ] Create `cmd/sync-selic.go` command
  - [ ] Fetch latest SELIC rates
  - [ ] Store in `selic_rates` table
  - [ ] Show sync status
- [ ] Add automatic sync check (warn if data > 30 days old)
- [ ] Display SELIC comparison in investment details
- [ ] Calculate CDI equivalence (common Brazilian fixed-income benchmark)

**Database schema for SELIC:**
```sql
CREATE TABLE selic_rates (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date DATE NOT NULL UNIQUE,
    rate REAL NOT NULL,        -- Annual rate (e.g., 13.75 for 13.75% p.a.)
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## Mid-term Goals (v0.5.0 - v0.7.0)

### Statistics & Analytics (v0.5.0)

- [ ] Create `cmd/stats.go`
- [ ] Total portfolio value
- [ ] Total returns (gross and net)
- [ ] Average return rate
- [ ] Total taxes paid
- [ ] Performance vs SELIC
- [ ] Portfolio composition by bank
- [ ] Time-based charts (if terminal supports)

### Export Command (v0.5.0)

- [ ] Create `cmd/export.go`
- [ ] Export to CSV format
- [ ] Export to Excel (.xlsx)
- [ ] Export to JSON
- [ ] Date range filtering

### Import Command (v0.6.0)

- [ ] Create `cmd/import.go`
- [ ] Import from CSV
- [ ] Import from Excel
- [ ] Validation and error handling
- [ ] Duplicate detection

### Configuration System (v0.6.0)

- [ ] Configuration file (`~/.finfolio/config.yaml`)
- [ ] Configurable database path
- [ ] Default currency and locale
- [ ] API preferences
- [ ] Theme/color customization

### Backup & Restore (v0.7.0)

- [ ] Create `cmd/backup.go`
- [ ] Create `cmd/restore.go`
- [ ] Automatic backup before destructive operations
- [ ] Backup to custom location

## Long-term Goals (v1.0.0+)

### Advanced Features

- [ ] Investment categories/tags
- [ ] Multi-currency support
- [ ] Recurring investments tracking
- [ ] Goal setting and tracking
- [ ] Notifications for upcoming redemptions
- [ ] Web dashboard (optional)
- [ ] Mobile companion app (optional)

### Other Investment Types

- [ ] Stocks tracking
- [ ] Dividends tracking
- [ ] Crypto portfolio
- [ ] International investments

### API & Integrations

- [ ] REST API for external tools
- [ ] Integration with banking APIs
- [ ] Integration with broker APIs
- [ ] Tax report generation (IRPF)

### Performance & Quality

- [ ] Comprehensive test coverage (>80%)
- [ ] Benchmarking suite
- [ ] Performance optimizations for large datasets
- [ ] Documentation improvements
- [ ] Localization (PT-BR, EN)

## Technical Debt & Improvements

### Code Quality

- [ ] Add unit tests for all packages
- [ ] Add integration tests
- [ ] Set up CI/CD pipeline (GitHub Actions)
- [ ] Add code coverage reporting
- [ ] Improve error messages
- [ ] Add logging system

### Documentation

- [ ] API documentation (GoDoc)
- [ ] User guide
- [ ] Video tutorials
- [ ] Architecture decision records (ADRs)

### Developer Experience

- [ ] Add pre-commit hooks
- [ ] Improve development setup docs
- [ ] Add debugging guide
- [ ] Create development Docker container

## Community & Distribution

- [ ] Homebrew tap for macOS installation
- [ ] AUR package for Arch Linux
- [ ] Snap/Flatpak for Linux
- [ ] Chocolatey package for Windows
- [ ] Docker image

## Ideas & Considerations

**Questions to explore:**
- Should we support multiple portfolios?
- Should we add encryption for sensitive data?
- Should we support team/shared portfolios?
- Should we add a TUI (Terminal UI) mode with full-screen interface?

**APIs to consider:**
- **Banco Central API**: https://dadosabertos.bcb.gov.br/
  - SELIC, CDI, IPCA (inflation), exchange rates
- **IBGE API**: Economic indicators
- **B3 API**: Stock market data (if expanding to stocks)

## Contributing

Want to work on something from this roadmap? Check out [CONTRIBUTING.md](CONTRIBUTING.md) for development guidelines.

Feel free to open an issue to discuss new features or suggest changes to this roadmap!
