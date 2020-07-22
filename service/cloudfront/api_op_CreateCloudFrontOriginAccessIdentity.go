// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request to create a new origin access identity (OAI). An origin access
// identity is a special CloudFront user that you can associate with Amazon
// S3 origins, so that you can secure all or just some of your Amazon S3 content.
// For more information, see Restricting Access to Amazon S3 Content by Using
// an Origin Access Identity (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html)
// in the Amazon CloudFront Developer Guide.
type CreateCloudFrontOriginAccessIdentityInput struct {
	_ struct{} `type:"structure" payload:"CloudFrontOriginAccessIdentityConfig"`

	// The current configuration information for the identity.
	//
	// CloudFrontOriginAccessIdentityConfig is a required field
	CloudFrontOriginAccessIdentityConfig *CloudFrontOriginAccessIdentityConfig `locationName:"CloudFrontOriginAccessIdentityConfig" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2020-05-31/"`
}

// String returns the string representation
func (s CreateCloudFrontOriginAccessIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCloudFrontOriginAccessIdentityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCloudFrontOriginAccessIdentityInput"}

	if s.CloudFrontOriginAccessIdentityConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("CloudFrontOriginAccessIdentityConfig"))
	}
	if s.CloudFrontOriginAccessIdentityConfig != nil {
		if err := s.CloudFrontOriginAccessIdentityConfig.Validate(); err != nil {
			invalidParams.AddNested("CloudFrontOriginAccessIdentityConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateCloudFrontOriginAccessIdentityInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.CloudFrontOriginAccessIdentityConfig != nil {
		v := s.CloudFrontOriginAccessIdentityConfig

		metadata := protocol.Metadata{XMLNamespaceURI: "http://cloudfront.amazonaws.com/doc/2020-05-31/"}
		e.SetFields(protocol.PayloadTarget, "CloudFrontOriginAccessIdentityConfig", v, metadata)
	}
	return nil
}

// The returned result of the corresponding request.
type CreateCloudFrontOriginAccessIdentityOutput struct {
	_ struct{} `type:"structure" payload:"CloudFrontOriginAccessIdentity"`

	// The origin access identity's information.
	CloudFrontOriginAccessIdentity *CloudFrontOriginAccessIdentity `type:"structure"`

	// The current version of the origin access identity created.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// The fully qualified URI of the new origin access identity just created.
	Location *string `location:"header" locationName:"Location" type:"string"`
}

// String returns the string representation
func (s CreateCloudFrontOriginAccessIdentityOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateCloudFrontOriginAccessIdentityOutput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.CloudFrontOriginAccessIdentity != nil {
		v := s.CloudFrontOriginAccessIdentity

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "CloudFrontOriginAccessIdentity", v, metadata)
	}
	return nil
}

const opCreateCloudFrontOriginAccessIdentity = "CreateCloudFrontOriginAccessIdentity2020_05_31"

// CreateCloudFrontOriginAccessIdentityRequest returns a request value for making API operation for
// Amazon CloudFront.
//
// Creates a new origin access identity. If you're using Amazon S3 for your
// origin, you can use an origin access identity to require users to access
// your content using a CloudFront URL instead of the Amazon S3 URL. For more
// information about how to use origin access identities, see Serving Private
// Content through CloudFront (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html)
// in the Amazon CloudFront Developer Guide.
//
//    // Example sending a request using CreateCloudFrontOriginAccessIdentityRequest.
//    req := client.CreateCloudFrontOriginAccessIdentityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudfront-2020-05-31/CreateCloudFrontOriginAccessIdentity
func (c *Client) CreateCloudFrontOriginAccessIdentityRequest(input *CreateCloudFrontOriginAccessIdentityInput) CreateCloudFrontOriginAccessIdentityRequest {
	op := &aws.Operation{
		Name:       opCreateCloudFrontOriginAccessIdentity,
		HTTPMethod: "POST",
		HTTPPath:   "/2020-05-31/origin-access-identity/cloudfront",
	}

	if input == nil {
		input = &CreateCloudFrontOriginAccessIdentityInput{}
	}

	req := c.newRequest(op, input, &CreateCloudFrontOriginAccessIdentityOutput{})

	return CreateCloudFrontOriginAccessIdentityRequest{Request: req, Input: input, Copy: c.CreateCloudFrontOriginAccessIdentityRequest}
}

// CreateCloudFrontOriginAccessIdentityRequest is the request type for the
// CreateCloudFrontOriginAccessIdentity API operation.
type CreateCloudFrontOriginAccessIdentityRequest struct {
	*aws.Request
	Input *CreateCloudFrontOriginAccessIdentityInput
	Copy  func(*CreateCloudFrontOriginAccessIdentityInput) CreateCloudFrontOriginAccessIdentityRequest
}

// Send marshals and sends the CreateCloudFrontOriginAccessIdentity API request.
func (r CreateCloudFrontOriginAccessIdentityRequest) Send(ctx context.Context) (*CreateCloudFrontOriginAccessIdentityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCloudFrontOriginAccessIdentityResponse{
		CreateCloudFrontOriginAccessIdentityOutput: r.Request.Data.(*CreateCloudFrontOriginAccessIdentityOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCloudFrontOriginAccessIdentityResponse is the response type for the
// CreateCloudFrontOriginAccessIdentity API operation.
type CreateCloudFrontOriginAccessIdentityResponse struct {
	*CreateCloudFrontOriginAccessIdentityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCloudFrontOriginAccessIdentity request.
func (r *CreateCloudFrontOriginAccessIdentityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
