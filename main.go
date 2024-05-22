package main

import (
	"database/sql"
	"fmt"
	"golang-database/cases"
	"golang-database/database"

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

	for {
		var choice int
		fmt.Println("Choose :")
		fmt.Println("1. Insert")
		fmt.Println("2. Update")
		fmt.Println("3. View")
		fmt.Println("4. View All Data")
		fmt.Println("5. Delete")
		fmt.Println("6. Exit")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			database.InsertCrimeData(db)
		case 2:
			database.UpdateCrimeStatus(db)
		case 3:
			cases.Case1(db)
			cases.Case2(db)
			cases.Case3(db)
			cases.Case4(db)
			cases.Case5(db)
		case 4:
			database.ViewAllData(db)
		case 5:
			database.DeleteCrimeData(db)
		case 6:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Invalid choice")

		}
	}

	// insert data
	// database.InsertCrimeData(db)

	// step 3

	// case 1 :
	// Retrieve criminal name data whose status is 'Open' and Criminal date > 2024-03-15
	// cases.Case1(db)

	// case 2
	// Take all criminal data and sort it by date of incident from the most recent and take only 3 pieces of data
	// cases.Case2(db)

	// case 3
	// Take criminal data, group it based on 'Open' status and Count
	// cases.Case3(db)

	// case 4
	// Take criminal data, group it based on 'Close' status and Count
	// cases.Case4(db)

	// case 5
	// Take criminal data based on the crime description that contains the sentence 'Mengambil’
	// cases.Case5(db)

	// Step 4. When the suspect is found to be innocent, update the criminal data to 'Close'
	// gunakan case when
	// database.UpdateCrimeStatus(db)

	// Step 5. When the suspect is truly innocent and his status has become 'Close', then remove him from the criminal list
	// gunakan case when
	// database.DeleteCrimeData(db)

}
