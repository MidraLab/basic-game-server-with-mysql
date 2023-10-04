// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

import "testing"

func Test_initOptions(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initOptions()
		})
	}
}

func Test_bitIsSet(t *testing.T) {
	type args struct {
		bits  []uint64
		index uint
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
			if got := bitIsSet(tt.args.bits, tt.args.index); got != tt.want {
				t.Errorf("bitIsSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_facilityList_Has(t *testing.T) {
	type fields struct {
		bits [4]uint64
	}
	type args struct {
		fs []facility
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &facilityList{
				bits: tt.fields.bits,
			}
			if got := s.Has(tt.args.fs...); got != tt.want {
				t.Errorf("facilityList.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResult_Has(t *testing.T) {
	type fields struct {
		bits [2]uint64
	}
	type args struct {
		fns []function
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queryResult{
				bits: tt.fields.bits,
			}
			if got := q.Has(tt.args.fns...); got != tt.want {
				t.Errorf("queryResult.Has() = %v, want %v", got, tt.want)
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
