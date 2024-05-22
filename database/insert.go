package database

import (
	"database/sql"
	"fmt"
)

func InsertCrimeData(db *sql.DB) {
	insertQuery := `
	insert into
	criminal_records (suspect_name,
	crime_description,
	crime_date )
	values 	('Mirza', 'Mengambil Botol Minum', '2024-04-29 15:05:36.208328'),
			('Epul', 'Mengambil Hati Wanita Tanpa Bilang Bilang', '2024-03-15 16:05:36.208328' ),
			('Herdi', 'Melakukan pencucian uang', '2024-06-20 10:05:36.208328'),
			('Atthar', 'Melakukan kecurangan saat bermain FIFA', '2024-07-25 17:05:36.208328'),
			('Ahmad', 'Melakukan Pelanggaran dengan memburu hewan yang di lindungi saat mendaki gunung (Perpu Kementrian Kehutanan)', '2024-04-19 15:05:36.208328');
	`

	_, err := db.Exec(insertQuery)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Insert successful")
}
