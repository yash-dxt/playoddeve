package game_test

import (
	"context"
	"log"
	"playoddeve/api/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndToEnd(t *testing.T) {
	test_game, err := game.StartGame(context.Background(), "yamsh")
	game_id := test_game.GameId

	assert.NoError(t, err)
	assert.NotNil(t, test_game.GameId)
	assert.Equal(t, test_game.GameStatus, game.STATUS_ONGOING)

	_, err = game.Toss(context.Background(), game_id, "yamsh", game.CHOICE_ODD, 2, game.TOSS_BAT)
	assert.NoError(t, err)

	test_game, err = game.GetGame(context.Background(), game_id)
	assert.Nil(t, err)

	runs := game.CalculateRunsForInnings(test_game.Innings1)
	assert.Equal(t, runs, 0)

	for i := 0; test_game.GameStatus != game.STATUS_FINISHED; i++ {

		test_game, err = game.UpdateDelivery(context.TODO(), game_id, "yamsh", 6)
		assert.Nil(t, err)
	}

	log.Println("first innings total: ", game.CalculateRunsForInnings(test_game.Innings1))
	log.Println("second innings total: ", game.CalculateRunsForInnings(test_game.Innings2))
	log.Println("winner: ", test_game.GameWinner)
	log.Println("game_id: ", test_game.GameId)

}
