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

func TestSimpleSchema_TypeName(t *testing.T) {
	type fields struct {
		Type             string
		Nullable         bool
		Format           string
		Items            *Items
		CollectionFormat string
		Default          interface{}
		Example          interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleSchema{
				Type:             tt.fields.Type,
				Nullable:         tt.fields.Nullable,
				Format:           tt.fields.Format,
				Items:            tt.fields.Items,
				CollectionFormat: tt.fields.CollectionFormat,
				Default:          tt.fields.Default,
				Example:          tt.fields.Example,
			}
			if got := s.TypeName(); got != tt.want {
				t.Errorf("SimpleSchema.TypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleSchema_ItemsTypeName(t *testing.T) {
	type fields struct {
		Type             string
		Nullable         bool
		Format           string
		Items            *Items
		CollectionFormat string
		Default          interface{}
		Example          interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleSchema{
				Type:             tt.fields.Type,
				Nullable:         tt.fields.Nullable,
				Format:           tt.fields.Format,
				Items:            tt.fields.Items,
				CollectionFormat: tt.fields.CollectionFormat,
				Default:          tt.fields.Default,
				Example:          tt.fields.Example,
			}
			if got := s.ItemsTypeName(); got != tt.want {
				t.Errorf("SimpleSchema.ItemsTypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewItems(t *testing.T) {
	tests := []struct {
		name string
		want *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewItems(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_Typed(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		tpe    string
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.Typed(tt.args.tpe, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.Typed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_AsNullable(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	tests := []struct {
		name   string
		fields fields
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.AsNullable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.AsNullable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_CollectionOf(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		items  *Items
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.CollectionOf(tt.args.items, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.CollectionOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_WithDefault(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		defaultValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.WithDefault(tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.WithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_WithMaxLength(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		max int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.WithMaxLength(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.WithMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_WithMinLength(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		min int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.WithMinLength(tt.args.min); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.WithMinLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_WithPattern(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.WithPattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.WithPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_WithMultipleOf(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		number float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.WithMultipleOf(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.WithMultipleOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_WithMaximum(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		max       float64
		exclusive bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.WithMaximum(tt.args.max, tt.args.exclusive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.WithMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_WithMinimum(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		min       float64
		exclusive bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.WithMinimum(tt.args.min, tt.args.exclusive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.WithMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_WithEnum(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.WithEnum(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.WithEnum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_WithMaxItems(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		size int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.WithMaxItems(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.WithMaxItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_WithMinItems(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		size int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.WithMinItems(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.WithMinItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_UniqueValues(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	tests := []struct {
		name   string
		fields fields
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.UniqueValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.UniqueValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_AllowDuplicates(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	tests := []struct {
		name   string
		fields fields
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.AllowDuplicates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.AllowDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_WithValidations(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
	}
	type args struct {
		val CommonValidations
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Items
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if got := i.WithValidations(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.WithValidations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
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
			i := &Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			if err := i.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Items.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestItems_MarshalJSON(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
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
			i := Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			got, err := i.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Items.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItems_JSONLookup(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
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
			i := Items{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
			}
			got, err := i.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Items.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Items.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
