package main


import (
	"log"
	"net/http"
	"github.com/urfave/negroni"
)

func main() {
	common.StartUp()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr: common.AppConf.Server,
		Handler: n,
	}

	log.Println("Listening...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

