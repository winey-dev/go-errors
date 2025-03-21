package errors

import (
	"errors"

	"github.com/winey-dev/go-errors/codes"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func GRPCError(err error) error {
	var Error *Error
	var grpcStatus *status.Status
	if errors.As(err, &Error) {
		grpcStatus = status.New(codes.ToGRPCCode(Error.code), Error.message)
		details := Error.Details()
		for _, detail := range details {
			grpcStatus, _ = grpcStatus.WithDetails(&errdetails.LocalizedMessage{Message: detail})
		}
		return grpcStatus.Err()
	} else {
		grpcStatus = status.New(codes.ToGRPCCode(codes.Internal), err.Error())
	}
	return grpcStatus.Err()
}
