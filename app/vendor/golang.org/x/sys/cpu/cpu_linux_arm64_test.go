// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

import "testing"

func Test_linuxKernelCanEmulateCPUID(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linuxKernelCanEmulateCPUID(); got != tt.want {
				t.Errorf("linuxKernelCanEmulateCPUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doinit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doinit()
		})
	}
}

func Test_isSet(t *testing.T) {
	type args struct {
		hwc   uint
		value uint
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
			if got := isSet(tt.args.hwc, tt.args.value); got != tt.want {
				t.Errorf("isSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
