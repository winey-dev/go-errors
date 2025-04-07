package tests

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/winey-dev/go-errors"
	"github.com/winey-dev/go-errors/codes"
)

func InsertExec() error {
	return &mysql.MySQLError{Number: 1062, Message: "Duplicate entry 'mss131' for key 'PRIMARY'"}
}

func RepositoryGetUser() (any, error) {
	err := InsertExec()
	return nil, errors.Wrapf(errors.SQLError(err), "insert exec failed. nf=%s", "mss131")
	//return nil, errors.SQLErrorf(sqlError, "Query() failed. nf=%s err=%v", "mss131", sqlError)
	//return nil, errors.SQLError(sqlError)
}

func UsecaseGetUser() (any, error) {
	if _, err := RepositoryGetUser(); err != nil {
		return nil, errors.Wrap(err, "RepositoryGetUser() failed.")
	}
	return nil, nil
}

var _, dErr = UsecaseGetUser()

func TestError(t *testing.T) {
	t.Run("일반 오류 로그 출력_1", func(t *testing.T) {
		log.Printf("%s", dErr)
	})
	t.Run("일반 오류 로그 출력_2", func(t *testing.T) {
		log.Printf("%v", dErr)
	})
	t.Run("파일 라인수와 현재 오류 출력", func(t *testing.T) {
		log.Printf("%#v", dErr)
	})
	t.Run("현재 오류와 오류 체인 정보 출력", func(t *testing.T) {
		log.Printf("%+v", dErr)
	})
	t.Run("파일 라인과 함께 현재 오류와 오류 체인 정보 출력", func(t *testing.T) {
		log.Printf("%+#v", dErr)
	})
}
func TestErrorWrap(t *testing.T) {
	t.Run("일반 오류 로그 출력_1", func(t *testing.T) {
		log.Printf("%s", errors.Wrapf(dErr, "usecase failed. nf=%s", "mss131"))
	})
	t.Run("일반 오류 로그 출력_2", func(t *testing.T) {
		log.Printf("%v", errors.Wrapf(dErr, "usecase failed. nf=%s", "mss131"))
	})
	t.Run("파일 라인수와 현재 오류 출력", func(t *testing.T) {
		log.Printf("%#v", errors.Wrapf(dErr, "usecase failed. nf=%s", "mss131"))
	})
	t.Run("현재 오류와 오류 체인 정보 출력", func(t *testing.T) {
		log.Printf("%+v", errors.Wrapf(dErr, "usecase failed. nf=%s", "mss131"))
	})
	t.Run("파일 라인과 함께 현재 오류와 오류 체인 정보 출력", func(t *testing.T) {
		log.Printf("%+#v", errors.Wrapf(dErr, "usecase failed. nf=%s", "mss131"))
	})

}

func TestHTTPError(t *testing.T) {
	testErr := errors.Wrapc(dErr, codes.NotFound, "usecase failed.")
	form := errors.HTTPError(testErr)
	t.Run("HTTP 오류 출력", func(t *testing.T) {
		dat, _ := json.Marshal(form)
		log.Printf("%s", string(dat))
	})
}

func TestGRPCError(t *testing.T) {
	testErr := errors.Wrapf(dErr, "user usecase get failed. nf=%s", "mss131")
	grpcErr := errors.GRPCError(testErr)
	t.Run("gRPC 오류 출력", func(t *testing.T) {
		log.Printf("%+v", grpcErr)
	})
}
