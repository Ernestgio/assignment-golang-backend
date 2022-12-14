package sentinelerrors

import "errors"

// business error
var ErrEmailAlreadyExists = errors.New("email already exist")
var ErrEmailNotExists = errors.New("email not exist")
var ErrInvalidPassword = errors.New("invalid password")

// standard error
var ErrInvalidRequestBody = errors.New("invalid request body")
var ErrNotFound = errors.New("not found")
