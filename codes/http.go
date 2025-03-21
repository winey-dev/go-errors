package codes

import "net/http"

func ToHTTPStatus(code Code) int {
	switch code {
	case OK:
		return http.StatusOK
	case NotUse:
		return http.StatusInternalServerError // 적절한 HTTP 상태 코드가 없으므로 InternalServerError로 매핑
	case Unknown:
		return http.StatusInternalServerError
	case InvalidArgument:
		return http.StatusBadRequest
	case InvalidRequest:
		return http.StatusBadRequest
	case NotFound:
		return http.StatusNotFound
	case Internal:
		return http.StatusInternalServerError
	case Unauthorized:
		return http.StatusUnauthorized
	case Forbidden:
		return http.StatusForbidden
	case Timeout:
		return http.StatusRequestTimeout
	case Deadlock:
		return http.StatusConflict
	case ConnectionRefused:
		return http.StatusServiceUnavailable
	case Disconnected:
		return http.StatusServiceUnavailable
	case ResourceExhausted:
		return http.StatusInsufficientStorage
	case TooManyRequests:
		return http.StatusTooManyRequests
	default:
		return http.StatusInternalServerError
	}
}
