package cases

import (
	"database/sql"
	"fmt"
	"golang-database/models"
	"log"
)

func Case2(db *sql.DB) {
	case2 := `
	select * from criminal_records 
	order by crime_date desc limit 3`

	rows, err := db.Query(case2)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Case 2 :")
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
