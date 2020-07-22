// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdatePublicKeyInput struct {
	_ struct{} `type:"structure" payload:"PublicKeyConfig"`

	// ID of the public key to be updated.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// The value of the ETag header that you received when retrieving the public
	// key to update. For example: E2QWRUHAPOMQZL.
	IfMatch *string `location:"header" locationName:"If-Match" type:"string"`

	// Request to update public key information.
	//
	// PublicKeyConfig is a required field
	PublicKeyConfig *PublicKeyConfig `locationName:"PublicKeyConfig" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2020-05-31/"`
}

// String returns the string representation
func (s UpdatePublicKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePublicKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePublicKeyInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.PublicKeyConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("PublicKeyConfig"))
	}
	if s.PublicKeyConfig != nil {
		if err := s.PublicKeyConfig.Validate(); err != nil {
			invalidParams.AddNested("PublicKeyConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePublicKeyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.IfMatch != nil {
		v := *s.IfMatch

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "If-Match", protocol.StringValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	if s.PublicKeyConfig != nil {
		v := s.PublicKeyConfig

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2020-05-31/"}
		e.SetFields(protocol.PayloadTarget, "PublicKeyConfig", v, metadata)
	}
	return nil
}

type UpdatePublicKeyOutput struct {
	_ struct{} `type:"structure" payload:"PublicKey"`

	// The current version of the update public key result. For example: E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// Return the results of updating the public key.
	PublicKey *PublicKey `type:"structure"`
}

// String returns the string representation
func (s UpdatePublicKeyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePublicKeyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.PublicKey != nil {
		v := s.PublicKey

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "PublicKey", v, metadata)
	}
	return nil
}

const opUpdatePublicKey = "UpdatePublicKey2020_05_31"

// UpdatePublicKeyRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Update public key information. Note that the only value you can change is
// the comment.
//
//    // Example sending a request using UpdatePublicKeyRequest.
//    req := client.UpdatePublicKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2020-05-31/UpdatePublicKey
func (c *Client) UpdatePublicKeyRequest(input *UpdatePublicKeyInput) UpdatePublicKeyRequest {
	op := &aws.Operation{
		Name:       opUpdatePublicKey,
		HTTPMethod: "PUT",
		HTTPPath:   "/2020-05-31/public-key/{Id}/config",
	}

	if input == nil {
		input = &UpdatePublicKeyInput{}
	}

	req := c.newRequest(op, input, &UpdatePublicKeyOutput{})

	return UpdatePublicKeyRequest{Request: req, Input: input, Copy: c.UpdatePublicKeyRequest}
}

// UpdatePublicKeyRequest is the request type for the
// UpdatePublicKey API operation.
type UpdatePublicKeyRequest struct {
	*aws.Request
	Input *UpdatePublicKeyInput
	Copy  func(*UpdatePublicKeyInput) UpdatePublicKeyRequest
}

// Send marshals and sends the UpdatePublicKey API request.
func (r UpdatePublicKeyRequest) Send(ctx context.Context) (*UpdatePublicKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePublicKeyResponse{
		UpdatePublicKeyOutput: r.Request.Data.(*UpdatePublicKeyOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePublicKeyResponse is the response type for the
// UpdatePublicKey API operation.
type UpdatePublicKeyResponse struct {
	*UpdatePublicKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePublicKey request.
func (r *UpdatePublicKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
