package e

import (
	"errors"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func MakeDatabaseError(err error, code codes.Code, description, message string) error {
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		st, _ := status.New(codes.Internal, "Cannnot query database").WithDetails(
			&errdetails.LocalizedMessage{
				Locale:  "vi",
				Message: "Không thể truy cập thông tin hệ thống!",
			},
		)
		return st.Err()
	}

	st, _ := status.New(code, message).WithDetails(
		&errdetails.LocalizedMessage{
			Locale:  "vi",
			Message: message,
		},
	)
	return st.Err()
}
