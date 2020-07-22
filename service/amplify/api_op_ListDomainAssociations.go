// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request structure for the list domain associations request.
type ListDomainAssociationsInput struct {
	_ struct{} `type:"structure"`

	// The unique ID for an Amplify app.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// The maximum number of records to list in a single response.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// A pagination token. Set to null to start listing apps from the start. If
	// non-null, a pagination token is returned in a result. Pass its value in here
	// to list more projects.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDomainAssociationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDomainAssociationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListDomainAssociationsInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
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
func (s ListDomainAssociationsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
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

// The result structure for the list domain association request.
type ListDomainAssociationsOutput struct {
	_ struct{} `type:"structure"`

	// A list of domain associations.
	//
	// DomainAssociations is a required field
	DomainAssociations []DomainAssociation `locationName:"domainAssociations" type:"list" required:"true"`

	// A pagination token. If non-null, a pagination token is returned in a result.
	// Pass its value in another request to retrieve more entries.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListDomainAssociationsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListDomainAssociationsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DomainAssociations != nil {
		v := s.DomainAssociations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "domainAssociations", metadata)
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

const opListDomainAssociations = "ListDomainAssociations"

// ListDomainAssociationsRequest returns a request value for making API operation for
// AWS Amplify.
//
// Returns the domain associations for an Amplify app.
//
//    // Example sending a request using ListDomainAssociationsRequest.
//    req := client.ListDomainAssociationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/ListDomainAssociations
func (c *Client) ListDomainAssociationsRequest(input *ListDomainAssociationsInput) ListDomainAssociationsRequest {
	op := &aws.Operation{
		Name:       opListDomainAssociations,
		HTTPMethod: "GET",
		HTTPPath:   "/apps/{appId}/domains",
	}

	if input == nil {
		input = &ListDomainAssociationsInput{}
	}

	req := c.newRequest(op, input, &ListDomainAssociationsOutput{})

	return ListDomainAssociationsRequest{Request: req, Input: input, Copy: c.ListDomainAssociationsRequest}
}

// ListDomainAssociationsRequest is the request type for the
// ListDomainAssociations API operation.
type ListDomainAssociationsRequest struct {
	*aws.Request
	Input *ListDomainAssociationsInput
	Copy  func(*ListDomainAssociationsInput) ListDomainAssociationsRequest
}

// Send marshals and sends the ListDomainAssociations API request.
func (r ListDomainAssociationsRequest) Send(ctx context.Context) (*ListDomainAssociationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListDomainAssociationsResponse{
		ListDomainAssociationsOutput: r.Request.Data.(*ListDomainAssociationsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListDomainAssociationsResponse is the response type for the
// ListDomainAssociations API operation.
type ListDomainAssociationsResponse struct {
	*ListDomainAssociationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListDomainAssociations request.
func (r *ListDomainAssociationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
