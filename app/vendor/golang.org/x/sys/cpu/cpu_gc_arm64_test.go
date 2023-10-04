// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gc
// +build gc

package cpu

import "testing"

func Test_getisar0(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getisar0(); got != tt.want {
				t.Errorf("getisar0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getisar1(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getisar1(); got != tt.want {
				t.Errorf("getisar1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getpfr0(t *testing.T) {
	tests := []struct {
		name string
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getpfr0(); got != tt.want {
				t.Errorf("getpfr0() = %v, want %v", got, tt.want)
			}
		})
	}
}
