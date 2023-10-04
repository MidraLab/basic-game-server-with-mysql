// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"hash"
	"reflect"
	"testing"
)

func TestNewHash(t *testing.T) {
	type args struct {
		h       hash.Hash
		space   UUID
		data    []byte
		version int
	}
	tests := []struct {
		name string
		args args
		want UUID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHash(tt.args.h, tt.args.space, tt.args.data, tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMD5(t *testing.T) {
	type args struct {
		space UUID
		data  []byte
	}
	tests := []struct {
		name string
		args args
		want UUID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMD5(tt.args.space, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSHA1(t *testing.T) {
	type args struct {
		space UUID
		data  []byte
	}
	tests := []struct {
		name string
		args args
		want UUID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSHA1(tt.args.space, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSHA1() = %v, want %v", got, tt.want)
			}
		})
	}
}
