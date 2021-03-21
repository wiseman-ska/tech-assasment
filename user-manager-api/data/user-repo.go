package data

import (
	"github.com/wiseman-ska/tech-assessment/user-manager-api/models"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	Col *mgo.Collection
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
func (repo *UserRepository) GetAllUsers()(u models.User)  {

	return
}

func (repo *UserRepository) GetUserById(id string) (u models.User, err error)  {
	return
}
