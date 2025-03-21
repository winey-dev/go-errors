package codes

import "github.com/lib/pq"

func FromPostgreSQLCode(code pq.ErrorCode) Code {
	switch code {
	case "23505": // unique_violation
		return InvalidRequest
	default:
		return Unknown
	}
}
