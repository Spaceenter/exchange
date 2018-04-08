// Package orderitem contains implementations of btree.Item interface.
package orderitem

import (
	"time"
)

type askItem struct {
	orderId   string
	timestamp time.Time
	value     float64
}

func (ai askItem) Less(than askItem) bool {
	if ai.value < than.value {
		return true
	}
	if ai.value > than.value {
		return false
	}
	return ai.timestamp.Before(than.timestamp)
}

type bidItem struct {
	orderId   string
	timestamp time.Time
	value     float64
}

func (bi bidItem) Less(than bidItem) bool {
	if bi.value < than.value {
		return true
	}
	if bi.value > than.value {
		return false
	}
	return bi.timestamp.After(than.timestamp)
}
