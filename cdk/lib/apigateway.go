package lib

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewApiGateway(scope constructs.Construct, lambdaFunc awslambda.Function) awsapigateway.RestApi {
	api := awsapigateway.NewRestApi(scope, jsii.String("ExpenseApi"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("ExpenseAPI"),
		DeployOptions: &awsapigateway.StageOptions{
			StageName: jsii.String("prod"),
		},
	})

	// `GET /expense/users` のエンドポイントを作成
	users := api.Root().AddResource(jsii.String("expense"), nil).AddResource(jsii.String("users"), nil)
	users.AddMethod(jsii.String("GET"), awsapigateway.NewLambdaIntegration(lambdaFunc, nil), nil)

	return api
}
