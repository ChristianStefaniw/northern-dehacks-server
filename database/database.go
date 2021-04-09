package database

import "go.mongodb.org/mongo-driver/mongo"

var Database *db

type db struct {
	Ndh *mongo.Database
}
