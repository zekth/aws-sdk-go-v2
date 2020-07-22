// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Allows you to filter your list of secrets.
type Filter struct {
	_ struct{} `type:"structure"`

	// Filters your list of secrets by a specific key.
	Key FilterNameStringType `type:"string" enum:"true"`

	// Filters your list of secrets by a specific value.
	Values []string `min:"1" type:"list"`
}

// String returns the string representation
func (s Filter) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Filter) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Filter"}
	if s.Values != nil && len(s.Values) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Values", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A structure that defines the rotation configuration for the secret.
type RotationRulesType struct {
	_ struct{} `type:"structure"`

	// Specifies the number of days between automatic scheduled rotations of the
	// secret.
	//
	// Secrets Manager schedules the next rotation when the previous one is complete.
	// Secrets Manager schedules the date by adding the rotation interval (number
	// of days) to the actual date of the last rotation. The service chooses the
	// hour within that 24-hour date window randomly. The minute is also chosen
	// somewhat randomly, but weighted towards the top of the hour and influenced
	// by a variety of factors that help distribute load.
	AutomaticallyAfterDays *int64 `min:"1" type:"long"`
}

// String returns the string representation
func (s RotationRulesType) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RotationRulesType) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RotationRulesType"}
	if s.AutomaticallyAfterDays != nil && *s.AutomaticallyAfterDays < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("AutomaticallyAfterDays", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A structure that contains the details about a secret. It does not include
// the encrypted SecretString and SecretBinary values. To get those values,
// use the GetSecretValue operation.
type SecretListEntry struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the secret.
	//
	// For more information about ARNs in Secrets Manager, see Policy Resources
	// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/reference_iam-permissions.html#iam-resources)
	// in the AWS Secrets Manager User Guide.
	ARN *string `min:"20" type:"string"`

	// The date and time when a secret was created.
	CreatedDate *time.Time `type:"timestamp"`

	// The date and time the deletion of the secret occurred. Not present on active
	// secrets. The secret can be recovered until the number of days in the recovery
	// window has passed, as specified in the RecoveryWindowInDays parameter of
	// the DeleteSecret operation.
	DeletedDate *time.Time `type:"timestamp"`

	// The user-provided description of the secret.
	Description *string `type:"string"`

	// The ARN or alias of the AWS KMS customer master key (CMK) used to encrypt
	// the SecretString and SecretBinary fields in each version of the secret. If
	// you don't provide a key, then Secrets Manager defaults to encrypting the
	// secret fields with the default KMS CMK, the key named awssecretsmanager,
	// for this account.
	KmsKeyId *string `type:"string"`

	// The last date that this secret was accessed. This value is truncated to midnight
	// of the date and therefore shows only the date, not the time.
	LastAccessedDate *time.Time `type:"timestamp"`

	// The last date and time that this secret was modified in any way.
	LastChangedDate *time.Time `type:"timestamp"`

	// The last date and time that the rotation process for this secret was invoked.
	LastRotatedDate *time.Time `type:"timestamp"`

	// The friendly name of the secret. You can use forward slashes in the name
	// to represent a path hierarchy. For example, /prod/databases/dbserver1 could
	// represent the secret for a server named dbserver1 in the folder databases
	// in the folder prod.
	Name *string `min:"1" type:"string"`

	// Returns the name of the service that created the secret.
	OwningService *string `min:"1" type:"string"`

	// Indicates whether automatic, scheduled rotation is enabled for this secret.
	RotationEnabled *bool `type:"boolean"`

	// The ARN of an AWS Lambda function invoked by Secrets Manager to rotate and
	// expire the secret either automatically per the schedule or manually by a
	// call to RotateSecret.
	RotationLambdaARN *string `type:"string"`

	// A structure that defines the rotation configuration for the secret.
	RotationRules *RotationRulesType `type:"structure"`

	// A list of all of the currently assigned SecretVersionStage staging labels
	// and the SecretVersionId attached to each one. Staging labels are used to
	// keep track of the different versions during the rotation process.
	//
	// A version that does not have any SecretVersionStage is considered deprecated
	// and subject to deletion. Such versions are not included in this list.
	SecretVersionsToStages map[string][]string `type:"map"`

	// The list of user-defined tags associated with the secret. To add tags to
	// a secret, use TagResource. To remove tags, use UntagResource.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s SecretListEntry) String() string {
	return awsutil.Prettify(s)
}

// A structure that contains information about one version of a secret.
type SecretVersionsListEntry struct {
	_ struct{} `type:"structure"`

	// The date and time this version of the secret was created.
	CreatedDate *time.Time `type:"timestamp"`

	// The date that this version of the secret was last accessed. Note that the
	// resolution of this field is at the date level and does not include the time.
	LastAccessedDate *time.Time `type:"timestamp"`

	// The unique version identifier of this version of the secret.
	VersionId *string `min:"32" type:"string"`

	// An array of staging labels that are currently associated with this version
	// of the secret.
	VersionStages []string `min:"1" type:"list"`
}

// String returns the string representation
func (s SecretVersionsListEntry) String() string {
	return awsutil.Prettify(s)
}

// A structure that contains information about a tag.
type Tag struct {
	_ struct{} `type:"structure"`

	// The key identifier, or name, of the tag.
	Key *string `min:"1" type:"string"`

	// The string value associated with the key of the tag.
	Value *string `type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Displays errors that occurred during validation of the resource policy.
type ValidationErrorsEntry struct {
	_ struct{} `type:"structure"`

	// Checks the name of the policy.
	CheckName *string `min:"1" type:"string"`

	// Displays error messages if validation encounters problems during validation
	// of the resource policy.
	ErrorMessage *string `type:"string"`
}

// String returns the string representation
func (s ValidationErrorsEntry) String() string {
	return awsutil.Prettify(s)
}
