package kt_precond

import "github.com/keytiles/lib-errorhandling-golang/v2/pkg/kt_errors"

func GetBaseValidationFaultBuilder(msgTpl string, source ...string) *kt_errors.FaultBuilder {
	return kt_errors.NewFaultBuilder(kt_errors.ValidationFault).
		WithSource(source...).
		WithErrorCodes(kt_errors.VALIDATION_ERRCODE_INVALID_VALUE).
		WithMessageTemplate(msgTpl)
}

func GetBasePublicValidationFaultBuilder(msgTpl string, source ...string) *kt_errors.FaultBuilder {
	return kt_errors.NewPublicFaultBuilder(kt_errors.ValidationFault).
		WithSource(source...).
		WithErrorCodes(kt_errors.VALIDATION_ERRCODE_INVALID_VALUE).
		WithMessageTemplate(msgTpl)
}

func GetBaseIllegalStateFaultBuilder(msgTpl string, source ...string) *kt_errors.FaultBuilder {
	return kt_errors.NewFaultBuilder(kt_errors.IllegalStateFault).
		WithSource(source...).
		WithErrorCodes(kt_errors.ILLEGALSTATE_ERRCODE_EXCPECTATION_FAILED).
		WithMessageTemplate(msgTpl)
}

func GetBasePublicIllegalStateFaultBuilder(msgTpl string, source ...string) *kt_errors.FaultBuilder {
	return kt_errors.NewFaultBuilder(kt_errors.IllegalStateFault).
		WithSource(source...).
		WithErrorCodes(kt_errors.ILLEGALSTATE_ERRCODE_EXCPECTATION_FAILED).
		WithMessageTemplate(msgTpl)
}
