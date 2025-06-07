package resultx

import (
	"context"
	"github.com/pkg/errors"
	"github.com/sword-demon/go-zero-im/pkg/xerr"
	"github.com/zeromicro/go-zero/core/logx"
	zerr "github.com/zeromicro/x/errors"
	"google.golang.org/grpc/status"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data`
}

func Success(data interface{}) *Response {
	return &Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}

func Fail(code int, err string) *Response {
	return &Response{
		Code: code,
		Msg:  err,
		Data: nil,
	}
}

func OKHandler(_ context.Context, v interface{}) any {
	return Success(v)
}

func ErrHandler(name string) func(ctx context.Context, err error) (int, any) {
	return func(ctx context.Context, err error) (int, any) {
		errCode := xerr.ServerCommonError
		errMsg := xerr.ErrMsg(errCode)

		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*zerr.CodeMsg); ok {
			errCode = e.Code
			errMsg = e.Msg
		} else {
			if gStatus, ok := status.FromError(causeErr); ok {
				errCode = int(gStatus.Code())
				errMsg = gStatus.Message()
			}
		}

		// 日志记录
		logx.WithContext(ctx).Errorf("[%s] err %v", name, err)

		return http.StatusBadRequest, Fail(errCode, errMsg)
	}
}
