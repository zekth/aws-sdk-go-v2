// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateNodeInput struct {
	_ struct{} `type:"structure"`

	// Engine attributes used for associating the node.
	//
	// Attributes accepted in a AssociateNode request for Chef
	//
	//    * CHEF_ORGANIZATION: The Chef organization with which the node is associated.
	//    By default only one organization named default can exist.
	//
	//    * CHEF_NODE_PUBLIC_KEY: A PEM-formatted public key. This key is required
	//    for the chef-client agent to access the Chef API.
	//
	// Attributes accepted in a AssociateNode request for Puppet
	//
	//    * PUPPET_NODE_CSR: A PEM-formatted certificate-signing request (CSR) that
	//    is created by the node.
	//
	// EngineAttributes is a required field
	EngineAttributes []EngineAttribute `type:"list" required:"true"`

	// The name of the node.
	//
	// NodeName is a required field
	NodeName *string `type:"string" required:"true"`

	// The name of the server with which to associate the node.
	//
	// ServerName is a required field
	ServerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateNodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateNodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateNodeInput"}

	if s.EngineAttributes == nil {
		invalidParams.Add(aws.NewErrParamRequired("EngineAttributes"))
	}

	if s.NodeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodeName"))
	}

	if s.ServerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServerName"))
	}
	if s.ServerName != nil && len(*s.ServerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateNodeOutput struct {
	_ struct{} `type:"structure"`

	// Contains a token which can be passed to the DescribeNodeAssociationStatus
	// API call to get the status of the association request.
	NodeAssociationStatusToken *string `type:"string"`
}

// String returns the string representation
func (s AssociateNodeOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateNode = "AssociateNode"

// AssociateNodeRequest returns a request value for making API operation for
// AWS OpsWorks CM.
//
// Associates a new node with the server. For more information about how to
// disassociate a node, see DisassociateNode.
//
// On a Chef server: This command is an alternative to knife bootstrap.
//
// Example (Chef): aws opsworks-cm associate-node --server-name MyServer --node-name
// MyManagedNode --engine-attributes "Name=CHEF_ORGANIZATION,Value=default"
// "Name=CHEF_NODE_PUBLIC_KEY,Value=public-key-pem"
//
// On a Puppet server, this command is an alternative to the puppet cert sign
// command that signs a Puppet node CSR.
//
// Example (Puppet): aws opsworks-cm associate-node --server-name MyServer --node-name
// MyManagedNode --engine-attributes "Name=PUPPET_NODE_CSR,Value=csr-pem"
//
// A node can can only be associated with servers that are in a HEALTHY state.
// Otherwise, an InvalidStateException is thrown. A ResourceNotFoundException
// is thrown when the server does not exist. A ValidationException is raised
// when parameters of the request are not valid. The AssociateNode API call
// can be integrated into Auto Scaling configurations, AWS Cloudformation templates,
// or the user data of a server's instance.
//
//    // Example sending a request using AssociateNodeRequest.
//    req := client.AssociateNodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/opsworkscm-2016-11-01/AssociateNode
func (c *Client) AssociateNodeRequest(input *AssociateNodeInput) AssociateNodeRequest {
	op := &aws.Operation{
		Name:       opAssociateNode,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateNodeInput{}
	}

	req := c.newRequest(op, input, &AssociateNodeOutput{})

	return AssociateNodeRequest{Request: req, Input: input, Copy: c.AssociateNodeRequest}
}

// AssociateNodeRequest is the request type for the
// AssociateNode API operation.
type AssociateNodeRequest struct {
	*aws.Request
	Input *AssociateNodeInput
	Copy  func(*AssociateNodeInput) AssociateNodeRequest
}

// Send marshals and sends the AssociateNode API request.
func (r AssociateNodeRequest) Send(ctx context.Context) (*AssociateNodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateNodeResponse{
		AssociateNodeOutput: r.Request.Data.(*AssociateNodeOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateNodeResponse is the response type for the
// AssociateNode API operation.
type AssociateNodeResponse struct {
	*AssociateNodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateNode request.
func (r *AssociateNodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
