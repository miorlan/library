package migration

import (
	"database/sql"
	"log"
)

func Migrate(db *sql.DB) {
	query := `
        CREATE TABLE IF NOT EXISTS songs (
            id SERIAL PRIMARY KEY,
            band VARCHAR(255) NOT NULL,
            song VARCHAR(255) NOT NULL,
            release_date VARCHAR(255),
            text TEXT,
            link VARCHAR(255)
        );
    `
	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
