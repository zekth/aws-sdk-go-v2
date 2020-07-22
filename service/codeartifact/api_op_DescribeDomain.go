// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeartifact

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeDomainInput struct {
	_ struct{} `type:"structure"`

	// A string that specifies the name of the requested domain.
	//
	// Domain is a required field
	Domain *string `location:"querystring" locationName:"domain" min:"2" type:"string" required:"true"`

	// The 12-digit account number of the AWS account that owns the domain. It does
	// not include dashes or spaces.
	DomainOwner *string `location:"querystring" locationName:"domain-owner" min:"12" type:"string"`
}

// String returns the string representation
func (s DescribeDomainInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDomainInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDomainInput"}

	if s.Domain == nil {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}
	if s.Domain != nil && len(*s.Domain) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Domain", 2))
	}
	if s.DomainOwner != nil && len(*s.DomainOwner) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainOwner", 12))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDomainInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Domain != nil {
		v := *s.Domain

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "domain", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DomainOwner != nil {
		v := *s.DomainOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "domain-owner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeDomainOutput struct {
	_ struct{} `type:"structure"`

	// Information about a domain. A domain is a container for repositories. When
	// you create a domain, it is empty until you add one or more repositories.
	Domain *DomainDescription `locationName:"domain" type:"structure"`
}

// String returns the string representation
func (s DescribeDomainOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeDomainOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Domain != nil {
		v := s.Domain

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "domain", v, metadata)
	}
	return nil
}

const opDescribeDomain = "DescribeDomain"

// DescribeDomainRequest returns a request value for making API operation for
// CodeArtifact.
//
// Returns a DomainDescription (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DomainDescription.html)
// object that contains information about the requested domain.
//
//    // Example sending a request using DescribeDomainRequest.
//    req := client.DescribeDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeartifact-2018-09-22/DescribeDomain
func (c *Client) DescribeDomainRequest(input *DescribeDomainInput) DescribeDomainRequest {
	op := &aws.Operation{
		Name:       opDescribeDomain,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/domain",
	}

	if input == nil {
		input = &DescribeDomainInput{}
	}

	req := c.newRequest(op, input, &DescribeDomainOutput{})

	return DescribeDomainRequest{Request: req, Input: input, Copy: c.DescribeDomainRequest}
}

// DescribeDomainRequest is the request type for the
// DescribeDomain API operation.
type DescribeDomainRequest struct {
	*aws.Request
	Input *DescribeDomainInput
	Copy  func(*DescribeDomainInput) DescribeDomainRequest
}

// Send marshals and sends the DescribeDomain API request.
func (r DescribeDomainRequest) Send(ctx context.Context) (*DescribeDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDomainResponse{
		DescribeDomainOutput: r.Request.Data.(*DescribeDomainOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDomainResponse is the response type for the
// DescribeDomain API operation.
type DescribeDomainResponse struct {
	*DescribeDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDomain request.
func (r *DescribeDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
