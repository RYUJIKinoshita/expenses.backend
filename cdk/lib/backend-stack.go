package lib

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

type BackendStackProps struct {
	awscdk.StackProps
}

type BackendStack struct {
	awscdk.Stack
	UserTable  awsdynamodb.Table
	LambdaFunc awslambda.Function
	Api        awsapigateway.RestApi
}

func NewBackendStack(scope constructs.Construct, id string, props *BackendStackProps) *BackendStack {
	stack := awscdk.NewStack(scope, &id, &props.StackProps)

	// DynamoDB
	userTable := NewDynamoDB(stack)

	// Lambda
	lambdaFunc := NewLambda(stack, userTable)

	// API Gateway
	api := NewApiGateway(stack, lambdaFunc)

	// CloudFront
	NewCloudFront(stack, api)

	return &BackendStack{
		Stack:     stack,
		UserTable: userTable,
		LambdaFunc: lambdaFunc,
		Api:       api,
	}
}
