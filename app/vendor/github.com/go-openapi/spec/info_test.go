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

func TestExtensions_Add(t *testing.T) {
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		e    Extensions
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Add(tt.args.key, tt.args.value)
		})
	}
}

func TestExtensions_GetString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		e     Extensions
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.e.GetString(tt.args.key)
			if got != tt.want {
				t.Errorf("Extensions.GetString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Extensions.GetString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestExtensions_GetInt(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		e     Extensions
		args  args
		want  int
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.e.GetInt(tt.args.key)
			if got != tt.want {
				t.Errorf("Extensions.GetInt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Extensions.GetInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestExtensions_GetBool(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		e     Extensions
		args  args
		want  bool
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.e.GetBool(tt.args.key)
			if got != tt.want {
				t.Errorf("Extensions.GetBool() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Extensions.GetBool() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestExtensions_GetStringSlice(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		e     Extensions
		args  args
		want  []string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.e.GetStringSlice(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Extensions.GetStringSlice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Extensions.GetStringSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestVendorExtensible_AddExtension(t *testing.T) {
	type fields struct {
		Extensions Extensions
	}
	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &VendorExtensible{
				Extensions: tt.fields.Extensions,
			}
			v.AddExtension(tt.args.key, tt.args.value)
		})
	}
}

func TestVendorExtensible_MarshalJSON(t *testing.T) {
	type fields struct {
		Extensions Extensions
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
			v := VendorExtensible{
				Extensions: tt.fields.Extensions,
			}
			got, err := v.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("VendorExtensible.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VendorExtensible.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVendorExtensible_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Extensions Extensions
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
			v := &VendorExtensible{
				Extensions: tt.fields.Extensions,
			}
			if err := v.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("VendorExtensible.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInfo_JSONLookup(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		InfoProps        InfoProps
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
			i := Info{
				VendorExtensible: tt.fields.VendorExtensible,
				InfoProps:        tt.fields.InfoProps,
			}
			got, err := i.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Info.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Info.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_MarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		InfoProps        InfoProps
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
			i := Info{
				VendorExtensible: tt.fields.VendorExtensible,
				InfoProps:        tt.fields.InfoProps,
			}
			got, err := i.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Info.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Info.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_UnmarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		InfoProps        InfoProps
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
			i := &Info{
				VendorExtensible: tt.fields.VendorExtensible,
				InfoProps:        tt.fields.InfoProps,
			}
			if err := i.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Info.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
