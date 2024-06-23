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

func TestJsonCompactify(t *testing.T) {
	in := `{"a": 1,
"b": 2}`
	out, err := JsonCompactify(in)
	if err != nil {
		t.Fatalf("error compactifying json, got error %v", err)
	}
	if out != `{"a":1,"b":2}` {
		t.Errorf("expected '{\"a\":1,\"b\":2}', got %s", out)
	}
}

func TestJsonPrettify(t *testing.T) {
	in := `{"a":1,"b":2}`
	out, err := JsonPrettify(in)
	if err != nil {
		t.Fatalf("error prettifying json, got error %v", err)
	}
	if out != `{
  "a": 1,
  "b": 2
}` {
		t.Errorf("expected '{\n  \"a\": 1,\n  \"b\": 2\n}', got %s", out)
	}
}
