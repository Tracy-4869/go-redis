package utils

type ResBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func InitResBody() ResBody {
	return ResBody{
		Code: 200,
		Message: "success",
		Data: nil,
	}
}
