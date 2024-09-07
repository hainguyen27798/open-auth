package response

// 200xx: success
// 400xx: bad request err

const (
	CodeSuccess          = 20000 // Success
	LoginSuccess         = 20001 // Login success
	LogoutSuccess        = 20002 // Logout success
	CreatedSuccess       = 20003 // Created success
	ErrBadRequest        = 40000 // Bad Request
	ErrCodeParamInvalid  = 40001 // Param invalid
	ErrCreateFailed      = 40003 // Create failed
	ErrInvalidOTP        = 40004 // OTP is invalid
	ErrSendEmailFailed   = 40005 // Send mail failed
	ErrCodeUserHasExists = 40006 // User has already exists
	ErrCodeLoginFailed   = 40007 // Login credential is incorrect
	ErrUnauthorized      = 40101 // Unauthorized
	ErrInvalidToken      = 40102 // Token is invalid
	ErrExpiredToken      = 40103 // Token is Expired
	ErrStolenToken       = 40104 // Token was stolen
	ErrNotFound          = 40400 // Resource not found
	ErrCodeUserNotExists = 40401 // User is not exists
	ErrJWTInternalError  = 50001 // JWT internal error
)

var CodeMsg = map[int]string{
	CodeSuccess:          "Success",
	LoginSuccess:         "Login Success",
	LogoutSuccess:        "Logout Success",
	CreatedSuccess:       "Created Success",
	ErrBadRequest:        "Bad Request",
	ErrCodeParamInvalid:  "Param invalid",
	ErrInvalidOTP:        "OTP is invalid",
	ErrSendEmailFailed:   "Send email failed",
	ErrCreateFailed:      "Create failed",
	ErrCodeUserHasExists: "User has already exists",
	ErrCodeLoginFailed:   "Login credential is incorrect",
	ErrUnauthorized:      "Unauthorized",
	ErrInvalidToken:      "Token is invalid",
	ErrExpiredToken:      "Token is expired",
	ErrStolenToken:       "Token is stolen",
	ErrNotFound:          "Resource not found",
	ErrCodeUserNotExists: "User not exists",
	ErrJWTInternalError:  "JWT internal error",
}
