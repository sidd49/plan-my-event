package repository

import (
	"errors"
	"fmt"

	models "example.com/plan-my-event/models"
	"example.com/plan-my-event/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *Mongodb) GetUserByEmail(email string) models.User {
	var user models.User
	// get the user by email from the db
	m.DB.Collection("users").FindOne(*m.ctx, bson.D{{Key: "email", Value: email}}).Decode(&user)
	return user
}

func (m *Mongodb) SaveMongo(u *models.User) error {
	// save the user in the db
	_, err := m.DB.Collection("users").InsertOne(*m.ctx, u)
	return err
}

func (m *Mongodb) ValidateMongoCredentials(u *models.User) error {
	var retrievedUser models.User
	// find the user by email
	err := m.DB.Collection("users").FindOne(*m.ctx, bson.D{{Key: "email", Value: u.Email}}).Decode(&retrievedUser)
	if err != nil {
		fmt.Println(err)
		return errors.New("credentials invalid")
	}
	if !utils.CheckPasswordHash(retrievedUser.Password, u.Password) {
		return errors.New("credentials invalid")
	}
	u.ID = retrievedUser.ID
	return nil
}

// func (p *Mysql) SaveUser(u *models.User) error {
// 	query := `
// 	INSERT INTO users(email, password) VALUES (?, ?)
// 	`
// 	stmt, err := p.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	hashPassword, err := utils.HashPassword(u.Password)
// 	if err != nil {
// 		return err
// 	}

// 	result, err := stmt.Exec(u.Email, hashPassword)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = result.LastInsertId()

// 	return err
// }

// func (p *Mysql) ValidateCredentials(u *models.User) error {
// 	query := "SELECT id, password FROM users WHERE email = ?"
// 	row := p.DB.QueryRow(query, u.Email)

// 	var retrievedPassword string
// 	var retrievedID int64
// 	err := row.Scan(&retrievedID, &retrievedPassword)

// 	if err != nil {
// 		return errors.New("credentials invalid")
// 	}
// 	if !utils.CheckPasswordHash(retrievedPassword, u.Password) {
// 		return errors.New("credentials invalid")
// 	}
// 	u.ID = retrievedID
// 	return nil
// }
