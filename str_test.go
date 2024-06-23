package str

import (
	"bytes"
	"testing"
)

func TestSplitIntoLines(t *testing.T) {
	str := "a\nb\nc"
	lines := SplitIntoLines(str)
	if len(lines) != 3 {
		t.Errorf("expected 3 lines, got %d", len(lines))
	}
	if lines[0] != "a" {
		t.Errorf("expected 'a', got %s", lines[0])
	}
	if lines[1] != "b" {
		t.Errorf("expected 'b', got %s", lines[1])
	}
	if lines[2] != "c" {
		t.Errorf("expected 'c', got %s", lines[2])
	}
}

func TestSplitBufIntoLines(t *testing.T) {
	buf := bytes.NewBufferString("a\nb\nc")
	lines := SplitBufIntoLines(buf)
	if len(lines) != 3 {
		t.Errorf("expected 3 lines, got %d", len(lines))
	}
	if lines[0] != "a" {
		t.Errorf("expected 'a', got %s", lines[0])
	}
	if lines[1] != "b" {
		t.Errorf("expected 'b', got %s", lines[1])
	}
	if lines[2] != "c" {
		t.Errorf("expected 'c', got %s", lines[2])
	}
}
