// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Minimal copy of x/sys/unix so the cpu package can make a
// system call on AIX without depending on x/sys/unix.
// (See golang.org/issue/32102)

//go:build aix && ppc64 && gc
// +build aix,ppc64,gc

package cpu

import (
	"reflect"
	"testing"
)

func Test_rawSyscall6(t *testing.T) {
	type args struct {
		trap  uintptr
		nargs uintptr
		a1    uintptr
		a2    uintptr
		a3    uintptr
		a4    uintptr
		a5    uintptr
		a6    uintptr
	}
	tests := []struct {
		name    string
		args    args
		wantR1  uintptr
		wantR2  uintptr
		wantErr errno
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR1, gotR2, gotErr := rawSyscall6(tt.args.trap, tt.args.nargs, tt.args.a1, tt.args.a2, tt.args.a3, tt.args.a4, tt.args.a5, tt.args.a6)
			if gotR1 != tt.wantR1 {
				t.Errorf("rawSyscall6() gotR1 = %v, want %v", gotR1, tt.wantR1)
			}
			if gotR2 != tt.wantR2 {
				t.Errorf("rawSyscall6() gotR2 = %v, want %v", gotR2, tt.wantR2)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("rawSyscall6() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func Test_syscall6(t *testing.T) {
	type args struct {
		trap  uintptr
		nargs uintptr
		a1    uintptr
		a2    uintptr
		a3    uintptr
		a4    uintptr
		a5    uintptr
		a6    uintptr
	}
	tests := []struct {
		name    string
		args    args
		wantR1  uintptr
		wantR2  uintptr
		wantErr errno
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR1, gotR2, gotErr := syscall6(tt.args.trap, tt.args.nargs, tt.args.a1, tt.args.a2, tt.args.a3, tt.args.a4, tt.args.a5, tt.args.a6)
			if gotR1 != tt.wantR1 {
				t.Errorf("syscall6() gotR1 = %v, want %v", gotR1, tt.wantR1)
			}
			if gotR2 != tt.wantR2 {
				t.Errorf("syscall6() gotR2 = %v, want %v", gotR2, tt.wantR2)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("syscall6() gotErr = %v, want %v", gotErr, tt.wantErr)
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
		wantE1 errno
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR1, gotE1 := callgetsystemcfg(tt.args.label)
			if gotR1 != tt.wantR1 {
				t.Errorf("callgetsystemcfg() gotR1 = %v, want %v", gotR1, tt.wantR1)
			}
			if !reflect.DeepEqual(gotE1, tt.wantE1) {
				t.Errorf("callgetsystemcfg() gotE1 = %v, want %v", gotE1, tt.wantE1)
			}
		})
	}
}
