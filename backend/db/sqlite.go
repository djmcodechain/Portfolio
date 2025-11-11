package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	_ "modernc.org/sqlite"
)

type DatabaseMap map[string]*sql.DB

// -----------------------------
// Init core registry database
// -----------------------------
func InitDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(fmt.Sprintf("failed to open database: %v", err))
	}

	if _, err := db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		panic(fmt.Sprintf("failed to enable foreign keys: %v", err))
	}

	runMigrations(db, filepath.Dir(dbPath)+"/migrations")

	return db
}

// -----------------------------
// Load domain DBs listed in registry
// -----------------------------
func LoadAllDatabases(registry *sql.DB) (DatabaseMap, error) {
	rows, err := registry.Query(`SELECT name, path FROM databases`)
	if err != nil {
		return nil, fmt.Errorf("failed to query database registry: %w", err)
	}
	defer rows.Close()

	dbs := make(DatabaseMap)

	for rows.Next() {
		var name, path string
		if err := rows.Scan(&name, &path); err != nil {
			return nil, fmt.Errorf("failed to scan database row: %w", err)
		}

		// Open domain DB
		db, err := sql.Open("sqlite", path)
		if err != nil {
			return nil, fmt.Errorf("failed to open domain DB %s: %w", name, err)
		}

		if _, err := db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
			return nil, fmt.Errorf("failed to enable FK for %s: %w", name, err)
		}

		// migrations folder = path/../migrations
		migrationDir := filepath.Join(filepath.Dir(path), "migrations")
		runMigrations(db, migrationDir)

		dbs[name] = db
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(dbs) == 0 {
		return nil, errors.New("no domain databases found in registry")
	}

	return dbs, nil
}

// -----------------------------
// Migration runner
// -----------------------------
func runMigrations(db *sql.DB, migrationsDir string) {
	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		panic(fmt.Sprintf("failed to read migrations directory: %v", err))
	}

	// Sort: 000_init.sql → 001_schema.sql → ...
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	for _, entry := range entries {
		if filepath.Ext(entry.Name()) != ".sql" {
			continue
		}

		path := filepath.Join(migrationsDir, entry.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			panic(fmt.Sprintf("failed to read migration %s: %v", entry.Name(), err))
		}

		if _, err := db.Exec(string(content)); err != nil {
			panic(fmt.Sprintf("migration %s failed: %v", entry.Name(), err))
		}

		fmt.Printf("Migration OK: %s (%s)\n", entry.Name(), migrationsDir)
	}
}
