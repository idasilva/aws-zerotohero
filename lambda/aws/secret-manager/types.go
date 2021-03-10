package secret_manager

import secret "github.com/aws/aws-sdk-go/service/secretsmanager"

const (

	//DecryptionFailure Manager can't decrypt the protected secret text using the provided KMS key.
	DecryptionFailure Manager = secret.ErrCodeDecryptionFailure

	//InternalService  an error occurred on the server side.
	InternalService Manager = secret.ErrCodeInternalServiceError

	//InvalidParameter  you provided an invalid value for a parameter.
	InvalidParameter Manager = secret.ErrCodeInvalidParameterException

	//InvalidRequest you provided a parameter value that is not valid for the current state of the resource.
	InvalidRequest Manager = secret.ErrCodeInvalidRequestException

	//ResourceNotFound  we can't find the resource that you asked for.
	ResourceNotFound Manager = secret.ErrCodeResourceNotFoundException
)

type Manager string
