package game_test

import (
	"math/rand"
	"playoddeve/api/game"
	"playoddeve/api/ptypes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculateRunsForInnings(t *testing.T) {

	tests := []struct {
		name    string
		innings []ptypes.Deliveries
		want    int
	}{
		{
			name:    "test 1: empty innings",
			innings: []ptypes.Deliveries{},
			want:    0,
		},
		{
			name: "test 2",
			innings: []ptypes.Deliveries{
				{Bat: 1, Bowl: 2},
				{Bat: 2, Bowl: 3},
				{Bat: 3, Bowl: 4},
				{Bat: 4, Bowl: 5},
				{Bat: 5, Bowl: 6},
			},
			want: 15,
		},
		{
			name: "test 3",
			innings: []ptypes.Deliveries{
				{Bat: 1, Bowl: 2},
				{Bat: 4, Bowl: 3},
				{Bat: 2, Bowl: 4},
				{Bat: 6, Bowl: 5},
				{Bat: 5, Bowl: 6},
			},
			want: 18,
		},
		{
			name: "test 4",
			innings: []ptypes.Deliveries{
				{Bat: 1, Bowl: 2},
				{Bat: 4, Bowl: 3},
				{Bat: 2, Bowl: 4},
				{Bat: 6, Bowl: 5},
				{Bat: 5, Bowl: 1},
				{Bat: 1, Bowl: 1},
			},
			want: 18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := game.CalculateRunsForInnings(tt.innings)
			if got != tt.want {
				t.Errorf("CalculateRunsForInnings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsP1TossWinner(t *testing.T) {
	assert.True(t, game.IsP1TossWinner(ptypes.Toss{
		ChoiceP1: game.CHOICE_ODD,
		ChoiceP2: game.CHOICE_EVE,
		NumP1:    1,
		NumP2:    2,
	}))

	assert.False(t, game.IsP1TossWinner(ptypes.Toss{
		ChoiceP1: game.CHOICE_ODD,
		ChoiceP2: game.CHOICE_EVE,
		NumP1:    2,
		NumP2:    2,
	}))
}

func TestRandomNumGenerator(t *testing.T) {
	// Set a seed value to generate predictable pseudo-random numbers
	rand.Seed(time.Now().UnixNano())

	// Define the min and max values for the random number generator
	min := 10
	max := 20

	// Generate 100 random numbers and check that they are within the specified range
	for i := 0; i < 100; i++ {
		num := game.RandomNumGenerator(max, min)
		if num < min || num > max {
			t.Errorf("RandomNumGenerator() returned %d, which is outside the range [%d, %d]", num, min, max)
		}
	}
}

func TestGetWinner(t *testing.T) {

	test_game := ptypes.Game{
		GameId: "test1",
		P1:     "aloo",
		P2:     "yash",
		Toss: ptypes.Toss{
			ChoiceP1: game.CHOICE_ODD,
			ChoiceP2: game.CHOICE_EVE,
			NumP1:    1,
			NumP2:    2,
		},
		TossWinnerChoice: game.TOSS_BAT,
		Innings1: []ptypes.Deliveries{
			{Bat: 6, Bowl: 2},
			{Bat: 6, Bowl: 2},
		},
		Innings2: []ptypes.Deliveries{
			{Bat: 1, Bowl: 2},
			{Bat: 1, Bowl: 1},
		},
	}
	winner1 := game.GetWinner(test_game)
	assert.Equal(t, "aloo", winner1)

	test_game.Innings2 = []ptypes.Deliveries{
		{Bat: 6, Bowl: 2},
		{Bat: 6, Bowl: 2},
		{Bat: 6, Bowl: 6},
	}

	winner2 := game.GetWinner(test_game)
	assert.Equal(t, "tie", winner2)

	test_game.Innings2 = []ptypes.Deliveries{
		{Bat: 6, Bowl: 2},
		{Bat: 6, Bowl: 2},
		{Bat: 6, Bowl: 3},
		{Bat: 6, Bowl: 4},
		{Bat: 6, Bowl: 1},
		{Bat: 6, Bowl: 1},
		{Bat: 6, Bowl: 1},
	}

	winner3 := game.GetWinner(test_game)
	assert.Equal(t, "yash", winner3)

}

func TestGetFirstInningsPlayer(t *testing.T) {

	test_game := ptypes.Game{
		GameId: "test1",
		P1:     "aloo",
		P2:     "yash",
		Toss: ptypes.Toss{
			ChoiceP1: game.CHOICE_ODD,
			ChoiceP2: game.CHOICE_EVE,
			NumP1:    1,
			NumP2:    2,
		},
		TossWinnerChoice: game.TOSS_BAT}

	assert.Equal(t, game.GetFirstInningsPlayer(test_game), "aloo")

	test_game.TossWinnerChoice = game.TOSS_FIELD
	assert.Equal(t, game.GetFirstInningsPlayer(test_game), "yash")

	test_game_2 := ptypes.Game{
		GameId: "test1",
		P1:     "aloo",
		P2:     "yash",
		Toss: ptypes.Toss{
			ChoiceP1: game.CHOICE_ODD,
			ChoiceP2: game.CHOICE_EVE,
			NumP1:    2,
			NumP2:    2,
		},
		TossWinnerChoice: game.TOSS_BAT}

	assert.Equal(t, game.GetFirstInningsPlayer(test_game_2), "yash")
	test_game_2.TossWinnerChoice = game.TOSS_FIELD
	assert.Equal(t, game.GetFirstInningsPlayer(test_game_2), "aloo")

}

func TestGetInnings(t *testing.T) {

	t1 := ptypes.Game{GameId: "test1",
		P1: "aloo",
		P2: "yash",
		Toss: ptypes.Toss{
			ChoiceP1: game.CHOICE_ODD,
			ChoiceP2: game.CHOICE_EVE,
			NumP1:    1,
			NumP2:    2,
		},
		TossWinnerChoice: game.TOSS_BAT}

	innings, deliveries := game.GetInnings(t1)

	assert.Equal(t, innings, 1)
	assert.Len(t, deliveries, 0)

	t2 := ptypes.Game{
		GameId: "test1",
		P1:     "aloo",
		P2:     "yash",
		Toss: ptypes.Toss{
			ChoiceP1: game.CHOICE_ODD,
			ChoiceP2: game.CHOICE_EVE,
			NumP1:    1,
			NumP2:    2,
		},
		TossWinnerChoice: game.TOSS_BAT,
		Innings1: []ptypes.Deliveries{
			{Bat: 6, Bowl: 2},
			{Bat: 6, Bowl: 6},
		},
		Innings2: []ptypes.Deliveries{
			{Bat: 1, Bowl: 2},
			{Bat: 1, Bowl: 1},
		},
	}

	innings, deliveries = game.GetInnings(t2)

	assert.Equal(t, innings, 2)
	assert.Len(t, deliveries, 2)
	assert.Equal(t, deliveries, t2.Innings2)

}
