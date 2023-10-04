// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

import (
	"reflect"
	"testing"
)

func Test_sysctl(t *testing.T) {
	type args struct {
		mib    []int32
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

func Test_sysctlNodes(t *testing.T) {
	type args struct {
		mib []int32
	}
	tests := []struct {
		name    string
		args    args
		want    []sysctlNode
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sysctlNodes(tt.args.mib)
			if (err != nil) != tt.wantErr {
				t.Errorf("sysctlNodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sysctlNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nametomib(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    []int32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nametomib(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("nametomib() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nametomib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sysctlCPUID(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    *aarch64SysctlCPUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sysctlCPUID(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("sysctlCPUID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sysctlCPUID() = %v, want %v", got, tt.want)
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
