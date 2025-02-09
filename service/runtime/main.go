package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/comail/colog"
)

var (
	dynamoDB *dynamodb.DynamoDB
)

func init() {

	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

}

func getDynamoDBClient() *dynamodb.DynamoDB {
	if dynamoDB != nil {
		return dynamoDB
	}

	var config *aws.Config

	if os.Getenv("AWS_LOCALSTACK") == "true" {
		config = &aws.Config{
			Region:   aws.String("ap-northeast-1"),
			Endpoint: aws.String("http://localhost:4566"),
		}
		log.Println("Using LocalStack DynamoDB at http://localhost:4566")
	} else {
		config = &aws.Config{
			Region: aws.String("ap-northeast-1"),
		}
		log.Println("Using AWS DynamoDB (ap-northeast-1)")
	}

	sess := session.Must(session.NewSession(config))
	dynamoDB = dynamodb.New(sess)

	return dynamoDB
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	table := os.Getenv("USERS")
	if table == "" {
		log.Println("Error: USERS environment variable is not set")
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Error: USERS environment variable is not set"}, nil
	}

	dynamoDB := getDynamoDBClient()

	input := &dynamodb.ScanInput{TableName: aws.String(table)}
	result, err := dynamoDB.Scan(input)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 500, Body: fmt.Sprintf("Error: %v", err)}, nil
	}

	simplifiedItems := []map[string]string{}

	for _, item := range result.Items {
		simpleItem := map[string]string{}
		for key, value := range item {
			if value.S != nil {
				simpleItem[key] = *value.S
			}
		}
		simplifiedItems = append(simplifiedItems, simpleItem)
	}
	itemsJSON, _ := json.Marshal(simplifiedItems)
	log.Printf("debug: items %s", itemsJSON)

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(itemsJSON)}, nil
}

func main() {
	lambda.Start(handler)
}
