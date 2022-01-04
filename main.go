package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Person struct {
	Name     string
	Email    string
	Company  string
	Position string
	Location string
}

func main() {
	fmt.Println("Hello, world!")

	// Capture connection properties
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "Peoplenect",
	}

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// Testing personByPosition
	people, err := personByPosition("Student")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("People found: %v\n", people)

	// Testing personByName
	person, err := personByName("Martin Heberling")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Person found: %v\n", person)

	// Testing addPerson
	personId, err := addPerson(Person{
		Name:    "Carlos Sainz",
		Email:   "carlossainz@ferrari.com",
		Company: "Ferrari",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of new added person: %v\n", personId)
}

// Query for multiple rows
func personByPosition(position string) ([]Person, error) {
	var people []Person

	rows, err := db.Query("SELECT * FROM people WHERE position = ?", position)
	if err != nil {
		return nil, fmt.Errorf("personByPosition %q: %v", position, err)
	}

	defer rows.Close()

	for rows.Next() {
		var person Person
		if err := rows.Scan(&person.Name, &person.Email, &person.Company, &person.Position, &person.Location); err != nil {
			return nil, fmt.Errorf("personByPosition %q: %v", position, err)
		}
		people = append(people, person)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("personByPosition %q: %v", position, err)
	}

	return people, nil
}

// Query for one row
func personByName(name string) (Person, error) {
	var person Person

	row := db.QueryRow("SELECT * FROM people WHERE name = ?", name)

	if err := row.Scan(&person.Name, &person.Email, &person.Company, &person.Position, &person.Location); err != nil {
		if err == sql.ErrNoRows {
			return person, fmt.Errorf("personByName %q: no such person", name)
		}
		return person, fmt.Errorf("personByName %q: %v", name, err)
	}
	return person, nil
}

func addPerson(person Person) (int64, error) {
	result, err := db.Exec("INSERT INTO people (name, email, company) VALUES (?, ?, ?)", person.Name, person.Email, person.Company)
	if err != nil {
		return 0, fmt.Errorf("addPerson: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addPerson: %v", err)
	}
	return id, nil
}
