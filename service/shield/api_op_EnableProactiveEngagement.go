// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type EnableProactiveEngagementInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableProactiveEngagementInput) String() string {
	return awsutil.Prettify(s)
}

type EnableProactiveEngagementOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s EnableProactiveEngagementOutput) String() string {
	return awsutil.Prettify(s)
}

const opEnableProactiveEngagement = "EnableProactiveEngagement"

// EnableProactiveEngagementRequest returns a request value for making API operation for
// AWS Shield.
//
// Authorizes the DDoS Response Team (DRT) to use email and phone to notify
// contacts about escalations to the DRT and to initiate proactive customer
// support.
//
//    // Example sending a request using EnableProactiveEngagementRequest.
//    req := client.EnableProactiveEngagementRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/EnableProactiveEngagement
func (c *Client) EnableProactiveEngagementRequest(input *EnableProactiveEngagementInput) EnableProactiveEngagementRequest {
	op := &aws.Operation{
		Name:       opEnableProactiveEngagement,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableProactiveEngagementInput{}
	}

	req := c.newRequest(op, input, &EnableProactiveEngagementOutput{})

	return EnableProactiveEngagementRequest{Request: req, Input: input, Copy: c.EnableProactiveEngagementRequest}
}

// EnableProactiveEngagementRequest is the request type for the
// EnableProactiveEngagement API operation.
type EnableProactiveEngagementRequest struct {
	*aws.Request
	Input *EnableProactiveEngagementInput
	Copy  func(*EnableProactiveEngagementInput) EnableProactiveEngagementRequest
}

// Send marshals and sends the EnableProactiveEngagement API request.
func (r EnableProactiveEngagementRequest) Send(ctx context.Context) (*EnableProactiveEngagementResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableProactiveEngagementResponse{
		EnableProactiveEngagementOutput: r.Request.Data.(*EnableProactiveEngagementOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableProactiveEngagementResponse is the response type for the
// EnableProactiveEngagement API operation.
type EnableProactiveEngagementResponse struct {
	*EnableProactiveEngagementOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableProactiveEngagement request.
func (r *EnableProactiveEngagementResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
