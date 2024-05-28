package database

import (
	"database/sql"
	"golang-database/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateCriminalRecord(c echo.Context) error {
	db := Connect()
	defer db.Close()

	var record models.CriminalRecord
	if err := c.Bind(&record); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	sqlStatement := `INSERT INTO criminal_records (suspect_name, crime_description, crime_date, crime_status, prison_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, record.SuspectName, record.CrimeDescription, record.CrimeDate, record.CrimeStatus, record.PrisonID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, record)
}

func GetCriminalRecords(c echo.Context) error {
	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM criminal_records")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var records []models.CriminalRecord
	for rows.Next() {
		var record models.CriminalRecord
		err := rows.Scan(&record.ID, &record.SuspectName, &record.CrimeDescription, &record.CrimeDate, &record.CrimeStatus, &record.PrisonID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		records = append(records, record)
	}

	if len(records) == 0 {
		return c.JSON(http.StatusOK, "You don't have any data")
	}

	return c.JSON(http.StatusOK, records)
}

func GetCriminalRecord(c echo.Context) error {
	db := Connect()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var record models.CriminalRecord
	sqlStatement := `SELECT * FROM criminal_records WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err = row.Scan(&record.ID, &record.SuspectName, &record.CrimeDescription, &record.CrimeDate, &record.CrimeStatus, &record.PrisonID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "Record not found")
		} else {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, record)
}

func UpdateCriminalRecord(c echo.Context) error {
	db := Connect()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var record models.CriminalRecord
	if err := c.Bind(&record); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	sqlStatement := `UPDATE criminal_records SET suspect_name=$1, crime_description=$2, crime_date=$3, crime_status=$4, prison_id=$5 WHERE id=$6`
	_, err = db.Exec(sqlStatement, record.SuspectName, record.CrimeDescription, record.CrimeDate, record.CrimeStatus, record.PrisonID, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, record)
}

func DeleteCriminalRecord(c echo.Context) error {
	db := Connect()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	sqlStatement := `DELETE FROM criminal_records WHERE id=$1`
	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
