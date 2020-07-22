// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type ListVirtualGatewaysInput struct {
	_ struct{} `type:"structure"`

	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	MeshOwner *string `location:"querystring" locationName:"meshOwner" min:"12" type:"string"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListVirtualGatewaysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVirtualGatewaysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVirtualGatewaysInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}
	if s.MeshOwner != nil && len(*s.MeshOwner) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshOwner", 12))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVirtualGatewaysInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.MeshOwner != nil {
		v := *s.MeshOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "meshOwner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListVirtualGatewaysOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `locationName:"nextToken" type:"string"`

	// VirtualGateways is a required field
	VirtualGateways []VirtualGatewayRef `locationName:"virtualGateways" type:"list" required:"true"`
}

// String returns the string representation
func (s ListVirtualGatewaysOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListVirtualGatewaysOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "nextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualGateways != nil {
		v := s.VirtualGateways

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "virtualGateways", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListVirtualGateways = "ListVirtualGateways"

// ListVirtualGatewaysRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Returns a list of existing virtual gateways in a service mesh.
//
//    // Example sending a request using ListVirtualGatewaysRequest.
//    req := client.ListVirtualGatewaysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/ListVirtualGateways
func (c *Client) ListVirtualGatewaysRequest(input *ListVirtualGatewaysInput) ListVirtualGatewaysRequest {
	op := &aws.Operation{
		Name:       opListVirtualGateways,
		HTTPMethod: "GET",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualGateways",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListVirtualGatewaysInput{}
	}

	req := c.newRequest(op, input, &ListVirtualGatewaysOutput{})

	return ListVirtualGatewaysRequest{Request: req, Input: input, Copy: c.ListVirtualGatewaysRequest}
}

// ListVirtualGatewaysRequest is the request type for the
// ListVirtualGateways API operation.
type ListVirtualGatewaysRequest struct {
	*aws.Request
	Input *ListVirtualGatewaysInput
	Copy  func(*ListVirtualGatewaysInput) ListVirtualGatewaysRequest
}

// Send marshals and sends the ListVirtualGateways API request.
func (r ListVirtualGatewaysRequest) Send(ctx context.Context) (*ListVirtualGatewaysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListVirtualGatewaysResponse{
		ListVirtualGatewaysOutput: r.Request.Data.(*ListVirtualGatewaysOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListVirtualGatewaysRequestPaginator returns a paginator for ListVirtualGateways.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListVirtualGatewaysRequest(input)
//   p := appmesh.NewListVirtualGatewaysRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListVirtualGatewaysPaginator(req ListVirtualGatewaysRequest) ListVirtualGatewaysPaginator {
	return ListVirtualGatewaysPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListVirtualGatewaysInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListVirtualGatewaysPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListVirtualGatewaysPaginator struct {
	aws.Pager
}

func (p *ListVirtualGatewaysPaginator) CurrentPage() *ListVirtualGatewaysOutput {
	return p.Pager.CurrentPage().(*ListVirtualGatewaysOutput)
}

// ListVirtualGatewaysResponse is the response type for the
// ListVirtualGateways API operation.
type ListVirtualGatewaysResponse struct {
	*ListVirtualGatewaysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListVirtualGateways request.
func (r *ListVirtualGatewaysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
