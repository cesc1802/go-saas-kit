// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsDuplicateTenantName(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DUPLICATE_TENANT_NAME.String() && e.Code == 400
}

func ErrorDuplicateTenantName(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DUPLICATE_TENANT_NAME.String(), fmt.Sprintf(format, args...))
}
