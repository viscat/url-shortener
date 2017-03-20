package redirect

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"
	"urlshortener/api"
)

type Handler struct {
	Store *api.UrlStore
}

var UrlNotFound api.Response = api.Response{
	Message: "Url not found",
}

// Obtains the url from its hash and do the redirect. If the url is not found, responds with a 404 status
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url, err := h.Store.Get(vars["hash"])
	if err == api.UrlNotExists {
		log.WithFields(log.Fields{
			"hash": vars["hash"],
		}).Info("Url not found by hash")
		data, _ := json.Marshal(UrlNotFound)
		http.Error(w, string(data), http.StatusNotFound)
		return
	}
	http.Redirect(w, r, string(url), http.StatusTemporaryRedirect)
}
