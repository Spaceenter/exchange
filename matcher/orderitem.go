// Package orderitem contains implementations of btree.Item interface.
package matcher

import (
	"time"

	"github.com/google/btree"
)

type askItem struct {
	orderId   string
	timestamp time.Time
	value     float64
}

func (ai askItem) Less(than btree.Item) bool {
	ti := than.(askItem)
	if ai.value < ti.value {
		return true
	}
	if ai.value > ti.value {
		return false
	}
	return ai.timestamp.Before(ti.timestamp)
}

type bidItem struct {
	orderId   string
	timestamp time.Time
	value     float64
}

func (bi bidItem) Less(than btree.Item) bool {
	ti := than.(bidItem)
	if bi.value < ti.value {
		return true
	}
	if bi.value > ti.value {
		return false
	}
	return bi.timestamp.After(ti.timestamp)
}
