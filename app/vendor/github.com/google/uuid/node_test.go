// Copyright 2016 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"reflect"
	"testing"
)

func TestNodeInterface(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NodeInterface(); got != tt.want {
				t.Errorf("NodeInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetNodeInterface(t *testing.T) {
	type args struct {
		name string
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
			if got := SetNodeInterface(tt.args.name); got != tt.want {
				t.Errorf("SetNodeInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setNodeInterface(t *testing.T) {
	type args struct {
		name string
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
			if got := setNodeInterface(tt.args.name); got != tt.want {
				t.Errorf("setNodeInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNodeID(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NodeID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NodeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetNodeID(t *testing.T) {
	type args struct {
		id []byte
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
			if got := SetNodeID(tt.args.id); got != tt.want {
				t.Errorf("SetNodeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUID_NodeID(t *testing.T) {
	tests := []struct {
		name string
		uuid UUID
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.uuid.NodeID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UUID.NodeID() = %v, want %v", got, tt.want)
			}
		})
	}
}
