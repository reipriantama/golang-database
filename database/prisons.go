package database

import (
	"database/sql"
	"fmt"
)

// InsertPrison menambahkan data penjara baru ke dalam database.
func InsertPrison(db *sql.DB) {
	var class, location string
	fmt.Print("Masukkan kelas penjara: ")
	fmt.Scan(&class)
	fmt.Print("Masukkan lokasi penjara: ")
	fmt.Scan(&location)

	query := `INSERT INTO prisons (class, location) VALUES ($1, $2)`
	_, err := db.Exec(query, class, location)
	if err != nil {
		fmt.Println("Kesalahan saat menambahkan data penjara:", err)
		return
	}
	fmt.Println("Data penjara berhasil ditambahkan")
}

// ViewAllPrisons mengambil dan menampilkan semua data penjara dari database.
func ViewAllPrisons(db *sql.DB) {
	query := `SELECT id, class, location FROM prisons`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Kesalahan saat mengambil data penjara:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Data Penjara:")
	for rows.Next() {
		var id int
		var class, location string
		err := rows.Scan(&id, &class, &location)
		if err != nil {
			fmt.Println("Kesalahan saat membaca data:", err)
			continue
		}
		fmt.Printf("ID: %d, Kelas: %s, Lokasi: %s\n", id, class, location)
	}
}
