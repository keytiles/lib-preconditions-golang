package kt_precond

import (
	"fmt"
	"strings"

	"github.com/keytiles/lib-errorhandling-golang/v2/pkg/kt_errors"
)

// Use in function to validate string parameter. Returns non-public `kt_errors.ValidationFault` if value is empty naming `paramName` in the error message and
// `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureStringParamNotEmpty(val string, paramName string, source ...string) *kt_errors.FaultBuilder {
	if strings.TrimSpace(val) == "" {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' can not be empty", paramName), source...)
	}
	return nil
}

// Use in function to validate string parameter. Returns public `kt_errors.ValidationFault` if value is empty naming `paramName` in the error message and
// `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamStringNotEmptyPub(val string, paramName string, source ...string) *kt_errors.FaultBuilder {
	if strings.TrimSpace(val) == "" {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' can not be empty", paramName), source...)
	}
	return nil
}

// Use in function to validate string parameter. Returns non-public `kt_errors.ValidationFault` if value does not start with given prefix naming `paramName` in
// the error message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureStringParamStartsWith(val string, prefix string, paramName string, source ...string) *kt_errors.FaultBuilder {
	if !strings.HasPrefix(val, prefix) {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' must start with '%s'", paramName, prefix), source...)
	}
	return nil
}

// Use in function to validate string parameter. Returns public `kt_errors.ValidationFault` if value does not start with given prefix naming `paramName` in the
// error message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamStringStartsWithPub(val string, prefix string, paramName string, source ...string) *kt_errors.FaultBuilder {
	if !strings.HasPrefix(val, prefix) {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' must start with '%s'", paramName, prefix), source...)
	}
	return nil
}

// Use in function to validate string parameter. Returns non-public `kt_errors.ValidationFault` if value does not end with given suffix naming `paramName` in
// the error message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureStringParamEndsWith(val string, suffix string, paramName string, source ...string) *kt_errors.FaultBuilder {
	if !strings.HasSuffix(val, suffix) {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' must end with '%s'", paramName, suffix), source...)
	}
	return nil
}

// Use in function to validate string parameter. Returns public `kt_errors.ValidationFault` if value does not end with given suffix naming `paramName` in the
// error message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamStringEndsWithPub(val string, suffix string, paramName string, source ...string) *kt_errors.FaultBuilder {
	if !strings.HasSuffix(val, suffix) {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' must end with '%s'", paramName, suffix), source...)
	}
	return nil
}

// Use in function to validate numerical parameter. Returns non-public `kt_errors.ValidationFault` if value is not greater than zero naming `paramName` in the
// error message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamGreaterZero[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64](
	val T, paramName string, source ...string,
) *kt_errors.FaultBuilder {
	if val <= 0 {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' must be > 0", paramName), source...)
	}
	return nil
}

// Use in function to validate numerical parameter. Returns public `kt_errors.ValidationFault` if value is not greater than zero naming `paramName` in the error
// message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamGreaterZeroPub[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64](val T, paramName string, source ...string,
) *kt_errors.FaultBuilder {
	if val <= 0 {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' must be > 0", paramName), source...)
	}
	return nil
}

// Use in function to validate numerical parameter. Returns non-public `kt_errors.ValidationFault` if value is less than zero naming `paramName` in the error
// message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamGreaterOrEqualZero[T int | int16 | int32 | int64 | float32 | float64](val T, paramName string, source ...string) *kt_errors.FaultBuilder {
	if val < 0 {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' must be >= 0", paramName), source...)
	}
	return nil
}

// Use in function to validate numerical parameter. Returns public `kt_errors.ValidationFault` if value is less than zero naming `paramName` in the error
// message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamGreaterOrEqualZeroPub[T int | int16 | int32 | int64 | float32 | float64](val T, paramName string, source ...string) *kt_errors.FaultBuilder {
	if val < 0 {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' must be >= 0", paramName), source...)
	}
	return nil
}

// Use in function to validate numerical parameter. Returns non-public `kt_errors.ValidationFault` if value is not less than zero naming `paramName` in the
// error message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamLessZero[T int | int16 | int32 | int64 | float32 | float64](val T, paramName string, source ...string) *kt_errors.FaultBuilder {
	if val >= 0 {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' must be < 0", paramName), source...)
	}
	return nil
}

// Use in function to validate numerical parameter. Returns public `kt_errors.ValidationFault` if value is not less than zero naming `paramName` in the error
// message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamLessZeroPub[T int | int16 | int32 | int64 | float32 | float64](val T, paramName string, source ...string) *kt_errors.FaultBuilder {
	if val >= 0 {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' must be < 0", paramName), source...)
	}
	return nil
}

// Use in function to validate numerical parameter. Returns non-public `kt_errors.ValidationFault` if value is greater than zero naming `paramName` in the error
// message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamLessOrEqualZero[T int | int16 | int32 | int64 | float32 | float64](val T, paramName string, source ...string) *kt_errors.FaultBuilder {
	if val > 0 {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' must be <= 0", paramName), source...)
	}
	return nil
}

// Use in function to validate numerical parameter. Returns public `kt_errors.ValidationFault` if value is greater than zero naming `paramName` in the error
// message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamLessOrEqualZeroPub[T int | int16 | int32 | int64 | float32 | float64](val T, paramName string, source ...string) *kt_errors.FaultBuilder {
	if val > 0 {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' must be <= 0", paramName), source...)
	}
	return nil
}

// Use in function to validate any pointer parameter. Returns non-public `kt_errors.ValidationFault` if pointer is Nil naming `paramName` in the error message
// and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamNonNil(ptr any, paramName string, source ...string) *kt_errors.FaultBuilder {
	if ptr == nil {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' is mandatory and can not be Nil", paramName), source...)
	}
	return nil
}

// Use in function to validate any pointer parameter. Returns public `kt_errors.ValidationFault` if pointer is Nil naming `paramName` in the error message and
// `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamNonNilPub(ptr any, paramName string, source ...string) *kt_errors.FaultBuilder {
	if ptr == nil {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' is mandatory and can not be Nil", paramName), source...)
	}
	return nil
}

// Use in function to validate any Slice or Map parameter. Returns non-public `kt_errors.ValidationFault` if value is empty naming `paramName` in the error
// message and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureParamNonEmpty[T []any | map[any]any](container T, paramName string, source ...string) *kt_errors.FaultBuilder {
	if len(container) == 0 {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' can not be empty", paramName), source...)
	}
	return nil
}

// Use in function to validate any Slice or Map parameter. Returns public `kt_errors.ValidationFault` if value is empty naming `paramName` in the error message
// and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureNonEmptyPub[T []any | map[any]any](container T, paramName string, source ...string) *kt_errors.FaultBuilder {
	if len(container) == 0 {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' can not be empty", paramName), source...)
	}
	return nil
}
