package kt_precond

import "github.com/keytiles/lib-errorhandling-golang/v2/pkg/kt_errors"

// Checking the given `expectation` which should be True. Returns non-public `kt_errors.IllegalStateFault` if it is False with `msgTpl` message template
// and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureState(expectation bool, msgTpl string, source ...string) *kt_errors.FaultBuilder {
	if !expectation {
		return GetBaseIllegalStateFaultBuilder(msgTpl, source...)
	}
	return nil
}

// Checking the given `expectation` which should be True. Returns public `kt_errors.IllegalStateFault` if it is False with `msgTpl` message template
// and `source` as source (if you pass multiple pieces they are concatenated with "." char)
func EnsureStatePub(expectation bool, msgTpl string, source ...string) *kt_errors.FaultBuilder {
	if !expectation {
		return GetBaseIllegalStateFaultBuilder(msgTpl, source...)
	}
	return nil
}
