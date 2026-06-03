package appErrors

import "errors"

var (
	ErrInvalidFieldLength      error = errors.New("wrong length of the field")
	ErrInvalidDepartmentNumber error = errors.New("invalid department number")
	ErrInvalidTime             error = errors.New("detected time travel!")
	ErrInvalidArguments        error = errors.New("invalid arguments")
	ErrInvalidMode             error = errors.New("invalid mode")
)
