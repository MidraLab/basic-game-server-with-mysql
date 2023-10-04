// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Recreate a getsystemcfg syscall handler instead of
// using the one provided by x/sys/unix to avoid having
// the dependency between them. (See golang.org/issue/32102)
// Moreover, this file will be used during the building of
// gccgo's libgo and thus must not used a CGo method.

//go:build aix && gccgo
// +build aix,gccgo

package cpu

import (
	"syscall"
	"testing"
)

func Test_gccgoGetsystemcfg(t *testing.T) {
	type args struct {
		label uint32
	}
	tests := []struct {
		name  string
		args  args
		wantR uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := gccgoGetsystemcfg(tt.args.label); gotR != tt.wantR {
				t.Errorf("gccgoGetsystemcfg() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func Test_callgetsystemcfg(t *testing.T) {
	type args struct {
		label int
	}
	tests := []struct {
		name   string
		args   args
		wantR1 uintptr
		wantE1 syscall.Errno
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR1, gotE1 := callgetsystemcfg(tt.args.label)
			if gotR1 != tt.wantR1 {
				t.Errorf("callgetsystemcfg() gotR1 = %v, want %v", gotR1, tt.wantR1)
			}
			if gotE1 != tt.wantE1 {
				t.Errorf("callgetsystemcfg() gotE1 = %v, want %v", gotE1, tt.wantE1)
			}
		})
	}
}
