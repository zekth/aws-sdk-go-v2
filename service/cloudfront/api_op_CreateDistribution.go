// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to create a new distribution.
type CreateDistributionInput struct {
	_ struct{} `type:"structure" payload:"DistributionConfig"`

	// The distribution's configuration information.
	//
	// DistributionConfig is a required field
	DistributionConfig *DistributionConfig `locationName:"DistributionConfig" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2020-05-31/"`
}

// String returns the string representation
func (s CreateDistributionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDistributionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDistributionInput"}

	if s.DistributionConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("DistributionConfig"))
	}
	if s.DistributionConfig != nil {
		if err := s.DistributionConfig.Validate(); err != nil {
			invalidParams.AddNested("DistributionConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDistributionInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.DistributionConfig != nil {
		v := s.DistributionConfig

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2020-05-31/"}
		e.SetFields(protocol.PayloadTarget, "DistributionConfig", v, metadata)
	}
	return nil
}

// The returned result of the corresponding request.
type CreateDistributionOutput struct {
	_ struct{} `type:"structure" payload:"Distribution"`

	// The distribution's information.
	Distribution *Distribution `type:"structure"`

	// The current version of the distribution created.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// The fully qualified URI of the new distribution resource just created.
	Location *string `location:"header" locationName:"Location" type:"string"`
}

// String returns the string representation
func (s CreateDistributionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateDistributionOutput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.Distribution != nil {
		v := s.Distribution

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "Distribution", v, metadata)
	}
	return nil
}

const opCreateDistribution = "CreateDistribution2020_05_31"

// CreateDistributionRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Creates a new web distribution. You create a CloudFront distribution to tell
// CloudFront where you want content to be delivered from, and the details about
// how to track and manage content delivery. Send a POST request to the /CloudFront
// API version/distribution/distribution ID resource.
//
// When you update a distribution, there are more required fields than when
// you create a distribution. When you update your distribution by using UpdateDistribution
// (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html),
// follow the steps included in the documentation to get the current configuration
// and then make your updates. This helps to make sure that you include all
// of the required fields. To view a summary, see Required Fields for Create
// Distribution and Update Distribution (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-overview-required-fields.html)
// in the Amazon CloudFront Developer Guide.
//
//    // Example sending a request using CreateDistributionRequest.
//    req := client.CreateDistributionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2020-05-31/CreateDistribution
func (c *Client) CreateDistributionRequest(input *CreateDistributionInput) CreateDistributionRequest {
	op := &aws.Operation{
		Name:       opCreateDistribution,
		HTTPMethod: "POST",
		HTTPPath:   "/2020-05-31/distribution",
	}

	if input == nil {
		input = &CreateDistributionInput{}
	}

	req := c.newRequest(op, input, &CreateDistributionOutput{})

	return CreateDistributionRequest{Request: req, Input: input, Copy: c.CreateDistributionRequest}
}

// CreateDistributionRequest is the request type for the
// CreateDistribution API operation.
type CreateDistributionRequest struct {
	*aws.Request
	Input *CreateDistributionInput
	Copy  func(*CreateDistributionInput) CreateDistributionRequest
}

// Send marshals and sends the CreateDistribution API request.
func (r CreateDistributionRequest) Send(ctx context.Context) (*CreateDistributionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDistributionResponse{
		CreateDistributionOutput: r.Request.Data.(*CreateDistributionOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDistributionResponse is the response type for the
// CreateDistribution API operation.
type CreateDistributionResponse struct {
	*CreateDistributionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDistribution request.
func (r *CreateDistributionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
