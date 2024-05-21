// database/update.go

package database

import (
	"database/sql"
	"log"
)

// UpdateCrimeStatus mengupdate status kejahatan sesuai dengan kondisi tertentu
func UpdateCrimeStatus(db *sql.DB) {
	updateQuery := `
        UPDATE criminal_records
        SET crime_status =
            CASE 
                WHEN suspect_name = 'Herdi' OR suspect_name = 'Ahmad' THEN 'Open'
                ELSE 'Close'
            END
    `

	_, err := db.Exec(updateQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Update successful")
}
