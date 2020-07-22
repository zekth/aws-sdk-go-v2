// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appconfig

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

type Application struct {
	_ struct{} `type:"structure"`

	// The description of the application.
	Description *string `type:"string"`

	// The application ID.
	Id *string `type:"string"`

	// The application name.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s Application) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Application) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A summary of a configuration profile.
type ConfigurationProfileSummary struct {
	_ struct{} `type:"structure"`

	// The application ID.
	ApplicationId *string `type:"string"`

	// The ID of the configuration profile.
	Id *string `type:"string"`

	// The URI location of the configuration.
	LocationUri *string `min:"1" type:"string"`

	// The name of the configuration profile.
	Name *string `min:"1" type:"string"`

	// The types of validators in the configuration profile.
	ValidatorTypes []ValidatorType `type:"list"`
}

// String returns the string representation
func (s ConfigurationProfileSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ConfigurationProfileSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LocationUri != nil {
		v := *s.LocationUri

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "LocationUri", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ValidatorTypes != nil {
		v := s.ValidatorTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "ValidatorTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// An object that describes a deployment event.
type DeploymentEvent struct {
	_ struct{} `type:"structure"`

	// A description of the deployment event. Descriptions include, but are not
	// limited to, the user account or the CloudWatch alarm ARN that initiated a
	// rollback, the percentage of hosts that received the deployment, or in the
	// case of an internal error, a recommendation to attempt a new deployment.
	Description *string `type:"string"`

	// The type of deployment event. Deployment event types include the start, stop,
	// or completion of a deployment; a percentage update; the start or stop of
	// a bake period; the start or completion of a rollback.
	EventType DeploymentEventType `type:"string" enum:"true"`

	// The date and time the event occurred.
	OccurredAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The entity that triggered the deployment event. Events can be triggered by
	// a user, AWS AppConfig, an Amazon CloudWatch alarm, or an internal error.
	TriggeredBy TriggeredBy `type:"string" enum:"true"`
}

// String returns the string representation
func (s DeploymentEvent) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeploymentEvent) MarshalFields(e protocol.FieldEncoder) error {
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.EventType) > 0 {
		v := s.EventType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EventType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.OccurredAt != nil {
		v := *s.OccurredAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OccurredAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if len(s.TriggeredBy) > 0 {
		v := s.TriggeredBy

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TriggeredBy", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type DeploymentStrategy struct {
	_ struct{} `type:"structure"`

	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes *int64 `type:"integer"`

	// The description of the deployment strategy.
	Description *string `type:"string"`

	// The amount of time AppConfig monitored for alarms before considering the
	// deployment to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes *int64 `type:"integer"`

	// The percentage of targets that received a deployed configuration during each
	// interval.
	GrowthFactor *float64 `min:"1" type:"float"`

	// The algorithm used to define how percentage grew over time.
	GrowthType GrowthType `type:"string" enum:"true"`

	// The deployment strategy ID.
	Id *string `type:"string"`

	// The name of the deployment strategy.
	Name *string `min:"1" type:"string"`

	// Save the deployment strategy to a Systems Manager (SSM) document.
	ReplicateTo ReplicateTo `type:"string" enum:"true"`
}

// String returns the string representation
func (s DeploymentStrategy) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeploymentStrategy) MarshalFields(e protocol.FieldEncoder) error {
	if s.DeploymentDurationInMinutes != nil {
		v := *s.DeploymentDurationInMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeploymentDurationInMinutes", protocol.Int64Value(v), metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FinalBakeTimeInMinutes != nil {
		v := *s.FinalBakeTimeInMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FinalBakeTimeInMinutes", protocol.Int64Value(v), metadata)
	}
	if s.GrowthFactor != nil {
		v := *s.GrowthFactor

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GrowthFactor", protocol.Float64Value(v), metadata)
	}
	if len(s.GrowthType) > 0 {
		v := s.GrowthType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GrowthType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ReplicateTo) > 0 {
		v := s.ReplicateTo

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ReplicateTo", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Information about the deployment.
type DeploymentSummary struct {
	_ struct{} `type:"structure"`

	// Time the deployment completed.
	CompletedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The name of the configuration.
	ConfigurationName *string `min:"1" type:"string"`

	// The version of the configuration.
	ConfigurationVersion *string `min:"1" type:"string"`

	// Total amount of time the deployment lasted.
	DeploymentDurationInMinutes *int64 `type:"integer"`

	// The sequence number of the deployment.
	DeploymentNumber *int64 `type:"integer"`

	// The amount of time AppConfig monitors for alarms before considering the deployment
	// to be complete and no longer eligible for automatic roll back.
	FinalBakeTimeInMinutes *int64 `type:"integer"`

	// The percentage of targets to receive a deployed configuration during each
	// interval.
	GrowthFactor *float64 `min:"1" type:"float"`

	// The algorithm used to define how percentage grows over time.
	GrowthType GrowthType `type:"string" enum:"true"`

	// The percentage of targets for which the deployment is available.
	PercentageComplete *float64 `min:"1" type:"float"`

	// Time the deployment started.
	StartedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// The state of the deployment.
	State DeploymentState `type:"string" enum:"true"`
}

// String returns the string representation
func (s DeploymentSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeploymentSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.CompletedAt != nil {
		v := *s.CompletedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "CompletedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if s.ConfigurationName != nil {
		v := *s.ConfigurationName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConfigurationName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationVersion != nil {
		v := *s.ConfigurationVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConfigurationVersion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DeploymentDurationInMinutes != nil {
		v := *s.DeploymentDurationInMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeploymentDurationInMinutes", protocol.Int64Value(v), metadata)
	}
	if s.DeploymentNumber != nil {
		v := *s.DeploymentNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "DeploymentNumber", protocol.Int64Value(v), metadata)
	}
	if s.FinalBakeTimeInMinutes != nil {
		v := *s.FinalBakeTimeInMinutes

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "FinalBakeTimeInMinutes", protocol.Int64Value(v), metadata)
	}
	if s.GrowthFactor != nil {
		v := *s.GrowthFactor

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GrowthFactor", protocol.Float64Value(v), metadata)
	}
	if len(s.GrowthType) > 0 {
		v := s.GrowthType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GrowthType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.PercentageComplete != nil {
		v := *s.PercentageComplete

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PercentageComplete", protocol.Float64Value(v), metadata)
	}
	if s.StartedAt != nil {
		v := *s.StartedAt

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "StartedAt",
			protocol.TimeValue{V: v, Format: "iso8601", QuotedFormatTime: true}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type Environment struct {
	_ struct{} `type:"structure"`

	// The application ID.
	ApplicationId *string `type:"string"`

	// The description of the environment.
	Description *string `type:"string"`

	// The environment ID.
	Id *string `type:"string"`

	// Amazon CloudWatch alarms monitored during the deployment.
	Monitors []Monitor `type:"list"`

	// The name of the environment.
	Name *string `min:"1" type:"string"`

	// The state of the environment. An environment can be in one of the following
	// states: READY_FOR_DEPLOYMENT, DEPLOYING, ROLLING_BACK, or ROLLED_BACK
	State EnvironmentState `type:"string" enum:"true"`
}

// String returns the string representation
func (s Environment) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Environment) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Monitors != nil {
		v := s.Monitors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Monitors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.State) > 0 {
		v := s.State

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "State", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Information about the configuration.
type HostedConfigurationVersionSummary struct {
	_ struct{} `type:"structure"`

	// The application ID.
	ApplicationId *string `type:"string"`

	// The configuration profile ID.
	ConfigurationProfileId *string `type:"string"`

	// A standard MIME type describing the format of the configuration content.
	// For more information, see Content-Type (https://docs.aws.amazon.com/https:/www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17).
	ContentType *string `min:"1" type:"string"`

	// A description of the configuration.
	Description *string `type:"string"`

	// The configuration version.
	VersionNumber *int64 `type:"integer"`
}

// String returns the string representation
func (s HostedConfigurationVersionSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s HostedConfigurationVersionSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ApplicationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ConfigurationProfileId != nil {
		v := *s.ConfigurationProfileId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ConfigurationProfileId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ContentType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionNumber != nil {
		v := *s.VersionNumber

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VersionNumber", protocol.Int64Value(v), metadata)
	}
	return nil
}

// Amazon CloudWatch alarms to monitor during the deployment process.
type Monitor struct {
	_ struct{} `type:"structure"`

	// ARN of the Amazon CloudWatch alarm.
	AlarmArn *string `min:"20" type:"string"`

	// ARN of an IAM role for AppConfig to monitor AlarmArn.
	AlarmRoleArn *string `min:"20" type:"string"`
}

// String returns the string representation
func (s Monitor) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Monitor) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Monitor"}
	if s.AlarmArn != nil && len(*s.AlarmArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("AlarmArn", 20))
	}
	if s.AlarmRoleArn != nil && len(*s.AlarmRoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("AlarmRoleArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Monitor) MarshalFields(e protocol.FieldEncoder) error {
	if s.AlarmArn != nil {
		v := *s.AlarmArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AlarmArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AlarmRoleArn != nil {
		v := *s.AlarmRoleArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AlarmRoleArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// A validator provides a syntactic or semantic check to ensure the configuration
// you want to deploy functions as intended. To validate your application configuration
// data, you provide a schema or a Lambda function that runs against the configuration.
// The configuration deployment or update can only proceed when the configuration
// data is valid.
type Validator struct {
	_ struct{} `type:"structure"`

	// Either the JSON Schema content or the Amazon Resource Name (ARN) of an AWS
	// Lambda function.
	//
	// Content is a required field
	Content *string `type:"string" required:"true" sensitive:"true"`

	// AppConfig supports validators of type JSON_SCHEMA and LAMBDA
	//
	// Type is a required field
	Type ValidatorType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s Validator) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Validator) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Validator"}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Validator) MarshalFields(e protocol.FieldEncoder) error {
	if s.Content != nil {
		v := *s.Content

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Content", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}
