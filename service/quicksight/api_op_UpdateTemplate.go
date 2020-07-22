// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateTemplateInput struct {
	_ struct{} `type:"structure"`

	// The ID of the AWS account that contains the template that you're updating.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The name for the template.
	Name *string `min:"1" type:"string"`

	// The entity that you are using as a source when you update the template. In
	// SourceEntity, you specify the type of object you're using as source: SourceTemplate
	// for a template or SourceAnalysis for an analysis. Both of these require an
	// Amazon Resource Name (ARN). For SourceTemplate, specify the ARN of the source
	// template. For SourceAnalysis, specify the ARN of the source analysis. The
	// SourceTemplate ARN can contain any AWS Account and any QuickSight-supported
	// AWS Region.
	//
	// Use the DataSetReferences entity within SourceTemplate or SourceAnalysis
	// to list the replacement datasets for the placeholders listed in the original.
	// The schema in each dataset must match its placeholder.
	//
	// SourceEntity is a required field
	SourceEntity *TemplateSourceEntity `type:"structure" required:"true"`

	// The ID for the template.
	//
	// TemplateId is a required field
	TemplateId *string `location:"uri" locationName:"TemplateId" min:"1" type:"string" required:"true"`

	// A description of the current template version that is being updated. Every
	// time you call UpdateTemplate, you create a new version of the template. Each
	// version of the template maintains a description of the version in the VersionDescription
	// field.
	VersionDescription *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTemplateInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.SourceEntity == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceEntity"))
	}

	if s.TemplateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateId"))
	}
	if s.TemplateId != nil && len(*s.TemplateId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateId", 1))
	}
	if s.VersionDescription != nil && len(*s.VersionDescription) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionDescription", 1))
	}
	if s.SourceEntity != nil {
		if err := s.SourceEntity.Validate(); err != nil {
			invalidParams.AddNested("SourceEntity", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTemplateInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SourceEntity != nil {
		v := s.SourceEntity

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "SourceEntity", v, metadata)
	}
	if s.VersionDescription != nil {
		v := *s.VersionDescription

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VersionDescription", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "TemplateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateTemplateOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the template.
	Arn *string `type:"string"`

	// The creation status of the template.
	CreationStatus ResourceStatus `type:"string" enum:"true"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The ID for the template.
	TemplateId *string `min:"1" type:"string"`

	// The ARN for the template, including the version information of the first
	// version.
	VersionArn *string `type:"string"`
}

// String returns the string representation
func (s UpdateTemplateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateTemplateOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.CreationStatus) > 0 {
		v := s.CreationStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TemplateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionArn != nil {
		v := *s.VersionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VersionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opUpdateTemplate = "UpdateTemplate"

// UpdateTemplateRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates a template from an existing Amazon QuickSight analysis or another
// template.
//
//    // Example sending a request using UpdateTemplateRequest.
//    req := client.UpdateTemplateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateTemplate
func (c *Client) UpdateTemplateRequest(input *UpdateTemplateInput) UpdateTemplateRequest {
	op := &aws.Operation{
		Name:       opUpdateTemplate,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/templates/{TemplateId}",
	}

	if input == nil {
		input = &UpdateTemplateInput{}
	}

	req := c.newRequest(op, input, &UpdateTemplateOutput{})

	return UpdateTemplateRequest{Request: req, Input: input, Copy: c.UpdateTemplateRequest}
}

// UpdateTemplateRequest is the request type for the
// UpdateTemplate API operation.
type UpdateTemplateRequest struct {
	*aws.Request
	Input *UpdateTemplateInput
	Copy  func(*UpdateTemplateInput) UpdateTemplateRequest
}

// Send marshals and sends the UpdateTemplate API request.
func (r UpdateTemplateRequest) Send(ctx context.Context) (*UpdateTemplateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTemplateResponse{
		UpdateTemplateOutput: r.Request.Data.(*UpdateTemplateOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTemplateResponse is the response type for the
// UpdateTemplate API operation.
type UpdateTemplateResponse struct {
	*UpdateTemplateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTemplate request.
func (r *UpdateTemplateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
