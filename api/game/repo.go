package game

import (
	"context"
	"os"
	"playoddeve/api/infra"
	"playoddeve/api/ptypes"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var gameTable = os.Getenv("TABLE_GAME")

func PutGame(ctx context.Context, game ptypes.Game) error {

	db, err := infra.GetDynamoDb()
	if err != nil {
		return err
	}

	item, err := attributevalue.MarshalMap(game)
	if err != nil {
		return err
	}

	_, err = db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &gameTable,
		Item:      item,
	})

	return err
}

func SaveToss(ctx context.Context, game_id string, toss ptypes.Toss, toss_winner_choice string) error {
	db, errClient := infra.GetDynamoDb()

	if errClient != nil {
		return errClient
	}

	upd := expression.
		Set(expression.Name("toss_winner_choice"), expression.Value(toss_winner_choice)).
		Set(expression.Name("toss"), expression.Value(toss))

	expr, errExpressionBuild := expression.NewBuilder().WithUpdate(upd).Build()

	if errExpressionBuild != nil {
		return errExpressionBuild
	}

	key, err := attributevalue.MarshalMap(ptypes.GameKey{
		GameId: game_id,
	})

	if err != nil {
		return err
	}

	input := dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 &gameTable,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	_, errQuery := db.UpdateItem(ctx, &input)

	return errQuery
}

func GetGame(ctx context.Context, game_id string) (ptypes.Game, error) {
	var game ptypes.Game
	db, _ := infra.GetDynamoDb()

	keyMap, err := attributevalue.MarshalMap(ptypes.GameKey{
		GameId: game_id,
	})

	if err != nil {
		return game, err
	}

	gameItem, err := db.GetItem(ctx, &dynamodb.GetItemInput{
		Key:       keyMap,
		TableName: &gameTable,
	})

	if err != nil {
		return game, err
	}

	err = attributevalue.UnmarshalMap(gameItem.Item, &game)

	return game, err
}
