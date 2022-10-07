package services

import (
	"testing"

	"github.com/irahardianto/service-pattern-go/interfaces/mocks"
	"github.com/irahardianto/service-pattern-go/models"

	"github.com/stretchr/testify/assert"
)

func TestGetScore(t *testing.T) {
	playerRepository := new(mocks.IPlayerRepository)

	player1 := models.PlayerModel{
		Id:    101,
		Name:  "Rafael",
		Score: 3,
	}

	player2 := models.PlayerModel{
		Id:    103,
		Name:  "Serena",
		Score: 1,
	}
	playerRepository.On("GetPlayerByName", "Rafael").Return(player1, nil)
	playerRepository.On("GetPlayerByName", "Serena").Return(player2, nil)

	playerService := PlayerService{playerRepository}

	expectedResult := "Forty-Fifteen"

	actualResult, err := playerService.GetScores("Rafael", "Serena")
	if err != nil {
		t.Errorf("unexpected err during score info, err %s", err)
	}

	assert.Equal(t, expectedResult, actualResult)
}

func TestFailing(t *testing.T) {
	playerRepository := new(mocks.IPlayerRepository)

	player1 := models.PlayerModel{
		Id:    101,
		Name:  "Rafael",
		Score: 3,
	}

	player2 := models.PlayerModel{
		Id:    103,
		Name:  "Serena",
		Score: 3,
	}

	playerRepository.On("GetPlayerByName", "Rafael").Return(player1, nil)
	playerRepository.On("GetPlayerByName", "Serena").Return(player2, nil)
	playerService := PlayerService{playerRepository}

	// should be "Deuce" to not  fail
	expectedResult := "Thirty-All"

	actualResult, err := playerService.GetScores("Rafael", "Serena")
	if err != nil {
		t.Errorf("unexpected err during score info, err %s", err)
	}

	assert.Equal(t, expectedResult, actualResult)
}
