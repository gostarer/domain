package base

type ResCode int

const (
	ResCodeOk                ResCode = 1000
	ResCodeValidationError   ResCode = 2000
	ResCodeRequestParamError ResCode = 2100
	ResCodeInnnerServerError ResCode = 5000
	ResCodeBizError          ResCode = 6000
)

type Code struct {
	Val int
	Msg string
}

type Res struct {
	Code    ResCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
