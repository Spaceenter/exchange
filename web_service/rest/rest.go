package rest

import "net/http"

type Interface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type WebService struct {
}

func New() *WebService {
	return &WebService{}
}
