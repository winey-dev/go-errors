package codes

import "google.golang.org/grpc/codes"

func ToGRPCCode(code Code) codes.Code {
	switch code {
	case OK:
		return codes.OK
	case NotUse:
		return codes.Unknown // 적절한 gRPC 코드가 없으므로 Unknown으로 매핑
	case Unknown:
		return codes.Unknown
	case InvalidArgument:
		return codes.InvalidArgument
	case InvalidRequest:
		return codes.InvalidArgument // InvalidRequest를 InvalidArgument로 매핑
	case NotFound:
		return codes.NotFound
	case Internal:
		return codes.Internal
	case Unauthorized:
		return codes.Unauthenticated
	case Forbidden:
		return codes.PermissionDenied
	case Timeout:
		return codes.DeadlineExceeded
	case Deadlock:
		return codes.Aborted
	case ConnectionRefused:
		return codes.Unavailable
	case Disconnected:
		return codes.Unavailable
	case ResourceExhausted:
		return codes.ResourceExhausted
	case TooManyRequests:
		return codes.ResourceExhausted // TooManyRequests를 ResourceExhausted로 매핑
	}
	return codes.Unknown
}
