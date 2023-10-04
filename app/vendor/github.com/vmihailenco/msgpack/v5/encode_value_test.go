package msgpack

import (
	"reflect"
	"testing"
)

func Test_getEncoder(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name string
		args args
		want encoderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEncoder(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEncoder(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name string
		args args
		want encoderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := _getEncoder(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_getEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ptrEncoderFunc(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name string
		args args
		want encoderFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ptrEncoderFunc(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ptrEncoderFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeCustomValuePtr(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := encodeCustomValuePtr(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeCustomValuePtr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeCustomValue(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := encodeCustomValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeCustomValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marshalValuePtr(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := marshalValuePtr(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("marshalValuePtr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marshalValue(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := marshalValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("marshalValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeBoolValue(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := encodeBoolValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeBoolValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeInterfaceValue(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := encodeInterfaceValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeInterfaceValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeErrorValue(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := encodeErrorValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeErrorValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_encodeUnsupportedValue(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := encodeUnsupportedValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("encodeUnsupportedValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_nilable(t *testing.T) {
	type args struct {
		kind reflect.Kind
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
			if got := nilable(tt.args.kind); got != tt.want {
				t.Errorf("nilable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_marshalBinaryValueAddr(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := marshalBinaryValueAddr(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("marshalBinaryValueAddr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marshalBinaryValue(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := marshalBinaryValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("marshalBinaryValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marshalTextValueAddr(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := marshalTextValueAddr(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("marshalTextValueAddr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_marshalTextValue(t *testing.T) {
	type args struct {
		e *Encoder
		v reflect.Value
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := marshalTextValue(tt.args.e, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("marshalTextValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
