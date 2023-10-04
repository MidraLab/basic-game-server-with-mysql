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

func TestNewTag(t *testing.T) {
	type args struct {
		name         string
		description  string
		externalDocs *ExternalDocumentation
	}
	tests := []struct {
		name string
		args args
		want Tag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTag(tt.args.name, tt.args.description, tt.args.externalDocs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTag_JSONLookup(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		TagProps         TagProps
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
			tr := Tag{
				VendorExtensible: tt.fields.VendorExtensible,
				TagProps:         tt.fields.TagProps,
			}
			got, err := tr.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tag.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tag.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTag_MarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		TagProps         TagProps
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
			tr := Tag{
				VendorExtensible: tt.fields.VendorExtensible,
				TagProps:         tt.fields.TagProps,
			}
			got, err := tr.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tag.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tag.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTag_UnmarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		TagProps         TagProps
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
			tr := &Tag{
				VendorExtensible: tt.fields.VendorExtensible,
				TagProps:         tt.fields.TagProps,
			}
			if err := tr.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Tag.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
