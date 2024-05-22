package cases

import (
	"database/sql"
	"fmt"
	"golang-database/models"
)

func Case1(db *sql.DB) {
	case1 := `
	select * from criminal_records 
	where crime_status = 'Open' 
	and crime_date > '2024-03-15'
	`

	rows, err := db.Query(case1)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	fmt.Println("Case 1 (Retrieve criminal name data whose status is 'Open' and Criminal date > 2024-03-15) :")
	for rows.Next() {

		var crimeInfo models.Criminal_records

		err := rows.Scan(&crimeInfo.Id, &crimeInfo.Suspect_name, &crimeInfo.Crime_description, &crimeInfo.Crime_date, &crimeInfo.Crime_status)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Id : %d, Suspect_name : %s, Crime_descripton :%s, Crime_date : %s, Crime_status : %s \n",
			crimeInfo.Id, crimeInfo.Suspect_name, crimeInfo.Crime_description, crimeInfo.Crime_date, crimeInfo.Crime_status)

	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}
}
