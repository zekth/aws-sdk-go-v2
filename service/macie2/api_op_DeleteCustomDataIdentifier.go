// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteCustomDataIdentifierInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCustomDataIdentifierInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCustomDataIdentifierInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCustomDataIdentifierInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteCustomDataIdentifierInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteCustomDataIdentifierOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCustomDataIdentifierOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteCustomDataIdentifierOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteCustomDataIdentifier = "DeleteCustomDataIdentifier"

// DeleteCustomDataIdentifierRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Soft deletes a custom data identifier.
//
//    // Example sending a request using DeleteCustomDataIdentifierRequest.
//    req := client.DeleteCustomDataIdentifierRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/DeleteCustomDataIdentifier
func (c *Client) DeleteCustomDataIdentifierRequest(input *DeleteCustomDataIdentifierInput) DeleteCustomDataIdentifierRequest {
	op := &aws.Operation{
		Name:       opDeleteCustomDataIdentifier,
		HTTPMethod: "DELETE",
		HTTPPath:   "/custom-data-identifiers/{id}",
	}

	if input == nil {
		input = &DeleteCustomDataIdentifierInput{}
	}

	req := c.newRequest(op, input, &DeleteCustomDataIdentifierOutput{})

	return DeleteCustomDataIdentifierRequest{Request: req, Input: input, Copy: c.DeleteCustomDataIdentifierRequest}
}

// DeleteCustomDataIdentifierRequest is the request type for the
// DeleteCustomDataIdentifier API operation.
type DeleteCustomDataIdentifierRequest struct {
	*aws.Request
	Input *DeleteCustomDataIdentifierInput
	Copy  func(*DeleteCustomDataIdentifierInput) DeleteCustomDataIdentifierRequest
}

// Send marshals and sends the DeleteCustomDataIdentifier API request.
func (r DeleteCustomDataIdentifierRequest) Send(ctx context.Context) (*DeleteCustomDataIdentifierResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCustomDataIdentifierResponse{
		DeleteCustomDataIdentifierOutput: r.Request.Data.(*DeleteCustomDataIdentifierOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCustomDataIdentifierResponse is the response type for the
// DeleteCustomDataIdentifier API operation.
type DeleteCustomDataIdentifierResponse struct {
	*DeleteCustomDataIdentifierOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCustomDataIdentifier request.
func (r *DeleteCustomDataIdentifierResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
