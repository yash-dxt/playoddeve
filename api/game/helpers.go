package game

import (
	"math/rand"
	"playoddeve/api/ptypes"
	"time"
)

func CalculateRunsForInnings(innings []ptypes.Deliveries) int {
	runs := 0
	for _, v := range innings {
		if v.Bat == v.Bowl {
			return runs
		}
		runs += v.Bat
	}
	return runs
}

func GetWinner(game ptypes.Game) string {

	innings1_runs := CalculateRunsForInnings(game.Innings1)
	innings2_runs := CalculateRunsForInnings(game.Innings2)

	first_innings_batter := GetFirstInningsPlayer(game)
	second_innings_batter := ""

	if game.P1 == first_innings_batter {
		second_innings_batter = game.P2
	} else {
		second_innings_batter = game.P1
	}

	if innings1_runs > innings2_runs {
		return first_innings_batter
	} else if innings2_runs > innings1_runs {
		return second_innings_batter
	} else {
		return RESULT_TIE
	}
}

func GetFirstInningsPlayer(game ptypes.Game) string {

	if IsP1TossWinner(game.Toss) {

		if game.TossWinnerChoice == TOSS_BAT {
			return game.P1
		} else {
			return game.P2
		}

	} else {

		if game.TossWinnerChoice == TOSS_BAT {
			return game.P2
		} else {
			return game.P1
		}

	}

}

// innings number & innings itself.
// An array of deliveries combine to become an innings
func GetInnings(game ptypes.Game) (int, []ptypes.Deliveries) {

	// checking the last ball of the innings.
	if len(game.Innings1) == 0 {
		return 1, []ptypes.Deliveries{}
	}
	firstInnings := game.Innings1[len(game.Innings1)-1]

	if firstInnings.Bat == firstInnings.Bowl {
		// this means that the first innigns has ended.
		// according to the rules of odd eve, same number = OUT!!

		return 2, game.Innings2
	}
	return 1, game.Innings1
}

func randomIntFromOneToSix() int {
	max := 6
	min := 1
	return RandomNumGenerator(max, min)
}

func getBatOrBall() string {
	num := RandomNumGenerator(2, 1)

	if num == 1 {
		return TOSS_BAT
	} else {
		return TOSS_FIELD
	}
}

func RandomNumGenerator(max int, min int) int {
	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(max-min+1) + min
	return num
}

func IsP1TossWinner(toss ptypes.Toss) bool {

	num_total := toss.NumP1 + toss.NumP2

	if toss.ChoiceP1 == CHOICE_EVE && num_total%2 == 0 {
		return true
	}

	if toss.ChoiceP1 == CHOICE_ODD && num_total%2 != 0 {
		return true
	}

	return false
}
