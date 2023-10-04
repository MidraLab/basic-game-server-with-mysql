// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swag

import "testing"

func TestIsFloat64AJSONInteger(t *testing.T) {
	type args struct {
		f float64
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
			if got := IsFloat64AJSONInteger(tt.args.f); got != tt.want {
				t.Errorf("IsFloat64AJSONInteger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertBool(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertBool(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertFloat32(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertFloat32(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertFloat32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertFloat64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertFloat64(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertInt8(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int8
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertInt8(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertInt8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertInt16(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int16
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertInt16(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertInt16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertInt32(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertInt32(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertInt32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertInt64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertInt64(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertUint8(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    uint8
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertUint8(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertUint8() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertUint16(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    uint16
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertUint16(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertUint16() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertUint32(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertUint32(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertUint32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertUint64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertUint64(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertUint64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatBool(t *testing.T) {
	type args struct {
		value bool
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
			if got := FormatBool(tt.args.value); got != tt.want {
				t.Errorf("FormatBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatFloat32(t *testing.T) {
	type args struct {
		value float32
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
			if got := FormatFloat32(tt.args.value); got != tt.want {
				t.Errorf("FormatFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatFloat64(t *testing.T) {
	type args struct {
		value float64
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
			if got := FormatFloat64(tt.args.value); got != tt.want {
				t.Errorf("FormatFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatInt8(t *testing.T) {
	type args struct {
		value int8
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
			if got := FormatInt8(tt.args.value); got != tt.want {
				t.Errorf("FormatInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatInt16(t *testing.T) {
	type args struct {
		value int16
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
			if got := FormatInt16(tt.args.value); got != tt.want {
				t.Errorf("FormatInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatInt32(t *testing.T) {
	type args struct {
		value int32
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
			if got := FormatInt32(tt.args.value); got != tt.want {
				t.Errorf("FormatInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatInt64(t *testing.T) {
	type args struct {
		value int64
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
			if got := FormatInt64(tt.args.value); got != tt.want {
				t.Errorf("FormatInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatUint8(t *testing.T) {
	type args struct {
		value uint8
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
			if got := FormatUint8(tt.args.value); got != tt.want {
				t.Errorf("FormatUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatUint16(t *testing.T) {
	type args struct {
		value uint16
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
			if got := FormatUint16(tt.args.value); got != tt.want {
				t.Errorf("FormatUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatUint32(t *testing.T) {
	type args struct {
		value uint32
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
			if got := FormatUint32(tt.args.value); got != tt.want {
				t.Errorf("FormatUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatUint64(t *testing.T) {
	type args struct {
		value uint64
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
			if got := FormatUint64(tt.args.value); got != tt.want {
				t.Errorf("FormatUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
