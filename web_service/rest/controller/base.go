package controller

import (
	"io"

	"github.com/CatOrTiger/exchange/web_service/rest"
)

// RequestInfo request information versioning controller and support mutiple protocals
type RequestInfo struct {
	Protocal   string //pbuf, json etc.
	APIVersion string //used for backward compatibility
}

//Query url query
type Query map[string]string

//ProcessFunc Defind process func
type ProcessFunc func(*rest.WebService, *RequestInfo, Query, io.ReadCloser) (Response, *ResponseError)
