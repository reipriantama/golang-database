package cases

import (
	"database/sql"
	"fmt"
	"golang-database/models"
)

func Case2(db *sql.DB) {
	case2 := `
	select * from criminal_records 
	order by crime_date desc limit 3`

	rows, err := db.Query(case2)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	fmt.Println("Case 2 (Take all criminal data and sort it by date of incident from the most recent and take only 3 pieces of data) :")
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
