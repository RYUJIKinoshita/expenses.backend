package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"expense.backend/cdk/lib"
)

type BackendStackProps struct {
	awscdk.StackProps
}

func NewBackendStack(scope constructs.Construct, id string, props *BackendStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// DynamoDB
	userTable := lib.NewDynamoDB(stack)

	// Lambda
	lambdaFunc := lib.NewLambda(stack, userTable)

	// API Gateway
	api := lib.NewApiGateway(stack, lambdaFunc)

	// CloudFront
	lib.NewCloudFront(stack, api)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)
	awscdk.Tags_Of(app).Add(jsii.String("System"), jsii.String("Expense"), nil)

	NewBackendStack(app, "BackendStack", &BackendStackProps{
		awscdk.StackProps{
			Env: &awscdk.Environment{
				Region: jsii.String("ap-northeast-1"),
			},
		},
	})

	app.Synth(nil)
}
