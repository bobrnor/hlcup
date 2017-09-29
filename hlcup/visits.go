package main

import (
	"database/sql"
	"fmt"
	"strings"
)

//easyjson:json
type Visit struct {
	Id         int64 `json:"id"`
	LocationId int64 `json:"location"`
	UserId     int64 `json:"user"`
	VisitedAt  int64 `json:"visited_at"`
	Mark       int64 `json:"mark"`
}

//easyjson:json
type VisitList struct {
	Visits []Visit `json:"visits"`
}

//easyjson:json
type VisitJson struct {
	Place     string `json:"place"`
	VisitedAt int64  `json:"visited_at"`
	Mark      int64  `json:"mark"`
}

func VisitByID(id interface{}) *Visit {
	row := db.QueryRow("SELECT * FROM visits WHERE id = ?", id)
	return scanVisitFromRow(row)
}

func FindVisits(userID interface{}, params map[string]interface{}) []VisitJson {
	sel := "SELECT locations.place, visits.visited_at, visits.mark"
	from := []string{"locations", "visits"}
	where := []string{"locations.id = visits.location", "visits.user = ?"}
	vals := []interface{}{userID}

	if fromDate, ok := params["fromDate"]; ok {
		where = append(where, "visits.visited_at > ?")
		vals = append(vals, fromDate)
	}

	if toDate, ok := params["toDate"]; ok {
		where = append(where, "visits.visited_at < ?")
		vals = append(vals, toDate)
	}

	if country, ok := params["country"]; ok {
		from = append(from, "countries")
		where = append(where, "locations.country = ?")
		vals = append(vals, country)
	}

	if toDistance, ok := params["toDistance"]; ok {
		where = append(where, "locations.distance < ?")
		vals = append(vals, toDistance)
	}

	sql := fmt.Sprintf("%s FROM %s WHERE %s ORDER BY visits.visited_at ASC",
		sel,
		strings.Join(from, ", "),
		strings.Join(where, " AND "))

	return queryVisitRows(sql, vals...)
}

func queryVisitRows(sql string, vals ...interface{}) []VisitJson {
	rows, err := db.Query(sql, vals...)
	if err != nil {
		return []VisitJson{}
	}
	defer rows.Close()

	result := []VisitJson{}
	for rows.Next() {
		visit := scanVisitFromRows(rows)
		if visit != nil {
			result = append(result, *visit)
		}
	}

	if rows.Err() != nil {
		return []VisitJson{}
	}

	return result
}

func scanVisitFromRows(rows *sql.Rows) *VisitJson {
	visit := VisitJson{}
	err := rows.Scan(
		&visit.Place,
		&visit.VisitedAt,
		&visit.Mark,
	)

	if err != nil {
		return nil
	}

	return &visit
}

func scanVisitFromRow(row *sql.Row) *Visit {
	visit := Visit{}
	err := row.Scan(
		&visit.Id,
		&visit.LocationId,
		&visit.UserId,
		&visit.VisitedAt,
		&visit.Mark,
	)

	if err != nil {
		return nil
	}

	return &visit
}

func UpdateVisit(userID interface{}, params map[string]interface{}) error {
	set := []string{}
	vals := []interface{}{}

	for k, v := range params {
		if v == nil {
			return nilInMapErr
		}

		if _, ok := v.(float64); !ok {
			return badInMapErr
		}

		set = append(set, fmt.Sprintf("%v = ?", k))
		vals = append(vals, v)
	}
	vals = append(vals, userID)

	sql := fmt.Sprintf("UPDATE visits SET %s WHERE id = ?", strings.Join(set, ", "))
	_, err := db.Exec(sql, vals...)
	return err
}

func SaveVisit(params map[string]interface{}) error {
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

		if _, ok := v.(float64); !ok {
			return badInMapErr
		}

		columns = append(columns, k)
		vals = append(vals, v)
		placeholders = append(placeholders, "?")
	}

	sql := fmt.Sprintf("INSERT INTO visits (%s) VALUES (%s)", strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	_, err := db.Exec(sql, vals...)
	return err
}
