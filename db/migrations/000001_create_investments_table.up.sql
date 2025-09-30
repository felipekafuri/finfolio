CREATE TABLE IF NOT EXISTS investments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    application_date TEXT NOT NULL,
    value REAL NOT NULL,
    bank TEXT NOT NULL,
    title TEXT NOT NULL,
    redemption_date TEXT NOT NULL,
    period_days INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
