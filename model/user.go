package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// User User
type User struct {
	ObjectID primitive.ObjectID `bson:"_id"`
	UserName string
	UserFace string
}

// UserCollection 获取集合
func UserCollection() *mongo.Collection {
	return Mongo().Database("blog").Collection("user")
}
