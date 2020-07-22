// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateThemeAliasInput struct {
	_ struct{} `type:"structure"`

	// The name of the theme alias that you want to update.
	//
	// AliasName is a required field
	AliasName *string `location:"uri" locationName:"AliasName" min:"1" type:"string" required:"true"`

	// The ID of the AWS account that contains the theme alias that you're updating.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the theme.
	//
	// ThemeId is a required field
	ThemeId *string `location:"uri" locationName:"ThemeId" min:"1" type:"string" required:"true"`

	// The version number of the theme that the alias should reference.
	//
	// ThemeVersionNumber is a required field
	ThemeVersionNumber *int64 `min:"1" type:"long" required:"true"`
}

// String returns the string representation
func (s UpdateThemeAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateThemeAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateThemeAliasInput"}

	if s.AliasName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AliasName"))
	}
	if s.AliasName != nil && len(*s.AliasName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AliasName", 1))
	}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.ThemeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThemeId"))
	}
	if s.ThemeId != nil && len(*s.ThemeId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThemeId", 1))
	}

	if s.ThemeVersionNumber == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThemeVersionNumber"))
	}
	if s.ThemeVersionNumber != nil && *s.ThemeVersionNumber < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("ThemeVersionNumber", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateThemeAliasInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ThemeVersionNumber != nil {
		v := *s.ThemeVersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ThemeVersionNumber", protocol.Int64Value(v), metadata)
	}
	if s.AliasName != nil {
		v := *s.AliasName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AliasName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThemeId != nil {
		v := *s.ThemeId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ThemeId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateThemeAliasOutput struct {
	_ struct{} `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// Information about the theme alias.
	ThemeAlias *ThemeAlias `type:"structure"`
}

// String returns the string representation
func (s UpdateThemeAliasOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateThemeAliasOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThemeAlias != nil {
		v := s.ThemeAlias

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "ThemeAlias", v, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opUpdateThemeAlias = "UpdateThemeAlias"

// UpdateThemeAliasRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates an alias of a theme.
//
//    // Example sending a request using UpdateThemeAliasRequest.
//    req := client.UpdateThemeAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateThemeAlias
func (c *Client) UpdateThemeAliasRequest(input *UpdateThemeAliasInput) UpdateThemeAliasRequest {
	op := &aws.Operation{
		Name:       opUpdateThemeAlias,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/themes/{ThemeId}/aliases/{AliasName}",
	}

	if input == nil {
		input = &UpdateThemeAliasInput{}
	}

	req := c.newRequest(op, input, &UpdateThemeAliasOutput{})

	return UpdateThemeAliasRequest{Request: req, Input: input, Copy: c.UpdateThemeAliasRequest}
}

// UpdateThemeAliasRequest is the request type for the
// UpdateThemeAlias API operation.
type UpdateThemeAliasRequest struct {
	*aws.Request
	Input *UpdateThemeAliasInput
	Copy  func(*UpdateThemeAliasInput) UpdateThemeAliasRequest
}

// Send marshals and sends the UpdateThemeAlias API request.
func (r UpdateThemeAliasRequest) Send(ctx context.Context) (*UpdateThemeAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateThemeAliasResponse{
		UpdateThemeAliasOutput: r.Request.Data.(*UpdateThemeAliasOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateThemeAliasResponse is the response type for the
// UpdateThemeAlias API operation.
type UpdateThemeAliasResponse struct {
	*UpdateThemeAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateThemeAlias request.
func (r *UpdateThemeAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
