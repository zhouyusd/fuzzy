package fuzzy

import (
	"fmt"
	"strconv"
)

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func unmarshalInt[T Integer](data []byte, target *T, bitSize int) error {
	raw, null, err := scalarText(data)
	if err != nil || null {
		return err
	}

	value, err := strconv.ParseInt(raw, 10, bitSize)
	if err != nil {
		return fmt.Errorf("无效的整数 %q: %w", raw, err)
	}
	*target = T(value)
	return nil
}

type Int int

func (v *Int) UnmarshalJSON(data []byte) error {
	return unmarshalInt(data, v, 0)
}

type Int8 int8

func (v *Int8) UnmarshalJSON(data []byte) error {
	return unmarshalInt(data, v, 8)
}

type Int16 int16

func (v *Int16) UnmarshalJSON(data []byte) error {
	return unmarshalInt(data, v, 16)
}

type Int32 int32

func (v *Int32) UnmarshalJSON(data []byte) error {
	return unmarshalInt(data, v, 32)
}

type Int64 int64

func (v *Int64) UnmarshalJSON(data []byte) error {
	return unmarshalInt(data, v, 64)
}
