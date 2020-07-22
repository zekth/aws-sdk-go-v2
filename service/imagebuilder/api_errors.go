// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

const (

	// ErrCodeCallRateLimitExceededException for service response error code
	// "CallRateLimitExceededException".
	//
	// You have exceeded the permitted request rate for the specific operation.
	ErrCodeCallRateLimitExceededException = "CallRateLimitExceededException"

	// ErrCodeException for service response error code
	// "Exception".
	//
	// These errors are usually caused by a client action, such as using an action
	// or resource on behalf of a user that doesn't have permissions to use the
	// action or resource, or specifying an invalid resource identifier.
	ErrCodeException = "Exception"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	//
	// You are not authorized to perform the requested operation.
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeIdempotentParameterMismatchException for service response error code
	// "IdempotentParameterMismatchException".
	//
	// You have specified a client token for an operation using parameter values
	// that differ from a previous request that used the same client token.
	ErrCodeIdempotentParameterMismatchException = "IdempotentParameterMismatchException"

	// ErrCodeInvalidPaginationTokenException for service response error code
	// "InvalidPaginationTokenException".
	//
	// You have provided an invalid pagination token in your request.
	ErrCodeInvalidPaginationTokenException = "InvalidPaginationTokenException"

	// ErrCodeInvalidParameterCombinationException for service response error code
	// "InvalidParameterCombinationException".
	//
	// You have specified two or more mutually exclusive parameters. Review the
	// error message for details.
	ErrCodeInvalidParameterCombinationException = "InvalidParameterCombinationException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// The specified parameter is invalid. Review the available parameters for the
	// API request.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValueException".
	//
	// The value that you provided for the specified parameter is invalid.
	ErrCodeInvalidParameterValueException = "InvalidParameterValueException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// You have made a request for an action that is not supported by the service.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeInvalidVersionNumberException for service response error code
	// "InvalidVersionNumberException".
	//
	// Your version number is out of bounds or does not follow the required syntax.
	ErrCodeInvalidVersionNumberException = "InvalidVersionNumberException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	//
	// The resource that you are trying to create already exists.
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceDependencyException for service response error code
	// "ResourceDependencyException".
	//
	// You have attempted to mutate or delete a resource with a dependency that
	// prohibits this action. See the error message for more details.
	ErrCodeResourceDependencyException = "ResourceDependencyException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The resource that you are trying to operate on is currently in use. Review
	// the message details and retry later.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// At least one of the resources referenced by your request does not exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceException for service response error code
	// "ServiceException".
	//
	// This exception is thrown when the service encounters an unrecoverable exception.
	ErrCodeServiceException = "ServiceException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// You have exceeded the number of permitted resources or operations for this
	// service. For service quotas, see EC2 Image Builder endpoints and quotas (https://docs.aws.amazon.com/general/latest/gr/imagebuilder.html#limits_imagebuilder).
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is unable to process your request at this time.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"
)
