package main


import (
	"github.com/urfave/negroni"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/commons"
	"github.com/wiseman-ska/tech-assessment/user-manager-api/controllers/routers"
	"log"
	"net/http"
)

func main() {
	commons.StartUp()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr: commons.AppConf.Server,
		Handler: n,
	}

	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

