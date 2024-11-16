package models

type LoginRequest struct {
	Email          string `bson:"email" json:"email"`
	HashedPassword string `bson:"hashed_password" json:"hashed_password"`
}
