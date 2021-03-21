package data_persistance

import (
	"github.com/wiseman-ska/tech-assessment/user-manager-api/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	Col *mgo.Collection
}

func (repo *UserRepository) Login(user models.User) (u *models.User, err error) {
	err = repo.Col.Find(bson.M{"email": user.Email}).One(&u)
	if err != nil {
		return
	}
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = &models.User{}
	}
	return
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass
	user.Password = ""
	err = repo.Col.Insert(&user)
	return err
}

func (repo *UserRepository) GetAllUsers() []models.User  {
	var users []models.User
	iter := repo.Col.Find(nil).Iter()
	result := models.User{}
	for iter.Next(&result) {
		users = append(users, result)
	}
	return users
}

func (repo *UserRepository) GetUserById(id *bson.ObjectId) (u models.User, err error)  {
	err = repo.Col.FindId(&id).One(&u)
	return
}

func (repo *UserRepository) UpdateUser(user *models.User) error  {
	err := repo.Col.Update(bson.M{"_id": user.Id},
		bson.M{
			"firstName": user.FirstName,
			"lastName": user.LastName,
			"email": user.Email,
			"mobileNumber": user.MobileNumber,
			"idNumber": user.MobileNumber,
			"physicalAddress": user.PhysicalAddress,
			"password": user.Password,

		})
	return err
}

func (repo *UserRepository) DeleteUser(id *bson.ObjectId) error {
	err := repo.Col.Remove(bson.M{"_id": id})
	return err
}
