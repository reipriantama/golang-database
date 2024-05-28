package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type CatatanKriminal struct {
	ID                 int    `json:"id"`
	NamaTersangka      string `json:"nama_tersangka"`
	DeskripsiKejahatan string `json:"deskripsi_kejahatan"`
	TanggalKejahatan   string `json:"tanggal_kejahatan"`
	StatusKejahatan    string `json:"status_kejahatan"`
	IDPenjara          int    `json:"id_penjara"`
}

type Penjara struct {
	ID     int    `json:"id"`
	Kelas  string `json:"kelas"`
	Lokasi string `json:"lokasi"`
}

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "crime_data"
)

func initDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d password=%s user=%s dbname=%s sslmode=disable",
		host, port, password, user, dbname)

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
}

func buatCatatanKriminal(c echo.Context) error {
	catatan := new(CatatanKriminal)
	if err := c.Bind(catatan); err != nil {
		return err
	}

	query := "INSERT INTO criminal_records (suspect_name, crime_description, crime_date, crime_status, prison_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	err := db.QueryRow(query, catatan.NamaTersangka, catatan.DeskripsiKejahatan, catatan.TanggalKejahatan, catatan.StatusKejahatan, catatan.IDPenjara).Scan(&catatan.ID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, catatan)
}

func dapatkanCatatanKriminal(c echo.Context) error {
	query := "SELECT * FROM criminal_records"
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	catatans := make([]CatatanKriminal, 0)
	for rows.Next() {
		var catatan CatatanKriminal
		err := rows.Scan(&catatan.ID, &catatan.NamaTersangka, &catatan.DeskripsiKejahatan, &catatan.TanggalKejahatan, &catatan.StatusKejahatan, &catatan.IDPenjara)
		if err != nil {
			return err
		}
		catatans = append(catatans, catatan)
	}
	return c.JSON(http.StatusOK, catatans)
}

func dapatkanSatuCatatanKriminal(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	query := "SELECT * FROM criminal_records WHERE id=$1"
	row := db.QueryRow(query, id)

	var catatan CatatanKriminal
	err := row.Scan(&catatan.ID, &catatan.NamaTersangka, &catatan.DeskripsiKejahatan, &catatan.TanggalKejahatan, &catatan.StatusKejahatan, &catatan.IDPenjara)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, catatan)
}

func updateCriminalRecord(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	catatan := new(CatatanKriminal)
	if err := c.Bind(catatan); err != nil {
		return err
	}

	query := "UPDATE criminal_records SET suspect_name=$1, crime_description=$2, crime_date=$3, crime_status=$4, prison_id=$5 WHERE id=$6"
	_, err := db.Exec(query, catatan.NamaTersangka, catatan.DeskripsiKejahatan, catatan.TanggalKejahatan, catatan.StatusKejahatan, catatan.IDPenjara, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, catatan)
}

func deleteCriminalRecord(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	query := "DELETE FROM criminal_records WHERE id=$1"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

// Implementasi CRUD untuk Penjara juga disini...

func main() {
	initDB()
	e := echo.New()
	e.POST("/catatan_kriminal", buatCatatanKriminal)
	e.GET("/catatan_kriminal", dapatkanCatatanKriminal)
	e.GET("/catatan_kriminal/:id", dapatkanSatuCatatanKriminal)
	e.PUT("/catatan_kriminal/:id", updateCriminalRecord)
	e.DELETE("/catatan_kriminal/:id", deleteCriminalRecord)

	e.Logger.Fatal(e.Start(":8080"))
}
