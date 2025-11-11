CREATE TABLE IF NOT EXISTS definitions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    key TEXT NOT NULL UNIQUE,
    defining_content TEXT NOT NULL,
    the_definition TEXT NOT NULL,
    cross_links_with TEXT
);
