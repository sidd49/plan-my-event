package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func New() *Mysql {
	return &Mysql{
		config: mysql.Config{
			User:   "root",
			Passwd: "root",
			Net:    "tcp",
			Addr:   "127.0.0.1:3306",
			DBName: "eventsApi",
		},
	}
}

func (p *Mysql) Init() {

	dbLoc, err := sql.Open("mysql", p.config.FormatDSN())
	p.DB = dbLoc
	if err != nil {
		panic("Could not connect to the database" + err.Error())
	}

	pingErr := p.DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Database connected")

	CreateTables(p)
}

func CreateTables(p *Mysql) {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	)
	`
	_, err := p.DB.Exec(createUsersTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create a table in DB")
	}
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		description VARCHAR(255) NOT NULL,
		location VARCHAR(255) NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = p.DB.Exec(createEventsTable)

	if err != nil {
		fmt.Println(err)
		panic("Could not create a table in DB")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = p.DB.Exec(createRegistrationsTable)
	if err != nil {
		fmt.Println(err)
		panic("Could not create registration table")
	}

}
