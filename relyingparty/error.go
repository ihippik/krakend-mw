package relyingparty

type err struct {
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
	tokenExpiredErr = &err{
		Code:   tokenExpired,
		ErrMsg: "token expired",
	}

	invalidUserIDErr = &err{
		Code:   invalidTokenClaims,
		ErrMsg: "invalid user id err",
	}

	invalidUserRoleErr = &err{
		Code:   invalidTokenClaims,
		ErrMsg: "user role not exists",
	}

	tokenNotExists = &err{
		Code:   invalidToken,
		ErrMsg: "invalid user id err",
	}

	accessDenied = &err{
		Code:   roleNotMatch,
		ErrMsg: "access for the role denied",
	}
)

func newErr(code string, msg string) *err {
	return &err{
		Code:   code,
		ErrMsg: msg,
	}
}