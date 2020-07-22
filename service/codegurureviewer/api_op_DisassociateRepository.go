// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codegurureviewer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DisassociateRepositoryInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the RepositoryAssociation (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html)
	// object.
	//
	// AssociationArn is a required field
	AssociationArn *string `location:"uri" locationName:"AssociationArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateRepositoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateRepositoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateRepositoryInput"}

	if s.AssociationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociationArn"))
	}
	if s.AssociationArn != nil && len(*s.AssociationArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AssociationArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateRepositoryInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AssociationArn != nil {
		v := *s.AssociationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AssociationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DisassociateRepositoryOutput struct {
	_ struct{} `type:"structure"`

	// Information about the disassociated repository.
	RepositoryAssociation *RepositoryAssociation `type:"structure"`
}

// String returns the string representation
func (s DisassociateRepositoryOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisassociateRepositoryOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RepositoryAssociation != nil {
		v := s.RepositoryAssociation

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "RepositoryAssociation", v, metadata)
	}
	return nil
}

const opDisassociateRepository = "DisassociateRepository"

// DisassociateRepositoryRequest returns a request value for making API operation for
// Amazon CodeGuru Reviewer.
//
// Removes the association between Amazon CodeGuru Reviewer and a repository.
//
//    // Example sending a request using DisassociateRepositoryRequest.
//    req := client.DisassociateRepositoryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguru-reviewer-2019-09-19/DisassociateRepository
func (c *Client) DisassociateRepositoryRequest(input *DisassociateRepositoryInput) DisassociateRepositoryRequest {
	op := &aws.Operation{
		Name:       opDisassociateRepository,
		HTTPMethod: "DELETE",
		HTTPPath:   "/associations/{AssociationArn}",
	}

	if input == nil {
		input = &DisassociateRepositoryInput{}
	}

	req := c.newRequest(op, input, &DisassociateRepositoryOutput{})

	return DisassociateRepositoryRequest{Request: req, Input: input, Copy: c.DisassociateRepositoryRequest}
}

// DisassociateRepositoryRequest is the request type for the
// DisassociateRepository API operation.
type DisassociateRepositoryRequest struct {
	*aws.Request
	Input *DisassociateRepositoryInput
	Copy  func(*DisassociateRepositoryInput) DisassociateRepositoryRequest
}

// Send marshals and sends the DisassociateRepository API request.
func (r DisassociateRepositoryRequest) Send(ctx context.Context) (*DisassociateRepositoryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateRepositoryResponse{
		DisassociateRepositoryOutput: r.Request.Data.(*DisassociateRepositoryOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateRepositoryResponse is the response type for the
// DisassociateRepository API operation.
type DisassociateRepositoryResponse struct {
	*DisassociateRepositoryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateRepository request.
func (r *DisassociateRepositoryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
