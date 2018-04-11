// Package orderitem contains implementations of btree.Item interface.
package matcher

import (
	"time"

	"github.com/google/btree"
)

// Ask items are reversely ranked according to the price.
// If two ask items have the same price, the early one is ranked higher.
type askItem struct {
	orderId   string
	timestamp time.Time
	price     float64
	volume    float64
}

func (ai askItem) Less(than btree.Item) bool {
	ti := than.(askItem)
	if ai.price > ti.price {
		return true
	}
	if ai.price < ti.price {
		return false
	}
	return ai.timestamp.Before(ti.timestamp)
}

// Bid items are ranked according to the price.
// If two bid items have the same price, the early one is ranked higher.
type bidItem struct {
	orderId   string
	timestamp time.Time
	price     float64
	volume    float64
}

func (bi bidItem) Less(than btree.Item) bool {
	ti := than.(bidItem)
	if bi.price < ti.price {
		return true
	}
	if bi.price > ti.price {
		return false
	}
	return bi.timestamp.After(ti.timestamp)
}
