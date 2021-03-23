package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	UsersCollection = "users"
)

type User struct {
	Id              bson.ObjectId `bson:"_id,omitempty" json:"id"`
	FirstName       string        `bson:"firstName" json:"firstName"`
	LastName        string        `bson:"lastName" json:"lastName"`
	Email           string        `bson:"email" json:"email"`
	MobileNumber    string        `bson:"mobileNumber" json:"mobileNumber"`
	IdNumber        string        `bson:"idNumber" json:"idNumber"`
	PhysicalAddress string        `bson:"physicalAddress" json:"physicalAddress"`
	Password        string        `bson:"password,omitempty" json:"password"`
	HashPassword    []byte        `bson:"hashPassword,omitempty" json:"hashPassword"`
}
