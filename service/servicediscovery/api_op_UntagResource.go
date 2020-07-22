// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UntagResourceInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource that you want to retrieve
	// tags for.
	//
	// ResourceARN is a required field
	ResourceARN *string `min:"1" type:"string" required:"true"`

	// The tag keys to remove from the specified resource.
	//
	// TagKeys is a required field
	TagKeys []string `type:"list" required:"true"`
}

// String returns the string representation
func (s UntagResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UntagResourceInput"}

	if s.ResourceARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceARN"))
	}
	if s.ResourceARN != nil && len(*s.ResourceARN) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceARN", 1))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UntagResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UntagResourceOutput) String() string {
	return awsutil.Prettify(s)
}

const opUntagResource = "UntagResource"

// UntagResourceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Removes one or more tags from the specified resource.
//
//    // Example sending a request using UntagResourceRequest.
//    req := client.UntagResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/UntagResource
func (c *Client) UntagResourceRequest(input *UntagResourceInput) UntagResourceRequest {
	op := &aws.Operation{
		Name:       opUntagResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UntagResourceInput{}
	}

	req := c.newRequest(op, input, &UntagResourceOutput{})

	return UntagResourceRequest{Request: req, Input: input, Copy: c.UntagResourceRequest}
}

// UntagResourceRequest is the request type for the
// UntagResource API operation.
type UntagResourceRequest struct {
	*aws.Request
	Input *UntagResourceInput
	Copy  func(*UntagResourceInput) UntagResourceRequest
}

// Send marshals and sends the UntagResource API request.
func (r UntagResourceRequest) Send(ctx context.Context) (*UntagResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UntagResourceResponse{
		UntagResourceOutput: r.Request.Data.(*UntagResourceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UntagResourceResponse is the response type for the
// UntagResource API operation.
type UntagResourceResponse struct {
	*UntagResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UntagResource request.
func (r *UntagResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
