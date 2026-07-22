package fuzzy

import "fmt"

type Bool bool

func (v *Bool) Value() bool {
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
