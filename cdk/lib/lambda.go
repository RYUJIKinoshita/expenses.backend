package lib

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewLambda(scope constructs.Construct, userTable awsdynamodb.Table) awslambda.Function {
	lambda := awslambda.NewFunction(scope, jsii.String("GetUsersLambda"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_PROVIDED_AL2(),
		// Runtime:      awslambda.Runtime_GO_1_X(),
		Handler: jsii.String("bootstrap"),
		// Handler:      jsii.String("main"),
		Code:        awslambda.AssetCode_FromAsset(jsii.String("./service/runtime"), nil),
		Environment: &map[string]*string{"USERS": userTable.TableName()},
	})

	// Lambda に DynamoDB へのアクセス権を付与
	userTable.GrantReadData(lambda)

	return lambda
}
