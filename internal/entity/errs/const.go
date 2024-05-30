package errs

const (
	ErrProcessingData = ErrTypeProcessingData("Failed to process data")
	ErrInvalidData    = ErrTypeInvalidData("Invalid Data")
	ErrUnauthorized   = ErrTypeUnauthorized("Unauthorized")
	ErrInvalidToken   = ErrTypeInvalidToken("Invalid Token")
	ErrExpiredToken   = ErrTypeExpiredToken("Expired Token")
	ErrForbidden      = ErrTypeForbidden("Forbidden")
	ErrNotFound       = ErrTypeNotFound("Not Found")
	ErrUnprocessable  = ErrTypeUnprocessable("Unprocessable")
	ErrInternal       = ErrTypeInternal("Internal Server Error")
)
