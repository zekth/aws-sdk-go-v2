// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// UpdateNFSFileShareInput
type UpdateNFSFileShareInput struct {
	_ struct{} `type:"structure"`

	// Refresh cache information.
	CacheAttributes *CacheAttributes `type:"structure"`

	// The list of clients that are allowed to access the file gateway. The list
	// must contain either valid IP addresses or valid CIDR blocks.
	ClientList []string `min:"1" type:"list"`

	// The default storage class for objects put into an Amazon S3 bucket by the
	// file gateway. The default value is S3_INTELLIGENT_TIERING. Optional.
	//
	// Valid Values: S3_STANDARD | S3_INTELLIGENT_TIERING | S3_STANDARD_IA | S3_ONEZONE_IA
	DefaultStorageClass *string `min:"5" type:"string"`

	// The Amazon Resource Name (ARN) of the file share to be updated.
	//
	// FileShareARN is a required field
	FileShareARN *string `min:"50" type:"string" required:"true"`

	// The name of the file share. Optional.
	//
	// FileShareName must be set if an S3 prefix name is set in LocationARN.
	FileShareName *string `min:"1" type:"string"`

	// A value that enables guessing of the MIME type for uploaded objects based
	// on file extensions. Set this value to true to enable MIME type guessing,
	// otherwise set to false. The default value is true.
	//
	// Valid Values: true | false
	GuessMIMETypeEnabled *bool `type:"boolean"`

	// Set to true to use Amazon S3 server-side encryption with your own AWS KMS
	// key, or false to use a key managed by Amazon S3. Optional.
	//
	// Valid Values: true | false
	KMSEncrypted *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used
	// for Amazon S3 server-side encryption. Storage Gateway does not support asymmetric
	// CMKs. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string `min:"7" type:"string"`

	// The default values for the file share. Optional.
	NFSFileShareDefaults *NFSFileShareDefaults `type:"structure"`

	// A value that sets the access control list (ACL) permission for objects in
	// the S3 bucket that a file gateway puts objects into. The default value is
	// private.
	ObjectACL ObjectACL `type:"string" enum:"true"`

	// A value that sets the write status of a file share. Set this value to true
	// to set the write status to read-only, otherwise set to false.
	//
	// Valid Values: true | false
	ReadOnly *bool `type:"boolean"`

	// A value that sets who pays the cost of the request and the cost associated
	// with data download from the S3 bucket. If this value is set to true, the
	// requester pays the costs; otherwise, the S3 bucket owner pays. However, the
	// S3 bucket owner always pays the cost of storing data.
	//
	// RequesterPays is a configuration for the S3 bucket that backs the file share,
	// so make sure that the configuration on the file share is the same as the
	// S3 bucket configuration.
	//
	// Valid Values: true | false
	RequesterPays *bool `type:"boolean"`

	// The user mapped to anonymous user.
	//
	// Valid values are the following:
	//
	//    * RootSquash: Only root is mapped to anonymous user.
	//
	//    * NoSquash: No one is mapped to anonymous user.
	//
	//    * AllSquash: Everyone is mapped to anonymous user.
	Squash *string `min:"5" type:"string"`
}

// String returns the string representation
func (s UpdateNFSFileShareInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateNFSFileShareInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateNFSFileShareInput"}
	if s.ClientList != nil && len(s.ClientList) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientList", 1))
	}
	if s.DefaultStorageClass != nil && len(*s.DefaultStorageClass) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("DefaultStorageClass", 5))
	}

	if s.FileShareARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("FileShareARN"))
	}
	if s.FileShareARN != nil && len(*s.FileShareARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("FileShareARN", 50))
	}
	if s.FileShareName != nil && len(*s.FileShareName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FileShareName", 1))
	}
	if s.KMSKey != nil && len(*s.KMSKey) < 7 {
		invalidParams.Add(aws.NewErrParamMinLen("KMSKey", 7))
	}
	if s.Squash != nil && len(*s.Squash) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("Squash", 5))
	}
	if s.NFSFileShareDefaults != nil {
		if err := s.NFSFileShareDefaults.Validate(); err != nil {
			invalidParams.AddNested("NFSFileShareDefaults", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// UpdateNFSFileShareOutput
type UpdateNFSFileShareOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the updated file share.
	FileShareARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s UpdateNFSFileShareOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateNFSFileShare = "UpdateNFSFileShare"

// UpdateNFSFileShareRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Updates a Network File System (NFS) file share. This operation is only supported
// in the file gateway type.
//
// To leave a file share field unchanged, set the corresponding input field
// to null.
//
// Updates the following file share setting:
//
//    * Default storage class for your S3 bucket
//
//    * Metadata defaults for your S3 bucket
//
//    * Allowed NFS clients for your file share
//
//    * Squash settings
//
//    * Write status of your file share
//
// To leave a file share field unchanged, set the corresponding input field
// to null. This operation is only supported in file gateways.
//
//    // Example sending a request using UpdateNFSFileShareRequest.
//    req := client.UpdateNFSFileShareRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateNFSFileShare
func (c *Client) UpdateNFSFileShareRequest(input *UpdateNFSFileShareInput) UpdateNFSFileShareRequest {
	op := &aws.Operation{
		Name:       opUpdateNFSFileShare,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateNFSFileShareInput{}
	}

	req := c.newRequest(op, input, &UpdateNFSFileShareOutput{})

	return UpdateNFSFileShareRequest{Request: req, Input: input, Copy: c.UpdateNFSFileShareRequest}
}

// UpdateNFSFileShareRequest is the request type for the
// UpdateNFSFileShare API operation.
type UpdateNFSFileShareRequest struct {
	*aws.Request
	Input *UpdateNFSFileShareInput
	Copy  func(*UpdateNFSFileShareInput) UpdateNFSFileShareRequest
}

// Send marshals and sends the UpdateNFSFileShare API request.
func (r UpdateNFSFileShareRequest) Send(ctx context.Context) (*UpdateNFSFileShareResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateNFSFileShareResponse{
		UpdateNFSFileShareOutput: r.Request.Data.(*UpdateNFSFileShareOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateNFSFileShareResponse is the response type for the
// UpdateNFSFileShare API operation.
type UpdateNFSFileShareResponse struct {
	*UpdateNFSFileShareOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateNFSFileShare request.
func (r *UpdateNFSFileShareResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
