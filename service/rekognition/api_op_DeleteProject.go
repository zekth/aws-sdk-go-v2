// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteProjectInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the project that you want to delete.
	//
	// ProjectArn is a required field
	ProjectArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteProjectInput"}

	if s.ProjectArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectArn"))
	}
	if s.ProjectArn != nil && len(*s.ProjectArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteProjectOutput struct {
	_ struct{} `type:"structure"`

	// The current status of the delete project operation.
	Status ProjectStatus `type:"string" enum:"true"`
}

// String returns the string representation
func (s DeleteProjectOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteProject = "DeleteProject"

// DeleteProjectRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Deletes an Amazon Rekognition Custom Labels project. To delete a project
// you must first delete all models associated with the project. To delete a
// model, see DeleteProjectVersion.
//
// This operation requires permissions to perform the rekognition:DeleteProject
// action.
//
//    // Example sending a request using DeleteProjectRequest.
//    req := client.DeleteProjectRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteProjectRequest(input *DeleteProjectInput) DeleteProjectRequest {
	op := &aws.Operation{
		Name:       opDeleteProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteProjectInput{}
	}

	req := c.newRequest(op, input, &DeleteProjectOutput{})

	return DeleteProjectRequest{Request: req, Input: input, Copy: c.DeleteProjectRequest}
}

// DeleteProjectRequest is the request type for the
// DeleteProject API operation.
type DeleteProjectRequest struct {
	*aws.Request
	Input *DeleteProjectInput
	Copy  func(*DeleteProjectInput) DeleteProjectRequest
}

// Send marshals and sends the DeleteProject API request.
func (r DeleteProjectRequest) Send(ctx context.Context) (*DeleteProjectResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteProjectResponse{
		DeleteProjectOutput: r.Request.Data.(*DeleteProjectOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteProjectResponse is the response type for the
// DeleteProject API operation.
type DeleteProjectResponse struct {
	*DeleteProjectOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteProject request.
func (r *DeleteProjectResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
