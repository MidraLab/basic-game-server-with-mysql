// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package buildutil provides utilities related to the go/build
// package in the standard library.
//
// All I/O is done via the build.Context file system interface, which must
// be concurrency-safe.
package buildutil

import (
	"go/build"
	"reflect"
	"testing"
)

func TestAllPackages(t *testing.T) {
	type args struct {
		ctxt *build.Context
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllPackages(tt.args.ctxt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AllPackages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEachPackage(t *testing.T) {
	type args struct {
		ctxt  *build.Context
		found func(importPath string, err error)
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ForEachPackage(tt.args.ctxt, tt.args.found)
		})
	}
}

func Test_allPackages(t *testing.T) {
	type args struct {
		ctxt *build.Context
		root string
		ch   chan<- item
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allPackages(tt.args.ctxt, tt.args.root, tt.args.ch)
		})
	}
}

func TestExpandPatterns(t *testing.T) {
	type args struct {
		ctxt     *build.Context
		patterns []string
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandPatterns(tt.args.ctxt, tt.args.patterns); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExpandPatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}
