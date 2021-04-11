package database

import (
	"context"

	"github.com/northern-dehacks/northern-dehacks-server/constants"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *db) NewsletterEmailExists(ctx context.Context, email string) (bool, error) {
	var result bson.M
	filter := bson.M{"email": email}
	collection := db.ndh.Collection("newsletter")

	if err := collection.FindOne(ctx, filter).Decode(&result); err != nil {
		if err.Error() == constants.NO_DOC_FOUND {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
