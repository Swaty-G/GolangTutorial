package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` //primitive.ObjectID is a unique identifier for a document in a MongoDB database; _id means the field is of type primitive.ObjectID and it is named _id in the database; omitempty is a tag to omit the field if it is empty; bson is a tag to specify the field name in the database; primitive is a package that provides a type to represent a unique identifier for a document in a MongoDB database
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
