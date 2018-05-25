package controller

//Response good response
type Response interface {
	IsJson() bool
}

type JsonResponse struct {
	Data map[string]string
}

func (res *JsonResponse) IsJson() bool {
	return true
}

//ResponseError  Response error
type ResponseError struct {
	Error   error
	Message string
	Code    int
}
