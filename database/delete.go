package database

import (
	"database/sql"
	"fmt"
)

func DeleteCrimeData(db *sql.DB) {
	deleteQuery := `
	delete from criminal_records 
	where crime_status = 'Close';	
	`
	_, err := db.Exec(deleteQuery)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Delete successful")
}
