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
	crime_date
	 )
	values 	('Mirza', 'Mengambil Botol Minum', '2024-04-29 15:05:36.208328'),
			('Epul', 'Mengambil Hati Wanita Tanpa Bilang Bilang', '2024-03-15 16:05:36.208328' ),
			('Herdi', 'Melakukan pencucian uang', '2024-06-20 10:05:36.208328'),
			('Atthar', 'Melakukan kecurangan saat bermain FIFA', '2024-07-25 17:05:36.208328'),
			('Ahmad', 'Melakukan Pelanggaran dengan memburu hewan yang di lindungi saat mendaki gunung (Perpu Kementrian Kehutanan)', '2024-04-19 15:05:36.208328'),
			('Asep', 'Kasus penggelapan dana bansos', '2024-04-29 15:05:36.208328'),
			('Udin', 'Kasus penggelapan pajak', '2024-11-15 16:05:36.208328'),
			('Bejo', 'Kasus KDRT', '2024-10-20 10:05:36.208328'),
			('Antony', 'Kasus Tawuran', '2024-04-25 17:05:36.208328'),
			('James', 'Kasus pencurian', '2024-01-19 15:05:36.208328'),
			('Jhon', 'Kasus Penjambretan', '2024-01-29 15:05:36.208328'),
			('Robert', 'Kasus begal', '2024-02-15 16:05:36.208328'),
			('Tukul', 'Komplotan judi online', '2024-09-20 10:05:36.208328'),
			('Bobi', 'Melakukan pembulian', '2024-08-25 17:05:36.208328'),
			('Men', 'Melakukan pencabulan', '2024-07-19 15:05:36.208328');
	`

	_, err := db.Exec(insertQuery)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Insert successful")
}
