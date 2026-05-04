package kt_precond

import (
	"fmt"
	"strings"

	"github.com/keytiles/lib-errorhandling-golang/v2/pkg/kt_errors"
)

// Use in function to validate string parameter. Returns non-public `kt_errors.ValidationFault` if value is empty.
func EnsureStringParamNotEmpty(val string, paramName string, source ...string) kt_errors.Fault {
	if strings.TrimSpace(val) == "" {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' can not be empty", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate string parameter. Returns public `kt_errors.ValidationFault` if value is empty.
func EnsureParamStringNotEmptyPub(val string, paramName string, source ...string) kt_errors.Fault {
	if strings.TrimSpace(val) == "" {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' can not be empty", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate string parameter. Returns non-public `kt_errors.ValidationFault` if value does not start with given prefix.
func EnsureStringParamStartsWith(val string, prefix string, paramName string, source ...string) kt_errors.Fault {
	if !strings.HasPrefix(val, prefix) {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' must start with '%s'", paramName, prefix), source...).Build()
	}
	return nil
}

// Use in function to validate string parameter. Returns public `kt_errors.ValidationFault` if value does not start with given prefix.
func EnsureParamStringStartsWithPub(val string, prefix string, paramName string, source ...string) kt_errors.Fault {
	if !strings.HasPrefix(val, prefix) {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' must start with '%s'", paramName, prefix), source...).Build()
	}
	return nil
}

// Use in function to validate string parameter. Returns non-public `kt_errors.ValidationFault` if value does not end with given suffix.
func EnsureStringParamEndsWith(val string, suffix string, paramName string, source ...string) kt_errors.Fault {
	if !strings.HasSuffix(val, suffix) {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' must end with '%s'", paramName, suffix), source...).Build()
	}
	return nil
}

// Use in function to validate string parameter. Returns public `kt_errors.ValidationFault` if value does not end with given suffix.
func EnsureParamStringEndsWithPub(val string, suffix string, paramName string, source ...string) kt_errors.Fault {
	if !strings.HasSuffix(val, suffix) {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' must end with '%s'", paramName, suffix), source...).Build()
	}
	return nil
}

// Use in function to validate numerical parameter. Returns non-public `kt_errors.ValidationFault` if value is not greater than zero.
func EnsureParamGreaterZero[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64](
	val T, paramName string, source ...string,
) kt_errors.Fault {
	if val <= 0 {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' must be > 0", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate numerical parameter. Returns public `kt_errors.ValidationFault` if value is not greater than zero.
func EnsureParamGreaterZeroPub[T int | int16 | int32 | int64 | uint | uint16 | uint32 | uint64 | float32 | float64](val T, paramName string, source ...string,
) kt_errors.Fault {
	if val <= 0 {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' must be > 0", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate numerical parameter. Returns non-public `kt_errors.ValidationFault` if value is less than zero.
func EnsureParamGreaterOrEqualZero[T int | int16 | int32 | int64 | float32 | float64](val T, paramName string, source ...string) kt_errors.Fault {
	if val < 0 {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' must be >= 0", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate numerical parameter. Returns public `kt_errors.ValidationFault` if value is less than zero.
func EnsureParamGreaterOrEqualZeroPub[T int | int16 | int32 | int64 | float32 | float64](val T, paramName string, source ...string) kt_errors.Fault {
	if val < 0 {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' must be >= 0", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate numerical parameter. Returns non-public `kt_errors.ValidationFault` if value is not less than zero.
func EnsureParamLessZero[T int | int16 | int32 | int64 | float32 | float64](val T, paramName string, source ...string) kt_errors.Fault {
	if val >= 0 {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' must be < 0", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate numerical parameter. Returns public `kt_errors.ValidationFault` if value is not less than zero.
func EnsureParamLessZeroPub[T int | int16 | int32 | int64 | float32 | float64](val T, paramName string, source ...string) kt_errors.Fault {
	if val >= 0 {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' must be < 0", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate numerical parameter. Returns non-public `kt_errors.ValidationFault` if value is greater than zero.
func EnsureParamLessOrEqualZero[T int | int16 | int32 | int64 | float32 | float64](val T, paramName string, source ...string) kt_errors.Fault {
	if val > 0 {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' must be <= 0", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate numerical parameter. Returns public `kt_errors.ValidationFault` if value is greater than zero.
func EnsureParamLessOrEqualZeroPub[T int | int16 | int32 | int64 | float32 | float64](val T, paramName string, source ...string) kt_errors.Fault {
	if val > 0 {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' must be <= 0", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate any pointer parameter. Returns non-public `kt_errors.ValidationFault` if pointer is Nil.
func EnsureParamNonNil[T *any](ptr T, paramName string, source ...string) kt_errors.Fault {
	if ptr == nil {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' is mandatory and can not be Nil", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate any pointer parameter. Returns public `kt_errors.ValidationFault` if pointer is Nil.
func EnsureParamNonNilPub[T *any](ptr T, paramName string, source ...string) kt_errors.Fault {
	if ptr == nil {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' is mandatory and can not be Nil", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate any Slice or Map parameter. Returns non-public `kt_errors.ValidationFault` if value is empty.
func EnsureParamNonEmpty[T []any | map[any]any](container T, paramName string, source ...string) kt_errors.Fault {
	if len(container) == 0 {
		return GetBaseValidationFaultBuilder(fmt.Sprintf("'%s' can not be empty", paramName), source...).Build()
	}
	return nil
}

// Use in function to validate any Slice or Map parameter. Returns public `kt_errors.ValidationFault` if value is empty.
func EnsureNonEmptyPub[T []any | map[any]any](container T, paramName string, source ...string) kt_errors.Fault {
	if len(container) == 0 {
		return GetBasePublicValidationFaultBuilder(fmt.Sprintf("'%s' can not be empty", paramName), source...).Build()
	}
	return nil
}
