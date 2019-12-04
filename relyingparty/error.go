package relyingparty

type Err struct {
	Code   string `json:"code"`
	ErrMsg string `json:"errMessage"`
}

const (
	invalidTokenClaims = "INVALID_TOKEN_CLAIMS"
	tokenExpired       = "TOKEN_EXPIRED"
	invalidToken       = "INVALID_TOKEN"
	roleNotMatch       = "ROLE_NOT_MATCH"
)

// custom errors.
var (
	tokenExpiredErr = &Err{
		Code:   tokenExpired,
		ErrMsg: "token expired",
	}

	invalidUserIDErr = &Err{
		Code:   invalidTokenClaims,
		ErrMsg: "invalid user id err",
	}

	invalidUserRoleErr = &Err{
		Code:   invalidTokenClaims,
		ErrMsg: "user role not exists",
	}

	accessDenied = &Err{
		Code:   roleNotMatch,
		ErrMsg: "access for the role denied",
	}
)

// newErr create new err with specific code & msg.
func newErr(code string, msg string) *Err {
	return &Err{
		Code:   code,
		ErrMsg: msg,
	}
}
