// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix
// +build aix

package cpu

import "testing"

func Test_archInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			archInit()
		})
	}
}

func Test_getsystemcfg(t *testing.T) {
	type args struct {
		label int
	}
	tests := []struct {
		name  string
		args  args
		wantN uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := getsystemcfg(tt.args.label); gotN != tt.wantN {
				t.Errorf("getsystemcfg() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
