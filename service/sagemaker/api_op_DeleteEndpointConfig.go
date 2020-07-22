// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

type DeleteEndpointConfigInput struct {
	_ struct{} `type:"structure"`

	// The name of the endpoint configuration that you want to delete.
	//
	// EndpointConfigName is a required field
	EndpointConfigName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteEndpointConfigInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEndpointConfigInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteEndpointConfigInput"}

	if s.EndpointConfigName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointConfigName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteEndpointConfigOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteEndpointConfigOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteEndpointConfig = "DeleteEndpointConfig"

// DeleteEndpointConfigRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Deletes an endpoint configuration. The DeleteEndpointConfig API deletes only
// the specified configuration. It does not delete endpoints created using the
// configuration.
//
// You must not delete an EndpointConfig in use by an endpoint that is live
// or while the UpdateEndpoint or CreateEndpoint operations are being performed
// on the endpoint. If you delete the EndpointConfig of an endpoint that is
// active or being created or updated you may lose visibility into the instance
// type the endpoint is using. The endpoint must be deleted in order to stop
// incurring charges.
//
//    // Example sending a request using DeleteEndpointConfigRequest.
//    req := client.DeleteEndpointConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteEndpointConfig
func (c *Client) DeleteEndpointConfigRequest(input *DeleteEndpointConfigInput) DeleteEndpointConfigRequest {
	op := &aws.Operation{
		Name:       opDeleteEndpointConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteEndpointConfigInput{}
	}

	req := c.newRequest(op, input, &DeleteEndpointConfigOutput{})
	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteEndpointConfigRequest{Request: req, Input: input, Copy: c.DeleteEndpointConfigRequest}
}

// DeleteEndpointConfigRequest is the request type for the
// DeleteEndpointConfig API operation.
type DeleteEndpointConfigRequest struct {
	*aws.Request
	Input *DeleteEndpointConfigInput
	Copy  func(*DeleteEndpointConfigInput) DeleteEndpointConfigRequest
}

// Send marshals and sends the DeleteEndpointConfig API request.
func (r DeleteEndpointConfigRequest) Send(ctx context.Context) (*DeleteEndpointConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteEndpointConfigResponse{
		DeleteEndpointConfigOutput: r.Request.Data.(*DeleteEndpointConfigOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteEndpointConfigResponse is the response type for the
// DeleteEndpointConfig API operation.
type DeleteEndpointConfigResponse struct {
	*DeleteEndpointConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteEndpointConfig request.
func (r *DeleteEndpointConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
