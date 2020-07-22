// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package support

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTrustedAdvisorCheckRefreshStatusesInput struct {
	_ struct{} `type:"structure"`

	// The IDs of the Trusted Advisor checks to get the status of.
	//
	// If you specify the check ID of a check that is automatically refreshed, you
	// might see an InvalidParameterValue error.
	//
	// CheckIds is a required field
	CheckIds []string `locationName:"checkIds" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeTrustedAdvisorCheckRefreshStatusesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTrustedAdvisorCheckRefreshStatusesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTrustedAdvisorCheckRefreshStatusesInput"}

	if s.CheckIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("CheckIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The statuses of the Trusted Advisor checks returned by the DescribeTrustedAdvisorCheckRefreshStatuses
// operation.
type DescribeTrustedAdvisorCheckRefreshStatusesOutput struct {
	_ struct{} `type:"structure"`

	// The refresh status of the specified Trusted Advisor checks.
	//
	// Statuses is a required field
	Statuses []TrustedAdvisorCheckRefreshStatus `locationName:"statuses" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeTrustedAdvisorCheckRefreshStatusesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeTrustedAdvisorCheckRefreshStatuses = "DescribeTrustedAdvisorCheckRefreshStatuses"

// DescribeTrustedAdvisorCheckRefreshStatusesRequest returns a request value for making API operation for
// AWS Support.
//
// Returns the refresh status of the AWS Trusted Advisor checks that have the
// specified check IDs. You can get the check IDs by calling the DescribeTrustedAdvisorChecks
// operation.
//
// Some checks are refreshed automatically, and you can't return their refresh
// statuses by using the DescribeTrustedAdvisorCheckRefreshStatuses operation.
// If you call this operation for these checks, you might see an InvalidParameterValue
// error.
//
//    * You must have a Business or Enterprise support plan to use the AWS Support
//    API.
//
//    * If you call the AWS Support API from an account that does not have a
//    Business or Enterprise support plan, the SubscriptionRequiredException
//    error message appears. For information about changing your support plan,
//    see AWS Support (http://aws.amazon.com/premiumsupport/).
//
//    // Example sending a request using DescribeTrustedAdvisorCheckRefreshStatusesRequest.
//    req := client.DescribeTrustedAdvisorCheckRefreshStatusesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/support-2013-04-15/DescribeTrustedAdvisorCheckRefreshStatuses
func (c *Client) DescribeTrustedAdvisorCheckRefreshStatusesRequest(input *DescribeTrustedAdvisorCheckRefreshStatusesInput) DescribeTrustedAdvisorCheckRefreshStatusesRequest {
	op := &aws.Operation{
		Name:       opDescribeTrustedAdvisorCheckRefreshStatuses,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTrustedAdvisorCheckRefreshStatusesInput{}
	}

	req := c.newRequest(op, input, &DescribeTrustedAdvisorCheckRefreshStatusesOutput{})

	return DescribeTrustedAdvisorCheckRefreshStatusesRequest{Request: req, Input: input, Copy: c.DescribeTrustedAdvisorCheckRefreshStatusesRequest}
}

// DescribeTrustedAdvisorCheckRefreshStatusesRequest is the request type for the
// DescribeTrustedAdvisorCheckRefreshStatuses API operation.
type DescribeTrustedAdvisorCheckRefreshStatusesRequest struct {
	*aws.Request
	Input *DescribeTrustedAdvisorCheckRefreshStatusesInput
	Copy  func(*DescribeTrustedAdvisorCheckRefreshStatusesInput) DescribeTrustedAdvisorCheckRefreshStatusesRequest
}

// Send marshals and sends the DescribeTrustedAdvisorCheckRefreshStatuses API request.
func (r DescribeTrustedAdvisorCheckRefreshStatusesRequest) Send(ctx context.Context) (*DescribeTrustedAdvisorCheckRefreshStatusesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeTrustedAdvisorCheckRefreshStatusesResponse{
		DescribeTrustedAdvisorCheckRefreshStatusesOutput: r.Request.Data.(*DescribeTrustedAdvisorCheckRefreshStatusesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeTrustedAdvisorCheckRefreshStatusesResponse is the response type for the
// DescribeTrustedAdvisorCheckRefreshStatuses API operation.
type DescribeTrustedAdvisorCheckRefreshStatusesResponse struct {
	*DescribeTrustedAdvisorCheckRefreshStatusesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeTrustedAdvisorCheckRefreshStatuses request.
func (r *DescribeTrustedAdvisorCheckRefreshStatusesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
