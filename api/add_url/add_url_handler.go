package add_url

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"urlshortener/api"
)

const KEY_LEN = 5

type Handler struct {
	Store *api.UrlStore
}

type AddUrlRequest struct {
	Url string `json:"url"`
}

type AddUrlResponse struct {
	Url string `json:"url"`
}

func (a Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := a.handle(w, r)

	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	}
}

var BucketFull = errors.New("Cannot add more urls")

//handles the request of adding a new url shortened
func (a Handler) handle(w http.ResponseWriter, r *http.Request) error {

	parsedUrl, err := a.getUrl(r)
	if err != nil {
		return err
	}

	key, _ := a.Store.ReverseGet(parsedUrl.String())

	if key != "" {
		a.respond(AddUrlResponse{Url: a.buildShortenedUrlFromKey(key)}, w)
		return nil
	}

	err = a.checkUrl(parsedUrl)
	if err != nil {
		return err
	}

	response, err := a.addUrl(parsedUrl.String())
	if err != nil {
		return err
	}

	a.respond(response, w)

	return nil
}

// Builds absolute url by its key
func (a Handler) buildShortenedUrlFromKey(key string) string {
	return os.Getenv("URLSHORTENER_DOMAIN") + "/" + key
}

// Adds url to the store and returns AddUrlResponse with the url shortened
func (a Handler) addUrl(url string) (AddUrlResponse, error) {

	key, err := a.generateUrlKey(url)
	if err != nil {
		return AddUrlResponse{}, err
	}
	err = a.Store.Add(key, url)
	if err != nil {
		return AddUrlResponse{}, err
	}
	return AddUrlResponse{Url: a.buildShortenedUrlFromKey(key)}, nil
}

// Generates a not used random string
func (a Handler) generateUrlKey(url string) (string, error) {

	var key string
	i := 0
	for {
		i++
		key = string(api.RandomBytes(KEY_LEN))
		if _, err := a.Store.Get(key); err != nil {
			break
		}

		if i >= api.NumCombinations()/2 {
			return "", BucketFull
		}
	}

	if i >= api.NumCombinations()/2 {
		log.WithFields(log.Fields{
			"loops": i,
			"max":   api.NumCombinations(),
		}).Warn("Looped more than 1/4 of possibles combinations")
	}

	log.WithFields(log.Fields{"loops": i}).Debug("Key generated")

	return key, nil
}

// Respond to the client
func (a Handler) respond(response AddUrlResponse, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	data, _ := json.Marshal(response)
	w.Write(data)
}

// Obtains url from json request and parses it to validate tha is a valid url
func (a Handler) getUrl(r *http.Request) (*url.URL, error) {

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var addUrlReq AddUrlRequest
	err = json.Unmarshal(b, &addUrlReq)
	if err != nil {
		return nil, err
	}
	parsedUrl, err := url.ParseRequestURI(addUrlReq.Url)
	if err != nil {
		return nil, err
	}

	return parsedUrl, err
}

// Checks that url works correctly (i.e. http request is successful)
func (a Handler) checkUrl(u *url.URL) error {

	resp, err := http.Get(u.String())
	if err != nil {
		return err
	}
	resp.Close = true
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("Url not works. Status code: %d", resp.StatusCode)
	}

	return nil
}
