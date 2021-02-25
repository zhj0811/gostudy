package test

import "testing"

func TestFormat(t *testing.T) {
	var a uint64 = 10000
	t.Logf("%d", int64(a))
	b := int64(a)
	t.Logf("%d", b)
	c := uint64(b)
	t.Logf("%d", c)
}
