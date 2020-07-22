// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetBackupPlanInput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a backup plan.
	//
	// BackupPlanId is a required field
	BackupPlanId *string `location:"uri" locationName:"backupPlanId" type:"string" required:"true"`

	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most
	// 1,024 bytes long. Version IDs cannot be edited.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s GetBackupPlanInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBackupPlanInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBackupPlanInput"}

	if s.BackupPlanId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupPlanId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBackupPlanInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BackupPlanId != nil {
		v := *s.BackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "backupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "versionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetBackupPlanOutput struct {
	_ struct{} `type:"structure"`

	// Specifies the body of a backup plan. Includes a BackupPlanName and one or
	// more sets of Rules.
	BackupPlan *BackupPlan `type:"structure"`

	// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for
	// example, arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50.
	BackupPlanArn *string `type:"string"`

	// Uniquely identifies a backup plan.
	BackupPlanId *string `type:"string"`

	// The date and time that a backup plan is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time `type:"timestamp"`

	// A unique string that identifies the request and allows failed requests to
	// be retried without the risk of executing the operation twice.
	CreatorRequestId *string `type:"string"`

	// The date and time that a backup plan is deleted, in Unix format and Coordinated
	// Universal Time (UTC). The value of DeletionDate is accurate to milliseconds.
	// For example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	DeletionDate *time.Time `type:"timestamp"`

	// The last time a job to back up resources was executed with this backup plan.
	// A date and time, in Unix format and Coordinated Universal Time (UTC). The
	// value of LastExecutionDate is accurate to milliseconds. For example, the
	// value 1516925490.087 represents Friday, January 26, 2018 12:11:30.087 AM.
	LastExecutionDate *time.Time `type:"timestamp"`

	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most
	// 1,024 bytes long. Version IDs cannot be edited.
	VersionId *string `type:"string"`
}

// String returns the string representation
func (s GetBackupPlanOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBackupPlanOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.BackupPlan != nil {
		v := s.BackupPlan

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "BackupPlan", v, metadata)
	}
	if s.BackupPlanArn != nil {
		v := *s.BackupPlanArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupPlanArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BackupPlanId != nil {
		v := *s.BackupPlanId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "BackupPlanId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreationDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.CreatorRequestId != nil {
		v := *s.CreatorRequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CreatorRequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeletionDate != nil {
		v := *s.DeletionDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeletionDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.LastExecutionDate != nil {
		v := *s.LastExecutionDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastExecutionDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VersionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetBackupPlan = "GetBackupPlan"

// GetBackupPlanRequest returns a request value for making API operation for
// AWS Backup.
//
// Returns the body of a backup plan in JSON format, in addition to plan metadata.
//
//    // Example sending a request using GetBackupPlanRequest.
//    req := client.GetBackupPlanRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/backup-2018-11-15/GetBackupPlan
func (c *Client) GetBackupPlanRequest(input *GetBackupPlanInput) GetBackupPlanRequest {
	op := &aws.Operation{
		Name:       opGetBackupPlan,
		HTTPMethod: "GET",
		HTTPPath:   "/backup/plans/{backupPlanId}/",
	}

	if input == nil {
		input = &GetBackupPlanInput{}
	}

	req := c.newRequest(op, input, &GetBackupPlanOutput{})

	return GetBackupPlanRequest{Request: req, Input: input, Copy: c.GetBackupPlanRequest}
}

// GetBackupPlanRequest is the request type for the
// GetBackupPlan API operation.
type GetBackupPlanRequest struct {
	*aws.Request
	Input *GetBackupPlanInput
	Copy  func(*GetBackupPlanInput) GetBackupPlanRequest
}

// Send marshals and sends the GetBackupPlan API request.
func (r GetBackupPlanRequest) Send(ctx context.Context) (*GetBackupPlanResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBackupPlanResponse{
		GetBackupPlanOutput: r.Request.Data.(*GetBackupPlanOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBackupPlanResponse is the response type for the
// GetBackupPlan API operation.
type GetBackupPlanResponse struct {
	*GetBackupPlanOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBackupPlan request.
func (r *GetBackupPlanResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
