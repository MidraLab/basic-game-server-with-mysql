// Go MySQL Driver - A MySQL-Driver for Go's database/sql package.
//
// Copyright 2020 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

//go:build gofuzz
// +build gofuzz

package mysql

import "testing"

func TestFuzz(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fuzz(tt.args.data); got != tt.want {
				t.Errorf("Fuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
