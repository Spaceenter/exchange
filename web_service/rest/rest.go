package rest

import (
	"github.com/CatOrTiger/exchange/store"
)

// WebService the main server could rename to `app`
type WebService struct {
	DB *store.Store //
	// redis *redis       // should support pool
}
