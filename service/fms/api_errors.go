// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fms

const (

	// ErrCodeInternalErrorException for service response error code
	// "InternalErrorException".
	//
	// The operation failed because of a system problem, even though the request
	// was valid. Retry your request.
	ErrCodeInternalErrorException = "InternalErrorException"

	// ErrCodeInvalidInputException for service response error code
	// "InvalidInputException".
	//
	// The parameters of the request were invalid.
	ErrCodeInvalidInputException = "InvalidInputException"

	// ErrCodeInvalidOperationException for service response error code
	// "InvalidOperationException".
	//
	// The operation failed because there was nothing to do or the operation wasn't
	// possible. For example, you might have submitted an AssociateAdminAccount
	// request for an account ID that was already set as the AWS Firewall Manager
	// administrator. Or you might have tried to access a Region that's disabled
	// by default, and that you need to enable for the Firewall Manager administrator
	// account and for AWS Organizations before you can access it.
	ErrCodeInvalidOperationException = "InvalidOperationException"

	// ErrCodeInvalidTypeException for service response error code
	// "InvalidTypeException".
	//
	// The value of the Type parameter is invalid.
	ErrCodeInvalidTypeException = "InvalidTypeException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The operation exceeds a resource limit, for example, the maximum number of
	// policy objects that you can create for an AWS account. For more information,
	// see Firewall Manager Limits (https://docs.aws.amazon.com/waf/latest/developerguide/fms-limits.html)
	// in the AWS WAF Developer Guide.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The specified resource was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)
