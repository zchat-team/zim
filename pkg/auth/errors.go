package auth

import (
	"errors"
)

var (
	// ErrMissingToken can be thrown by follow
	// if authing with a HTTP header, the Auth header needs to be set
	// if authing with URL Query, the query token variable is empty
	// if authing with a cookie, the token cookie is empty
	// if authing with parameter in path, the parameter in path is empty
	ErrMissingToken = errors.New("auth token is empty")
	// ErrInvalidAuthHeader indicates auth header is invalid
	ErrInvalidAuthHeader = errors.New("auth header is invalid")
	ErrEncodingToken     = errors.New("error encoding the token")
	ErrInvalidToken      = errors.New("invalid token provided")
	ErrTokenExpired      = errors.New("token expired")
)
