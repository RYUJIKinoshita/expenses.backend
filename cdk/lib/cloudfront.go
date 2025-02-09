package lib

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfrontorigins"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewCloudFront(scope constructs.Construct, api awsapigateway.RestApi) {
	// CloudFront ディストリビューションを作成
	cloudFront := awscloudfront.NewDistribution(scope, jsii.String("ExpenseApiDistribution"), &awscloudfront.DistributionProps{
		DefaultBehavior: &awscloudfront.BehaviorOptions{
			Origin:      awscloudfrontorigins.NewRestApiOrigin(api, &awscloudfrontorigins.RestApiOriginProps{}),
			CachePolicy: awscloudfront.CachePolicy_CACHING_DISABLED(),
		},
	})

	// CloudFront の URL を出力
	awscdk.NewCfnOutput(scope, jsii.String("CloudFrontURL"), &awscdk.CfnOutputProps{
		Value: cloudFront.DistributionDomainName(),
	})
}
