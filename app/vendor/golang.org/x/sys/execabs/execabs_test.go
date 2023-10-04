// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package execabs is a drop-in replacement for os/exec
// that requires PATH lookups to find absolute paths.
// That is, execabs.Command("cmd") runs the same PATH lookup
// as exec.Command("cmd"), but if the result is a path
// which is relative, the Run and Start methods will report
// an error instead of running the executable.
//
// See https://blog.golang.org/path-security for more information
// about when it may be necessary or appropriate to use this package.
package execabs

import (
	"context"
	"os/exec"
	"reflect"
	"testing"
)

func Test_relError(t *testing.T) {
	type args struct {
		file string
		path string
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
			if err := relError(tt.args.file, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("relError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLookPath(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LookPath(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("LookPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LookPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fixCmd(t *testing.T) {
	type args struct {
		name string
		cmd  *exec.Cmd
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixCmd(tt.args.name, tt.args.cmd)
		})
	}
}

func TestCommandContext(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
		arg  []string
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommandContext(tt.args.ctx, tt.args.name, tt.args.arg...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommandContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand(t *testing.T) {
	type args struct {
		name string
		arg  []string
	}
	tests := []struct {
		name string
		args args
		want *exec.Cmd
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Command(tt.args.name, tt.args.arg...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Command() = %v, want %v", got, tt.want)
			}
		})
	}
}
