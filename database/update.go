package database

import (
	"database/sql"
	"fmt"
)

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
		fmt.Println(err)
	}

	fmt.Println("Update successful")
}
