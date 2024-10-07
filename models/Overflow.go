package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Overflow model
type Overflow struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
	Username string             `bson:"username"`
	Comments []Comment          `bson:"comments"`
}
