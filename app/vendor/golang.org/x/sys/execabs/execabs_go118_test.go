// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !go1.19
// +build !go1.19

package execabs

import (
	"os/exec"
	"testing"
)

func Test_isGo119ErrDot(t *testing.T) {
	type args struct {
		err error
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
			if got := isGo119ErrDot(tt.args.err); got != tt.want {
				t.Errorf("isGo119ErrDot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isGo119ErrFieldSet(t *testing.T) {
	type args struct {
		cmd *exec.Cmd
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
			if got := isGo119ErrFieldSet(tt.args.cmd); got != tt.want {
				t.Errorf("isGo119ErrFieldSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
