package main

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/metaphi-org/go-infra-sdk/helpers"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambdaV2

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOriginFunc = func(origin string) bool {
		return strings.HasSuffix(origin, ".awen.finance") ||
			origin == "http://localhost:3000"
	}

	router.Use(cors.New(config))
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.GET("/hello", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"msg": "s"})
	})

	router.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"msg": "helloworld"})
	})

	router.GET(("/hello/ip"), func(c *gin.Context) {
		data, err := http.Get("https://icanhazip.com/")
		if err != nil {
			helpers.HandleError(*c, err)
			return
		}
		body, err := io.ReadAll(data.Body)
		if err != nil {
			helpers.HandleError(*c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"ip": strings.TrimSpace(string(body))})
	})

	ginLambda = ginadapter.NewV2(router)
	lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, request)
}

// && ls && export PATH=$PATH:$(pwd)/go/bin && go version
