package database

import "database/sql"

func ClearPostgresSchema(db *sql.DB) error {
	_, err := db.Exec(`
		DROP SCHEMA public CASCADE;
		CREATE SCHEMA public;
		GRANT ALL ON SCHEMA public TO postgres;
		GRANT ALL ON SCHEMA public TO public;
	`)
	return err
}
