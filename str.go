package str

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
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

func JsonCompactify(in string) (string, error) {
	var obj interface{}
	err := json.Unmarshal([]byte(in), &obj)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %w", err)
	}

	bs, err := json.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("failed to marshal json: %w", err)
	}

	return string(bs), nil
}

func JsonPrettify(in string) (string, error) {
	var obj interface{}
	err := json.Unmarshal([]byte(in), &obj)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %w", err)
	}

	bs, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal json: %w", err)
	}

	return string(bs), nil
}
