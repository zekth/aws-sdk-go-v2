module github.com/aws/aws-sdk-go-v2/service/timestreamquery

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.16.2
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.9
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.3
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.3
	github.com/aws/smithy-go v1.11.2
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery => ../../service/internal/endpoint-discovery/
