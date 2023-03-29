package main

import (
	"context"
	"strconv"

	"example.com/character"

	pb "github.com/LinkedInLearning/beginner-s-guide-to-go-Proto-protocol-buffer-go-4378006/go/character"
)

type CharacterServer struct {
	pb.UnimplementedCharacterServiceServer
}

func newServer() *CharacterServer {
	s := &CharacterServer{}

	return s
}

func (cs *CharacterServer) GetCharacters(ctx context.Context, request *pb.AllCharactersRequest) (*pb.AllCharactersResponse, error) {
	characters, err := character.GetCharacters()
	if err != nil {
		return nil, err
	}

	var results []*pb.Result
	for _, character := range characters {
		i, _ := strconv.ParseInt(character.ID, 10, 32)

		id := int32(i)

		result := &pb.Result{
			Character: &pb.Character{
				Id: id,
				Name: character.Name,
				Category: character.Category,
				Bio: character.Bio,
				Description: character.Description,
			},
		}

		results = append(results, result)
	}

	return &pb.AllCharactersResponse{
		Header: request.GetHeader(),
		Results: results,
	}, nil
}

func (cs *CharacterServer) GetCharacterById(ctx context.Context, request *pb.GetCharacterRequest) (*pb.GetCharacterResponse, error) {
	response, err := character.GetCharacterById(strconv.Itoa(int(request.GetCharacterId())))
	if err != nil {
		return nil, err
	}

	i, _ := strconv.ParseInt(response.ID, 10, 32)

	id := int32(i)
	return &pb.GetCharacterResponse{
		Header: request.GetHeader(),
		Result: &pb.Result{
			Character: &pb.Character{
				Id: id,
				Name: response.Name,
				Category: response.Category,
				Bio: response.Bio,
				Description: response.Description,
			},
		},
	}, nil
}