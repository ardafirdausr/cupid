package errs

type ErrTypeProcessingData string

func (s ErrTypeProcessingData) Error() string { return string(s) }

type ErrTypeInvalidData string

func (s ErrTypeInvalidData) Error() string { return string(s) }

type ErrTypeUnauthorized string

func (s ErrTypeUnauthorized) Error() string { return string(s) }

type ErrTypeInvalidToken string

func (s ErrTypeInvalidToken) Error() string { return string(s) }

type ErrTypeExpiredToken string

func (s ErrTypeExpiredToken) Error() string { return string(s) }

type ErrTypeForbidden string

func (s ErrTypeForbidden) Error() string { return string(s) }

type ErrTypeNotFound string

func (s ErrTypeNotFound) Error() string { return string(s) }

type ErrTypeUnprocessable string

func (s ErrTypeUnprocessable) Error() string { return string(s) }

type ErrTypeInternal string

func (s ErrTypeInternal) Error() string { return string(s) }
