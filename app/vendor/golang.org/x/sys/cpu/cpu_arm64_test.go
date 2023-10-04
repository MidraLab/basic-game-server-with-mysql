// Copyright 2019 The Go Authors. All rights reserved.
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

func Test_setMinimalFeatures(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setMinimalFeatures()
		})
	}
}

func Test_readARM64Registers(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readARM64Registers()
		})
	}
}

func Test_parseARM64SystemRegisters(t *testing.T) {
	type args struct {
		isar0 uint64
		isar1 uint64
		pfr0  uint64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseARM64SystemRegisters(tt.args.isar0, tt.args.isar1, tt.args.pfr0)
		})
	}
}

func Test_extractBits(t *testing.T) {
	type args struct {
		data  uint64
		start uint
		end   uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractBits(tt.args.data, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("extractBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
