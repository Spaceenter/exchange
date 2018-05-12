package controller

import (
	"io"
	"reflect"
	"testing"

	"github.com/CatOrTiger/exchange/web_service/rest"
)

func TestGetUser(t *testing.T) {
	type args struct {
		server *rest.WebService
		info   *RequestInfo
		query  Query
		body   io.ReadCloser
	}
	tests := []struct {
		name  string
		args  args
		want  Response
		want1 *ResponseError
	}{
		// TODO: Add test cases.

		//create a mock server,
		//fake reqquest info
		//...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetUser(tt.args.server, tt.args.info, tt.args.query, tt.args.body)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	type args struct {
		server *rest.WebService
		info   *RequestInfo
		query  Query
		body   io.ReadCloser
	}
	tests := []struct {
		name  string
		args  args
		want  Response
		want1 *ResponseError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetUsers(tt.args.server, tt.args.info, tt.args.query, tt.args.body)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetUsers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		server *rest.WebService
		info   *RequestInfo
		query  Query
		body   io.ReadCloser
	}
	tests := []struct {
		name  string
		args  args
		want  Response
		want1 *ResponseError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CreateUser(tt.args.server, tt.args.info, tt.args.query, tt.args.body)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CreateUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
