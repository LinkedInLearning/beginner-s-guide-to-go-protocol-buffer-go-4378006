package character

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type App struct{}

func (a *App) Initialize() {
	http.HandleFunc("/characters", a.getCharacters)
	http.HandleFunc("/characters/", a.getCharacterById)
}

func (a *App) Run() {
	fmt.Println("Server started and listening on port :9007")
	log.Fatal(http.ListenAndServe(":9007", nil))
}

func (a *App) getCharacters(w http.ResponseWriter, r *http.Request) {
	characters, err := GetCharacters()
	if err != nil {
		log.Printf("GetCharacters failed with error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, characters)
}

func (a *App) getCharacterById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/characters/")

	character, err := GetCharacterById(id)
	if err != nil {
		log.Printf("GetCharactersById failed with error: %s\n", err.Error())
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
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