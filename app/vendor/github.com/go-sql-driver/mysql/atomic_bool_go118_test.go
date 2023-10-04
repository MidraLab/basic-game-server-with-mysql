// Go MySQL Driver - A MySQL-Driver for Go's database/sql package.
//
// Copyright 2022 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.
//go:build !go1.19
// +build !go1.19

package mysql

import "testing"

func Test_atomicBool_Load(t *testing.T) {
	tests := []struct {
		name string
		ab   *atomicBool
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ab.Load(); got != tt.want {
				t.Errorf("atomicBool.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_atomicBool_Store(t *testing.T) {
	type args struct {
		value bool
	}
	tests := []struct {
		name string
		ab   *atomicBool
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ab.Store(tt.args.value)
		})
	}
}

func Test_atomicBool_Swap(t *testing.T) {
	type args struct {
		value bool
	}
	tests := []struct {
		name string
		ab   *atomicBool
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ab.Swap(tt.args.value); got != tt.want {
				t.Errorf("atomicBool.Swap() = %v, want %v", got, tt.want)
			}
		})
	}
}
