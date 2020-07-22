// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteHumanTaskUiInput struct {
	_ struct{} `type:"structure"`

	// The name of the human task user interface (work task template) you want to
	// delete.
	//
	// HumanTaskUiName is a required field
	HumanTaskUiName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteHumanTaskUiInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteHumanTaskUiInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteHumanTaskUiInput"}

	if s.HumanTaskUiName == nil {
		invalidParams.Add(aws.NewErrParamRequired("HumanTaskUiName"))
	}
	if s.HumanTaskUiName != nil && len(*s.HumanTaskUiName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HumanTaskUiName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteHumanTaskUiOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteHumanTaskUiOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteHumanTaskUi = "DeleteHumanTaskUi"

// DeleteHumanTaskUiRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Use this operation to delete a worker task template (HumanTaskUi).
//
// To see a list of human task user interfaces (work task templates) in your
// account, use . When you delete a worker task template, it no longer appears
// when you call ListHumanTaskUis.
//
//    // Example sending a request using DeleteHumanTaskUiRequest.
//    req := client.DeleteHumanTaskUiRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DeleteHumanTaskUi
func (c *Client) DeleteHumanTaskUiRequest(input *DeleteHumanTaskUiInput) DeleteHumanTaskUiRequest {
	op := &aws.Operation{
		Name:       opDeleteHumanTaskUi,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteHumanTaskUiInput{}
	}

	req := c.newRequest(op, input, &DeleteHumanTaskUiOutput{})

	return DeleteHumanTaskUiRequest{Request: req, Input: input, Copy: c.DeleteHumanTaskUiRequest}
}

// DeleteHumanTaskUiRequest is the request type for the
// DeleteHumanTaskUi API operation.
type DeleteHumanTaskUiRequest struct {
	*aws.Request
	Input *DeleteHumanTaskUiInput
	Copy  func(*DeleteHumanTaskUiInput) DeleteHumanTaskUiRequest
}

// Send marshals and sends the DeleteHumanTaskUi API request.
func (r DeleteHumanTaskUiRequest) Send(ctx context.Context) (*DeleteHumanTaskUiResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteHumanTaskUiResponse{
		DeleteHumanTaskUiOutput: r.Request.Data.(*DeleteHumanTaskUiOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteHumanTaskUiResponse is the response type for the
// DeleteHumanTaskUi API operation.
type DeleteHumanTaskUiResponse struct {
	*DeleteHumanTaskUiOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteHumanTaskUi request.
func (r *DeleteHumanTaskUiResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
