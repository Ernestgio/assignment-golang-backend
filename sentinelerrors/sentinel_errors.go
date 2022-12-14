package sentinelerrors

import "errors"

// business error
var ErrEmailAlreadyExists = errors.New("email already exist")
var ErrEmailNotExists = errors.New("email not exist")
var ErrInvalidPassword = errors.New("invalid password")

// standard error
var ErrInternalServerError = errors.New("internal server error")
var ErrInvalidToken = errors.New("invalid token")
var ErrInvalidRequestBody = errors.New("invalid request body")
var ErrNotFound = errors.New("not found")
