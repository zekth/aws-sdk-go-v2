// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfig

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateHostedConfigurationVersionInput struct {
	_ struct{} `type:"structure" payload:"Content"`

	// The application ID.
	//
	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"ApplicationId" type:"string" required:"true"`

	// The configuration profile ID.
	//
	// ConfigurationProfileId is a required field
	ConfigurationProfileId *string `location:"uri" locationName:"ConfigurationProfileId" type:"string" required:"true"`

	// The content of the configuration or the configuration data.
	//
	// Content is a required field
	Content []byte `type:"blob" required:"true" sensitive:"true"`

	// A standard MIME type describing the format of the configuration content.
	// For more information, see Content-Type (https://docs.aws.amazon.com/https:/www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	//
	// ContentType is a required field
	ContentType *string `location:"header" locationName:"Content-Type" min:"1" type:"string" required:"true"`

	// A description of the configuration.
	Description *string `location:"header" locationName:"Description" type:"string"`

	// An optional locking token used to prevent race conditions from overwriting
	// configuration updates when creating a new version. To ensure your data is
	// not overwritten when creating multiple hosted configuration versions in rapid
	// succession, specify the version of the latest hosted configuration version.
	LatestVersionNumber *int64 `location:"header" locationName:"Latest-Version-Number" type:"integer"`
}

// String returns the string representation
func (s CreateHostedConfigurationVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateHostedConfigurationVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateHostedConfigurationVersionInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.ConfigurationProfileId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfigurationProfileId"))
	}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}

	if s.ContentType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContentType"))
	}
	if s.ContentType != nil && len(*s.ContentType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContentType", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateHostedConfigurationVersionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LatestVersionNumber != nil {
		v := *s.LatestVersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Latest-Version-Number", protocol.Int64Value(v), metadata)
	}
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationProfileId != nil {
		v := *s.ConfigurationProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ConfigurationProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Content != nil {
		v := s.Content

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Content", protocol.BytesStream(v), metadata)
	}
	return nil
}

type CreateHostedConfigurationVersionOutput struct {
	_ struct{} `type:"structure" payload:"Content"`

	// The application ID.
	ApplicationId *string `location:"header" locationName:"Application-Id" type:"string"`

	// The configuration profile ID.
	ConfigurationProfileId *string `location:"header" locationName:"Configuration-Profile-Id" type:"string"`

	// The content of the configuration or the configuration data.
	Content []byte `type:"blob" sensitive:"true"`

	// A standard MIME type describing the format of the configuration content.
	// For more information, see Content-Type (https://docs.aws.amazon.com/https:/www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	ContentType *string `location:"header" locationName:"Content-Type" min:"1" type:"string"`

	// A description of the configuration.
	Description *string `location:"header" locationName:"Description" type:"string"`

	// The configuration version.
	VersionNumber *int64 `location:"header" locationName:"Version-Number" type:"integer"`
}

// String returns the string representation
func (s CreateHostedConfigurationVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateHostedConfigurationVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Application-Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationProfileId != nil {
		v := *s.ConfigurationProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Configuration-Profile-Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Version-Number", protocol.Int64Value(v), metadata)
	}
	if s.Content != nil {
		v := s.Content

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "Content", protocol.BytesStream(v), metadata)
	}
	return nil
}

const opCreateHostedConfigurationVersion = "CreateHostedConfigurationVersion"

// CreateHostedConfigurationVersionRequest returns a request value for making API operation for
// Amazon AppConfig.
//
// Create a new configuration in the AppConfig configuration store.
//
//    // Example sending a request using CreateHostedConfigurationVersionRequest.
//    req := client.CreateHostedConfigurationVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appconfig-2019-10-09/CreateHostedConfigurationVersion
func (c *Client) CreateHostedConfigurationVersionRequest(input *CreateHostedConfigurationVersionInput) CreateHostedConfigurationVersionRequest {
	op := &aws.Operation{
		Name:       opCreateHostedConfigurationVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/applications/{ApplicationId}/configurationprofiles/{ConfigurationProfileId}/hostedconfigurationversions",
	}

	if input == nil {
		input = &CreateHostedConfigurationVersionInput{}
	}

	req := c.newRequest(op, input, &CreateHostedConfigurationVersionOutput{})

	return CreateHostedConfigurationVersionRequest{Request: req, Input: input, Copy: c.CreateHostedConfigurationVersionRequest}
}

// CreateHostedConfigurationVersionRequest is the request type for the
// CreateHostedConfigurationVersion API operation.
type CreateHostedConfigurationVersionRequest struct {
	*aws.Request
	Input *CreateHostedConfigurationVersionInput
	Copy  func(*CreateHostedConfigurationVersionInput) CreateHostedConfigurationVersionRequest
}

// Send marshals and sends the CreateHostedConfigurationVersion API request.
func (r CreateHostedConfigurationVersionRequest) Send(ctx context.Context) (*CreateHostedConfigurationVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateHostedConfigurationVersionResponse{
		CreateHostedConfigurationVersionOutput: r.Request.Data.(*CreateHostedConfigurationVersionOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateHostedConfigurationVersionResponse is the response type for the
// CreateHostedConfigurationVersion API operation.
type CreateHostedConfigurationVersionResponse struct {
	*CreateHostedConfigurationVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateHostedConfigurationVersion request.
func (r *CreateHostedConfigurationVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
