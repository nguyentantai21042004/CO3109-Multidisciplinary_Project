package response

const (
	// Default depth of stack trace
	defaultStackTraceDepth = 32
	// DefaultErrorMessage is the default error message.
	DefaultErrorMessage = "Something went wrong"
	// ValidationErrorCode is the validation error code.
	ValidationErrorCode = 400
	// ValidationErrorMsg is the validation error message.
	ValidationErrorMsg = "Validation error"
)

// permission
const (
	PermissionErrorMsg = "You don't have permission to do this"
	// PermissionErrorCode is the permission error code.
	PermissionErrorCode = 403
)

// payment
const (
	PaymentErrorMsg = "Payment error"
	// PaymentErrorCode is the payment error code.
	PaymentErrorCode = 403
)
