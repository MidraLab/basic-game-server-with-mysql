// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spec

import (
	"reflect"
	"testing"
)

func TestPathItem_JSONLookup(t *testing.T) {
	type fields struct {
		Refable          Refable
		VendorExtensible VendorExtensible
		PathItemProps    PathItemProps
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PathItem{
				Refable:          tt.fields.Refable,
				VendorExtensible: tt.fields.VendorExtensible,
				PathItemProps:    tt.fields.PathItemProps,
			}
			got, err := p.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathItem.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathItem.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPathItem_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Refable          Refable
		VendorExtensible VendorExtensible
		PathItemProps    PathItemProps
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PathItem{
				Refable:          tt.fields.Refable,
				VendorExtensible: tt.fields.VendorExtensible,
				PathItemProps:    tt.fields.PathItemProps,
			}
			if err := p.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("PathItem.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPathItem_MarshalJSON(t *testing.T) {
	type fields struct {
		Refable          Refable
		VendorExtensible VendorExtensible
		PathItemProps    PathItemProps
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := PathItem{
				Refable:          tt.fields.Refable,
				VendorExtensible: tt.fields.VendorExtensible,
				PathItemProps:    tt.fields.PathItemProps,
			}
			got, err := p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("PathItem.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathItem.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
