// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fms

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutProtocolsListInput struct {
	_ struct{} `type:"structure"`

	// The details of the AWS Firewall Manager protocols list to be created.
	//
	// ProtocolsList is a required field
	ProtocolsList *ProtocolsListData `type:"structure" required:"true"`

	// The tags associated with the resource.
	TagList []Tag `type:"list"`
}

// String returns the string representation
func (s PutProtocolsListInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutProtocolsListInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutProtocolsListInput"}

	if s.ProtocolsList == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProtocolsList"))
	}
	if s.ProtocolsList != nil {
		if err := s.ProtocolsList.Validate(); err != nil {
			invalidParams.AddNested("ProtocolsList", err.(aws.ErrInvalidParams))
		}
	}
	if s.TagList != nil {
		for i, v := range s.TagList {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "TagList", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutProtocolsListOutput struct {
	_ struct{} `type:"structure"`

	// The details of the AWS Firewall Manager protocols list.
	ProtocolsList *ProtocolsListData `type:"structure"`

	// The Amazon Resource Name (ARN) of the protocols list.
	ProtocolsListArn *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PutProtocolsListOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutProtocolsList = "PutProtocolsList"

// PutProtocolsListRequest returns a request value for making API operation for
// Firewall Management Service.
//
// Creates an AWS Firewall Manager protocols list.
//
//    // Example sending a request using PutProtocolsListRequest.
//    req := client.PutProtocolsListRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/fms-2018-01-01/PutProtocolsList
func (c *Client) PutProtocolsListRequest(input *PutProtocolsListInput) PutProtocolsListRequest {
	op := &aws.Operation{
		Name:       opPutProtocolsList,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutProtocolsListInput{}
	}

	req := c.newRequest(op, input, &PutProtocolsListOutput{})

	return PutProtocolsListRequest{Request: req, Input: input, Copy: c.PutProtocolsListRequest}
}

// PutProtocolsListRequest is the request type for the
// PutProtocolsList API operation.
type PutProtocolsListRequest struct {
	*aws.Request
	Input *PutProtocolsListInput
	Copy  func(*PutProtocolsListInput) PutProtocolsListRequest
}

// Send marshals and sends the PutProtocolsList API request.
func (r PutProtocolsListRequest) Send(ctx context.Context) (*PutProtocolsListResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutProtocolsListResponse{
		PutProtocolsListOutput: r.Request.Data.(*PutProtocolsListOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutProtocolsListResponse is the response type for the
// PutProtocolsList API operation.
type PutProtocolsListResponse struct {
	*PutProtocolsListOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutProtocolsList request.
func (r *PutProtocolsListResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
