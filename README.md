# Welcome to your CDK Go project!

This is a blank project for CDK development with Go.

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests

cd /expenses.backend/service/runtime
GOARCH=amd64 GOOS=linux go build -o bootstrap main.go

cdk deploy --profile backend
cdk destroy --profile backend

aws dynamodb describe-table --table-name Users --profile backend

## Local

cd localstack
./localstack.exe start -d

aws dynamodb list-tables --endpoint-url http://localhost:4566

cd /expenses.backend/service/runtime
go test -v



