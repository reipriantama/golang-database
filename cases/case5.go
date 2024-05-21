package cases

import (
	"database/sql"
	"fmt"
	"golang-database/models"
	"log"
)

func Case5(db *sql.DB) {
	case5 := `
	select * 
	from criminal_records  
	where crime_description ilike '%mengambil%'`

	rows, err := db.Query(case5)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Case 5 :")
	for rows.Next() {
		var crimeInfo models.Criminal_records

		err := rows.Scan(&crimeInfo.Id, &crimeInfo.Suspect_name, &crimeInfo.Crime_description, &crimeInfo.Crime_date, &crimeInfo.Crime_status)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Id : %d, Suspect_name : %s, Crime_descripton :%s, Crime_date : %s, Crime_status : %s \n",
			crimeInfo.Id, crimeInfo.Suspect_name, crimeInfo.Crime_description, crimeInfo.Crime_date, crimeInfo.Crime_status)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
