// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.
//
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hex is an efficient hexadecimal implementation for Golang.
package hex

import (
	"reflect"
	"testing"
)

func TestInvalidByteError_Error(t *testing.T) {
	tests := []struct {
		name string
		e    InvalidByteError
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("InvalidByteError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodedLen(t *testing.T) {
	type args struct {
		n int
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
			if got := EncodedLen(tt.args.n); got != tt.want {
				t.Errorf("EncodedLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodedLen(t *testing.T) {
	type args struct {
		n int
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
			if got := DecodedLen(tt.args.n); got != tt.want {
				t.Errorf("DecodedLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	type args struct {
		dst []byte
		src []byte
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
			if got := Encode(tt.args.dst, tt.args.src); got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeUpper(t *testing.T) {
	type args struct {
		dst []byte
		src []byte
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
			if got := EncodeUpper(tt.args.dst, tt.args.src); got != tt.want {
				t.Errorf("EncodeUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeToString(t *testing.T) {
	type args struct {
		src []byte
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
			if got := EncodeToString(tt.args.src); got != tt.want {
				t.Errorf("EncodeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncodeUpperToString(t *testing.T) {
	type args struct {
		src []byte
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
			if got := EncodeUpperToString(tt.args.src); got != tt.want {
				t.Errorf("EncodeUpperToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawEncodeToString(t *testing.T) {
	type args struct {
		src   []byte
		alpha []byte
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
			if got := RawEncodeToString(tt.args.src, tt.args.alpha); got != tt.want {
				t.Errorf("RawEncodeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustDecodeString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustDecodeString(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustDecodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeGeneric(t *testing.T) {
	type args struct {
		dst   []byte
		src   []byte
		alpha []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodeGeneric(tt.args.dst, tt.args.src, tt.args.alpha)
		})
	}
}

func Test_decodeGeneric(t *testing.T) {
	type args struct {
		dst []byte
		src []byte
	}
	tests := []struct {
		name  string
		args  args
		want  uint64
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := decodeGeneric(tt.args.dst, tt.args.src)
			if got != tt.want {
				t.Errorf("decodeGeneric() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("decodeGeneric() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_fromHexChar(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name  string
		args  args
		want  byte
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := fromHexChar(tt.args.c)
			if got != tt.want {
				t.Errorf("fromHexChar() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("fromHexChar() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
