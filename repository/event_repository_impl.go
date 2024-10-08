package repository

import (
	"log"

	models "example.com/plan-my-event/models"
	"example.com/plan-my-event/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *Mongodb) SaveMongoEvent(events []models.Event) error {
	// convert the object id generated by mongo driver to string
	var insertEvents []interface{}
	for _, e := range events {
		eventID, err := utils.ConvertIDtoString(primitive.NewObjectID().String())
		if err != nil {
			log.Println("Error in converting eventID in proper format : " + err.Error())
			return err
		}
		e.ID = eventID
		insertEvents = append(insertEvents, e)
	}
	// save event to database
	_, err := m.DB.Collection("events").InsertMany(*m.ctx, insertEvents)

	return err
}

func (m *Mongodb) GetAllMongoEvents() ([]models.Event, error) {
	var events []models.Event
	// fetch all the events from the database
	cur, err := m.DB.Collection("events").Find(*m.ctx, bson.D{{}}, nil)
	if err != nil {
		log.Println("Error in fetching events from db : " + err.Error())
		return nil, err
	}
	// convert the data received from mongo to the model
	for cur.Next(*m.ctx) {
		var event models.Event
		err = cur.Decode(&event)
		if err != nil {
			log.Println("Error in decoding event from db : " + err.Error())
			return nil, err
		}
		events = append(events, event)

	}

	err = cur.Close(*m.ctx)
	return events, err
}

func (m *Mongodb) GetMongoEventByID(id string) (*models.Event, error) {
	var event models.Event
	// fetch the event with the id from db
	err := m.DB.Collection("events").FindOne(*m.ctx, bson.D{{Key: "id", Value: id}}).Decode(&event)
	if err != nil {
		log.Println("Error in getting the event with id : " + id + " : " + err.Error())
		return nil, err
	}
	return &event, nil
}

func (m *Mongodb) UpdateMongo(e *models.Event) error {
	// update the event in the db
	result := m.DB.Collection("events").FindOneAndReplace(*m.ctx, bson.D{{Key: "id", Value: e.ID}}, e)
	if result.Err() != nil {
		log.Println("Error in updating the event : " + e.ID + " : " + result.Err().Error())
		return result.Err()
	}
	return nil
}

func (m *Mongodb) DeleteMongoEvent(e *models.Event) error {
	// delete the event from the db
	_, err := m.DB.Collection("events").DeleteOne(*m.ctx, bson.D{{Key: "id", Value: e.ID}})
	return err
}

func (m *Mongodb) MongoRegister(e *models.Event, userID string) error {
	var register models.Registrations
	register.Event = *e
	register.UserID = userID
	// register the user for the event in the db
	_, err := m.DB.Collection("registrations").InsertOne(*m.ctx, register)
	return err
}

func (m *Mongodb) CancelMongoRegistration(eventID string, userID string) (int64, error) {
	// cancel the event registration of the user from db
	result, err := m.DB.Collection("registrations").DeleteOne(*m.ctx, bson.D{{Key: "event.id", Value: eventID}, {Key: "userid", Value: userID}})
	return result.DeletedCount, err
}

// func (p *Mysql) SaveEvent(e *models.Event) error {
// 	query := `
// 	INSERT INTO events(name, description, location, dateTime, user_id)
// 	VALUES(?, ?, ?, ?, ?)
// 	`
// 	stmt, err := p.DB.Prepare(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
// 	if err != nil {
// 		return err
// 	}
// 	e.ID, err = result.LastInsertId()

// 	return err
// }

// func (p *Mysql) GetAllEvents() ([]models.Event, error) {
// 	query := "SELECT * FROM events"
// 	rows, err := p.DB.Query(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var events []models.Event

// 	for rows.Next() {
// 		var event models.Event
// 		rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
// 		events = append(events, event)
// 	}
// 	return events, nil
// }

// func (p *Mysql) GetEventByID(id int64) (*models.Event, error) {
// 	query := "SELECT * FROM events where id =?"
// 	row := p.DB.QueryRow(query, id)
// 	var event models.Event
// 	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &event, nil
// }

// func (p *Mysql) Update(e *models.Event) error {
// 	query := `
// 	UPDATE events
// 	SET name = ?, description = ?, location = ?, datetime = ?
// 	WHERE ID = ?
// 	`
// 	stmt, err := p.DB.Prepare(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
// 	return err
// }

// func (p *Mysql) Delete(e *models.Event) error {
// 	query := `
// 	DELETE FROM events
// 	WHERE ID = ?
// 	`
// 	stmt, err := p.DB.Prepare(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec(e.ID)
// 	return err
// }

// func (p *Mysql) Register(e *models.Event, userID int64) error {
// 	query := "INSERT INTO registrations(event_id, user_id) VALUES (?,?)"

// 	stmt, err := p.DB.Prepare(query)
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(e.ID, userID)
// 	return err
// }

// func (p *Mysql) CancelRegistration(e *models.Event, userID int64) error {
// 	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

// 	stmt, err := p.DB.Prepare(query)
// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(e.ID, userID)
// 	return err
// }
