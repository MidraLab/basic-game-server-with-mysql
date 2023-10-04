// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typeparams

import (
	"go/types"
	"reflect"
	"testing"
)

func TestStructuralTerms(t *testing.T) {
	type args struct {
		tparam *TypeParam
	}
	tests := []struct {
		name    string
		args    args
		want    []*Term
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StructuralTerms(tt.args.tparam)
			if (err != nil) != tt.wantErr {
				t.Errorf("StructuralTerms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StructuralTerms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInterfaceTermSet(t *testing.T) {
	type args struct {
		iface *types.Interface
	}
	tests := []struct {
		name    string
		args    args
		want    []*Term
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InterfaceTermSet(tt.args.iface)
			if (err != nil) != tt.wantErr {
				t.Errorf("InterfaceTermSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InterfaceTermSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnionTermSet(t *testing.T) {
	type args struct {
		union *Union
	}
	tests := []struct {
		name    string
		args    args
		want    []*Term
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnionTermSet(tt.args.union)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnionTermSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnionTermSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeTermSet(t *testing.T) {
	type args struct {
		typ types.Type
	}
	tests := []struct {
		name    string
		args    args
		want    []*Term
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := computeTermSet(tt.args.typ)
			if (err != nil) != tt.wantErr {
				t.Errorf("computeTermSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeTermSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indentf(t *testing.T) {
	type args struct {
		depth  int
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indentf(tt.args.depth, tt.args.format, tt.args.args...)
		})
	}
}

func Test_computeTermSetInternal(t *testing.T) {
	type args struct {
		t     types.Type
		seen  map[types.Type]*termSet
		depth int
	}
	tests := []struct {
		name    string
		args    args
		wantRes *termSet
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, err := computeTermSetInternal(tt.args.t, tt.args.seen, tt.args.depth)
			if (err != nil) != tt.wantErr {
				t.Errorf("computeTermSetInternal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("computeTermSetInternal() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_under(t *testing.T) {
	type args struct {
		t types.Type
	}
	tests := []struct {
		name string
		args args
		want types.Type
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := under(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("under() = %v, want %v", got, tt.want)
			}
		})
	}
}
