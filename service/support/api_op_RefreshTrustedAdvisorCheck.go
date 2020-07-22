// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RefreshTrustedAdvisorCheckInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the Trusted Advisor check to refresh. Note: Specifying
	// the check ID of a check that is automatically refreshed causes an InvalidParameterValue
	// error.
	//
	// CheckId is a required field
	CheckId *string `locationName:"checkId" type:"string" required:"true"`
}

// String returns the string representation
func (s RefreshTrustedAdvisorCheckInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RefreshTrustedAdvisorCheckInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RefreshTrustedAdvisorCheckInput"}

	if s.CheckId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CheckId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The current refresh status of a Trusted Advisor check.
type RefreshTrustedAdvisorCheckOutput struct {
	_ struct{} `type:"structure"`

	// The current refresh status for a check, including the amount of time until
	// the check is eligible for refresh.
	//
	// Status is a required field
	Status *TrustedAdvisorCheckRefreshStatus `locationName:"status" type:"structure" required:"true"`
}

// String returns the string representation
func (s RefreshTrustedAdvisorCheckOutput) String() string {
	return awsutil.Prettify(s)
}

const opRefreshTrustedAdvisorCheck = "RefreshTrustedAdvisorCheck"

// RefreshTrustedAdvisorCheckRequest returns a request value for making API operation for
// AWS Support.
//
// Refreshes the AWS Trusted Advisor check that you specify using the check
// ID. You can get the check IDs by calling the DescribeTrustedAdvisorChecks
// operation.
//
// Some checks are refreshed automatically. If you call the RefreshTrustedAdvisorCheck
// operation to refresh them, you might see the InvalidParameterValue error.
//
// The response contains a TrustedAdvisorCheckRefreshStatus object.
//
//    * You must have a Business or Enterprise support plan to use the AWS Support
//    API.
//
//    * If you call the AWS Support API from an account that does not have a
//    Business or Enterprise support plan, the SubscriptionRequiredException
//    error message appears. For information about changing your support plan,
//    see AWS Support (http://aws.amazon.com/premiumsupport/).
//
//    // Example sending a request using RefreshTrustedAdvisorCheckRequest.
//    req := client.RefreshTrustedAdvisorCheckRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/RefreshTrustedAdvisorCheck
func (c *Client) RefreshTrustedAdvisorCheckRequest(input *RefreshTrustedAdvisorCheckInput) RefreshTrustedAdvisorCheckRequest {
	op := &aws.Operation{
		Name:       opRefreshTrustedAdvisorCheck,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RefreshTrustedAdvisorCheckInput{}
	}

	req := c.newRequest(op, input, &RefreshTrustedAdvisorCheckOutput{})

	return RefreshTrustedAdvisorCheckRequest{Request: req, Input: input, Copy: c.RefreshTrustedAdvisorCheckRequest}
}

// RefreshTrustedAdvisorCheckRequest is the request type for the
// RefreshTrustedAdvisorCheck API operation.
type RefreshTrustedAdvisorCheckRequest struct {
	*aws.Request
	Input *RefreshTrustedAdvisorCheckInput
	Copy  func(*RefreshTrustedAdvisorCheckInput) RefreshTrustedAdvisorCheckRequest
}

// Send marshals and sends the RefreshTrustedAdvisorCheck API request.
func (r RefreshTrustedAdvisorCheckRequest) Send(ctx context.Context) (*RefreshTrustedAdvisorCheckResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RefreshTrustedAdvisorCheckResponse{
		RefreshTrustedAdvisorCheckOutput: r.Request.Data.(*RefreshTrustedAdvisorCheckOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RefreshTrustedAdvisorCheckResponse is the response type for the
// RefreshTrustedAdvisorCheck API operation.
type RefreshTrustedAdvisorCheckResponse struct {
	*RefreshTrustedAdvisorCheckOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RefreshTrustedAdvisorCheck request.
func (r *RefreshTrustedAdvisorCheckResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
