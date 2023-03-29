package customer

import (
	"context"
	"log"
	"strconv"

	"github.com/google/uuid"

	cs "github.com/LinkedInLearning/beginner-s-guide-to-go-Proto-protocol-buffer-go-4378006/go/character"
	header "github.com/LinkedInLearning/beginner-s-guide-to-go-Proto-protocol-buffer-go-4378006/go/header"
)

func RpcGetCharacters(client cs.CharacterServiceClient) ([]Character, error) {
	log.Println("Getting GetCharacters through gRPC")

	request := &cs.AllCharactersRequest {
		Header: &header.Header{
			Span: &header.UUID {
				Value: uuid.New().String(),
			},
		},
		Query: &cs.Empty{},
	}

	response, err := client.GetCharacters(context.Background(), request)
	if err != nil {
		return nil, err
	}

	results := response.GetResults()
	var characters []Character
	for _, result := range results {
		character := Character {
			ID: strconv.Itoa(int(result.GetCharacter().GetId())),
			Name: result.GetCharacter().GetName(),
			Category: result.GetCharacter().GetCategory(),
			Bio: result.GetCharacter().GetBio(),
			Description: result.GetCharacter().GetDescription(),
		}

		characters = append(characters, character)
	}
	return characters, nil
}

func RpcGetCharacterById(client cs.CharacterServiceClient, id int) (Character, error) {
	log.Println("Getting GetCharacterById through grpc")

	request := &cs.GetCharacterRequest{
		Header: &header.Header{
			Span: &header.UUID {
				Value: uuid.New().String(),
			},
		},
		CharacterId: int32(id),
	}

	response, err := client.GetCharacterById(context.Background(), request)
	if err != nil {
		return Character{}, err
	}

	result := response.GetResult().GetCharacter()
	character := Character {
			ID: strconv.Itoa(int(result.GetId())),
			Name: result.GetName(),
			Category: result.GetCategory(),
			Bio: result.GetBio(),
			Description: result.GetDescription(),
	}

	return character, nil
}
