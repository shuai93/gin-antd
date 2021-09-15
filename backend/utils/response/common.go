package response

import "time"

const (
	SUCCESS       = 200
	ERROR         = 500
	InvalidParams = 400

	ErrorPermission = 40001

	ErrorExistTag        = 10001
	ErrorNotExistTag     = 10002
	ErrorNotExistArticle = 10003

	ErrorAuthCheckTokenFail    = 20001
	ErrorAuthCheckTokenTimeout = 20002
	ErrorAuthToken             = 20003
	ErrorAuth                  = 20004
)

var Msg = map[int]string{
	SUCCESS:                    "ok",
	ERROR:                      "fail",
	InvalidParams:              "请求参数错误",
	ErrorExistTag:              "已存在该标签名称",
	ErrorNotExistTag:           "该标签不存在",
	ErrorNotExistArticle:       "该文章不存在",
	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",
	ErrorAuth:                  "Token错误",

	ErrorPermission: "权限获取失败",
}

// Response 团队基础序列化器
type Response struct {
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Msg       string      `json:"msg"`
	TimeStamp int64       `json:"timestamp"`
}

func (response Response) SuccessResponse() Response {
	response.TimeStamp = time.Now().Unix()
	response.Code = SUCCESS
	response.Msg = Msg[SUCCESS]
	return response
}

func (response Response) ErrorResponse(message string) Response {
	response.TimeStamp = time.Now().Unix()
	response.Code = ERROR
	if message != "" {
		response.Msg = message
	} else {
		response.Msg = Msg[ERROR]
	}
	return response
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}
