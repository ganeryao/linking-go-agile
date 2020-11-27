package constants

import "errors"

// Errors that can occur during message handling.
var (
	ErrWrongValueType = errors.New("protos: convert on wrong type value")
)
