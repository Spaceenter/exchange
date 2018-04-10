// Package orderitem contains implementations of btree.Item interface.
package matcher

import (
	"time"

	"github.com/google/btree"
)

// Ask items are reversely ranked according to the values.
// If two ask items have the same value, the early one is ranked higher.
type askItem struct {
	orderId   string
	timestamp time.Time
	value     float64
}

func (ai askItem) Less(than btree.Item) bool {
	ti := than.(askItem)
	if ai.value > ti.value {
		return true
	}
	if ai.value < ti.value {
		return false
	}
	return ai.timestamp.Before(ti.timestamp)
}

// Bid items are ranked according to the values.
// If two bid items have the same value, the early one is ranked higher.
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
