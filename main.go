package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "belajar"
)

type Criminal_records struct {
	id                int
	suspect_name      string
	crime_description string
	crime_date        string
	crime_status      string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Connected!")

	// step 3

	//case1
	query1 := `
	select * from criminal_records 
	where crime_status = 'Open' 
	and crime_date > '2024-03-15'
	`

	rows, err := db.Query(query1)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Case 1 :")
	for rows.Next() {

		var crime_data = Criminal_records{}

		err := rows.Scan(&crime_data.id, &crime_data.suspect_name, &crime_data.crime_description, &crime_data.crime_date, &crime_data.crime_status)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Id : %d, suspect_name : %s, crime_descripton :%s, crime_date : %s, crime_status : %s \n",
			crime_data.id, crime_data.suspect_name, crime_data.crime_description, crime_data.crime_date, crime_data.crime_status)

	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// case 2

	// case 3
	// case 4
	// case 5

	// Step 4. When the suspect is found to be innocent, update the criminal data to 'Close'
	// gunakan case when

	// Step 5. When the suspect is truly innocent and his status has become 'Close', then remove him from the criminal list
	// gunakan case when

}
