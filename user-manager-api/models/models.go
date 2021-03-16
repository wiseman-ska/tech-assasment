package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	UsersCollection = "users"
)

type User struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName       string             `bson:"firstName" json:"firstName"`
	LastName        string             `bson:"LastName" json:"lastName"`
	Email           string             `bson:"email" json:"email"`
	MobileNumber    string             `bson:"mobileNumber" json:"mobileNumber"`
	IdNumber        string             `bson:"idNumber" json:"idNumber"`
	PhysicalAddress string             `bson:"physicalAddress" json:"physicalAddress"`
	Password        string             `bson:"password,omitempty" json:"password"`
	HashPassword    []byte             `bson:"hashPassword,omitempty" json:"hashPassword"`
}