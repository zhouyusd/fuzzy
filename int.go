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

func NewInt(i int) *Int {
	return new(Int(i))
}

func (v *Int) Value() int {
	if v == nil {
		return 0
	}
	return int(*v)
}

func (v *Int) UnmarshalJSON(data []byte) error {
	return unmarshalInt(data, v, 0)
}

type Int8 int8

func NewInt8(i int8) *Int8 {
	return new(Int8(i))
}

func (v *Int8) Value() int8 {
	if v == nil {
		return 0
	}
	return int8(*v)
}

func (v *Int8) UnmarshalJSON(data []byte) error {
	return unmarshalInt(data, v, 8)
}

type Int16 int16

func NewInt16(i int16) *Int16 {
	return new(Int16(i))
}

func (v *Int16) Value() int16 {
	if v == nil {
		return 0
	}
	return int16(*v)
}

func (v *Int16) UnmarshalJSON(data []byte) error {
	return unmarshalInt(data, v, 16)
}

type Int32 int32

func NewInt32(i int32) *Int32 {
	return new(Int32(i))
}

func (v *Int32) Value() int32 {
	if v == nil {
		return 0
	}
	return int32(*v)
}

func (v *Int32) UnmarshalJSON(data []byte) error {
	return unmarshalInt(data, v, 32)
}

type Int64 int64

func NewInt64(i int64) *Int64 {
	return new(Int64(i))
}

func (v *Int64) Value() int64 {
	if v == nil {
		return 0
	}
	return int64(*v)
}

func (v *Int64) UnmarshalJSON(data []byte) error {
	return unmarshalInt(data, v, 64)
}
