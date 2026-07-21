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

func (v *Uint) UnmarshalJSON(data []byte) error {
	return unmarshalUint(data, v, 0)
}

type Uint8 uint8

func (v *Uint8) UnmarshalJSON(data []byte) error {
	return unmarshalUint(data, v, 8)
}

type Uint16 uint16

func (v *Uint16) UnmarshalJSON(data []byte) error {
	return unmarshalUint(data, v, 16)
}

type Uint32 uint32

func (v *Uint32) UnmarshalJSON(data []byte) error {
	return unmarshalUint(data, v, 32)
}

type Uint64 uint64

func (v *Uint64) UnmarshalJSON(data []byte) error {
	return unmarshalUint(data, v, 64)
}
