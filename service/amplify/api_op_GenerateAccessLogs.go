// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request structure for the generate access logs request.
type GenerateAccessLogsInput struct {
	_ struct{} `type:"structure"`

	// The unique ID for an Amplify app.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// The name of the domain.
	//
	// DomainName is a required field
	DomainName *string `locationName:"domainName" type:"string" required:"true"`

	// The time at which the logs should end. The time range specified is inclusive
	// of the end time.
	EndTime *time.Time `locationName:"endTime" type:"timestamp"`

	// The time at which the logs should start. The time range specified is inclusive
	// of the start time.
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`
}

// String returns the string representation
func (s GenerateAccessLogsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateAccessLogsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GenerateAccessLogsInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GenerateAccessLogsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DomainName != nil {
		v := *s.DomainName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "domainName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.EndTime != nil {
		v := *s.EndTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "endTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.StartTime != nil {
		v := *s.StartTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result structure for the generate access logs request.
type GenerateAccessLogsOutput struct {
	_ struct{} `type:"structure"`

	// The pre-signed URL for the requested access logs.
	LogUrl *string `locationName:"logUrl" type:"string"`
}

// String returns the string representation
func (s GenerateAccessLogsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GenerateAccessLogsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.LogUrl != nil {
		v := *s.LogUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "logUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGenerateAccessLogs = "GenerateAccessLogs"

// GenerateAccessLogsRequest returns a request value for making API operation for
// AWS Amplify.
//
// Returns the website access logs for a specific time range using a presigned
// URL.
//
//    // Example sending a request using GenerateAccessLogsRequest.
//    req := client.GenerateAccessLogsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GenerateAccessLogs
func (c *Client) GenerateAccessLogsRequest(input *GenerateAccessLogsInput) GenerateAccessLogsRequest {
	op := &aws.Operation{
		Name:       opGenerateAccessLogs,
		HTTPMethod: "POST",
		HTTPPath:   "/apps/{appId}/accesslogs",
	}

	if input == nil {
		input = &GenerateAccessLogsInput{}
	}

	req := c.newRequest(op, input, &GenerateAccessLogsOutput{})

	return GenerateAccessLogsRequest{Request: req, Input: input, Copy: c.GenerateAccessLogsRequest}
}

// GenerateAccessLogsRequest is the request type for the
// GenerateAccessLogs API operation.
type GenerateAccessLogsRequest struct {
	*aws.Request
	Input *GenerateAccessLogsInput
	Copy  func(*GenerateAccessLogsInput) GenerateAccessLogsRequest
}

// Send marshals and sends the GenerateAccessLogs API request.
func (r GenerateAccessLogsRequest) Send(ctx context.Context) (*GenerateAccessLogsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GenerateAccessLogsResponse{
		GenerateAccessLogsOutput: r.Request.Data.(*GenerateAccessLogsOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GenerateAccessLogsResponse is the response type for the
// GenerateAccessLogs API operation.
type GenerateAccessLogsResponse struct {
	*GenerateAccessLogsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GenerateAccessLogs request.
func (r *GenerateAccessLogsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
