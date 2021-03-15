package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter creates all endpoint for chat app.
func NewRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/signup", postSignup).Methods(http.MethodPost)
	router.HandleFunc("/signin", postSignin).Methods(http.MethodPost)

	router.HandleFunc("/page/{page_id}", getPage).Methods(http.MethodGet)
	router.HandleFunc("/page/{page_id}", putPage).Methods(http.MethodPut)
	router.HandleFunc("/page/{page_id}", deletePage).Methods(http.MethodDelete)

	router.HandleFunc("/page/{page_id}/move", postMovePage).Methods(http.MethodPost)

	return router
}

type postSignupRequestModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// postSignup handles request for a new account creation.
func postSignup(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// postSignin handles login request for existing user.
func postSignin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

type link struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type getPageResponseModel struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Links    []link `json:"links"`
}

// getPage handles request for page info.
func getPage(w http.ResponseWriter, r *http.Request) {
	m := getPageResponseModel{Title: "Awesome nails", Subtitle: "", Links: []link{link{Name: "vk", Url: "https://vk.com/id0"}}}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type putPageRequestModel struct {
	IsPublic bool   `json:"is-public"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Links    []link `json:"links"`
}

// putPage writes page info.
func putPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// deletePage deletes page.
func deletePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

type postMovePageRequestModel struct {
	IsPublic bool `json:"is-public"`
	PageId string `json:"page-id"`
}

// postMovePage moves page to new address.
func postMovePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
