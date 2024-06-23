package str

import (
	"bufio"
	"bytes"
)

func SplitIntoLines(str string) []string {
	return SplitBufIntoLines(bytes.NewBufferString(str))
}

func SplitBufIntoLines(buffer *bytes.Buffer) []string {
	var lines []string
	scanner := bufio.NewScanner(buffer)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
