// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteGatewayRouteInput struct {
	_ struct{} `type:"structure"`

	// GatewayRouteName is a required field
	GatewayRouteName *string `location:"uri" locationName:"gatewayRouteName" min:"1" type:"string" required:"true"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	MeshOwner *string `location:"querystring" locationName:"meshOwner" min:"12" type:"string"`

	// VirtualGatewayName is a required field
	VirtualGatewayName *string `location:"uri" locationName:"virtualGatewayName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteGatewayRouteInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteGatewayRouteInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteGatewayRouteInput"}

	if s.GatewayRouteName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayRouteName"))
	}
	if s.GatewayRouteName != nil && len(*s.GatewayRouteName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayRouteName", 1))
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

	if s.VirtualGatewayName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualGatewayName"))
	}
	if s.VirtualGatewayName != nil && len(*s.VirtualGatewayName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualGatewayName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGatewayRouteInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.GatewayRouteName != nil {
		v := *s.GatewayRouteName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "gatewayRouteName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualGatewayName != nil {
		v := *s.VirtualGatewayName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "virtualGatewayName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeshOwner != nil {
		v := *s.MeshOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "meshOwner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteGatewayRouteOutput struct {
	_ struct{} `type:"structure" payload:"GatewayRoute"`

	// An object that represents a gateway route returned by a describe operation.
	//
	// GatewayRoute is a required field
	GatewayRoute *GatewayRouteData `locationName:"gatewayRoute" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteGatewayRouteOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteGatewayRouteOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.GatewayRoute != nil {
		v := s.GatewayRoute

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "gatewayRoute", v, metadata)
	}
	return nil
}

const opDeleteGatewayRoute = "DeleteGatewayRoute"

// DeleteGatewayRouteRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Deletes an existing gateway route.
//
//    // Example sending a request using DeleteGatewayRouteRequest.
//    req := client.DeleteGatewayRouteRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DeleteGatewayRoute
func (c *Client) DeleteGatewayRouteRequest(input *DeleteGatewayRouteInput) DeleteGatewayRouteRequest {
	op := &aws.Operation{
		Name:       opDeleteGatewayRoute,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualGateway/{virtualGatewayName}/gatewayRoutes/{gatewayRouteName}",
	}

	if input == nil {
		input = &DeleteGatewayRouteInput{}
	}

	req := c.newRequest(op, input, &DeleteGatewayRouteOutput{})

	return DeleteGatewayRouteRequest{Request: req, Input: input, Copy: c.DeleteGatewayRouteRequest}
}

// DeleteGatewayRouteRequest is the request type for the
// DeleteGatewayRoute API operation.
type DeleteGatewayRouteRequest struct {
	*aws.Request
	Input *DeleteGatewayRouteInput
	Copy  func(*DeleteGatewayRouteInput) DeleteGatewayRouteRequest
}

// Send marshals and sends the DeleteGatewayRoute API request.
func (r DeleteGatewayRouteRequest) Send(ctx context.Context) (*DeleteGatewayRouteResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGatewayRouteResponse{
		DeleteGatewayRouteOutput: r.Request.Data.(*DeleteGatewayRouteOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGatewayRouteResponse is the response type for the
// DeleteGatewayRoute API operation.
type DeleteGatewayRouteResponse struct {
	*DeleteGatewayRouteOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGatewayRoute request.
func (r *DeleteGatewayRouteResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
