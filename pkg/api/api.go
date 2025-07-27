package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"shorterUrl/pkg/db"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type API struct {
	router *mux.Router
}

func New() *API {
	api := API{router: mux.NewRouter()}
	api.endpoints()
	return &api
}

func (api *API) Router() *mux.Router {
	return api.router
}

type ShortRes struct {
	ShortUrl string `json:"shortUrl"`
	Error    string `json:"error,omitempty"`
}

func (api *API) endpoints() {
	api.router.Use(logMiddleware)
	api.router.HandleFunc("/{url:.+}", api.getShortUrl).Methods(http.MethodGet)

}

func genCode() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

// Получение короткой ссылки
func (api *API) getShortUrl(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	url_l := params["url"]
	if url_l == "" {
		err := json.NewEncoder(w).Encode("Url have to exist")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	url_s := genCode()
	_, err := db.DB.Exec("INSERT INTO urls (long_url, short_code) VALUES (?, ?)", url_l, url_s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	answer := "http://" + r.Host + "/" + url_s
	err = json.NewEncoder(w).Encode(answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (api *API) redirectUrl(w http.ResponseWriter, r *http.Request) {

}
