package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	_ "modernc.org/sqlite"
)

func InitDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(fmt.Sprintf("failed to open database: %v", err))
	}

	_, err = db.Exec(`PRAGMA foreign_keys = ON;`)
	if err != nil {
		panic(fmt.Sprintf("failed to enable foreign keys: %v", err))
	}

	runMigrations(db)
	return db
}

func runMigrations(db *sql.DB) {
	migrationsDir := "./internal/db/migrations"

	entries, err := os.ReadDir(migrationsDir)
	if err != nil {
		panic(fmt.Sprintf("failed to read migrations directory: %v", err))
	}

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

		_, err = db.Exec(string(content))
		if err != nil {
			panic(fmt.Sprintf("migration %s failed: %v", entry.Name(), err))
		}

		fmt.Printf("Migration OK: %s\n", entry.Name())
	}
}

func LoadAllDatabases(registry *sql.DB) map[string]*sql.DB {
	rows, _ := registry.Query("SELECT name, path FROM databases")

	dbs := make(map[string]*sql.DB)

	for rows.Next() {
		var name, path string
		rows.Scan(&name, &path)

		db, err := sql.Open("sqlite", path)
		if err != nil {
			panic(err)
		}

		dbs[name] = db
	}

	return dbs
}
