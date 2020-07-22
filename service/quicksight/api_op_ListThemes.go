// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListThemesInput struct {
	_ struct{} `type:"structure"`

	// The ID of the AWS account that contains the themes that you're listing.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The maximum number of results to be returned per request.
	MaxResults *int64 `location:"querystring" locationName:"max-results" min:"1" type:"integer"`

	// The token for the next set of results, or null if there are no more results.
	NextToken *string `location:"querystring" locationName:"next-token" type:"string"`

	// The type of themes that you want to list. Valid options include the following:
	//
	//    * ALL (default)- Display all existing themes.
	//
	//    * CUSTOM - Display only the themes created by people using Amazon QuickSight.
	//
	//    * QUICKSIGHT - Display only the starting themes defined by QuickSight.
	Type ThemeType `location:"querystring" locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListThemesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListThemesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListThemesInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListThemesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "max-results", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "next-token", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type ListThemesOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of results, or null if there are no more results.
	NextToken *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// Information about the themes in the list.
	ThemeSummaryList []ThemeSummary `type:"list"`
}

// String returns the string representation
func (s ListThemesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListThemesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ThemeSummaryList != nil {
		v := s.ThemeSummaryList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ThemeSummaryList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opListThemes = "ListThemes"

// ListThemesRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Lists all the themes in the current AWS account.
//
//    // Example sending a request using ListThemesRequest.
//    req := client.ListThemesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/ListThemes
func (c *Client) ListThemesRequest(input *ListThemesInput) ListThemesRequest {
	op := &aws.Operation{
		Name:       opListThemes,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/themes",
	}

	if input == nil {
		input = &ListThemesInput{}
	}

	req := c.newRequest(op, input, &ListThemesOutput{})

	return ListThemesRequest{Request: req, Input: input, Copy: c.ListThemesRequest}
}

// ListThemesRequest is the request type for the
// ListThemes API operation.
type ListThemesRequest struct {
	*aws.Request
	Input *ListThemesInput
	Copy  func(*ListThemesInput) ListThemesRequest
}

// Send marshals and sends the ListThemes API request.
func (r ListThemesRequest) Send(ctx context.Context) (*ListThemesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListThemesResponse{
		ListThemesOutput: r.Request.Data.(*ListThemesOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListThemesResponse is the response type for the
// ListThemes API operation.
type ListThemesResponse struct {
	*ListThemesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListThemes request.
func (r *ListThemesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
