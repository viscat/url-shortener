package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"urlshortener/api"
	"urlshortener/api/add_url"
	"urlshortener/api/redirect"
)

type ApiCommand struct {
}

//Initializes API and start listening for requests
func (a ApiCommand) Execute(args []string) error {
	log.Info("Opening url store")
	db, err := api.OpenDB()
	if err != nil {
		return err
	}
	urlStore, err := api.NewUrlStore(db)
	if err != nil {
		panic(err)
	}

	log.Info("Url store open. Initializing API")
	r := mux.NewRouter()
	r.Handle("/add", &add_url.Handler{Store: urlStore}).Methods(http.MethodPut)
	r.Handle("/favicon.ico", http.NotFoundHandler())
	r.Handle("/{hash}", &redirect.Handler{Store: urlStore}).Methods(http.MethodGet)

	log.Info("API initialized. Waiting for requests...")
	return http.ListenAndServe(":"+os.Getenv("URLSHORTENER_PORT"), r)
}
