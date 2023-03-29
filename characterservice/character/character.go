package character

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

type Character struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Bio         string `json:"bio"`
	Description string `json:"description"`
}

func GetCharacters() ([]Character, error) {
	var characters []Character

	url := "https://bigstarcollectibles.com/api/characters/all"
	log.Printf("url: %s\n", url)
	
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Get Characters response failed")
	}

	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		return nil, errors.New("Get Characers response - get body failed")
	}

	err3 := json.Unmarshal(body, &characters)
	if err3 != nil {
		return nil, errors.New("Get Characters response - unmarshaling failed")
	}

	return characters, nil
}

func GetCharacterById(id string) (*Character, error) {
	var characters []Character

	url := "https://bigstarcollectibles.com/api/characters/" + id
	log.Printf("url: %s\n", url)

	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("response failed")
	}

	body, err2 := io.ReadAll(response.Body)
	if err2 != nil {
		return nil, errors.New("get body failed")
	}

	err3 := json.Unmarshal(body, &characters)
	if err3 != nil {
		return nil, errors.New("unmarshaling failed")
	}

	return &characters[0], nil
}