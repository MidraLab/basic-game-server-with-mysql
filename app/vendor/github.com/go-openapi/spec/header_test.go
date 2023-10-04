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

func TestResponseHeader(t *testing.T) {
	tests := []struct {
		name string
		want *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResponseHeader(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResponseHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_WithDescription(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		description string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.WithDescription(tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.WithDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_Typed(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		tpe    string
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.Typed(tt.args.tpe, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.Typed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_CollectionOf(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		items  *Items
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.CollectionOf(tt.args.items, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.CollectionOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_WithDefault(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		defaultValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.WithDefault(tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.WithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_WithMaxLength(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		max int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.WithMaxLength(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.WithMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_WithMinLength(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		min int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.WithMinLength(tt.args.min); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.WithMinLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_WithPattern(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.WithPattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.WithPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_WithMultipleOf(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		number float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.WithMultipleOf(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.WithMultipleOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_WithMaximum(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		max       float64
		exclusive bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.WithMaximum(tt.args.max, tt.args.exclusive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.WithMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_WithMinimum(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		min       float64
		exclusive bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.WithMinimum(tt.args.min, tt.args.exclusive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.WithMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_WithEnum(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.WithEnum(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.WithEnum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_WithMaxItems(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		size int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.WithMaxItems(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.WithMaxItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_WithMinItems(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		size int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.WithMinItems(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.WithMinItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_UniqueValues(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	tests := []struct {
		name   string
		fields fields
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.UniqueValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.UniqueValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_AllowDuplicates(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	tests := []struct {
		name   string
		fields fields
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.AllowDuplicates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.AllowDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_WithValidations(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
	}
	type args struct {
		val CommonValidations
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if got := h.WithValidations(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.WithValidations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_MarshalJSON(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
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
			h := Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			got, err := h.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Header.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeader_UnmarshalJSON(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
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
			h := &Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			if err := h.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Header.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHeader_JSONLookup(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		HeaderProps       HeaderProps
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
			h := Header{
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				HeaderProps:       tt.fields.HeaderProps,
			}
			got, err := h.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Header.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Header.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
