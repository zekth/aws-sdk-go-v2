// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateAppInput struct {
	_ struct{} `type:"structure"`

	// The name of the app.
	//
	// AppName is a required field
	AppName *string `type:"string" required:"true"`

	// The type of app.
	//
	// AppType is a required field
	AppType AppType `type:"string" required:"true" enum:"true"`

	// The domain ID.
	//
	// DomainId is a required field
	DomainId *string `type:"string" required:"true"`

	// The instance type and the Amazon Resource Name (ARN) of the SageMaker image
	// created on the instance.
	ResourceSpec *ResourceSpec `type:"structure"`

	// Each tag consists of a key and an optional value. Tag keys must be unique
	// per resource.
	Tags []Tag `type:"list"`

	// The user profile name.
	//
	// UserProfileName is a required field
	UserProfileName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateAppInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAppInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAppInput"}

	if s.AppName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppName"))
	}
	if len(s.AppType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AppType"))
	}

	if s.DomainId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainId"))
	}

	if s.UserProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserProfileName"))
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

type CreateAppOutput struct {
	_ struct{} `type:"structure"`

	// The App's Amazon Resource Name (ARN).
	AppArn *string `type:"string"`
}

// String returns the string representation
func (s CreateAppOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateApp = "CreateApp"

// CreateAppRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Creates a running App for the specified UserProfile. Supported Apps are JupyterServer
// and KernelGateway. This operation is automatically invoked by Amazon SageMaker
// Studio upon access to the associated Domain, and when new kernel configurations
// are selected by the user. A user may have multiple Apps active simultaneously.
//
//    // Example sending a request using CreateAppRequest.
//    req := client.CreateAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/CreateApp
func (c *Client) CreateAppRequest(input *CreateAppInput) CreateAppRequest {
	op := &aws.Operation{
		Name:       opCreateApp,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAppInput{}
	}

	req := c.newRequest(op, input, &CreateAppOutput{})

	return CreateAppRequest{Request: req, Input: input, Copy: c.CreateAppRequest}
}

// CreateAppRequest is the request type for the
// CreateApp API operation.
type CreateAppRequest struct {
	*aws.Request
	Input *CreateAppInput
	Copy  func(*CreateAppInput) CreateAppRequest
}

// Send marshals and sends the CreateApp API request.
func (r CreateAppRequest) Send(ctx context.Context) (*CreateAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateAppResponse{
		CreateAppOutput: r.Request.Data.(*CreateAppOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateAppResponse is the response type for the
// CreateApp API operation.
type CreateAppResponse struct {
	*CreateAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateApp request.
func (r *CreateAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
