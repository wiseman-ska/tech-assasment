package routers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/commons"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/controllers"
)

func SetupUserRoutes(router *mux.Router) *mux.Router {
	userRouter := mux.NewRouter()
	userRouter.HandleFunc("/users/login", controllers.UserLoginHandler).Methods("POST")
	userRouter.HandleFunc("/users/create", controllers.UserRegisterHandler).Methods("POST")
	userRouter.HandleFunc("/api/v1/users/update/{id}", controllers.UserRegisterHandler).Methods("POST")
	userRouter.HandleFunc("/api/v1/users/delete/{id}", controllers.UserRegisterHandler).Methods("DELETE")
	userRouter.HandleFunc("/api/v1/users/all", controllers.UserRegisterHandler).Methods("GET")
	userRouter.HandleFunc("/api/v1/users/{id}", controllers.UserRegisterHandler).Methods("GET")

	router.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(commons.Authorize), negroni.Wrap(userRouter),
	))
	return router
}
