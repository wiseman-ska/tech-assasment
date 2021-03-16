package routers

import "github.com/gorilla/mux"

func SetupUserRoutes(router *mux.Router) *mux.Router {
	userRouter := mux.NewRouter()
	userRouter.HandleFunc("/users/login", controllers.UserLoginHandler).Methods("POST")
	userRouter.HandleFunc("/users/create", controllers.UserRegisterHandler).Methods("POST")
	userRouter.HandleFunc("/api/v1/users/update", controllers.UserRegisterHandler).Methods("POST")
	userRouter.HandleFunc("/api/v1/users/delete", controllers.UserRegisterHandler).Methods("DELETE")
	userRouter.HandleFunc("/api/v1/users", controllers.UserRegisterHandler).Methods("GET")
	userRouter.HandleFunc("/api/v1/users/{id}", controllers.UserRegisterHandler).Methods("GET")

	router.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(commons.Authorize),
		negroni.Wrap(userRouter),
	))
	return router
}
