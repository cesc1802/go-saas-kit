// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsConfirmPasswordMismatch(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONFIRM_PASSWORD_MISMATCH.String() && e.Code == 400
}

func ErrorConfirmPasswordMismatch(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_CONFIRM_PASSWORD_MISMATCH.String(), fmt.Sprintf(format, args...))
}

func IsPasswordInsufficientStrength(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PASSWORD_INSUFFICIENT_STRENGTH.String() && e.Code == 400
}

func ErrorPasswordInsufficientStrength(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_PASSWORD_INSUFFICIENT_STRENGTH.String(), fmt.Sprintf(format, args...))
}

func IsInvalidPassword(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_PASSWORD.String() && e.Code == 400
}

func ErrorInvalidPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_INVALID_PASSWORD.String(), fmt.Sprintf(format, args...))
}

func IsDuplicateUsername(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DUPLICATE_USERNAME.String() && e.Code == 400
}

func ErrorDuplicateUsername(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DUPLICATE_USERNAME.String(), fmt.Sprintf(format, args...))
}

func IsDuplicateEmail(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DUPLICATE_EMAIL.String() && e.Code == 400
}

func ErrorDuplicateEmail(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DUPLICATE_EMAIL.String(), fmt.Sprintf(format, args...))
}

func IsDuplicatePhone(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DUPLICATE_PHONE.String() && e.Code == 400
}

func ErrorDuplicatePhone(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DUPLICATE_PHONE.String(), fmt.Sprintf(format, args...))
}
