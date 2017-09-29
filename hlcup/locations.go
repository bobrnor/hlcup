package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

//easyjson:json
type Location struct {
	Id       int64  `json:"id"`
	Place    string `json:"place"`
	Country  string `json:"country"`
	City     string `json:"city"`
	Distance int64  `json:"distance"`
}

//easyjson:json
type LocationList struct {
	Locations []Location `json:"locations"`
}

func LocationByID(id interface{}) *Location {
	row := db.QueryRow("SELECT * FROM locations WHERE id = ?", id)
	return scanLocationFromRow(row)
}

func LocationAvg(id interface{}, params map[string]interface{}) float64 {
	visitSQL := "SELECT ROUND(AVG(mark), 5) FROM visits WHERE location = ?"
	vals := []interface{}{id}

	if fromDate, ok := params["fromDate"]; ok {
		visitSQL += " AND visited_at > ?"
		vals = append(vals, fromDate)
	}

	if toDate, ok := params["toDate"]; ok {
		visitSQL += " AND visited_at < ?"
		vals = append(vals, toDate)
	}

	userSQL := "SELECT id FROM users"
	useUserSQL := false
	if fromAge, ok := params["fromAge"].(int64); ok {
		userSQL += " WHERE birth_date < ?"
		useUserSQL = true

		now := time.Now().UTC()
		minBirthday := time.Date(now.Year()-int(fromAge), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.UTC)
		vals = append(vals, minBirthday.Unix())
	}

	if toAge, ok := params["toAge"].(int64); ok {
		if useUserSQL {
			userSQL += " AND birth_date > ?"
		} else {
			useUserSQL = true
			userSQL += " WHERE birth_date > ?"
		}

		now := time.Now().UTC()
		maxBirthday := time.Date(now.Year()-int(toAge), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.UTC)
		vals = append(vals, maxBirthday.Unix())
	}

	if gender, ok := params["gender"]; ok {
		if useUserSQL {
			userSQL += " AND gender = ?"
		} else {
			useUserSQL = true
			userSQL += " WHERE gender = ?"
		}
		vals = append(vals, gender)
	}

	if useUserSQL {
		visitSQL += " AND user IN (" + userSQL + ")"
	}

	var avg *float64
	row := db.QueryRow(visitSQL, vals...)
	err := row.Scan(&avg)
	if err != nil || avg == nil {
		return float64(0.0)
	}
	return *avg
}

func scanLocationFromRow(row *sql.Row) *Location {
	location := Location{}
	err := row.Scan(
		&location.Id,
		&location.Place,
		&location.Country,
		&location.City,
		&location.Distance,
	)

	if err != nil {
		log.Println(err)
		return nil
	}

	return &location
}

func UpdateLocation(locationID interface{}, params map[string]interface{}) error {
	set := []string{}
	vals := []interface{}{}

	for k, v := range params {
		if v == nil {
			return nilInMapErr
		}

		if k == "distance" {
			if _, ok := v.(float64); !ok {
				return badInMapErr
			}
		} else {
			if _, ok := v.(string); !ok {
				return badInMapErr
			}
		}

		set = append(set, fmt.Sprintf("%v = ?", k))
		vals = append(vals, v)
	}
	vals = append(vals, locationID)

	sql := fmt.Sprintf("UPDATE locations SET %s WHERE id = ?", strings.Join(set, ", "))
	_, err := db.Exec(sql, vals...)
	return err
}

func SaveLocation(params map[string]interface{}) error {
	columns := []string{}
	vals := []interface{}{}
	placeholders := []string{}

	if len(params) != 5 {
		return notEnoughErr
	}

	for k, v := range params {
		if v == nil {
			return nilInMapErr
		}

		if k == "distance" || k == "id" {
			if _, ok := v.(float64); !ok {
				return badInMapErr
			}
		} else {
			if _, ok := v.(string); !ok {
				return badInMapErr
			}
		}

		columns = append(columns, k)
		vals = append(vals, v)
		placeholders = append(placeholders, "?")
	}

	sql := fmt.Sprintf("INSERT INTO locations (%s) VALUES (%s)", strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	_, err := db.Exec(sql, vals...)
	return err
}
