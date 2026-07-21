package fuzzy

import (
	"fmt"
	"math"
	"strconv"
)

type Float interface {
	~float32 | ~float64
}

func unmarshalFloat[T Float](data []byte, target *T, bitSize int) error {
	raw, null, err := scalarText(data)
	if err != nil || null {
		return err
	}
	if !isJSONNumber(raw) {
		return fmt.Errorf("无效的浮点数 %q: 不是有效的 JSON 数字", raw)
	}

	value, err := strconv.ParseFloat(raw, bitSize)
	if err != nil {
		return fmt.Errorf("无效的浮点数 %q: %w", raw, err)
	}
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return fmt.Errorf("无效的浮点数 %q: JSON 不支持非有限数", raw)
	}
	*target = T(value)
	return nil
}

type Float32 float32

func (v *Float32) UnmarshalJSON(data []byte) error {
	return unmarshalFloat(data, v, 32)
}

type Float64 float64

func (v *Float64) UnmarshalJSON(data []byte) error {
	return unmarshalFloat(data, v, 64)
}
