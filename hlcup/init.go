package main

import (
	"context"
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"sync"

	"github.com/mailru/easyjson"
)

func initEnv() {
	initDB()
	populateDB()
}

func initDB() {
	log.Print("init db")
	if database, err := sql.Open("mysql", "root@unix(/var/run/mysqld/mysqld.sock)/hlcup?parseTime=true"); err != nil {
		panic(err)
	} else {
		db = database
	}

	db.SetMaxOpenConns(4)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(10 * time.Second)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		if err := db.PingContext(ctx); err != nil {
			cancel()
			<-time.Tick(time.Second)
			continue
		}
		cancel()
		return
	}
}

func populateDB() {
	log.Print("populate db")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := filepath.Walk("/tmp/data/", populateUsers); err != nil {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := filepath.Walk("/tmp/data/", populateLocations); err != nil {
			panic(err)
		}
	}()

	wg.Wait()

	if err := filepath.Walk("/tmp/data/", populateVisits); err != nil {
		panic(err)
	}
}

func populateUsers(path string, f os.FileInfo, err error) error {
	if f.IsDir() || !strings.Contains(path, "users") {
		return nil
	}

	log.Printf("populate users from %s", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	users := UserList{}
	if err := easyjson.Unmarshal(data, &users); err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT IGNORE INTO users (id, email, first_name, last_name, gender, birth_date) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	for i := range users.Users {
		user := &users.Users[i]
		_, err := stmt.Exec(user.Id,
			user.Email,
			user.FirstName,
			user.LastName,
			user.Gender,
			user.BirthDate)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func populateLocations(path string, f os.FileInfo, err error) error {
	if f.IsDir() || !strings.Contains(path, "locations") {
		return nil
	}

	log.Printf("populate locations from %s", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	locations := LocationList{}
	if err := easyjson.Unmarshal(data, &locations); err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT IGNORE INTO locations (id, place, country, city, distance) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	for i := range locations.Locations {
		location := &locations.Locations[i]
		_, err := stmt.Exec(location.Id,
			location.Place,
			location.Country,
			location.City,
			location.Distance)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func populateVisits(path string, f os.FileInfo, err error) error {
	if f.IsDir() || !strings.Contains(path, "visits") {
		return nil
	}

	log.Printf("populate visits from %s", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	visits := VisitList{}
	if err := easyjson.Unmarshal(data, &visits); err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT IGNORE INTO visits (id, location, user, visited_at, mark) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	for i := range visits.Visits {
		visit := &visits.Visits[i]
		_, err := stmt.Exec(visit.Id,
			visit.LocationId,
			visit.UserId,
			visit.VisitedAt,
			visit.Mark)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
