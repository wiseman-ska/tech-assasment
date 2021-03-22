package data_persistance

import (
	"fmt"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/commons"
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
	var hpass []byte
	var err error
	var savedUser *models.User
	if user.Password != "" {
		hpass, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
	}else {
		err := fmt.Errorf("password not provided")
		return err
	}
	if user.IdNumber != "" {
		if commons.IsValidSAIdNumber(user.IdNumber) {
			savedUser, _ = repo.GetUserByIdNumber(user.IdNumber)
			if savedUser != nil {
				err := fmt.Errorf("user already exists in the system")
				return err
			}
		}else {
			err := fmt.Errorf("invalid SA ID number")
			return err
		}
	}else {
		err := fmt.Errorf("id number not provided")
		return err
	}
	if user.MobileNumber != "" {
		if savedUser != nil {
			if user.MobileNumber == savedUser.MobileNumber {
				err := fmt.Errorf("mobile number already exists in the system")
				return err
			}
		}
	}
	if user.FirstName == "" {
		err := fmt.Errorf("first name not provided")
		return err
	}
	if user.LastName == "" {
		err := fmt.Errorf("last name not provided")
		return err
	}
	user.HashPassword = hpass
	user.Password = ""
	err = repo.Col.Insert(&user)
	return err
}

func (repo *UserRepository) GetAllUsers() []*models.User  {
	users := make([]*models.User, 0)
	iter := repo.Col.Find(bson.M{}).Iter()
	result := new(models.User)
	for iter.Next(result) {
		users = append(users, result)
		result = new(models.User)
	}
	return users
}

func (repo *UserRepository) GetUserById(id *bson.ObjectId) (u models.User, err error)  {
	err = repo.Col.FindId(&id).One(&u)
	return
}

func (repo *UserRepository) GetUserByIdNumber(idNumber string) (u *models.User, err error)  {
	err = repo.Col.Find(bson.M{"idNumber": idNumber}).One(&u)
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
		})
	return err
}

func (repo *UserRepository) DeleteUser(id *bson.ObjectId) error {
	err := repo.Col.Remove(bson.M{"_id": id})
	return err
}
