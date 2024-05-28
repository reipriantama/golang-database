package main

import (
	"database/sql"
	"fmt"
	"golang-database/cases"
	"golang-database/database"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "crime_data"
)

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

	fmt.Println("Berhasil Terhubung!")

	for {
		var pilihan int
		fmt.Println("Pilih :")
		fmt.Println("1. Tambah Data Kejahatan")
		fmt.Println("2. Perbarui Status Kejahatan")
		fmt.Println("3. Lihat Kasus")
		fmt.Println("4. Lihat Semua Data Kejahatan")
		fmt.Println("5. Hapus Data Kejahatan")
		fmt.Println("6. Tambah Data Penjara")
		fmt.Println("7. Lihat Semua Penjara")
		fmt.Println("8. Keluar")
		fmt.Scan(&pilihan)

		switch pilihan {
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
			database.InsertPrison(db)
		case 7:
			database.ViewAllPrisons(db)
		case 8:
			fmt.Println("Keluar")
			return
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}
