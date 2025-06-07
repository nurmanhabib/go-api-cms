package database

import (
	"database/sql"
	"fmt"
)

func ClearMySQLDatabase(db *sql.DB, dbName string) error {
	// Disable foreign key checks
	if _, err := db.Exec("SET FOREIGN_KEY_CHECKS = 0"); err != nil {
		return err
	}

	// Get all tables
	rows, err := db.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = ?", dbName)
	if err != nil {
		return err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return err
		}
		tables = append(tables, table)
	}

	// Drop all tables
	for _, t := range tables {
		if _, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS `%s`", t)); err != nil {
			return err
		}
	}

	// Re-enable foreign key checks
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	return err
}
