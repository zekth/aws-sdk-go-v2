// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateSubscriptionInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateSubscriptionInput) String() string {
	return awsutil.Prettify(s)
}

type CreateSubscriptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateSubscriptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSubscription = "CreateSubscription"

// CreateSubscriptionRequest returns a request value for making API operation for
// AWS Shield.
//
// Activates AWS Shield Advanced for an account.
//
// When you initally create a subscription, your subscription is set to be automatically
// renewed at the end of the existing subscription period. You can change this
// by submitting an UpdateSubscription request.
//
//    // Example sending a request using CreateSubscriptionRequest.
//    req := client.CreateSubscriptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/CreateSubscription
func (c *Client) CreateSubscriptionRequest(input *CreateSubscriptionInput) CreateSubscriptionRequest {
	op := &aws.Operation{
		Name:       opCreateSubscription,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSubscriptionInput{}
	}

	req := c.newRequest(op, input, &CreateSubscriptionOutput{})

	return CreateSubscriptionRequest{Request: req, Input: input, Copy: c.CreateSubscriptionRequest}
}

// CreateSubscriptionRequest is the request type for the
// CreateSubscription API operation.
type CreateSubscriptionRequest struct {
	*aws.Request
	Input *CreateSubscriptionInput
	Copy  func(*CreateSubscriptionInput) CreateSubscriptionRequest
}

// Send marshals and sends the CreateSubscription API request.
func (r CreateSubscriptionRequest) Send(ctx context.Context) (*CreateSubscriptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSubscriptionResponse{
		CreateSubscriptionOutput: r.Request.Data.(*CreateSubscriptionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSubscriptionResponse is the response type for the
// CreateSubscription API operation.
type CreateSubscriptionResponse struct {
	*CreateSubscriptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSubscription request.
func (r *CreateSubscriptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
