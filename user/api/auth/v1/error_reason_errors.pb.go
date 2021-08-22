// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsInvalidCredentials(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_CREDENTIALS.String() && e.Code == 404
}

func ErrorInvalidCredentials(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_INVALID_CREDENTIALS.String(), fmt.Sprintf(format, args...))
}
