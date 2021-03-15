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

	router.HandleFunc("/public/{page_id}", getPublicPage).Methods(http.MethodGet)
	//router.HandleFunc("/public/{page_id}", putPublicPage).Methods(http.MethodPut)
	//router.HandleFunc("/public/{page_id}", deletePublicPage).Methods(http.MethodDelete)

	router.HandleFunc("/private/{page_id}", getPrivatePage).Methods(http.MethodGet)
	//router.HandleFunc("/private/{page_id}", putPrivatePage).Methods(http.MethodPut)
	//router.HandleFunc("/private/{page_id}", deletePrivatePage).Methods(http.MethodDelete)

	//router.HandleFunc("/public/{page_id}/move", postMovePublicPage).Methods(http.MethodPost)
	//router.HandleFunc("/private/{page_id}/move", postMovePrivatePage).Methods(http.MethodPost)

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
func getPage(w http.ResponseWriter, r *http.Request, isPublic bool) {
	m := getPageResponseModel{Title: "Awesome nails", Subtitle: "", Links: []link{link{Name: "vk", Url: "https://vk.com/id0"}}}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// getPublicPage handles request for public page info.
func getPublicPage(w http.ResponseWriter, r *http.Request) {
	getPage(w, r, true)
}

// getPrivatePage handles request for private page info.
func getPrivatePage(w http.ResponseWriter, r *http.Request) {
	getPage(w, r, false)
}
