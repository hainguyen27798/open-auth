package response

const (
	ErrCodeSuccess       = 20001 // Success
	ErrCodeParamInvalid  = 2003  // Param invalid
	ErrInvalidToken      = 30001 // Token is invalid
	ErrInvalidOTP        = 30002 // OTP is invalid
	ErrSendEmailFailed   = 30003 // Send mail failed
	ErrCreateFailed      = 40001 // Create failed
	ErrCodeUserHasExists = 50001 // User has already exists
)

var ErrCodeMsg = map[int]string{
	ErrCodeSuccess:       "Success",
	ErrCodeParamInvalid:  "Param invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrInvalidOTP:        "OTP is invalid",
	ErrSendEmailFailed:   "Send email failed",
	ErrCreateFailed:      "Create failed",
	ErrCodeUserHasExists: "User has already exists",
}
