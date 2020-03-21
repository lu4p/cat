package rtftxt

import (
	"testing"
)

func TestPushPeekPopLen(t *testing.T) {
	var s stack
	const h = "hello"
	s.Push(h)
	if s.size != 1 || s.top.value != h || s.top.next != nil {
		t.Error("expected valid value", s.size, s.top)
	}

	if s.Len() != 1 {
		t.Error("expected correct length")
	}

	if v := s.Peek(); v != h || s.size != 1 || s.top.value != h || s.top.next != nil {
		t.Error("expected same value and no size reduction")
	}

	if v := s.Pop(); v != h {
		t.Error("expected pushed value", v)
	}

	if v := s.Pop(); v != "" {
		t.Error("expected empty value")
	}

	if v := s.Peek(); v != "" || s.size != 0 || s.top != nil {
		t.Error("expected nil top and 0 size")
	}
}
