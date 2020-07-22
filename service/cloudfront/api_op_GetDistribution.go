// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to get a distribution's information.
type GetDistributionInput struct {
	_ struct{} `type:"structure"`

	// The distribution's ID. If the ID is empty, an empty distribution configuration
	// is returned.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDistributionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDistributionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDistributionInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDistributionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Id", protocol.StringValue(v), metadata)
	}
	return nil
}

// The returned result of the corresponding request.
type GetDistributionOutput struct {
	_ struct{} `type:"structure" payload:"Distribution"`

	// The distribution's information.
	Distribution *Distribution `type:"structure"`

	// The current version of the distribution's information. For example: E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`
}

// String returns the string representation
func (s GetDistributionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDistributionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ETag != nil {
		v := *s.ETag

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "ETag", protocol.StringValue(v), metadata)
	}
	if s.Distribution != nil {
		v := s.Distribution

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "Distribution", v, metadata)
	}
	return nil
}

const opGetDistribution = "GetDistribution2020_05_31"

// GetDistributionRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Get the information about a distribution.
//
//    // Example sending a request using GetDistributionRequest.
//    req := client.GetDistributionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2020-05-31/GetDistribution
func (c *Client) GetDistributionRequest(input *GetDistributionInput) GetDistributionRequest {
	op := &aws.Operation{
		Name:       opGetDistribution,
		HTTPMethod: "GET",
		HTTPPath:   "/2020-05-31/distribution/{Id}",
	}

	if input == nil {
		input = &GetDistributionInput{}
	}

	req := c.newRequest(op, input, &GetDistributionOutput{})

	return GetDistributionRequest{Request: req, Input: input, Copy: c.GetDistributionRequest}
}

// GetDistributionRequest is the request type for the
// GetDistribution API operation.
type GetDistributionRequest struct {
	*aws.Request
	Input *GetDistributionInput
	Copy  func(*GetDistributionInput) GetDistributionRequest
}

// Send marshals and sends the GetDistribution API request.
func (r GetDistributionRequest) Send(ctx context.Context) (*GetDistributionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDistributionResponse{
		GetDistributionOutput: r.Request.Data.(*GetDistributionOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDistributionResponse is the response type for the
// GetDistribution API operation.
type GetDistributionResponse struct {
	*GetDistributionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDistribution request.
func (r *GetDistributionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
