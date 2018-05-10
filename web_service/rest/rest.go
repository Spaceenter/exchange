package rest

import (
	"net/http"

	"github.com/catortiger/exchange/store"
)

type Interface interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type WebService struct {
	store *store.Store
}

func New(store *store.Store) *WebService {
	return &WebService{store: store}
}
