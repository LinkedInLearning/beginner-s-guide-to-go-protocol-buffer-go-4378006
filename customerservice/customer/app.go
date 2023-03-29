package customer

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cs "github.com/LinkedInLearning/beginner-s-guide-to-go-Proto-protocol-buffer-go-4378006/go/character"
)

type App struct {
	DB     *sql.DB
	client cs.CharacterServiceClient
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
	a.client = InitializeCharacterClient()
	http.HandleFunc("/signup", a.signup)
	http.HandleFunc("/login", a.login)
	http.HandleFunc("/characters", a.getCharacters)
	http.HandleFunc("/characters/", a.getCharacterById)
	http.HandleFunc("/charactersrpc", a.getCharactersRPC)
	http.HandleFunc("/charactersrpc/", a.getCharacterByIdRPC)
}

func InitializeCharacterClient() cs.CharacterServiceClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	fmt.Println("Dialing characterservice grpc at localhost:9008")
	conn, err := grpc.Dial("localhost:9008", opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	fmt.Println("Connection established")
	client := cs.NewCharacterServiceClient(conn)

	return client
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

func (a *App) getCharactersRPC(w http.ResponseWriter, r *http.Request) {
	characters, err := RpcGetCharacters(a.client)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	respondWithJSON(w, http.StatusOK, characters)
}

func (a *App) getCharacterByIdRPC(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/charactersrpc/")
	characterId, _ := strconv.Atoi(id)
	fmt.Printf("characterId: %d", characterId)

	character, err := RpcGetCharacterById(a.client, characterId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
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
