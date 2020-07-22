// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request structure for the list apps request.
type ListAppsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of records to list in a single response.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// A pagination token. If non-null, the pagination token is returned in a result.
	// Pass its value in another request to retrieve more entries.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListAppsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAppsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAppsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAppsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "maxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The result structure for an Amplify app list request.
type ListAppsOutput struct {
	_ struct{} `type:"structure"`

	// A list of Amplify apps.
	//
	// Apps is a required field
	Apps []App `locationName:"apps" type:"list" required:"true"`

	// A pagination token. Set to null to start listing apps from start. If non-null,
	// the pagination token is returned in a result. Pass its value in here to list
	// more projects.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListAppsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListAppsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Apps != nil {
		v := s.Apps

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "apps", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListApps = "ListApps"

// ListAppsRequest returns a request value for making API operation for
// AWS Amplify.
//
// Returns a list of the existing Amplify apps.
//
//    // Example sending a request using ListAppsRequest.
//    req := client.ListAppsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/ListApps
func (c *Client) ListAppsRequest(input *ListAppsInput) ListAppsRequest {
	op := &aws.Operation{
		Name:       opListApps,
		HTTPMethod: "GET",
		HTTPPath:   "/apps",
	}

	if input == nil {
		input = &ListAppsInput{}
	}

	req := c.newRequest(op, input, &ListAppsOutput{})

	return ListAppsRequest{Request: req, Input: input, Copy: c.ListAppsRequest}
}

// ListAppsRequest is the request type for the
// ListApps API operation.
type ListAppsRequest struct {
	*aws.Request
	Input *ListAppsInput
	Copy  func(*ListAppsInput) ListAppsRequest
}

// Send marshals and sends the ListApps API request.
func (r ListAppsRequest) Send(ctx context.Context) (*ListAppsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAppsResponse{
		ListAppsOutput: r.Request.Data.(*ListAppsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListAppsResponse is the response type for the
// ListApps API operation.
type ListAppsResponse struct {
	*ListAppsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListApps request.
func (r *ListAppsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
