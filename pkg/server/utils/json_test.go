package utils_test

import (
	"testing"

	"github.com/nodejayes/streaming-ui-server/server/utils"
)

type (
	testTableData[TValue, TExpect any] struct {
		Name              string
		Value             TValue
		Data              []byte
		Expect            TExpect
		CompareTestStruct func(v1, v2 testStruct) bool
	}
	testStruct struct {
		BoolValue       bool       `json:"BoolValue"`
		Int8Value       int8       `json:"Int8Value"`
		Int16Value      int16      `json:"Int16Value"`
		Int32Value      int32      `json:"Int32Value"`
		Int64Value      int64      `json:"Int64Value"`
		IntValue        int        `json:"IntValue"`
		Float32Value    float32    `json:"Float32Value"`
		Float64Value    float64    `json:"Float64Value"`
		Complex64Value  complex64  `json:"Complex64Value"`
		Complex128Value complex128 `json:"Complex128Value"`
		StringValue     string     `json:"StringValue"`
	}
)

var testTable = []testTableData[any, any]{
	{Name: "Test Bool Value", Data: []byte("true"), Value: bool(false), Expect: true},
	{Name: "Test Integer8 Value", Data: []byte("12"), Value: int8(0), Expect: int8(12)},
	{Name: "Test Integer16 Value", Data: []byte("12"), Value: int16(0), Expect: int16(12)},
	{Name: "Test Integer32 Value", Data: []byte("12"), Value: int32(0), Expect: int32(12)},
	{Name: "Test Integer64 Value", Data: []byte("12"), Value: int64(0), Expect: int64(12)},
	{Name: "Test Integer Value", Data: []byte("12"), Value: int(0), Expect: 12},
	{Name: "Test Float32 Value", Data: []byte("12.5"), Value: float32(0), Expect: float32(12.5)},
	{Name: "Test Float64 Value", Data: []byte("12.5"), Value: float64(0), Expect: 12.5},
	{Name: "Test Complex64 Value", Data: []byte(`{"real":3.0,"imagine":4.0}`), Value: complex64(0), Expect: complex64(complex(3.0, 4.0))},
	{Name: "Test Complex128 Value", Data: []byte(`{"real":3.0,"imagine":4.0}`), Value: complex128(0), Expect: complex(3.0, 4.0)},
	{Name: "Test String Value", Data: []byte(`"Hallo"`), Value: string(""), Expect: "Hallo"},
	{Name: "Test Struct Value", Data: []byte(`{
		"BoolValue":true,
		"Int8Value": 12,
		"Int16Value": 12,
		"Int32Value": 12,
		"Int64Value": 12,
		"IntValue": 12,
		"Float32Value": 12.5,
		"Float64Value": 12.5,
		"Complex64Value": 12,
		"Complex128Value": 12,
		"StringValue": "Hallo"
		}`), Value: testStruct{
		BoolValue:       false,
		Int8Value:       0,
		Int16Value:      0,
		Int32Value:      0,
		Int64Value:      0,
		IntValue:        0,
		Float32Value:    0,
		Float64Value:    0,
		Complex64Value:  0,
		Complex128Value: 0,
		StringValue:     "",
	}, Expect: testStruct{
		BoolValue:       true,
		Int8Value:       12,
		Int16Value:      12,
		Int32Value:      12,
		Int64Value:      12,
		IntValue:        12,
		Float32Value:    12.5,
		Float64Value:    12.5,
		Complex64Value:  12,
		Complex128Value: 12,
		StringValue:     "Hallo",
	}, CompareTestStruct: func(v1, v2 testStruct) bool {
		return v1.BoolValue == v2.BoolValue &&
			v1.Int8Value == v2.Int8Value &&
			v1.Int16Value == v2.Int16Value &&
			v1.Int32Value == v2.Int32Value &&
			v1.Int64Value == v2.Int64Value &&
			v1.IntValue == v2.IntValue &&
			v1.Float32Value == v2.Float32Value &&
			v1.Float64Value == v2.Float64Value &&
			v1.Complex64Value == v2.Complex64Value &&
			v1.Complex128Value == v2.Complex128Value &&
			v1.StringValue == v2.StringValue
	}},
}

func TestUnmarshal_Types(t *testing.T) {
	for _, testData := range testTable {
		t.Run(testData.Name, func(t *testing.T) {
			err := utils.Unmarshal(testData.Data, &testData.Value)
			if err != nil {
				t.Error(err)
			}
			switch tv1 := testData.Expect.(type) {
			case testStruct:
				switch tv2 := testData.Value.(type) {
				case testStruct:
					if !testData.CompareTestStruct(tv1, tv2) {
						t.Errorf("expect value %v not equals %v", testData.Expect, testData.Value)
					}
					return
				}
				return
			default:
				if testData.Value != testData.Expect {
					t.Errorf("expect value %v not equals %v", testData.Expect, testData.Value)
				}
				return
			}
		})
	}
}

func TestUnmarshal_Nil(t *testing.T) {
	var target int
	err := utils.Unmarshal([]byte("1"), &target)
	if err != nil {
		t.Error(err)
	}
	if target != 1 {
		t.Errorf("expect target to be 1 but was %v", target)
	}
}
