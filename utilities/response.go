package utilities

type ResponseJson struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessResponse(returnData *ResponseJson) {
	returnData.Msg = "Success"
	returnData.Data = nil
}

func NewSuccessResponse(returnData *ResponseJson, data interface{}) {
	returnData.Msg = "Success"
	returnData.Data = data
}

func ErrorResponse(returnData *ResponseJson, msg string) {
	returnData.Msg = msg
	returnData.Data = nil
}
