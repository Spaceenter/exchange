package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/CatOrTiger/exchange/web_service/rest"
)

//CreateUser ...
func CreateUser(server *rest.WebService, info *RequestInfo, query Query, body io.ReadCloser) (Response, *ResponseError) {
	// // Decode data
	// decoder := json.NewDecoder(body)

	// var newUser struct {
	// 	UserName       string
	// 	HashedPassword string
	// 	Email          string
	// 	PhoneNumber    string
	// }

	// err := decoder.Decode(&newUser)
	// log.Println(err)
	// if nil != err {
	// 	return nil, &ResponseError{err, "decode failed", http.StatusBadRequest}
	// }

	// responseData := map[string]map[string]interface{}{
	// 	"properties": map[string]interface{}{},
	// 	"results":    map[string]interface{}{},
	// }

	test := map[string]string{
		"test":  "abc",
		"test2": "bcd",
	}
	// return responseData, nil
	res := &JsonResponse{Data: test}
	return res, nil
}

//GetUsers ...
func GetUsers(server *rest.WebService, info *RequestInfo, query Query, body io.ReadCloser) (Response, *ResponseError) {

	// Decode data
	decoder := json.NewDecoder(body)

	var newUser struct {
		UserName       string
		HashedPassword string
		Email          string
		PhoneNumber    string
	}

	err := decoder.Decode(&newUser)
	log.Println(err)
	if nil != err {
		return nil, &ResponseError{err, "decode failed", http.StatusBadRequest}
	}

	// responseData := map[string]map[string]interface{}{
	// 	"properties": map[string]interface{}{},
	// 	"results":    map[string]interface{}{},
	// }

	test := map[string]string{
		"test":  "abc",
		"test2": "bcd",
	}
	// return responseData, nil
	res := &JsonResponse{Data: test}
	return res, nil
}

//GetUser ...
func GetUser(server *rest.WebService, info *RequestInfo, query Query, body io.ReadCloser) (Response, *ResponseError) {

	if len(query) == 0 {
		return nil, &ResponseError{nil, "bad request", http.StatusBadRequest}
	}

	var key string
	var ok bool

	if key, ok = query["key"]; !ok {
		return nil, &ResponseError{nil, "bad request", http.StatusBadRequest}
	}

	test := map[string]string{
		"test":  "abc",
		"test2": "bcd",
		"key":   key,
	}
	// return responseData, nil
	res := &JsonResponse{Data: test}
	return res, nil
}
