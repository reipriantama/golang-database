package cases

import (
	"database/sql"
	"fmt"
)

func Case3(db *sql.DB) {
	case3 := `
	select count(*) 
	from criminal_records 
	where crime_status in ('Open')`

	var count int
	err := db.QueryRow(case3).Scan(&count)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Case 3 (Take criminal data, group it based on 'Open' status and Count) :")
	fmt.Printf("Total open criminal records: %d\n", count)
}
