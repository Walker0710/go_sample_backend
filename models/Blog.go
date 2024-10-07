package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Blog model
type Blog struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
	Author   string             `bson:"author"`
	Comments []Comment          `bson:"comments"`
}

// Comment model
type Comment struct {
	Username string `bson:"username"`
	Content  string `bson:"content"`
}
