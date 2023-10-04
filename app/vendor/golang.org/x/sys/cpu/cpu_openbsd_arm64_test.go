// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

import (
	"syscall"
	"testing"
)

func Test_syscall_syscall6(t *testing.T) {
	type args struct {
		fn uintptr
		a1 uintptr
		a2 uintptr
		a3 uintptr
		a4 uintptr
		a5 uintptr
		a6 uintptr
	}
	tests := []struct {
		name    string
		args    args
		wantR1  uintptr
		wantR2  uintptr
		wantErr syscall.Errno
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR1, gotR2, gotErr := syscall_syscall6(tt.args.fn, tt.args.a1, tt.args.a2, tt.args.a3, tt.args.a4, tt.args.a5, tt.args.a6)
			if gotR1 != tt.wantR1 {
				t.Errorf("syscall_syscall6() gotR1 = %v, want %v", gotR1, tt.wantR1)
			}
			if gotR2 != tt.wantR2 {
				t.Errorf("syscall_syscall6() gotR2 = %v, want %v", gotR2, tt.wantR2)
			}
			if gotErr != tt.wantErr {
				t.Errorf("syscall_syscall6() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func Test_sysctl(t *testing.T) {
	type args struct {
		mib    []uint32
		old    *byte
		oldlen *uintptr
		new    *byte
		newlen uintptr
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sysctl(tt.args.mib, tt.args.old, tt.args.oldlen, tt.args.new, tt.args.newlen); (err != nil) != tt.wantErr {
				t.Errorf("sysctl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sysctlUint64(t *testing.T) {
	type args struct {
		mib []uint32
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
			got, got1 := sysctlUint64(tt.args.mib)
			if got != tt.want {
				t.Errorf("sysctlUint64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("sysctlUint64() got1 = %v, want %v", got1, tt.want1)
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
