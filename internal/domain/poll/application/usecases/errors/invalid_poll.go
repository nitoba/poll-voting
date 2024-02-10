package errors

import "errors"

var ErrInvalidPoll = errors.New("invalid poll, polls must have at least 2 options")
