package response

// 200xx: success
// 400xx: bad request err

const (
	CodeSuccess          = 20000 // Success
	ErrBadRequest        = 40000 // Bad Request
	ErrCodeParamInvalid  = 40001 // Param invalid
	ErrInvalidToken      = 40002 // Token is invalid
	ErrCreateFailed      = 40003 // Create failed
	ErrInvalidOTP        = 40004 // OTP is invalid
	ErrSendEmailFailed   = 40005 // Send mail failed
	ErrCodeUserHasExists = 40006 // User has already exists
)

var CodeMsg = map[int]string{
	CodeSuccess:          "Success",
	ErrBadRequest:        "Bad Request",
	ErrCodeParamInvalid:  "Param invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrInvalidOTP:        "OTP is invalid",
	ErrSendEmailFailed:   "Send email failed",
	ErrCreateFailed:      "Create failed",
	ErrCodeUserHasExists: "User has already exists",
}
