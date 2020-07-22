// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateSnapshotFromVolumeRecoveryPointInput struct {
	_ struct{} `type:"structure"`

	// Textual description of the snapshot that appears in the Amazon EC2 console,
	// Elastic Block Store snapshots panel in the Description field, and in the
	// AWS Storage Gateway snapshot Details pane, Description field.
	//
	// SnapshotDescription is a required field
	SnapshotDescription *string `min:"1" type:"string" required:"true"`

	// A list of up to 50 tags that can be assigned to a snapshot. Each tag is a
	// key-value pair.
	//
	// Valid characters for key and value are letters, spaces, and numbers representable
	// in UTF-8 format, and the following special characters: + - = . _ : / @. The
	// maximum length of a tag's key is 128 characters, and the maximum length for
	// a tag's value is 256.
	Tags []Tag `type:"list"`

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the DescribeStorediSCSIVolumes
	// operation to return to retrieve the TargetARN for specified VolumeARN.
	//
	// VolumeARN is a required field
	VolumeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSnapshotFromVolumeRecoveryPointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSnapshotFromVolumeRecoveryPointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSnapshotFromVolumeRecoveryPointInput"}

	if s.SnapshotDescription == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotDescription"))
	}
	if s.SnapshotDescription != nil && len(*s.SnapshotDescription) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SnapshotDescription", 1))
	}

	if s.VolumeARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeARN"))
	}
	if s.VolumeARN != nil && len(*s.VolumeARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("VolumeARN", 50))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSnapshotFromVolumeRecoveryPointOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the snapshot.
	SnapshotId *string `type:"string"`

	// The Amazon Resource Name (ARN) of the iSCSI volume target. Use the DescribeStorediSCSIVolumes
	// operation to return to retrieve the TargetARN for specified VolumeARN.
	VolumeARN *string `min:"50" type:"string"`

	// The time the volume was created from the recovery point.
	VolumeRecoveryPointTime *string `type:"string"`
}

// String returns the string representation
func (s CreateSnapshotFromVolumeRecoveryPointOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSnapshotFromVolumeRecoveryPoint = "CreateSnapshotFromVolumeRecoveryPoint"

// CreateSnapshotFromVolumeRecoveryPointRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Initiates a snapshot of a gateway from a volume recovery point. This operation
// is only supported in the cached volume gateway type.
//
// A volume recovery point is a point in time at which all data of the volume
// is consistent and from which you can create a snapshot. To get a list of
// volume recovery point for cached volume gateway, use ListVolumeRecoveryPoints.
//
// In the CreateSnapshotFromVolumeRecoveryPoint request, you identify the volume
// by providing its Amazon Resource Name (ARN). You must also provide a description
// for the snapshot. When the gateway takes a snapshot of the specified volume,
// the snapshot and its description appear in the AWS Storage Gateway console.
// In response, the gateway returns you a snapshot ID. You can use this snapshot
// ID to check the snapshot progress or later use it when you want to create
// a volume from a snapshot.
//
// To list or delete a snapshot, you must use the Amazon EC2 API. For more information,
// see DescribeSnapshots (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html)
// or DeleteSnapshot (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DeleteSnapshot.html)
// in the Amazon Elastic Compute Cloud API Reference.
//
//    // Example sending a request using CreateSnapshotFromVolumeRecoveryPointRequest.
//    req := client.CreateSnapshotFromVolumeRecoveryPointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/CreateSnapshotFromVolumeRecoveryPoint
func (c *Client) CreateSnapshotFromVolumeRecoveryPointRequest(input *CreateSnapshotFromVolumeRecoveryPointInput) CreateSnapshotFromVolumeRecoveryPointRequest {
	op := &aws.Operation{
		Name:       opCreateSnapshotFromVolumeRecoveryPoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSnapshotFromVolumeRecoveryPointInput{}
	}

	req := c.newRequest(op, input, &CreateSnapshotFromVolumeRecoveryPointOutput{})

	return CreateSnapshotFromVolumeRecoveryPointRequest{Request: req, Input: input, Copy: c.CreateSnapshotFromVolumeRecoveryPointRequest}
}

// CreateSnapshotFromVolumeRecoveryPointRequest is the request type for the
// CreateSnapshotFromVolumeRecoveryPoint API operation.
type CreateSnapshotFromVolumeRecoveryPointRequest struct {
	*aws.Request
	Input *CreateSnapshotFromVolumeRecoveryPointInput
	Copy  func(*CreateSnapshotFromVolumeRecoveryPointInput) CreateSnapshotFromVolumeRecoveryPointRequest
}

// Send marshals and sends the CreateSnapshotFromVolumeRecoveryPoint API request.
func (r CreateSnapshotFromVolumeRecoveryPointRequest) Send(ctx context.Context) (*CreateSnapshotFromVolumeRecoveryPointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSnapshotFromVolumeRecoveryPointResponse{
		CreateSnapshotFromVolumeRecoveryPointOutput: r.Request.Data.(*CreateSnapshotFromVolumeRecoveryPointOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSnapshotFromVolumeRecoveryPointResponse is the response type for the
// CreateSnapshotFromVolumeRecoveryPoint API operation.
type CreateSnapshotFromVolumeRecoveryPointResponse struct {
	*CreateSnapshotFromVolumeRecoveryPointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSnapshotFromVolumeRecoveryPoint request.
func (r *CreateSnapshotFromVolumeRecoveryPointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
