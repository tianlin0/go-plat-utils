package httputil

import (
	"github.com/tianlin0/go-plat-utils/conv"
	"net/http"
)

// CommResponse 接口返回值
type CommResponse struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Now     string      `json:"now,omitempty"`
	Env     string      `json:"env,omitempty"` //环境
	Time    int64       `json:"time,omitempty"`
	LogId   string      `json:"logId,omitempty"`
	TraceId string      `json:"traceId,omitempty"`
	Params  interface{} `json:"params,omitempty"`
	Debug   interface{} `json:"debug,omitempty"`
	Data    interface{} `json:"data"`
}

// PageModel 分页结构输出
type PageModel struct {
	Count     int64       `json:"count"`
	PageNow   int         `json:"pageNow,omitempty"`
	PageStart int         `json:"pageStart,omitempty"`
	PageSize  int         `json:"pageSize,omitempty"`
	PageTotal int         `json:"pageTotal,omitempty"`
	DataList  interface{} `json:"dataList"`
}

// WithNowTime 获取通用的返回格式
func (comm *CommResponse) withNowTime() *CommResponse {
	comm.Now = conv.FormatFromUnixTime() //当前时间
	return comm
}

// WriteCommResponse 将通用返回设置到response，输出到客户端
func WriteCommResponse(respWriter http.ResponseWriter, comm *CommResponse, statusCode ...int) error {
	response := comm.withNowTime()

	contentType := "Content-Type"
	respWriter.Header().Set(contentType, "application/json; charset=utf-8")

	respStr := conv.String(response)
	respByte := []byte(respStr)

	oneStatusCode := http.StatusOK
	if len(statusCode) > 0 {
		oneStatusCode = statusCode[0]
	}
	respWriter.WriteHeader(oneStatusCode)

	_, err := respWriter.Write(respByte)

	return err
}

// GetErrorResponse 系统获取错误码和错误信息
func GetErrorResponse(allErrorMap map[int64]string, errorCode int64, err ...error) *CommResponse {
	respError := &CommResponse{}

	respError.Code = errorCode

	if len(err) > 0 {
		respError.Message = err[0].Error()
	}

	if allErrorMap != nil {
		if errorMsg, ok := allErrorMap[errorCode]; ok {
			if respError.Message == "" {
				respError.Message = conv.String(errorMsg)
			}
			return respError
		}
	}

	if respError.Message == "" {
		respError.Message = "系统错误"
	}

	return respError
}
