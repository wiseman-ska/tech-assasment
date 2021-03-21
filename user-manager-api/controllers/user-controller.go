package controllers

import (
	"encoding/json"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/commons"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/data"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/models"
	"net/http"
)

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		commons.DisplayAppError(w,
			err,
			"Invalid user data",
			500,
		)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	userCol := context.Collection(models.UsersCollection)
	userRepo := &data.UserRepository{Col: userCol}
	_ = userRepo.CreateUser(user)
	user.HashPassword = nil
	if resp, err := json.Marshal(UserResource{Data: *user}); err != nil {
		commons.DisplayAppError(w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(resp)
	}

}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		commons.DisplayAppError(w,
			err,
			"Invalid login data",
			500,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	userCol := context.Collection(models.UsersCollection)
	userRepo := &data.UserRepository{Col: userCol}
	if user, err := userRepo.Login(loginUser); err != nil {
		commons.DisplayAppError(w,
			err,
			"Invalid login credentials",
			401,
		)
		return
	} else {
		token, err := commons.GenerateToken(user.Email, "member")
		if err != nil {
			commons.DisplayAppError(w,
				err,
				"Error while generating the access token",
				500,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		user.HashPassword = nil
		authUser := AuthUserModel{
			User:  user,
			Token: token,
		}
		resp, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			commons.DisplayAppError(w,
				err,
				"An unexpected error has occured",
				500,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}

}

func UserUpdateHandler(w http.ResponseWriter, r *http.Request) {

}

func UserDeleteHandler(w http.ResponseWriter, r *http.Request) {

}

func UserGetUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func UserGetUserHandler(w http.ResponseWriter, r *http.Request) {

}
