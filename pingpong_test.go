package pingpong

import (
	"testing"
)

func TestPingPong(t *testing.T) {
	ping, err := NewPingPong([]string{"a", "b", "c"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"a", "b", "c", "b", "a", "b", "c", "b", "a", "b"}
	for i, exp := range expected {
		if got := ping.Next(); got != exp {
			t.Errorf("Step %d: expected %q, got %q", i, exp, got)
		}
	}
}

func TestNewPingPong_EmptySlice(t *testing.T) {
	_, err := NewPingPong([]string{})
	if err == nil {
		t.Fatal("expected an error for empty slice, got nil")
	}
}
