package unit_tests

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/commons"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/controllers"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserData struct {
	Data *models.User `json:"data"`
}

func TestUserRegistration(t *testing.T) {
	expectedStatus := http.StatusCreated
	data := &UserData{
		Data: &models.User{
			FirstName:       "Wiseman",
			LastName:        "Fibonacci",
			Email:           "wiseman@mail.com",
			MobileNumber:    "0807069540",
			IdNumber:        "9412216126086",
			PhysicalAddress: "Coral Sands Eco-Estate, 95 Westlake Avenue, Sandton, 1609",
			Password:        "12345",
		},
	}
	reqBody, _ := commons.JSONMarshal(data, true)
	req, err := http.NewRequest("POST", "/users/create", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.UserRegisterHandler)
	handler.ServeHTTP(rr, req)
	/*test scenarios*/
	resp := new(controllers.UserResource)
	json.NewDecoder(rr.Body).Decode(&resp)
	assert.Equal(t, expectedStatus, rr.Code, "status code must be 201 for both")
	assert.Equal(t, data.Data.FirstName, resp.Data.FirstName, "they must be equal")
	assert.Equal(t, data.Data.LastName, resp.Data.LastName, "they must be equal")
	assert.Equal(t, data.Data.IdNumber, resp.Data.IdNumber, "they must be equal")
}

func TestUserRegistrationReqField(t *testing.T) {

}

func TestGetUsers(t *testing.T)  {

}

func TestGetUserById(t *testing.T)  {

}

func TestUpdateUser(t *testing.T)  {

}




