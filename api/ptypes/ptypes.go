package ptypes

type Toss struct {
	ChoiceP1 string `json:"choice_p1" dynamodbav:"choice_p1"` // either one of odd or eve.
	NumP1    int    `json:"num_p1" dynamodbav:"num_p1"`       // in the range 1 - 6
	ChoiceP2 string `json:"choice_p2" dynamodbav:"choice_p2"`
	NumP2    int    `json:"num_p2" dynamodbav:"num_p2"`
}

type Game struct {
	GameId           string       `json:"game_id" dynamodbav:"game_id"`
	P1               string       `json:"p1" dynamodbav:"p1"`
	P2               string       `json:"p2" dynamodbav:"p2"`
	Toss             Toss         `json:"toss" dynamodbav:"toss"`
	TossWinnerChoice string       `json:"toss_winner_choice" dynamodbav:"toss_winner_choice"`
	Innings1         []Deliveries `json:"innings1" dynamodbav:"innings1"`
	Innings2         []Deliveries `json:"innings2" dynamodbav:"innings2"`
	GameStatus       string       `json:"game_status" dynamodbav:"game_status"` // ongoing/finished/quit
	GameWinner       string       `json:"game_winner" dynamodbav:"game_winner"`
}

type Deliveries struct {
	Bat  int `json:"bat" dynamodbav:"bat"`
	Bowl int `json:"bowl" dynamodbav:"bowl"`
}

type GameKey struct {
	GameId string `dynamodbav:"game_id"`
}
