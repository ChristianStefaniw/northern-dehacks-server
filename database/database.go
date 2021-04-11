package database

import "go.mongodb.org/mongo-driver/mongo"

var Database *db

type db struct {
	ndh *mongo.Database
}
