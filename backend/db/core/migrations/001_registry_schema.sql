CREATE TABLE IF NOT EXISTS databases (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    path TEXT NOT NULL,
    purpose TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);

INSERT OR IGNORE INTO databases (name, path, purpose) VALUES
    ("content",     "./backend/db/content/content.db",        "content storage"),
    ("definitions", "./backend/db/definitions/definitions.db","definition store"),
    ("tokens",      "./backend/db/tokens/tokens.db",          "design tokens");
