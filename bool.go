package fuzzy

import "fmt"

type Bool bool

func NewBool(b bool) *Bool {
	return new(Bool(b))
}

func (v *Bool) Value() bool {
	if v == nil {
		return false
	}
	return bool(*v)
}

func (v *Bool) UnmarshalJSON(data []byte) error {
	raw, null, err := scalarText(data)
	if err != nil || null {
		return err
	}
	if raw != "true" && raw != "false" {
		return fmt.Errorf("无效的布尔值 %q: 只接受 true 或 false", raw)
	}

	*v = raw == "true"
	return nil
}

func (v *Bool) Scan(src any) error {
	if src == nil {
		*v = false
		return nil
	}

	switch b := src.(type) {
	case bool:
		*v = Bool(b)
	case int64:
		*v = b != 0
	case []byte:
		*v = string(b) == "1" || string(b) == "true"
	default:
		return fmt.Errorf("cannot scan %T into fuzzy.Bool", src)
	}

	return nil
}
