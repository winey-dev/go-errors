package errors

import (
	"database/sql"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	"github.com/winey-dev/go-errors/codes"
)

// SQLError wraps a SQL error with a default message.
func SQLError(err error) error {

	if err == nil {
		return nil
	}

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return newError(codes.NotFound, sql.ErrNoRows.Error(), nil, 2)
	case errors.Is(err, sql.ErrTxDone):
		return newError(codes.Disconnected, sql.ErrTxDone.Error(), nil, 2)
	case errors.Is(err, sql.ErrConnDone):
		return newError(codes.Disconnected, sql.ErrConnDone.Error(), nil, 2)
	}
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return newError(codes.FromMYSQLCode(int(mysqlErr.Number)), mysqlErr.Message, nil, 2)
	}

	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		return newError(pgErrorToCode(pgErr.Code), pgErr.Message, nil, 2)
	}

	return newError(codes.Unknown, err.Error(), nil, 2)

}

func pgErrorToCode(code pq.ErrorCode) codes.Code {
	switch code {
	case "23505": // unique_violation
		return codes.InvalidRequest
	default:
		return codes.Unknown
	}
}
