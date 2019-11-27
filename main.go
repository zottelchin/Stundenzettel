package main

import (
	"database/sql"
	"strconv"

	"codeberg.org/momar/logg"

	parcelServe "github.com/moqmar/parcel-serve"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var database = initDB()

func main() {
	r := gin.Default()
	r.GET("/api/v1/:year/:month/:day", tagAbrufen)
	r.POST("/api/v1/:year/:month/:day", tagSpeichern)
	r.DELETE("/api/v1/:id", deleteByID)
	parcelServe.Serve("frontend", r, AssetNames(), MustAsset)
	r.Run(":8899")
}

type Schicht struct {
	ID     int
	Beginn string
	Ende   string
	Pause  string
}

func tagSpeichern(c *gin.Context) {
	year, _ := strconv.Atoi(c.Param("year"))
	month, _ := strconv.Atoi(c.Param("month"))
	day, _ := strconv.Atoi(c.Param("day"))
	s := Schicht{}
	c.BindJSON(&s)
	err := save(year, month, day, s)
	if err != nil {
		c.String(500, "%s", err)
		return
	}
	c.Status(201)
}

func tagAbrufen(c *gin.Context) {
	year, _ := strconv.Atoi(c.Param("year"))
	month, _ := strconv.Atoi(c.Param("month"))
	day, _ := strconv.Atoi(c.Param("day"))
	result, err := load(year, month, day)
	if err != nil {
		logg.Error("%s", err)
	}
	c.JSON(200, result)
}

func save(jahr int, monat int, tag int, v Schicht) error {
	_, err := database.Exec("Insert into Schichten (Tag, Monat, Jahr, Beginn, Ende, Pause) Values (?, ?, ?, ?, ?, ?)", tag, monat, jahr, v.Beginn, v.Ende, v.Pause)
	if err != nil {
		return err
	}
	return nil
}

func load(jahr int, monat int, tag int) ([]Schicht, error) {
	res := []Schicht{}
	rows, err := database.Query("Select id, Beginn, Ende, Pause From Schichten Where Jahr = ? And Monat = ? And Tag = ?", jahr, monat, tag)
	if err != nil {
		return []Schicht{}, err
	}
	for rows.Next() {
		tmp := Schicht{}
		err = rows.Scan(&tmp.ID, &tmp.Beginn, &tmp.Ende, &tmp.Pause)
		if err != nil {
			logg.Error("%s", err)
		}
		res = append(res, tmp)
	}
	return res, nil
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/schichten.db")
	if err != nil {
		logg.Error("%s", err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Schichten (" +
		"id INTEGER PRIMARY KEY AUTOINCREMENT," +
		"Beginn TEXT," +
		"Ende TEXT," +
		"Pause TEXT," +
		"Jahr Text," +
		"Monat Text," +
		"Tag Text);")
	if err != nil {
		logg.Error("%s", err)
	}
	return db
}

func deleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err !=  nil || id < 0 {
		c.Status(500)
		return
	}
	_, err = database.Exec("Delete From Schichten where id = ?", id)
	if err != nil {
		c.Status(500)
		return
	}
	c.Status(200)
}