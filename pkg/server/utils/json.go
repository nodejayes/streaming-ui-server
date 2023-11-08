package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type complexData struct {
	Real    float64 `json:"real"`
	Imagine float64 `json:"imagine"`
}

func unmarshalInt8(data []byte) (int8, error) {
	var v int8
	err := json.Unmarshal(data, &v)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func unmarshalInt16(data []byte) (int16, error) {
	var v int16
	err := json.Unmarshal(data, &v)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func unmarshalInt32(data []byte) (int32, error) {
	var v int32
	err := json.Unmarshal(data, &v)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func unmarshalInt64(data []byte) (int64, error) {
	var v int64
	err := json.Unmarshal(data, &v)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func unmarshalInt(data []byte) (int, error) {
	var v int
	err := json.Unmarshal(data, &v)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func unmarshalComplex64(data []byte) (complex64, error) {
	var v complexData
	err := json.Unmarshal(data, &v)
	if err != nil {
		return 0, err
	}
	return complex64(complex(v.Real, v.Imagine)), nil
}

func unmarshalComplex128(data []byte) (complex128, error) {
	var v complexData
	err := json.Unmarshal(data, &v)
	if err != nil {
		return 0, err
	}
	return complex(v.Real, v.Imagine), nil
}

func unmarshalFloat32(data []byte) (float32, error) {
	var v float32
	err := json.Unmarshal(data, &v)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func Unmarshal(data []byte, target any) error {
	targetV := reflect.ValueOf(target)
	if targetV.Kind() != reflect.Pointer || targetV.IsNil() {
		return fmt.Errorf("invalid input type must be a pointer that is not nil")
	}

	tv := targetV.Elem()
	switch v := tv.Interface().(type) {
	case int8:
		v, err := unmarshalInt8(data)
		tv.Set(reflect.ValueOf(v))
		return err
	case int16:
		v, err := unmarshalInt16(data)
		tv.Set(reflect.ValueOf(v))
		return err
	case int32:
		v, err := unmarshalInt32(data)
		tv.Set(reflect.ValueOf(v))
		return err
	case int64:
		v, err := unmarshalInt64(data)
		tv.Set(reflect.ValueOf(v))
		return err
	case int:
		v, err := unmarshalInt(data)
		tv.Set(reflect.ValueOf(v))
		return err
	case complex64:
		v, err := unmarshalComplex64(data)
		tv.Set(reflect.ValueOf(v))
		return err
	case complex128:
		v, err := unmarshalComplex128(data)
		tv.Set(reflect.ValueOf(v))
		return err
	case float32:
		v, err := unmarshalFloat32(data)
		tv.Set(reflect.ValueOf(v))
		return err
	}

	return json.Unmarshal(data, &target)
}
