package fuzzy

import (
	"fmt"
	"strconv"
)

type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func unmarshalUint[T UnsignedInteger](data []byte, target *T, bitSize int) error {
	raw, null, err := scalarText(data)
	if err != nil || null {
		return err
	}

	value, err := strconv.ParseUint(raw, 10, bitSize)
	if err != nil {
		return fmt.Errorf("无效的无符号整数 %q: %w", raw, err)
	}
	*target = T(value)
	return nil
}

type Uint uint

func NewUint(ui uint) *Uint {
	return new(Uint(ui))
}

func (v *Uint) Value() uint {
	if v == nil {
		return 0
	}
	return uint(*v)
}

func (v *Uint) UnmarshalJSON(data []byte) error {
	return unmarshalUint(data, v, 0)
}

type Uint8 uint8

func NewUint8(ui uint8) *Uint8 {
	return new(Uint8(ui))
}

func (v *Uint8) Value() uint8 {
	if v == nil {
		return 0
	}
	return uint8(*v)
}

func (v *Uint8) UnmarshalJSON(data []byte) error {
	return unmarshalUint(data, v, 8)
}

type Uint16 uint16

func NewUint16(ui uint16) *Uint16 {
	return new(Uint16(ui))
}

func (v *Uint16) Value() uint16 {
	if v == nil {
		return 0
	}
	return uint16(*v)
}

func (v *Uint16) UnmarshalJSON(data []byte) error {
	return unmarshalUint(data, v, 16)
}

type Uint32 uint32

func NewUint32(ui uint32) *Uint32 {
	return new(Uint32(ui))
}

func (v *Uint32) Value() uint32 {
	if v == nil {
		return 0
	}
	return uint32(*v)
}

func (v *Uint32) UnmarshalJSON(data []byte) error {
	return unmarshalUint(data, v, 32)
}

type Uint64 uint64

func NewUint64(ui uint64) *Uint64 {
	return new(Uint64(ui))
}

func (v *Uint64) Value() uint64 {
	if v == nil {
		return 0
	}
	return uint64(*v)
}

func (v *Uint64) UnmarshalJSON(data []byte) error {
	return unmarshalUint(data, v, 64)
}
