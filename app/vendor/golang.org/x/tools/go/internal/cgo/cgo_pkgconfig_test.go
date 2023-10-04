// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

import (
	"go/build"
	"reflect"
	"testing"
)

func Test_pkgConfig(t *testing.T) {
	type args struct {
		mode string
		pkgs []string
	}
	tests := []struct {
		name      string
		args      args
		wantFlags []string
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFlags, err := pkgConfig(tt.args.mode, tt.args.pkgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("pkgConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFlags, tt.wantFlags) {
				t.Errorf("pkgConfig() = %v, want %v", gotFlags, tt.wantFlags)
			}
		})
	}
}

func Test_pkgConfigFlags(t *testing.T) {
	type args struct {
		p *build.Package
	}
	tests := []struct {
		name       string
		args       args
		wantCflags []string
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCflags, err := pkgConfigFlags(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("pkgConfigFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCflags, tt.wantCflags) {
				t.Errorf("pkgConfigFlags() = %v, want %v", gotCflags, tt.wantCflags)
			}
		})
	}
}
