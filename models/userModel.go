package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


//User is the model that governs all notes objects retrived or inserted into the DB
type User struct {
    ID            primitive.ObjectID `bson:"_id"`
    UserName     string            `json:"user_name"`
    Password      string            `json:"Password"`
    Email         string            `json:"email"`
    Phone         string            `json:"phone"`
    Created_at    time.Time         `json:"created_at"`
    Updated_at    time.Time         `json:"updated_at"`
}

type UserSignUp struct{
    UserName     string            `json:"user_name"`
	Email         string            `json:"email"`
    Phone         string            `json:"phone"`
    Password      string            `json:"Password"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}