package repository

import (
	"context"
	"database/sql"

	models "example.com/plan-my-event/models"
	"github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mysql struct {
	DB     *sql.DB
	config mysql.Config
}

type Mongodb struct {
	DB  *mongo.Database
	uri string
	ctx *context.Context
}

type Repository interface {
	IUserRepository
	IEventRepository
}

type IUserRepository interface {
	SaveMongo(*models.User) error
	ValidateMongoCredentials(*models.User) error
	GetUserByEmail(string) models.User
}

type IEventRepository interface {
	SaveMongoEvent([]models.Event) error
	GetAllMongoEvents() ([]models.Event, error)
	GetMongoEventByID(string) (*models.Event, error)
	UpdateMongo(*models.Event) error
	DeleteMongoEvent(*models.Event) error
	MongoRegister(*models.Event, string) error
	CancelMongoRegistration(string, string) (int64, error)
}
