package swag

import (
	"reflect"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.v); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringValue(t *testing.T) {
	type args struct {
		v *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringValue(tt.args.v); got != tt.want {
				t.Errorf("StringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSlice(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []*string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringValueSlice(t *testing.T) {
	type args struct {
		src []*string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringValueSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringMap(t *testing.T) {
	type args struct {
		src map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]*string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringValueMap(t *testing.T) {
	type args struct {
		src map[string]*string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringValueMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBool(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want *bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args.v); got != tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolValue(t *testing.T) {
	type args struct {
		v *bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolValue(tt.args.v); got != tt.want {
				t.Errorf("BoolValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolSlice(t *testing.T) {
	type args struct {
		src []bool
	}
	tests := []struct {
		name string
		args args
		want []*bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolValueSlice(t *testing.T) {
	type args struct {
		src []*bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolValueSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolMap(t *testing.T) {
	type args struct {
		src map[string]bool
	}
	tests := []struct {
		name string
		args args
		want map[string]*bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolValueMap(t *testing.T) {
	type args struct {
		src map[string]*bool
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolValueMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.v); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntValue(t *testing.T) {
	type args struct {
		v *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntValue(tt.args.v); got != tt.want {
				t.Errorf("IntValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntSlice(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []*int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntValueSlice(t *testing.T) {
	type args struct {
		src []*int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntValueSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntMap(t *testing.T) {
	type args struct {
		src map[string]int
	}
	tests := []struct {
		name string
		args args
		want map[string]*int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntValueMap(t *testing.T) {
	type args struct {
		src map[string]*int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntValueMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	type args struct {
		v int32
	}
	tests := []struct {
		name string
		args args
		want *int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32(tt.args.v); got != tt.want {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Value(t *testing.T) {
	type args struct {
		v *int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Value(tt.args.v); got != tt.want {
				t.Errorf("Int32Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Slice(t *testing.T) {
	type args struct {
		src []int32
	}
	tests := []struct {
		name string
		args args
		want []*int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Slice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32ValueSlice(t *testing.T) {
	type args struct {
		src []*int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32ValueSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32ValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Map(t *testing.T) {
	type args struct {
		src map[string]int32
	}
	tests := []struct {
		name string
		args args
		want map[string]*int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Map(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32ValueMap(t *testing.T) {
	type args struct {
		src map[string]*int32
	}
	tests := []struct {
		name string
		args args
		want map[string]int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32ValueMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32ValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want *int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64(tt.args.v); got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Value(t *testing.T) {
	type args struct {
		v *int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Value(tt.args.v); got != tt.want {
				t.Errorf("Int64Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Slice(t *testing.T) {
	type args struct {
		src []int64
	}
	tests := []struct {
		name string
		args args
		want []*int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Slice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ValueSlice(t *testing.T) {
	type args struct {
		src []*int64
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64ValueSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64ValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Map(t *testing.T) {
	type args struct {
		src map[string]int64
	}
	tests := []struct {
		name string
		args args
		want map[string]*int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Map(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64ValueMap(t *testing.T) {
	type args struct {
		src map[string]*int64
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64ValueMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64ValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want *uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16(tt.args.v); got != tt.want {
				t.Errorf("Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Value(t *testing.T) {
	type args struct {
		v *uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16Value(tt.args.v); got != tt.want {
				t.Errorf("Uint16Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Slice(t *testing.T) {
	type args struct {
		src []uint16
	}
	tests := []struct {
		name string
		args args
		want []*uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16Slice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16ValueSlice(t *testing.T) {
	type args struct {
		src []*uint16
	}
	tests := []struct {
		name string
		args args
		want []uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16ValueSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16ValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16Map(t *testing.T) {
	type args struct {
		src map[string]uint16
	}
	tests := []struct {
		name string
		args args
		want map[string]*uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16Map(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16ValueMap(t *testing.T) {
	type args struct {
		src map[string]*uint16
	}
	tests := []struct {
		name string
		args args
		want map[string]uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint16ValueMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16ValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	type args struct {
		v uint
	}
	tests := []struct {
		name string
		args args
		want *uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint(tt.args.v); got != tt.want {
				t.Errorf("Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintValue(t *testing.T) {
	type args struct {
		v *uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintValue(tt.args.v); got != tt.want {
				t.Errorf("UintValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintSlice(t *testing.T) {
	type args struct {
		src []uint
	}
	tests := []struct {
		name string
		args args
		want []*uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintValueSlice(t *testing.T) {
	type args struct {
		src []*uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintValueSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintMap(t *testing.T) {
	type args struct {
		src map[string]uint
	}
	tests := []struct {
		name string
		args args
		want map[string]*uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintValueMap(t *testing.T) {
	type args struct {
		src map[string]*uint
	}
	tests := []struct {
		name string
		args args
		want map[string]uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UintValueMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UintValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name string
		args args
		want *uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32(tt.args.v); got != tt.want {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Value(t *testing.T) {
	type args struct {
		v *uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32Value(tt.args.v); got != tt.want {
				t.Errorf("Uint32Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Slice(t *testing.T) {
	type args struct {
		src []uint32
	}
	tests := []struct {
		name string
		args args
		want []*uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32Slice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32ValueSlice(t *testing.T) {
	type args struct {
		src []*uint32
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32ValueSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32ValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Map(t *testing.T) {
	type args struct {
		src map[string]uint32
	}
	tests := []struct {
		name string
		args args
		want map[string]*uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32Map(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32ValueMap(t *testing.T) {
	type args struct {
		src map[string]*uint32
	}
	tests := []struct {
		name string
		args args
		want map[string]uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint32ValueMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32ValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want *uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64(tt.args.v); got != tt.want {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Value(t *testing.T) {
	type args struct {
		v *uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64Value(tt.args.v); got != tt.want {
				t.Errorf("Uint64Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Slice(t *testing.T) {
	type args struct {
		src []uint64
	}
	tests := []struct {
		name string
		args args
		want []*uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64Slice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64ValueSlice(t *testing.T) {
	type args struct {
		src []*uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64ValueSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64ValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64Map(t *testing.T) {
	type args struct {
		src map[string]uint64
	}
	tests := []struct {
		name string
		args args
		want map[string]*uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64Map(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64ValueMap(t *testing.T) {
	type args struct {
		src map[string]*uint64
	}
	tests := []struct {
		name string
		args args
		want map[string]uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uint64ValueMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64ValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	type args struct {
		v float32
	}
	tests := []struct {
		name string
		args args
		want *float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32(tt.args.v); got != tt.want {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Value(t *testing.T) {
	type args struct {
		v *float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32Value(tt.args.v); got != tt.want {
				t.Errorf("Float32Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Slice(t *testing.T) {
	type args struct {
		src []float32
	}
	tests := []struct {
		name string
		args args
		want []*float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32Slice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32ValueSlice(t *testing.T) {
	type args struct {
		src []*float32
	}
	tests := []struct {
		name string
		args args
		want []float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32ValueSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32ValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Map(t *testing.T) {
	type args struct {
		src map[string]float32
	}
	tests := []struct {
		name string
		args args
		want map[string]*float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32Map(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32ValueMap(t *testing.T) {
	type args struct {
		src map[string]*float32
	}
	tests := []struct {
		name string
		args args
		want map[string]float32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32ValueMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32ValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want *float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64(tt.args.v); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Value(t *testing.T) {
	type args struct {
		v *float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Value(tt.args.v); got != tt.want {
				t.Errorf("Float64Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Slice(t *testing.T) {
	type args struct {
		src []float64
	}
	tests := []struct {
		name string
		args args
		want []*float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Slice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64ValueSlice(t *testing.T) {
	type args struct {
		src []*float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64ValueSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64ValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Map(t *testing.T) {
	type args struct {
		src map[string]float64
	}
	tests := []struct {
		name string
		args args
		want map[string]*float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Map(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64ValueMap(t *testing.T) {
	type args struct {
		src map[string]*float64
	}
	tests := []struct {
		name string
		args args
		want map[string]float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64ValueMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64ValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTime(t *testing.T) {
	type args struct {
		v time.Time
	}
	tests := []struct {
		name string
		args args
		want *time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Time(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeValue(t *testing.T) {
	type args struct {
		v *time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeValue(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeSlice(t *testing.T) {
	type args struct {
		src []time.Time
	}
	tests := []struct {
		name string
		args args
		want []*time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeValueSlice(t *testing.T) {
	type args struct {
		src []*time.Time
	}
	tests := []struct {
		name string
		args args
		want []time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeValueSlice(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeValueSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeMap(t *testing.T) {
	type args struct {
		src map[string]time.Time
	}
	tests := []struct {
		name string
		args args
		want map[string]*time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeValueMap(t *testing.T) {
	type args struct {
		src map[string]*time.Time
	}
	tests := []struct {
		name string
		args args
		want map[string]time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeValueMap(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeValueMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
