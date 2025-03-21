package codes

// Code 타입은 오류 코드를 나타내는 정수형 타입입니다.
type Code int

const (
	// 성공을 나타내는 코드
	OK Code = 0
	// 코드를 출력하고 싶지 않을때 사용하는 코드
	NotUse Code = 1
	// 알 수 없는 오류 코드
	Unknown Code = 2
	// 잘못된 인자를 나타내는 코드
	InvalidArgument Code = 3
	// 잘못된 요청을 나타내는 코드
	InvalidRequest Code = 4
	// 리소스를 찾을 수 없음을 나타내는 코드
	NotFound Code = 5
	// 내부 서버 오류를 나타내는 코드
	Internal Code = 6
	// 인증되지 않은 접근을 나타내는 코드
	Unauthorized Code = 7
	// 접근이 금지됨을 나타내는 코드
	Forbidden Code = 8
	// 요청 시간이 초과되었음을 나타내는 코드
	Timeout Code = 9
	// 교착 상태를 나타내는 코드
	Deadlock Code = 10
	// 연결이 거부되었음을 나타내는 코드
	ConnectionRefused Code = 11
	// 연결이 끊어졌음을 나타내는 코드
	Disconnected Code = 12
	// 리소스가 고갈되었음을 나타내는 코드
	ResourceExhausted Code = 13
	// 너무 많은 요청을 나타내는 코드
	TooManyRequests Code = 14
)

// codeToString 맵은 각 오류 코드를 문자열로 매핑합니다.
var codeToString = map[Code]string{
	OK:                "OK",
	NotUse:            "NotUse",
	Unknown:           "Unknown",
	InvalidArgument:   "InvalidArgument",
	InvalidRequest:    "InvalidRequest",
	NotFound:          "NotFound",
	Internal:          "Internal",
	Unauthorized:      "Unauthorized",
	Forbidden:         "Forbidden",
	Timeout:           "Timeout",
	Deadlock:          "Deadlock",
	ConnectionRefused: "ConnectionRefused",
	Disconnected:      "Disconnected",
	ResourceExhausted: "ResourceExhausted",
	TooManyRequests:   "TooManyRequests",
}

func (c Code) String() string {
	if c, ok := codeToString[c]; ok {
		return c
	}
	return "Unknown"
}
