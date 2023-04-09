package game

import (
	"context"
	"errors"
	"log"
	"playoddeve/api/ptypes"

	"github.com/google/uuid"
)

const PLAYER_BOT = "bot"

const STATUS_ONGOING = "ongoing"
const STATUS_FINISHED = "finished"

const CHOICE_ODD = "odd"
const CHOICE_EVE = "even"

const TOSS_BAT = "bat"
const TOSS_FIELD = "field"

const RESULT_TIE = "tie"

func GetUserGame(ctx context.Context, game_id string) (ptypes.Game, error) {
	return GetGame(ctx, "16c42cd0-6198-4365-b811-73de88dbaecc")
}

func QuitGame() {

}

func StartGame(ctx context.Context, username string) (ptypes.Game, error) {

	game_id := uuid.NewString()
	game := ptypes.Game{
		GameId:     game_id,
		P1:         username,
		P2:         PLAYER_BOT,
		GameStatus: STATUS_ONGOING,
	}

	err := PutGame(ctx, game)

	return game, err
}

func Toss(ctx context.Context, game_id string, username string,
	user_choice string, user_num int, bat_or_bowl string) (bool, error) {

	// made these variables so that in the future
	// replacement of p2 by bot is easier.
	choice_p1 := user_choice
	num_p1 := user_num
	username_p1 := username

	choice_p2 := ""
	num_p2 := randomIntFromOneToSix()
	username_p2 := PLAYER_BOT

	if user_choice == CHOICE_EVE {
		choice_p2 = CHOICE_ODD
	} else {
		choice_p2 = CHOICE_EVE
	}

	toss := ptypes.Toss{
		ChoiceP1: choice_p1,
		ChoiceP2: choice_p2,
		NumP1:    num_p1,
		NumP2:    num_p2,
	}

	winner_username := ""
	winner_bat_or_bowl := ""

	if IsP1TossWinner(toss) {
		winner_username = username_p1
		winner_bat_or_bowl = bat_or_bowl
	} else {
		winner_username = username_p2
		winner_bat_or_bowl = getBatOrBall()
	}

	err := SaveToss(ctx, game_id, toss, winner_bat_or_bowl)

	return winner_username == username, err
}

func UpdateDelivery(ctx context.Context, game_id string, username string, user_input int) (ptypes.Game, error) {

	game, err := GetGame(ctx, game_id)

	if err != nil {
		return ptypes.Game{}, err
	}

	// stage 1: check if toss has happened & game exists.
	if game.GameId == "" || len(game.TossWinnerChoice) == 0 {
		// throw some error saying that toss hasn't happened or game hasn't been created.
		log.Print("game doesn't exist or toss hasn't happened")

		return ptypes.Game{}, errors.New("game doesn't exist or toss hasn't happened")
	}

	// stage 2: check if status is finished.
	if game.GameStatus == STATUS_FINISHED {
		// return that the innings is finished.
		log.Print("game already finished")

		return ptypes.Game{}, errors.New("game has already finished")
	}

	// stage 3: check which innings is ongoing.
	innings_no, innings := GetInnings(game)
	p2_input := randomIntFromOneToSix() // currently bot is generating it.

	if IsP1Batting(game, innings_no) {

		innings = append(innings, ptypes.Deliveries{
			Bat:  user_input,
			Bowl: p2_input, // can be the reverse, doesn't really matter.
		})

	} else {

		innings = append(innings, ptypes.Deliveries{
			Bat:  p2_input,
			Bowl: user_input, // can be the reverse, doesn't really matter.
		})
	}

	if innings_no == 2 {
		game.Innings2 = innings

		last_delivery := innings[len(innings)-1]

		if last_delivery.Bat == last_delivery.Bowl ||
			CalculateRunsForInnings(game.Innings2) > CalculateRunsForInnings(game.Innings1) {

			game.GameStatus = STATUS_FINISHED
			game.GameWinner = GetWinner(game)
		}
	} else {
		game.Innings1 = innings
		game.GameStatus = STATUS_ONGOING
	}

	err = PutGame(ctx, game)
	return game, err
}
