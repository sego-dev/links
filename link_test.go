package link

import (
	"testing"
)

func TestGetShort(t *testing.T) {
	var short = GetShort("http://example.com")
	assertNotEqual(t, short, "http://example.com", "")
}

func TestGetOriginal(t *testing.T) {
	var short = GetShort("http://example.com")
	var original, e = GetOriginal(short)
	assertEqual(t, e, nil, "error with error")
	assertEqual(t, original, "http://example.com", "")
}
