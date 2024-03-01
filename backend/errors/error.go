package errors

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	pb_common "pchat/pb/common"
)

type PError struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func (e *PError) Error() string {
	data, err := jsoniter.Marshal(e)
	if err != nil {
		return e.Message
	}
	return string(data)
}

func New(code int64, message string) *PError {
	return &PError{
		code,
		message,
	}
}

func ToPError(err error) *PError {
	var pe *PError
	if errors.As(err, &pe) {
		return pe
	}
	pe = &PError{}
	err = jsoniter.UnmarshalFromString(err.Error(), &pe)
	if err != nil {
		pe.Code = ERR_COMMON_UNKNOWN
		pe.Message = err.Error()
	}
	return pe
}

func (err *PError) ToCommonError() *pb_common.ErrorResponse {
	return &pb_common.ErrorResponse{
		Message: err.Message,
		Code:    err.Code,
	}
}

func ToCommonError(err error) *pb_common.ErrorResponse {
	return ToPError(err).ToCommonError()
}
