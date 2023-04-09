package infra

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var (
	awsCfg *aws.Config
)

var awsRegion = os.Getenv("AWS_REGION")

func GetAwsConfig() *aws.Config {
	if awsCfg != nil {
		return awsCfg
	}
	newAwsCfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("could not get awsconfig")
	}
	newAwsCfg.Region = awsRegion
	awsCfg = &newAwsCfg
	return awsCfg
}

var db *dynamodb.Client

func GetDynamoDb() (*dynamodb.Client, error) {
	if db != nil {
		return db, nil
	}
	db = dynamodb.NewFromConfig(*GetAwsConfig())
	return db, nil
}
