// Package orderitem contains implementations of btree.Item interface.
package matcher

import (
	"time"

	"github.com/google/btree"
)

// Ask items are reversely ranked according to the price.
// Bid items are ranked according to the price.
// If two items have the same price, the early one is ranked higher.
type orderItem struct {
	orderId   string
	orderTime time.Time
	isSell    bool
	price     float64
	volume    float64
}

func (o orderItem) Less(than btree.Item) bool {
	t := than.(orderItem)
	if o.price != t.price {
		return (o.price > t.price) == o.isSell
	}
	if !o.orderTime.Equal(t.orderTime) {
		return o.orderTime.Before(t.orderTime)
	}
	return o.orderId < t.orderId
}
