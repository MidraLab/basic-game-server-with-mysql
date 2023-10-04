// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (386 || amd64 || amd64p32) && gccgo
// +build 386 amd64 amd64p32
// +build gccgo

package cpu

import "testing"

func Test_gccgoGetCpuidCount(t *testing.T) {
	type args struct {
		eaxArg uint32
		ecxArg uint32
		eax    *uint32
		ebx    *uint32
		ecx    *uint32
		edx    *uint32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gccgoGetCpuidCount(tt.args.eaxArg, tt.args.ecxArg, tt.args.eax, tt.args.ebx, tt.args.ecx, tt.args.edx)
		})
	}
}

func Test_cpuid(t *testing.T) {
	type args struct {
		eaxArg uint32
		ecxArg uint32
	}
	tests := []struct {
		name    string
		args    args
		wantEax uint32
		wantEbx uint32
		wantEcx uint32
		wantEdx uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEax, gotEbx, gotEcx, gotEdx := cpuid(tt.args.eaxArg, tt.args.ecxArg)
			if gotEax != tt.wantEax {
				t.Errorf("cpuid() gotEax = %v, want %v", gotEax, tt.wantEax)
			}
			if gotEbx != tt.wantEbx {
				t.Errorf("cpuid() gotEbx = %v, want %v", gotEbx, tt.wantEbx)
			}
			if gotEcx != tt.wantEcx {
				t.Errorf("cpuid() gotEcx = %v, want %v", gotEcx, tt.wantEcx)
			}
			if gotEdx != tt.wantEdx {
				t.Errorf("cpuid() gotEdx = %v, want %v", gotEdx, tt.wantEdx)
			}
		})
	}
}

func Test_gccgoXgetbv(t *testing.T) {
	type args struct {
		eax *uint32
		edx *uint32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gccgoXgetbv(tt.args.eax, tt.args.edx)
		})
	}
}

func Test_xgetbv(t *testing.T) {
	tests := []struct {
		name    string
		wantEax uint32
		wantEdx uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEax, gotEdx := xgetbv()
			if gotEax != tt.wantEax {
				t.Errorf("xgetbv() gotEax = %v, want %v", gotEax, tt.wantEax)
			}
			if gotEdx != tt.wantEdx {
				t.Errorf("xgetbv() gotEdx = %v, want %v", gotEdx, tt.wantEdx)
			}
		})
	}
}

func Test_darwinSupportsAVX512(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := darwinSupportsAVX512(); got != tt.want {
				t.Errorf("darwinSupportsAVX512() = %v, want %v", got, tt.want)
			}
		})
	}
}
