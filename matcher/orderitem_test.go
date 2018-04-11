package matcher

import (
	"testing"
	"time"
)

func TestAskItemLess(t *testing.T) {
	item := askItem{"a", time.Unix(0, 10), 2.3, 2}
	for _, c := range []struct {
		than askItem
		want bool
	}{
		{askItem{"b", time.Unix(0, 8), 2.1, 2}, true},
		{askItem{"b", time.Unix(0, 8), 2.9, 2}, false},
		{askItem{"b", time.Unix(0, 8), 2.3, 2}, false},
		{askItem{"b", time.Unix(0, 12), 2.3, 2}, true},
	} {
		if got := item.Less(c.than); got != c.want {
			t.Errorf("Less(%v) = %t, want %t", c.than, got, c.want)
		}
	}
}

func TestBidItemLess(t *testing.T) {
	item := bidItem{"a", time.Unix(0, 10), 2.3, 2}
	for _, c := range []struct {
		than bidItem
		want bool
	}{
		{bidItem{"b", time.Unix(0, 8), 2.1, 2}, false},
		{bidItem{"b", time.Unix(0, 8), 2.9, 2}, true},
		{bidItem{"b", time.Unix(0, 8), 2.3, 2}, true},
		{bidItem{"b", time.Unix(0, 12), 2.3, 2}, false},
	} {
		if got := item.Less(c.than); got != c.want {
			t.Errorf("Less(%v) = %t, want %t", c.than, got, c.want)
		}
	}
}
