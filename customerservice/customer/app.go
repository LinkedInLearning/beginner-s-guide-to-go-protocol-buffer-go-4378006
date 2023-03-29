package customer

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB *sql.DB
}

type Character struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Bio         string `json:"bio"`
	Description string `json:"description"`
}

func (a *App) Initialize() {
	db, err := sql.Open("sqlite3", "./sqlcustomer.db")
	if err != nil {
		log.Printf("Opening database failed for error: %s\n", err.Error())
		db = InitializeDB()
	}

	a.DB = db

	http.HandleFunc("/signup", a.signup)
	http.HandleFunc("/login", a.login)
	http.HandleFunc("/characters", a.getCharacters)
	http.HandleFunc("/characters/", a.getCharacterById)
}

func (a *App) Run() {
	fmt.Println("Server started and listening on port :9005")
	log.Fatal(http.ListenAndServe(":9005", nil))
}

func (a *App) signup(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)

	var c Customer
	json.Unmarshal(reqBody, &c)
	log.Printf("Customer: %+v\n", c)

	if c.existingUser(a.DB) {
		respondWithError(w, http.StatusUnauthorized, "User already exists, unable to create your account.")
		return
	} else {
		if err := c.signup(a.DB); err != nil {
			log.Printf("signup failed with error: %s\n", err.Error())
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	respondWithJSON(w, http.StatusOK, c)
}

func (a *App) login(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)

	var c Customer
	json.Unmarshal(reqBody, &c)
	log.Printf("Customer: %+v\n", c)

	loginSuccess, err := c.login(a.DB)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Can't find user "+c.Username)
	}

	if !loginSuccess {
		respondWithError(w, http.StatusUnauthorized, "Login failed for user "+c.Username)
	}

	respondWithJSON(w, http.StatusOK, c)
}

func (a *App) getCharacters(w http.ResponseWriter, r *http.Request) {
	var characters []Character

	url := "http://localhost:9007/characters"
	log.Printf("url: %s\n", url)

	response, err := http.Get(url)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to get characters")
	}

	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to read body")
	}

	err3 := json.Unmarshal(body, &characters)
	if err3 != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to unmarshal")
	}

	respondWithJSON(w, http.StatusOK, characters)
}

func (a *App) getCharacterById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/characters/")

	url := "http://localhost:9007/characters/" + id
	log.Printf("url: %s\n", url)

	response, err := http.Get(url)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to get characters rpc")
	}

	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to read body rpc")
	}

	var character Character
	err3 := json.Unmarshal(body, &character)
	if err3 != nil {
		respondWithError(w, http.StatusBadRequest, "Failed to unmarshal rpc")
	}

	respondWithJSON(w, http.StatusOK, character)
}

// Helper Functions
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
