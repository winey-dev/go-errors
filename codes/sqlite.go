package codes

func FromSQLITECode(code int) Code {
	switch code {
	case 0:
		return OK
	case 1:
		return Unknown
	case 2:
		return Internal
	case 3:
		return InvalidArgument
	case 4:
		return InvalidRequest
	case 5:
		return NotFound
	case 6:
		return Internal
	case 7:
		return Unauthorized
	case 8:
		return Forbidden
	case 9:
		return Timeout
	case 10:
		return Deadlock
	case 11:
		return ConnectionRefused
	case 12:
		return Disconnected
	case 13:
		return ResourceExhausted
	case 14:
		return TooManyRequests
	default:
		return Unknown
	}
}
