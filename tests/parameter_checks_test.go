package tests

import (
	"testing"

	"github.com/keytiles/lib-errorhandling-golang/v2/pkg/kt_errors"
	"github.com/keytiles/lib-preconditions-golang/pkg/kt_precond"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	name string
}

type TestIface interface {
	GetName() string
}

func Test_nil_check(t *testing.T) {
	var fb *kt_errors.FaultBuilder

	// ---- GIVEN
	var structPtr *TestStruct
	// ---- WHEN
	fb = kt_precond.EnsureParamNonNil(structPtr, "structPtr", "tests")
	// ---- THEN
	assert.NotNil(t, fb)

	// ---- GIVEN
	var ifaceParam TestIface
	// ---- WHEN
	fb = kt_precond.EnsureParamNonNil(ifaceParam, "ifaceParam", "tests")
	// ---- THEN
	assert.NotNil(t, fb)

	// ---- GIVEN
	var mapParam map[string]string
	// ---- WHEN
	fb = kt_precond.EnsureParamNonNil(mapParam, "mapParam", "tests")
	// ---- THEN
	assert.NotNil(t, fb)

	// ---- GIVEN
	var sliceParam []string
	// ---- WHEN
	fb = kt_precond.EnsureParamNonNil(sliceParam, "sliceParam", "tests")
	// ---- THEN
	assert.NotNil(t, fb)

}
