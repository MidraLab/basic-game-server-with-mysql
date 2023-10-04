// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import "testing"

func Test_randomBits(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			randomBits(tt.args.b)
		})
	}
}

func Test_xtob(t *testing.T) {
	type args struct {
		x1 byte
		x2 byte
	}
	tests := []struct {
		name  string
		args  args
		want  byte
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := xtob(tt.args.x1, tt.args.x2)
			if got != tt.want {
				t.Errorf("xtob() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("xtob() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
