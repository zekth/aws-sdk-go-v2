// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateUserProfileInput struct {
	_ struct{} `type:"structure"`

	// The ID of the associated Domain.
	//
	// DomainId is a required field
	DomainId *string `type:"string" required:"true"`

	// A specifier for the type of value specified in SingleSignOnUserValue. Currently,
	// the only supported value is "UserName". If the Domain's AuthMode is SSO,
	// this field is required. If the Domain's AuthMode is not SSO, this field cannot
	// be specified.
	SingleSignOnUserIdentifier *string `type:"string"`

	// The username of the associated AWS Single Sign-On User for this UserProfile.
	// If the Domain's AuthMode is SSO, this field is required, and must match a
	// valid username of a user in your directory. If the Domain's AuthMode is not
	// SSO, this field cannot be specified.
	SingleSignOnUserValue *string `type:"string"`

	// Each tag consists of a key and an optional value. Tag keys must be unique
	// per resource.
	Tags []Tag `type:"list"`

	// A name for the UserProfile.
	//
	// UserProfileName is a required field
	UserProfileName *string `type:"string" required:"true"`

	// A collection of settings.
	UserSettings *UserSettings `type:"structure"`
}

// String returns the string representation
func (s CreateUserProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUserProfileInput"}

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
	if s.UserSettings != nil {
		if err := s.UserSettings.Validate(); err != nil {
			invalidParams.AddNested("UserSettings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateUserProfileOutput struct {
	_ struct{} `type:"structure"`

	// The user profile Amazon Resource Name (ARN).
	UserProfileArn *string `type:"string"`
}

// String returns the string representation
func (s CreateUserProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateUserProfile = "CreateUserProfile"

// CreateUserProfileRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Creates a user profile. A user profile represents a single user within a
// domain, and is the main way to reference a "person" for the purposes of sharing,
// reporting, and other user-oriented features. This entity is created when
// a user onboards to Amazon SageMaker Studio. If an administrator invites a
// person by email or imports them from SSO, a user profile is automatically
// created. A user profile is the primary holder of settings for an individual
// user and has a reference to the user's private Amazon Elastic File System
// (EFS) home directory.
//
//    // Example sending a request using CreateUserProfileRequest.
//    req := client.CreateUserProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/CreateUserProfile
func (c *Client) CreateUserProfileRequest(input *CreateUserProfileInput) CreateUserProfileRequest {
	op := &aws.Operation{
		Name:       opCreateUserProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateUserProfileInput{}
	}

	req := c.newRequest(op, input, &CreateUserProfileOutput{})

	return CreateUserProfileRequest{Request: req, Input: input, Copy: c.CreateUserProfileRequest}
}

// CreateUserProfileRequest is the request type for the
// CreateUserProfile API operation.
type CreateUserProfileRequest struct {
	*aws.Request
	Input *CreateUserProfileInput
	Copy  func(*CreateUserProfileInput) CreateUserProfileRequest
}

// Send marshals and sends the CreateUserProfile API request.
func (r CreateUserProfileRequest) Send(ctx context.Context) (*CreateUserProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUserProfileResponse{
		CreateUserProfileOutput: r.Request.Data.(*CreateUserProfileOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUserProfileResponse is the response type for the
// CreateUserProfile API operation.
type CreateUserProfileResponse struct {
	*CreateUserProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUserProfile request.
func (r *CreateUserProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
