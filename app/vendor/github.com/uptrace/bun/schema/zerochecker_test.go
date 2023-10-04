package schema

import (
	"reflect"
	"testing"
)

func Test_isZero(t *testing.T) {
	type args struct {
		v interface{}
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
			if got := isZero(tt.args.v); got != tt.want {
				t.Errorf("isZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zeroChecker(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name string
		args args
		want IsZeroerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zeroChecker(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zeroChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addrChecker(t *testing.T) {
	type args struct {
		fn IsZeroerFunc
	}
	tests := []struct {
		name string
		args args
		want IsZeroerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addrChecker(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addrChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isZeroInterface(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := isZeroInterface(tt.args.v); got != tt.want {
				t.Errorf("isZeroInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isZeroDriverValue(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := isZeroDriverValue(tt.args.v); got != tt.want {
				t.Errorf("isZeroDriverValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isZeroLen(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := isZeroLen(tt.args.v); got != tt.want {
				t.Errorf("isZeroLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNil(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := isNil(tt.args.v); got != tt.want {
				t.Errorf("isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isZeroBool(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := isZeroBool(tt.args.v); got != tt.want {
				t.Errorf("isZeroBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isZeroInt(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := isZeroInt(tt.args.v); got != tt.want {
				t.Errorf("isZeroInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isZeroUint(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := isZeroUint(tt.args.v); got != tt.want {
				t.Errorf("isZeroUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isZeroFloat(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := isZeroFloat(tt.args.v); got != tt.want {
				t.Errorf("isZeroFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isZeroBytes(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := isZeroBytes(tt.args.v); got != tt.want {
				t.Errorf("isZeroBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_notZero(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := notZero(tt.args.v); got != tt.want {
				t.Errorf("notZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
