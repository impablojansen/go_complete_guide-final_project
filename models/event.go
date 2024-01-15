package models

import (
	"database/sql"
	"example.com/final_project_-_REST_API/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event = []Event{}

func (e Event) Save() error {
	var stmt *sql.Stmt
	var err error
	var result sql.Result
	var id int64

	query := `INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err = db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err = result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	var events []Event
	query := "SELECT * FROM events"

	// Exec could be used here.
	//But normally you use Query for gets and Exec when you want to change something in the database.

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}