package apperr

type ErrorCode int

const (
	ErrBadReqeust ErrorCode = iota
	ErrNotFound
	ErrUnauthorized
	ErrForbidden
	ErrInternalError
)

func (e ErrorCode) String() string {
	switch e {
	case ErrBadReqeust:
		return "BadRequest"
	case ErrNotFound:
		return "NotFound"
	case ErrUnauthorized:
		return "Unauthorized"
	case ErrForbidden:
		return "Forbidden"
	case ErrInternalError:
		return "InternalServerError"
	default:
		return "InternalServerError"
	}
}
