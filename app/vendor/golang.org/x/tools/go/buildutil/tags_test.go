// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buildutil

import (
	"reflect"
	"testing"
)

func TestTagsFlag_Set(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		v       *TagsFlag
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.Set(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("TagsFlag.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTagsFlag_Get(t *testing.T) {
	tests := []struct {
		name string
		v    *TagsFlag
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TagsFlag.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitQuotedFields(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitQuotedFields(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitQuotedFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitQuotedFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTagsFlag_String(t *testing.T) {
	tests := []struct {
		name string
		v    *TagsFlag
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.String(); got != tt.want {
				t.Errorf("TagsFlag.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSpaceByte(t *testing.T) {
	type args struct {
		c byte
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
			if got := isSpaceByte(tt.args.c); got != tt.want {
				t.Errorf("isSpaceByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
