package common

import "errors"

type InternalError struct {
	Details string
}

func (e *InternalError) Error() string {
	return e.Details
}

func IsErrInternal(err error) bool {
	var checkErr *InternalError
	return errors.As(err, &checkErr)
}

type BadRequestError struct {
	Details string
}

func (e *BadRequestError) Error() string {
	return e.Details
}

func IsErrBadRequest(err error) bool {
	var checkErr *InternalError
	return errors.As(err, &checkErr)
}
