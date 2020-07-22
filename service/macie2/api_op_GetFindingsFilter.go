// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetFindingsFilterInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetFindingsFilterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFindingsFilterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFindingsFilterInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFindingsFilterInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Provides information about the criteria and other settings for a findings
// filter.
type GetFindingsFilterOutput struct {
	_ struct{} `type:"structure"`

	// The action to perform on findings that meet the filter criteria. To suppress
	// (automatically archive) findings that meet the criteria, set this value to
	// ARCHIVE. Valid values are:
	Action FindingsFilterAction `locationName:"action" type:"string" enum:"true"`

	Arn *string `locationName:"arn" type:"string"`

	Description *string `locationName:"description" type:"string"`

	// Specifies, as a map, one or more property-based conditions that filter the
	// results of a query for findings.
	FindingCriteria *FindingCriteria `locationName:"findingCriteria" type:"structure"`

	Id *string `locationName:"id" type:"string"`

	Name *string `locationName:"name" type:"string"`

	Position *int64 `locationName:"position" type:"integer"`

	// A string-to-string map of key-value pairs that specifies the tags (keys and
	// values) for a classification job, custom data identifier, findings filter,
	// or member account.
	Tags map[string]string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s GetFindingsFilterOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFindingsFilterOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Action) > 0 {
		v := s.Action

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "action", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FindingCriteria != nil {
		v := s.FindingCriteria

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "findingCriteria", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Position != nil {
		v := *s.Position

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "position", protocol.Int64Value(v), metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opGetFindingsFilter = "GetFindingsFilter"

// GetFindingsFilterRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Retrieves information about the criteria and other settings for a findings
// filter.
//
//    // Example sending a request using GetFindingsFilterRequest.
//    req := client.GetFindingsFilterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/GetFindingsFilter
func (c *Client) GetFindingsFilterRequest(input *GetFindingsFilterInput) GetFindingsFilterRequest {
	op := &aws.Operation{
		Name:       opGetFindingsFilter,
		HTTPMethod: "GET",
		HTTPPath:   "/findingsfilters/{id}",
	}

	if input == nil {
		input = &GetFindingsFilterInput{}
	}

	req := c.newRequest(op, input, &GetFindingsFilterOutput{})

	return GetFindingsFilterRequest{Request: req, Input: input, Copy: c.GetFindingsFilterRequest}
}

// GetFindingsFilterRequest is the request type for the
// GetFindingsFilter API operation.
type GetFindingsFilterRequest struct {
	*aws.Request
	Input *GetFindingsFilterInput
	Copy  func(*GetFindingsFilterInput) GetFindingsFilterRequest
}

// Send marshals and sends the GetFindingsFilter API request.
func (r GetFindingsFilterRequest) Send(ctx context.Context) (*GetFindingsFilterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFindingsFilterResponse{
		GetFindingsFilterOutput: r.Request.Data.(*GetFindingsFilterOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFindingsFilterResponse is the response type for the
// GetFindingsFilter API operation.
type GetFindingsFilterResponse struct {
	*GetFindingsFilterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFindingsFilter request.
func (r *GetFindingsFilterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
