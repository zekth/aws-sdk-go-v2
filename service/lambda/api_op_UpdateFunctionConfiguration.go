// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateFunctionConfigurationInput struct {
	_ struct{} `type:"structure"`

	// A dead letter queue configuration that specifies the queue or topic where
	// Lambda sends asynchronous events when they fail processing. For more information,
	// see Dead Letter Queues (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq).
	DeadLetterConfig *DeadLetterConfig `type:"structure"`

	// A description of the function.
	Description *string `type:"string"`

	// Environment variables that are accessible from function code during execution.
	Environment *Environment `type:"structure"`

	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []FileSystemConfig `type:"list"`

	// The name of the Lambda function.
	//
	// Name formats
	//
	//    * Function name - my-function.
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//    * Partial ARN - 123456789012:function:my-function.
	//
	// The length constraint applies only to the full ARN. If you specify only the
	// function name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// The name of the method within your code that Lambda calls to execute your
	// function. The format includes the file name. It can also include namespaces
	// and other qualifiers, depending on the runtime. For more information, see
	// Programming Model (https://docs.aws.amazon.com/lambda/latest/dg/programming-model-v2.html).
	Handler *string `type:"string"`

	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt
	// your function's environment variables. If it's not provided, AWS Lambda uses
	// a default service key.
	KMSKeyArn *string `type:"string"`

	// A list of function layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
	// to add to the function's execution environment. Specify each layer by its
	// ARN, including the version.
	Layers []string `type:"list"`

	// The amount of memory that your function has access to. Increasing the function's
	// memory also increases its CPU allocation. The default value is 128 MB. The
	// value must be a multiple of 64 MB.
	MemorySize *int64 `min:"128" type:"integer"`

	// Only update the function if the revision ID matches the ID that's specified.
	// Use this option to avoid modifying a function that has changed since you
	// last read it.
	RevisionId *string `type:"string"`

	// The Amazon Resource Name (ARN) of the function's execution role.
	Role *string `type:"string"`

	// The identifier of the function's runtime (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html).
	Runtime Runtime `type:"string" enum:"true"`

	// The amount of time that Lambda allows a function to run before stopping it.
	// The default is 3 seconds. The maximum allowed value is 900 seconds.
	Timeout *int64 `min:"1" type:"integer"`

	// Set Mode to Active to sample and trace a subset of incoming requests with
	// AWS X-Ray.
	TracingConfig *TracingConfig `type:"structure"`

	// For network connectivity to AWS resources in a VPC, specify a list of security
	// groups and subnets in the VPC. When you connect a function to a VPC, it can
	// only access resources and the internet through that VPC. For more information,
	// see VPC Settings (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).
	VpcConfig *VpcConfig `type:"structure"`
}

// String returns the string representation
func (s UpdateFunctionConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFunctionConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFunctionConfigurationInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}
	if s.MemorySize != nil && *s.MemorySize < 128 {
		invalidParams.Add(aws.NewErrParamMinValue("MemorySize", 128))
	}
	if s.Timeout != nil && *s.Timeout < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Timeout", 1))
	}
	if s.FileSystemConfigs != nil {
		for i, v := range s.FileSystemConfigs {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "FileSystemConfigs", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFunctionConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.DeadLetterConfig != nil {
		v := s.DeadLetterConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DeadLetterConfig", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Environment != nil {
		v := s.Environment

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Environment", v, metadata)
	}
	if s.FileSystemConfigs != nil {
		v := s.FileSystemConfigs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "FileSystemConfigs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Handler != nil {
		v := *s.Handler

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Handler", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.KMSKeyArn != nil {
		v := *s.KMSKeyArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KMSKeyArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Layers != nil {
		v := s.Layers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Layers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.MemorySize != nil {
		v := *s.MemorySize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MemorySize", protocol.Int64Value(v), metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Role != nil {
		v := *s.Role

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Role", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Runtime) > 0 {
		v := s.Runtime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Runtime", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Timeout != nil {
		v := *s.Timeout

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Timeout", protocol.Int64Value(v), metadata)
	}
	if s.TracingConfig != nil {
		v := s.TracingConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TracingConfig", v, metadata)
	}
	if s.VpcConfig != nil {
		v := s.VpcConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "VpcConfig", v, metadata)
	}
	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Details about a function's configuration.
type UpdateFunctionConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string `type:"string"`

	// The size of the function's deployment package, in bytes.
	CodeSize *int64 `type:"long"`

	// The function's dead letter queue.
	DeadLetterConfig *DeadLetterConfig `type:"structure"`

	// The function's description.
	Description *string `type:"string"`

	// The function's environment variables.
	Environment *EnvironmentResponse `type:"structure"`

	// Connection settings for an Amazon EFS file system.
	FileSystemConfigs []FileSystemConfig `type:"list"`

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string `type:"string"`

	// The name of the function.
	FunctionName *string `min:"1" type:"string"`

	// The function that Lambda calls to begin executing your function.
	Handler *string `type:"string"`

	// The KMS key that's used to encrypt the function's environment variables.
	// This key is only returned if you've configured a customer managed CMK.
	KMSKeyArn *string `type:"string"`

	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string `type:"string"`

	// The status of the last update that was performed on the function. This is
	// first set to Successful after function creation completes.
	LastUpdateStatus LastUpdateStatus `type:"string" enum:"true"`

	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string `type:"string"`

	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode LastUpdateStatusReasonCode `type:"string" enum:"true"`

	// The function's layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).
	Layers []Layer `type:"list"`

	// For Lambda@Edge functions, the ARN of the master function.
	MasterArn *string `type:"string"`

	// The memory that's allocated to the function.
	MemorySize *int64 `min:"128" type:"integer"`

	// The latest updated revision of the function or alias.
	RevisionId *string `type:"string"`

	// The function's execution role.
	Role *string `type:"string"`

	// The runtime environment for the Lambda function.
	Runtime Runtime `type:"string" enum:"true"`

	// The current state of the function. When the state is Inactive, you can reactivate
	// the function by invoking it.
	State State `type:"string" enum:"true"`

	// The reason for the function's current state.
	StateReason *string `type:"string"`

	// The reason code for the function's current state. When the code is Creating,
	// you can't invoke or modify the function.
	StateReasonCode StateReasonCode `type:"string" enum:"true"`

	// The amount of time in seconds that Lambda allows a function to run before
	// stopping it.
	Timeout *int64 `min:"1" type:"integer"`

	// The function's AWS X-Ray tracing configuration.
	TracingConfig *TracingConfigResponse `type:"structure"`

	// The version of the Lambda function.
	Version *string `min:"1" type:"string"`

	// The function's networking configuration.
	VpcConfig *VpcConfigResponse `type:"structure"`
}

// String returns the string representation
func (s UpdateFunctionConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateFunctionConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CodeSha256 != nil {
		v := *s.CodeSha256

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CodeSha256", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CodeSize != nil {
		v := *s.CodeSize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CodeSize", protocol.Int64Value(v), metadata)
	}
	if s.DeadLetterConfig != nil {
		v := s.DeadLetterConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DeadLetterConfig", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Environment != nil {
		v := s.Environment

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Environment", v, metadata)
	}
	if s.FileSystemConfigs != nil {
		v := s.FileSystemConfigs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "FileSystemConfigs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.FunctionArn != nil {
		v := *s.FunctionArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FunctionName != nil {
		v := *s.FunctionName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FunctionName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Handler != nil {
		v := *s.Handler

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Handler", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.KMSKeyArn != nil {
		v := *s.KMSKeyArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "KMSKeyArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastModified != nil {
		v := *s.LastModified

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastModified", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.LastUpdateStatus) > 0 {
		v := s.LastUpdateStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastUpdateStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.LastUpdateStatusReason != nil {
		v := *s.LastUpdateStatusReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastUpdateStatusReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.LastUpdateStatusReasonCode) > 0 {
		v := s.LastUpdateStatusReasonCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LastUpdateStatusReasonCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Layers != nil {
		v := s.Layers

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Layers", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.MasterArn != nil {
		v := *s.MasterArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MasterArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MemorySize != nil {
		v := *s.MemorySize

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MemorySize", protocol.Int64Value(v), metadata)
	}
	if s.RevisionId != nil {
		v := *s.RevisionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RevisionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Role != nil {
		v := *s.Role

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Role", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Runtime) > 0 {
		v := s.Runtime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Runtime", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StateReason != nil {
		v := *s.StateReason

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StateReason", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.StateReasonCode) > 0 {
		v := s.StateReasonCode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StateReasonCode", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Timeout != nil {
		v := *s.Timeout

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Timeout", protocol.Int64Value(v), metadata)
	}
	if s.TracingConfig != nil {
		v := s.TracingConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "TracingConfig", v, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VpcConfig != nil {
		v := s.VpcConfig

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "VpcConfig", v, metadata)
	}
	return nil
}

const opUpdateFunctionConfiguration = "UpdateFunctionConfiguration"

// UpdateFunctionConfigurationRequest returns a request value for making API operation for
// AWS Lambda.
//
// Modify the version-specific settings of a Lambda function.
//
// When you update a function, Lambda provisions an instance of the function
// and its supporting resources. If your function connects to a VPC, this process
// can take a minute. During this time, you can't modify the function, but you
// can still invoke it. The LastUpdateStatus, LastUpdateStatusReason, and LastUpdateStatusReasonCode
// fields in the response from GetFunctionConfiguration indicate when the update
// is complete and the function is processing events with the new configuration.
// For more information, see Function States (https://docs.aws.amazon.com/lambda/latest/dg/functions-states.html).
//
// These settings can vary between versions of a function and are locked when
// you publish a version. You can't modify the configuration of a published
// version, only the unpublished version.
//
// To configure function concurrency, use PutFunctionConcurrency. To grant invoke
// permissions to an account or AWS service, use AddPermission.
//
//    // Example sending a request using UpdateFunctionConfigurationRequest.
//    req := client.UpdateFunctionConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/UpdateFunctionConfiguration
func (c *Client) UpdateFunctionConfigurationRequest(input *UpdateFunctionConfigurationInput) UpdateFunctionConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateFunctionConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/2015-03-31/functions/{FunctionName}/configuration",
	}

	if input == nil {
		input = &UpdateFunctionConfigurationInput{}
	}

	req := c.newRequest(op, input, &UpdateFunctionConfigurationOutput{})

	return UpdateFunctionConfigurationRequest{Request: req, Input: input, Copy: c.UpdateFunctionConfigurationRequest}
}

// UpdateFunctionConfigurationRequest is the request type for the
// UpdateFunctionConfiguration API operation.
type UpdateFunctionConfigurationRequest struct {
	*aws.Request
	Input *UpdateFunctionConfigurationInput
	Copy  func(*UpdateFunctionConfigurationInput) UpdateFunctionConfigurationRequest
}

// Send marshals and sends the UpdateFunctionConfiguration API request.
func (r UpdateFunctionConfigurationRequest) Send(ctx context.Context) (*UpdateFunctionConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateFunctionConfigurationResponse{
		UpdateFunctionConfigurationOutput: r.Request.Data.(*UpdateFunctionConfigurationOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateFunctionConfigurationResponse is the response type for the
// UpdateFunctionConfiguration API operation.
type UpdateFunctionConfigurationResponse struct {
	*UpdateFunctionConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateFunctionConfiguration request.
func (r *UpdateFunctionConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
