package appErrors

import "errors"

var (
	ErrInvalidFieldLength      error = errors.New("wrong length of the field")
	ErrInvalidDepartmentNumber error = errors.New("department has same incorrect number")
	ErrInvalidTime             error = errors.New("detected time travel!")
	ErrInvalidArguments        error = errors.New("invalid arguments")
)
