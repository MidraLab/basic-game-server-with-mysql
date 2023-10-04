// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

import (
	"reflect"
	"testing"
)

func Test_littleEndian_Uint32(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		l    littleEndian
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := littleEndian{}
			if got := l.Uint32(tt.args.b); got != tt.want {
				t.Errorf("littleEndian.Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_littleEndian_Uint64(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		l    littleEndian
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := littleEndian{}
			if got := l.Uint64(tt.args.b); got != tt.want {
				t.Errorf("littleEndian.Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_Uint32(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		b    bigEndian
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bigEndian{}
			if got := b.Uint32(tt.args.b); got != tt.want {
				t.Errorf("bigEndian.Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bigEndian_Uint64(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		b    bigEndian
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := bigEndian{}
			if got := b.Uint64(tt.args.b); got != tt.want {
				t.Errorf("bigEndian.Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hostByteOrder(t *testing.T) {
	tests := []struct {
		name string
		want byteOrder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hostByteOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hostByteOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
