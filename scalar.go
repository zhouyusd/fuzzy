package fuzzy

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// scalarText returns the textual representation of a JSON scalar. Quoted
// values are decoded as JSON strings so escape sequences are handled correctly.
func scalarText(data []byte) (text string, null bool, err error) {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return "", false, fmt.Errorf("无效的 JSON 值: 输入为空")
	}
	if bytes.Equal(data, []byte("null")) {
		return "", true, nil
	}
	if data[0] != '"' {
		return string(data), false, nil
	}

	if err := json.Unmarshal(data, &text); err != nil {
		return "", false, fmt.Errorf("无效的 JSON 字符串: %w", err)
	}
	return text, false, nil
}

func isJSONNumber(text string) bool {
	if text == "" || string(bytes.TrimSpace([]byte(text))) != text {
		return false
	}

	var number json.Number
	return json.Unmarshal([]byte(text), &number) == nil && number.String() == text
}
