package cases

import (
	"database/sql"
	"fmt"
)

func Case4(db *sql.DB) {
	case3 := `
	select count(*) 
	from criminal_records 
	where crime_status in ('Close')`

	var count int
	err := db.QueryRow(case3).Scan(&count)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Case 4 (Take criminal data, group it based on 'Close' status and Count) :")
	fmt.Printf("Total close criminal records: %d\n", count)
}
