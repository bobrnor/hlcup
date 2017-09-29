package main

import (
	"database/sql"
	"fmt"
	"strings"
)

var (
	nilInMapErr  = fmt.Errorf("nil in map")
	badInMapErr  = fmt.Errorf("bad in map")
	notEnoughErr = fmt.Errorf("not enough in map")
)

//easyjson:json
type User struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	BirthDate int64  `json:"birth_date"`
}

//easyjson:json
type UserList struct {
	Users []User `json:"users"`
}

func UserByID(id interface{}) *User {
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	return scanUserFromRow(row)
}

func UpdateUser(userID interface{}, params map[string]interface{}) error {
	set := []string{}
	vals := []interface{}{}

	for k, v := range params {
		if v == nil {
			return nilInMapErr
		}

		if k == "birth_date" {
			if _, ok := v.(float64); !ok {
				return badInMapErr
			}
		} else {
			if _, ok := v.(string); !ok {
				return badInMapErr
			} else if k == "gender" && v != "m" && v != "f" {
				return badInMapErr
			}
		}

		set = append(set, fmt.Sprintf("%v = ?", k))
		vals = append(vals, v)
	}
	vals = append(vals, userID)

	sql := fmt.Sprintf("UPDATE users SET %s WHERE id = ?", strings.Join(set, ", "))
	_, err := db.Exec(sql, vals...)
	return err
}

func SaveUser(params map[string]interface{}) error {
	columns := []string{}
	vals := []interface{}{}
	placeholders := []string{}

	if len(params) != 6 {
		return notEnoughErr
	}

	for k, v := range params {
		if v == nil {
			return nilInMapErr
		}

		if k == "birth_date" || k == "id" {
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

	sql := fmt.Sprintf("INSERT INTO users (%s) VALUES (%s)", strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	_, err := db.Exec(sql, vals...)
	return err
}

func scanUserFromRow(row *sql.Row) *User {
	user := User{}
	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Gender,
		&user.BirthDate,
	)

	if err != nil {
		return nil
	}
	return &user
}
