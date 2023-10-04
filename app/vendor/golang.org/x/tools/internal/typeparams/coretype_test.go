// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typeparams

import (
	"go/types"
	"reflect"
	"testing"
)

func TestCoreType(t *testing.T) {
	type args struct {
		T types.Type
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
			if got := CoreType(tt.args.T); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CoreType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NormalTerms(t *testing.T) {
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
			got, err := _NormalTerms(tt.args.typ)
			if (err != nil) != tt.wantErr {
				t.Errorf("_NormalTerms() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_NormalTerms() = %v, want %v", got, tt.want)
			}
		})
	}
}
