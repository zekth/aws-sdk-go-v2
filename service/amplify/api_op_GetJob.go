// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The request structure for the get job request.
type GetJobInput struct {
	_ struct{} `type:"structure"`

	// The unique ID for an Amplify app.
	//
	// AppId is a required field
	AppId *string `location:"uri" locationName:"appId" min:"1" type:"string" required:"true"`

	// The branch name for the job.
	//
	// BranchName is a required field
	BranchName *string `location:"uri" locationName:"branchName" min:"1" type:"string" required:"true"`

	// The unique ID for the job.
	//
	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetJobInput"}

	if s.AppId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AppId"))
	}
	if s.AppId != nil && len(*s.AppId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AppId", 1))
	}

	if s.BranchName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BranchName"))
	}
	if s.BranchName != nil && len(*s.BranchName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BranchName", 1))
	}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AppId != nil {
		v := *s.AppId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "appId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BranchName != nil {
		v := *s.BranchName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "branchName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetJobOutput struct {
	_ struct{} `type:"structure"`

	// Describes an execution job for an Amplify app.
	//
	// Job is a required field
	Job *Job `locationName:"job" type:"structure" required:"true"`
}

// String returns the string representation
func (s GetJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Job != nil {
		v := s.Job

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "job", v, metadata)
	}
	return nil
}

const opGetJob = "GetJob"

// GetJobRequest returns a request value for making API operation for
// AWS Amplify.
//
// Returns a job for a branch of an Amplify app.
//
//    // Example sending a request using GetJobRequest.
//    req := client.GetJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/GetJob
func (c *Client) GetJobRequest(input *GetJobInput) GetJobRequest {
	op := &aws.Operation{
		Name:       opGetJob,
		HTTPMethod: "GET",
		HTTPPath:   "/apps/{appId}/branches/{branchName}/jobs/{jobId}",
	}

	if input == nil {
		input = &GetJobInput{}
	}

	req := c.newRequest(op, input, &GetJobOutput{})

	return GetJobRequest{Request: req, Input: input, Copy: c.GetJobRequest}
}

// GetJobRequest is the request type for the
// GetJob API operation.
type GetJobRequest struct {
	*aws.Request
	Input *GetJobInput
	Copy  func(*GetJobInput) GetJobRequest
}

// Send marshals and sends the GetJob API request.
func (r GetJobRequest) Send(ctx context.Context) (*GetJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetJobResponse{
		GetJobOutput: r.Request.Data.(*GetJobOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetJobResponse is the response type for the
// GetJob API operation.
type GetJobResponse struct {
	*GetJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetJob request.
func (r *GetJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
