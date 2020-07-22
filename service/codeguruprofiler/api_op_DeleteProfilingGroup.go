// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeguruprofiler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The structure representing the deleteProfilingGroupRequest.
type DeleteProfilingGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the profiling group to delete.
	//
	// ProfilingGroupName is a required field
	ProfilingGroupName *string `location:"uri" locationName:"profilingGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteProfilingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProfilingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteProfilingGroupInput"}

	if s.ProfilingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfilingGroupName"))
	}
	if s.ProfilingGroupName != nil && len(*s.ProfilingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfilingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteProfilingGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ProfilingGroupName != nil {
		v := *s.ProfilingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "profilingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The structure representing the deleteProfilingGroupResponse.
type DeleteProfilingGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteProfilingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteProfilingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteProfilingGroup = "DeleteProfilingGroup"

// DeleteProfilingGroupRequest returns a request value for making API operation for
// Amazon CodeGuru Profiler.
//
// Deletes a profiling group.
//
//    // Example sending a request using DeleteProfilingGroupRequest.
//    req := client.DeleteProfilingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguruprofiler-2019-07-18/DeleteProfilingGroup
func (c *Client) DeleteProfilingGroupRequest(input *DeleteProfilingGroupInput) DeleteProfilingGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteProfilingGroup,
		HTTPMethod: "DELETE",
		HTTPPath:   "/profilingGroups/{profilingGroupName}",
	}

	if input == nil {
		input = &DeleteProfilingGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteProfilingGroupOutput{})

	return DeleteProfilingGroupRequest{Request: req, Input: input, Copy: c.DeleteProfilingGroupRequest}
}

// DeleteProfilingGroupRequest is the request type for the
// DeleteProfilingGroup API operation.
type DeleteProfilingGroupRequest struct {
	*aws.Request
	Input *DeleteProfilingGroupInput
	Copy  func(*DeleteProfilingGroupInput) DeleteProfilingGroupRequest
}

// Send marshals and sends the DeleteProfilingGroup API request.
func (r DeleteProfilingGroupRequest) Send(ctx context.Context) (*DeleteProfilingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteProfilingGroupResponse{
		DeleteProfilingGroupOutput: r.Request.Data.(*DeleteProfilingGroupOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteProfilingGroupResponse is the response type for the
// DeleteProfilingGroup API operation.
type DeleteProfilingGroupResponse struct {
	*DeleteProfilingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteProfilingGroup request.
func (r *DeleteProfilingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
