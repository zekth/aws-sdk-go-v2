// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the Amazon Resource Name (ARN) of the gateway.
type DescribeBandwidthRateLimitInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBandwidthRateLimitInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBandwidthRateLimitInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBandwidthRateLimitInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the following fields:
type DescribeBandwidthRateLimitOutput struct {
	_ struct{} `type:"structure"`

	// The average download bandwidth rate limit in bits per second. This field
	// does not appear in the response if the download rate limit is not set.
	AverageDownloadRateLimitInBitsPerSec *int64 `min:"102400" type:"long"`

	// The average upload bandwidth rate limit in bits per second. This field does
	// not appear in the response if the upload rate limit is not set.
	AverageUploadRateLimitInBitsPerSec *int64 `min:"51200" type:"long"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s DescribeBandwidthRateLimitOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeBandwidthRateLimit = "DescribeBandwidthRateLimit"

// DescribeBandwidthRateLimitRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns the bandwidth rate limits of a gateway. By default, these limits
// are not set, which means no bandwidth rate limiting is in effect. This operation
// is supported for the stored volume, cached volume and tape gateway types.
//
// This operation only returns a value for a bandwidth rate limit only if the
// limit is set. If no limits are set for the gateway, then this operation returns
// only the gateway ARN in the response body. To specify which gateway to describe,
// use the Amazon Resource Name (ARN) of the gateway in your request.
//
//    // Example sending a request using DescribeBandwidthRateLimitRequest.
//    req := client.DescribeBandwidthRateLimitRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeBandwidthRateLimit
func (c *Client) DescribeBandwidthRateLimitRequest(input *DescribeBandwidthRateLimitInput) DescribeBandwidthRateLimitRequest {
	op := &aws.Operation{
		Name:       opDescribeBandwidthRateLimit,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBandwidthRateLimitInput{}
	}

	req := c.newRequest(op, input, &DescribeBandwidthRateLimitOutput{})

	return DescribeBandwidthRateLimitRequest{Request: req, Input: input, Copy: c.DescribeBandwidthRateLimitRequest}
}

// DescribeBandwidthRateLimitRequest is the request type for the
// DescribeBandwidthRateLimit API operation.
type DescribeBandwidthRateLimitRequest struct {
	*aws.Request
	Input *DescribeBandwidthRateLimitInput
	Copy  func(*DescribeBandwidthRateLimitInput) DescribeBandwidthRateLimitRequest
}

// Send marshals and sends the DescribeBandwidthRateLimit API request.
func (r DescribeBandwidthRateLimitRequest) Send(ctx context.Context) (*DescribeBandwidthRateLimitResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeBandwidthRateLimitResponse{
		DescribeBandwidthRateLimitOutput: r.Request.Data.(*DescribeBandwidthRateLimitOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeBandwidthRateLimitResponse is the response type for the
// DescribeBandwidthRateLimit API operation.
type DescribeBandwidthRateLimitResponse struct {
	*DescribeBandwidthRateLimitOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeBandwidthRateLimit request.
func (r *DescribeBandwidthRateLimitResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
