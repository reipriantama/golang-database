package main

import (
	"database/sql"
	"fmt"

	"golang-database/cases"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "plabs.id"
	// password = "postgres"
	dbname = "crime_data"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

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

	// case 1 :
	// Retrieve criminal name data whose status is 'Open' and Criminal date > 2024-03-15
	cases.Case1(db)

	// case 2
	// Take all criminal data and sort it by date of incident from the most recent and take only 3 pieces of data
	cases.Case2(db)

	// case 3
	// Take criminal data, group it based on 'Open' status and Count
	cases.Case3(db)

	// case 4
	// Take criminal data, group it based on 'Close' status and Count
	cases.Case4(db)

	// case 5
	// Take criminal data based on the crime description that contains the sentence 'Mengambil’
	cases.Case5(db)

	// Step 4. When the suspect is found to be innocent, update the criminal data to 'Close'
	// gunakan case when
	// database.UpdateCrimeStatus(db)

	// Step 5. When the suspect is truly innocent and his status has become 'Close', then remove him from the criminal list
	// gunakan case when

}
