// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"io"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want UUID
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewString(); got != tt.want {
				t.Errorf("NewString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRandom(t *testing.T) {
	tests := []struct {
		name    string
		want    UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRandom()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRandom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRandomFromReader(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRandomFromReader(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRandomFromReader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRandomFromReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newRandomFromPool(t *testing.T) {
	tests := []struct {
		name    string
		want    UUID
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newRandomFromPool()
			if (err != nil) != tt.wantErr {
				t.Errorf("newRandomFromPool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRandomFromPool() = %v, want %v", got, tt.want)
			}
		})
	}
}
