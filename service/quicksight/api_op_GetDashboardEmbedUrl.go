// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetDashboardEmbedUrlInput struct {
	_ struct{} `type:"structure"`

	// The ID for the AWS account that contains the dashboard that you're embedding.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the dashboard, also added to the IAM policy.
	//
	// DashboardId is a required field
	DashboardId *string `location:"uri" locationName:"DashboardId" min:"1" type:"string" required:"true"`

	// The authentication method that the user uses to sign in.
	//
	// IdentityType is a required field
	IdentityType IdentityType `location:"querystring" locationName:"creds-type" type:"string" required:"true" enum:"true"`

	// Remove the reset button on the embedded dashboard. The default is FALSE,
	// which enables the reset button.
	ResetDisabled *bool `location:"querystring" locationName:"reset-disabled" type:"boolean"`

	// How many minutes the session is valid. The session lifetime must be 15-600
	// minutes.
	SessionLifetimeInMinutes *int64 `location:"querystring" locationName:"session-lifetime" min:"15" type:"long"`

	// Remove the undo/redo button on the embedded dashboard. The default is FALSE,
	// which enables the undo/redo button.
	UndoRedoDisabled *bool `location:"querystring" locationName:"undo-redo-disabled" type:"boolean"`

	// The Amazon QuickSight user's Amazon Resource Name (ARN), for use with QUICKSIGHT
	// identity type. You can use this for any Amazon QuickSight users in your account
	// (readers, authors, or admins) authenticated as one of the following:
	//
	//    * Active Directory (AD) users or group members
	//
	//    * Invited nonfederated users
	//
	//    * IAM users and IAM role-based sessions authenticated through Federated
	//    Single Sign-On using SAML, OpenID Connect, or IAM federation.
	UserArn *string `location:"querystring" locationName:"user-arn" type:"string"`
}

// String returns the string representation
func (s GetDashboardEmbedUrlInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDashboardEmbedUrlInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDashboardEmbedUrlInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DashboardId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardId"))
	}
	if s.DashboardId != nil && len(*s.DashboardId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DashboardId", 1))
	}
	if len(s.IdentityType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("IdentityType"))
	}
	if s.SessionLifetimeInMinutes != nil && *s.SessionLifetimeInMinutes < 15 {
		invalidParams.Add(aws.NewErrParamMinValue("SessionLifetimeInMinutes", 15))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDashboardEmbedUrlInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DashboardId != nil {
		v := *s.DashboardId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DashboardId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.IdentityType) > 0 {
		v := s.IdentityType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "creds-type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ResetDisabled != nil {
		v := *s.ResetDisabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "reset-disabled", protocol.BoolValue(v), metadata)
	}
	if s.SessionLifetimeInMinutes != nil {
		v := *s.SessionLifetimeInMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "session-lifetime", protocol.Int64Value(v), metadata)
	}
	if s.UndoRedoDisabled != nil {
		v := *s.UndoRedoDisabled

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "undo-redo-disabled", protocol.BoolValue(v), metadata)
	}
	if s.UserArn != nil {
		v := *s.UserArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "user-arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetDashboardEmbedUrlOutput struct {
	_ struct{} `type:"structure"`

	// A single-use URL that you can put into your server-side webpage to embed
	// your dashboard. This URL is valid for 5 minutes. The API provides the URL
	// with an auth_code value that enables one (and only one) sign-on to a user
	// session that is valid for 10 hours.
	EmbedUrl *string `type:"string" sensitive:"true"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s GetDashboardEmbedUrlOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDashboardEmbedUrlOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.EmbedUrl != nil {
		v := *s.EmbedUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EmbedUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opGetDashboardEmbedUrl = "GetDashboardEmbedUrl"

// GetDashboardEmbedUrlRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Generates a URL and authorization code that you can embed in your web server
// code. Before you use this command, make sure that you have configured the
// dashboards and permissions.
//
// Currently, you can use GetDashboardEmbedURL only from the server, not from
// the user's browser. The following rules apply to the combination of URL and
// authorization code:
//
//    * They must be used together.
//
//    * They can be used one time only.
//
//    * They are valid for 5 minutes after you run this command.
//
//    * The resulting user session is valid for 10 hours.
//
// For more information, see Embedding Amazon QuickSight Dashboards (https://docs.aws.amazon.com/quicksight/latest/user/embedding-dashboards.html)
// in the Amazon QuickSight User Guide or Embedding Amazon QuickSight Dashboards
// (https://docs.aws.amazon.com/quicksight/latest/APIReference/qs-dev-embedded-dashboards.html)
// in the Amazon QuickSight API Reference.
//
//    // Example sending a request using GetDashboardEmbedUrlRequest.
//    req := client.GetDashboardEmbedUrlRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/GetDashboardEmbedUrl
func (c *Client) GetDashboardEmbedUrlRequest(input *GetDashboardEmbedUrlInput) GetDashboardEmbedUrlRequest {
	op := &aws.Operation{
		Name:       opGetDashboardEmbedUrl,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/dashboards/{DashboardId}/embed-url",
	}

	if input == nil {
		input = &GetDashboardEmbedUrlInput{}
	}

	req := c.newRequest(op, input, &GetDashboardEmbedUrlOutput{})

	return GetDashboardEmbedUrlRequest{Request: req, Input: input, Copy: c.GetDashboardEmbedUrlRequest}
}

// GetDashboardEmbedUrlRequest is the request type for the
// GetDashboardEmbedUrl API operation.
type GetDashboardEmbedUrlRequest struct {
	*aws.Request
	Input *GetDashboardEmbedUrlInput
	Copy  func(*GetDashboardEmbedUrlInput) GetDashboardEmbedUrlRequest
}

// Send marshals and sends the GetDashboardEmbedUrl API request.
func (r GetDashboardEmbedUrlRequest) Send(ctx context.Context) (*GetDashboardEmbedUrlResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDashboardEmbedUrlResponse{
		GetDashboardEmbedUrlOutput: r.Request.Data.(*GetDashboardEmbedUrlOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDashboardEmbedUrlResponse is the response type for the
// GetDashboardEmbedUrl API operation.
type GetDashboardEmbedUrlResponse struct {
	*GetDashboardEmbedUrlOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDashboardEmbedUrl request.
func (r *GetDashboardEmbedUrlResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
