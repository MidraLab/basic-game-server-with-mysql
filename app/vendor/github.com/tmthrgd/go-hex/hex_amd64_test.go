// Copyright 2016 Tom Thorogood. All rights reserved.
// Use of this source code is governed by a
// Modified BSD License license that can be found in
// the LICENSE file.

//go:build amd64 && !gccgo && !appengine
// +build amd64,!gccgo,!appengine

package hex

import "testing"

func TestRawEncode(t *testing.T) {
	type args struct {
		dst   []byte
		src   []byte
		alpha []byte
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
			if got := RawEncode(tt.args.dst, tt.args.src, tt.args.alpha); got != tt.want {
				t.Errorf("RawEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		dst []byte
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.dst, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encodeAVX(t *testing.T) {
	type args struct {
		dst   *byte
		src   *byte
		len   uint64
		alpha *byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodeAVX(tt.args.dst, tt.args.src, tt.args.len, tt.args.alpha)
		})
	}
}

func Test_encodeSSE(t *testing.T) {
	type args struct {
		dst   *byte
		src   *byte
		len   uint64
		alpha *byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodeSSE(tt.args.dst, tt.args.src, tt.args.len, tt.args.alpha)
		})
	}
}

func Test_decodeAVX(t *testing.T) {
	type args struct {
		dst *byte
		src *byte
		len uint64
	}
	tests := []struct {
		name   string
		args   args
		wantN  uint64
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, gotOk := decodeAVX(tt.args.dst, tt.args.src, tt.args.len)
			if gotN != tt.wantN {
				t.Errorf("decodeAVX() gotN = %v, want %v", gotN, tt.wantN)
			}
			if gotOk != tt.wantOk {
				t.Errorf("decodeAVX() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_decodeSSE(t *testing.T) {
	type args struct {
		dst *byte
		src *byte
		len uint64
	}
	tests := []struct {
		name   string
		args   args
		wantN  uint64
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, gotOk := decodeSSE(tt.args.dst, tt.args.src, tt.args.len)
			if gotN != tt.wantN {
				t.Errorf("decodeSSE() gotN = %v, want %v", gotN, tt.wantN)
			}
			if gotOk != tt.wantOk {
				t.Errorf("decodeSSE() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
