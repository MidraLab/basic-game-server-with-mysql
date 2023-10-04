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

func TestPaths_JSONLookup(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		Paths            map[string]PathItem
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
			p := Paths{
				VendorExtensible: tt.fields.VendorExtensible,
				Paths:            tt.fields.Paths,
			}
			got, err := p.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Paths.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Paths.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaths_UnmarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		Paths            map[string]PathItem
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
			p := &Paths{
				VendorExtensible: tt.fields.VendorExtensible,
				Paths:            tt.fields.Paths,
			}
			if err := p.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Paths.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPaths_MarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		Paths            map[string]PathItem
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
			p := Paths{
				VendorExtensible: tt.fields.VendorExtensible,
				Paths:            tt.fields.Paths,
			}
			got, err := p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Paths.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Paths.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
