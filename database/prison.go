package database

import (
	"database/sql"
	"golang-database/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreatePrison(c echo.Context) error {
	db := Connect()
	defer db.Close()

	var prison models.Prison
	if err := c.Bind(&prison); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	sqlStatement := `INSERT INTO prisons (name, location) VALUES ($1, $2)`
	_, err := db.Exec(sqlStatement, prison.Name, prison.Location)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, prison)
}

func GetPrisons(c echo.Context) error {
	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM prisons")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	var prisons []models.Prison
	for rows.Next() {
		var prison models.Prison
		err := rows.Scan(&prison.ID, &prison.Name, &prison.Location)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		prisons = append(prisons, prison)
	}

	if len(prisons) == 0 {
		return c.JSON(http.StatusOK, "You don't have any data")
	}

	return c.JSON(http.StatusOK, prisons)
}

func GetPrison(c echo.Context) error {
	db := Connect()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var prison models.Prison
	sqlStatement := `SELECT * FROM prisons WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err = row.Scan(&prison.ID, &prison.Name, &prison.Location)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, "Prison not found")
		} else {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, prison)
}

func UpdatePrison(c echo.Context) error {
	db := Connect()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var prison models.Prison
	if err := c.Bind(&prison); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	sqlStatement := `UPDATE prisons SET name=$1, location=$2 WHERE id=$3`
	_, err = db.Exec(sqlStatement, prison.Name, prison.Location, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, prison)
}

func DeletePrison(c echo.Context) error {
	db := Connect()
	defer db.Close()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	sqlStatement := `DELETE FROM prisons WHERE id=$1`
	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
