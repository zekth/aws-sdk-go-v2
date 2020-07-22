// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreatePublicKeyInput struct {
	_ struct{} `type:"structure" payload:"PublicKeyConfig"`

	// The request to add a public key to CloudFront.
	//
	// PublicKeyConfig is a required field
	PublicKeyConfig *PublicKeyConfig `locationName:"PublicKeyConfig" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2020-05-31/"`
}

// String returns the string representation
func (s CreatePublicKeyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePublicKeyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePublicKeyInput"}

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
func (s CreatePublicKeyInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.PublicKeyConfig != nil {
		v := s.PublicKeyConfig

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2020-05-31/"}
		e.SetFields(protocol.PayloadTarget, "PublicKeyConfig", v, metadata)
	}
	return nil
}

type CreatePublicKeyOutput struct {
	_ struct{} `type:"structure" payload:"PublicKey"`

	// The current version of the public key. For example: E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// The fully qualified URI of the new public key resource just created.
	Location *string `location:"header" locationName:"Location" type:"string"`

	// Returned when you add a public key.
	PublicKey *PublicKey `type:"structure"`
}

// String returns the string representation
func (s CreatePublicKeyOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreatePublicKeyOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.Location != nil {
		v := *s.Location

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Location", protocol.StringValue(v), metadata)
	}
	if s.PublicKey != nil {
		v := s.PublicKey

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "PublicKey", v, metadata)
	}
	return nil
}

const opCreatePublicKey = "CreatePublicKey2020_05_31"

// CreatePublicKeyRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Add a new public key to CloudFront to use, for example, for field-level encryption.
// You can add a maximum of 10 public keys with one AWS account.
//
//    // Example sending a request using CreatePublicKeyRequest.
//    req := client.CreatePublicKeyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2020-05-31/CreatePublicKey
func (c *Client) CreatePublicKeyRequest(input *CreatePublicKeyInput) CreatePublicKeyRequest {
	op := &aws.Operation{
		Name:       opCreatePublicKey,
		HTTPMethod: "POST",
		HTTPPath:   "/2020-05-31/public-key",
	}

	if input == nil {
		input = &CreatePublicKeyInput{}
	}

	req := c.newRequest(op, input, &CreatePublicKeyOutput{})

	return CreatePublicKeyRequest{Request: req, Input: input, Copy: c.CreatePublicKeyRequest}
}

// CreatePublicKeyRequest is the request type for the
// CreatePublicKey API operation.
type CreatePublicKeyRequest struct {
	*aws.Request
	Input *CreatePublicKeyInput
	Copy  func(*CreatePublicKeyInput) CreatePublicKeyRequest
}

// Send marshals and sends the CreatePublicKey API request.
func (r CreatePublicKeyRequest) Send(ctx context.Context) (*CreatePublicKeyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePublicKeyResponse{
		CreatePublicKeyOutput: r.Request.Data.(*CreatePublicKeyOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePublicKeyResponse is the response type for the
// CreatePublicKey API operation.
type CreatePublicKeyResponse struct {
	*CreatePublicKeyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePublicKey request.
func (r *CreatePublicKeyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
