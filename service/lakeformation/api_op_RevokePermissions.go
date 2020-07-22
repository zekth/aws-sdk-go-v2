// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lakeformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RevokePermissionsInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string `min:"1" type:"string"`

	// The permissions revoked to the principal on the resource. For information
	// about permissions, see Security and Access Control to Metadata and Data (https://docs-aws.amazon.com/lake-formation/latest/dg/security-data-access.html).
	//
	// Permissions is a required field
	Permissions []Permission `type:"list" required:"true"`

	// Indicates a list of permissions for which to revoke the grant option allowing
	// the principal to pass permissions to other principals.
	PermissionsWithGrantOption []Permission `type:"list"`

	// The principal to be revoked permissions on the resource.
	//
	// Principal is a required field
	Principal *DataLakePrincipal `type:"structure" required:"true"`

	// The resource to which permissions are to be revoked.
	//
	// Resource is a required field
	Resource *Resource `type:"structure" required:"true"`
}

// String returns the string representation
func (s RevokePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RevokePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RevokePermissionsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.Permissions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Permissions"))
	}

	if s.Principal == nil {
		invalidParams.Add(aws.NewErrParamRequired("Principal"))
	}

	if s.Resource == nil {
		invalidParams.Add(aws.NewErrParamRequired("Resource"))
	}
	if s.Principal != nil {
		if err := s.Principal.Validate(); err != nil {
			invalidParams.AddNested("Principal", err.(aws.ErrInvalidParams))
		}
	}
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			invalidParams.AddNested("Resource", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RevokePermissionsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RevokePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opRevokePermissions = "RevokePermissions"

// RevokePermissionsRequest returns a request value for making API operation for
// AWS Lake Formation.
//
// Revokes permissions to the principal to access metadata in the Data Catalog
// and data organized in underlying data storage such as Amazon S3.
//
//    // Example sending a request using RevokePermissionsRequest.
//    req := client.RevokePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lakeformation-2017-03-31/RevokePermissions
func (c *Client) RevokePermissionsRequest(input *RevokePermissionsInput) RevokePermissionsRequest {
	op := &aws.Operation{
		Name:       opRevokePermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RevokePermissionsInput{}
	}

	req := c.newRequest(op, input, &RevokePermissionsOutput{})

	return RevokePermissionsRequest{Request: req, Input: input, Copy: c.RevokePermissionsRequest}
}

// RevokePermissionsRequest is the request type for the
// RevokePermissions API operation.
type RevokePermissionsRequest struct {
	*aws.Request
	Input *RevokePermissionsInput
	Copy  func(*RevokePermissionsInput) RevokePermissionsRequest
}

// Send marshals and sends the RevokePermissions API request.
func (r RevokePermissionsRequest) Send(ctx context.Context) (*RevokePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RevokePermissionsResponse{
		RevokePermissionsOutput: r.Request.Data.(*RevokePermissionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RevokePermissionsResponse is the response type for the
// RevokePermissions API operation.
type RevokePermissionsResponse struct {
	*RevokePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RevokePermissions request.
func (r *RevokePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
