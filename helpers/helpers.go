package helpers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BytesToMongoID(b []byte) (id primitive.ObjectID) {
	copy(id[:], b)
	return
}
