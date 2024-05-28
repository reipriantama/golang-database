package database

import (
	"database/sql"
	"fmt"
	"golang-database/models"
)

func ViewAllData(db *sql.DB) {
	viewData := `
	select * from criminal_records
	`

	rows, err := db.Query(viewData)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	fmt.Println("View all data :")
	for rows.Next() {
		var crimeInfo models.Criminal_records

		err := rows.Scan(&crimeInfo.Id, &crimeInfo.Suspect_name, &crimeInfo.Crime_description, &crimeInfo.Crime_date, &crimeInfo.Crime_status, &crimeInfo.Prison_id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Id : %d, Suspect_name : %s, Crime_descripton :%s, Crime_date : %s, Crime_status : %s, Prison_id : %d \n",
			crimeInfo.Id, crimeInfo.Suspect_name, crimeInfo.Crime_description, crimeInfo.Crime_date, crimeInfo.Crime_status, crimeInfo.Prison_id)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}
}
