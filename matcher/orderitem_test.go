package matcher

import (
	"testing"
	"time"
)

func TestOrderItemLess(t *testing.T) {
	askItem := orderItem{"a", time.Unix(0, 10), true, 2.3, 2}
	bidItem := orderItem{"b", time.Unix(0, 10), false, 2.3, 2}
	for _, c := range []struct {
		item orderItem
		than orderItem
		want bool
	}{
		{askItem, orderItem{"c", time.Unix(0, 8), true, 2.1, 2}, true},
		{askItem, orderItem{"c", time.Unix(0, 8), true, 2.9, 2}, false},
		{askItem, orderItem{"c", time.Unix(0, 8), true, 2.3, 2}, false},
		{askItem, orderItem{"c", time.Unix(0, 12), true, 2.3, 2}, true},
		{askItem, orderItem{"c", time.Unix(0, 10), true, 2.3, 2}, true},
		{bidItem, orderItem{"c", time.Unix(0, 8), false, 2.1, 2}, false},
		{bidItem, orderItem{"c", time.Unix(0, 8), false, 2.9, 2}, true},
		{bidItem, orderItem{"c", time.Unix(0, 8), false, 2.3, 2}, false},
		{bidItem, orderItem{"c", time.Unix(0, 12), false, 2.3, 2}, true},
		{bidItem, orderItem{"c", time.Unix(0, 10), false, 2.3, 2}, true},
	} {
		if got := c.item.Less(c.than); got != c.want {
			t.Errorf("Less(%v) = %t, want %t", c.than, got, c.want)
		}
	}
}
