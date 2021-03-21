package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/commons"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/data-persistance"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		commons.DisplayAppError(w,
			err,
			"Invalid user data",
			http.StatusInternalServerError,
		)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.Close()
	userCol := context.Collection(models.UsersCollection)
	userRepo := &data_persistance.UserRepository{Col: userCol}
	_ = userRepo.CreateUser(user)
	user.HashPassword = nil
	if resp, err := json.Marshal(UserResource{Data: *user}); err != nil {
		commons.DisplayAppError(w,
			err,
			"An unexpected error has occurred",
			http.StatusInternalServerError,
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
			http.StatusInternalServerError,
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
	userRepo := &data_persistance.UserRepository{Col: userCol}
	if user, err := userRepo.Login(loginUser); err != nil {
		commons.DisplayAppError(w,
			err,
			"Invalid login credentials",
			http.StatusUnauthorized,
		)
		return
	} else {
		token, err := commons.GenerateToken(user.Email, "member")
		if err != nil {
			commons.DisplayAppError(w,
				err,
				"Error while generating the access token",
				http.StatusInternalServerError,
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
				"An unexpected error has occurred",
				http.StatusInternalServerError,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(resp)
	}
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	userCol := context.Collection(models.UsersCollection)
	userRepo := &data_persistance.UserRepository{Col: userCol}
	users := userRepo.GetAllUsers()
	resp, err := json.Marshal(UsersResource{Data: users})
	if err != nil {
		commons.DisplayAppError(w,
			err,
			"An unexpected error has occurred",
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	context := NewContext()
	defer context.Close()
	userCol := context.Collection(models.UsersCollection)
	userRepo := &data_persistance.UserRepository{Col: userCol}
	user, err := userRepo.GetUserById(&id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
			return
		}else {
			commons.DisplayAppError(w,
				err,
				"An unexpected error has occurred",
				http.StatusInternalServerError,
			)
			return
		}
	}
	resp, err := json.Marshal(user)
	if err != nil {
		commons.DisplayAppError(w,
			err,
			"An unexpected error has occurred",
			http.StatusInternalServerError,
		)
		return
	}else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}
}

func UserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var dataResource UserResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		commons.DisplayAppError(w,
			err,
			"Invalid user data",
			http.StatusInternalServerError,
		)
		return
	}
	user := &dataResource.Data
	user.Id = id
	context := NewContext()
	defer context.Close()
	userCol := context.Collection(models.UsersCollection)
	userRepo := &data_persistance.UserRepository{Col: userCol}
	if err := userRepo.UpdateUser(user); err != nil {
		commons.DisplayAppError(w,
			err,
			"An unexpected error has occurred",
			http.StatusInternalServerError,
		)
		return
	}else {
		w.WriteHeader(http.StatusOK)
	}
}

func UserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	context := NewContext()
	defer context.Close()
	userCol := context.Collection(models.UsersCollection)
	userRepo := &data_persistance.UserRepository{Col: userCol}
	if err := userRepo.DeleteUser(&id); err != nil {
		commons.DisplayAppError(w,
			err,
			"An unexpected error has occurred",
			http.StatusInternalServerError,
		)
		return
	}else {
		w.WriteHeader(http.StatusOK)
	}
}
