// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateHttpNamespaceInput struct {
	_ struct{} `type:"structure"`

	// A unique string that identifies the request and that allows failed CreateHttpNamespace
	// requests to be retried without the risk of executing the operation twice.
	// CreatorRequestId can be any unique string, for example, a date/time stamp.
	CreatorRequestId *string `type:"string" idempotencyToken:"true"`

	// A description for the namespace.
	Description *string `type:"string"`

	// The name that you want to assign to this namespace.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The tags to add to the namespace. Each tag consists of a key and an optional
	// value, both of which you define. Tag keys can have a maximum character length
	// of 128 characters, and tag values can have a maximum length of 256 characters.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateHttpNamespaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHttpNamespaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHttpNamespaceInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateHttpNamespaceOutput struct {
	_ struct{} `type:"structure"`

	// A value that you can use to determine whether the request completed successfully.
	// To get the status of the operation, see GetOperation (https://docs.aws.amazon.com/cloud-map/latest/api/API_GetOperation.html).
	OperationId *string `type:"string"`
}

// String returns the string representation
func (s CreateHttpNamespaceOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateHttpNamespace = "CreateHttpNamespace"

// CreateHttpNamespaceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Creates an HTTP namespace. Service instances that you register using an HTTP
// namespace can be discovered using a DiscoverInstances request but can't be
// discovered using DNS.
//
// For the current limit on the number of namespaces that you can create using
// the same AWS account, see AWS Cloud Map Limits (https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html)
// in the AWS Cloud Map Developer Guide.
//
//    // Example sending a request using CreateHttpNamespaceRequest.
//    req := client.CreateHttpNamespaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/CreateHttpNamespace
func (c *Client) CreateHttpNamespaceRequest(input *CreateHttpNamespaceInput) CreateHttpNamespaceRequest {
	op := &aws.Operation{
		Name:       opCreateHttpNamespace,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateHttpNamespaceInput{}
	}

	req := c.newRequest(op, input, &CreateHttpNamespaceOutput{})

	return CreateHttpNamespaceRequest{Request: req, Input: input, Copy: c.CreateHttpNamespaceRequest}
}

// CreateHttpNamespaceRequest is the request type for the
// CreateHttpNamespace API operation.
type CreateHttpNamespaceRequest struct {
	*aws.Request
	Input *CreateHttpNamespaceInput
	Copy  func(*CreateHttpNamespaceInput) CreateHttpNamespaceRequest
}

// Send marshals and sends the CreateHttpNamespace API request.
func (r CreateHttpNamespaceRequest) Send(ctx context.Context) (*CreateHttpNamespaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateHttpNamespaceResponse{
		CreateHttpNamespaceOutput: r.Request.Data.(*CreateHttpNamespaceOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateHttpNamespaceResponse is the response type for the
// CreateHttpNamespace API operation.
type CreateHttpNamespaceResponse struct {
	*CreateHttpNamespaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateHttpNamespace request.
func (r *CreateHttpNamespaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
