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

func TestQueryParam(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryParam(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeaderParam(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeaderParam(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeaderParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPathParam(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PathParam(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBodyParam(t *testing.T) {
	type args struct {
		name   string
		schema *Schema
	}
	tests := []struct {
		name string
		args args
		want *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BodyParam(tt.args.name, tt.args.schema); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BodyParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormDataParam(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormDataParam(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FormDataParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileParam(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileParam(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimpleArrayParam(t *testing.T) {
	type args struct {
		name string
		tpe  string
		fmt  string
	}
	tests := []struct {
		name string
		args args
		want *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SimpleArrayParam(tt.args.name, tt.args.tpe, tt.args.fmt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SimpleArrayParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParamRef(t *testing.T) {
	type args struct {
		uri string
	}
	tests := []struct {
		name string
		args args
		want *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParamRef(tt.args.uri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParamRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_JSONLookup(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
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
			p := Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			got, err := p.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parameter.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithDescription(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		description string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithDescription(tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_Named(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.Named(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.Named() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithLocation(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		in string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithLocation(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_Typed(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		tpe    string
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.Typed(tt.args.tpe, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.Typed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_CollectionOf(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		items  *Items
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.CollectionOf(tt.args.items, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.CollectionOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithDefault(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		defaultValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithDefault(tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_AllowsEmptyValues(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	tests := []struct {
		name   string
		fields fields
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.AllowsEmptyValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.AllowsEmptyValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_NoEmptyValues(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	tests := []struct {
		name   string
		fields fields
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.NoEmptyValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.NoEmptyValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_AsOptional(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	tests := []struct {
		name   string
		fields fields
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.AsOptional(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.AsOptional() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_AsRequired(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	tests := []struct {
		name   string
		fields fields
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.AsRequired(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.AsRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithMaxLength(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		max int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithMaxLength(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithMinLength(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		min int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithMinLength(tt.args.min); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithMinLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithPattern(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithPattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithMultipleOf(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		number float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithMultipleOf(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithMultipleOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithMaximum(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		max       float64
		exclusive bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithMaximum(tt.args.max, tt.args.exclusive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithMinimum(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		min       float64
		exclusive bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithMinimum(tt.args.min, tt.args.exclusive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithEnum(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithEnum(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithEnum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithMaxItems(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		size int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithMaxItems(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithMaxItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithMinItems(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		size int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithMinItems(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithMinItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_UniqueValues(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	tests := []struct {
		name   string
		fields fields
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.UniqueValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.UniqueValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_AllowDuplicates(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	tests := []struct {
		name   string
		fields fields
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.AllowDuplicates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.AllowDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_WithValidations(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
	}
	type args struct {
		val CommonValidations
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if got := p.WithValidations(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.WithValidations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParameter_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
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
			p := &Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			if err := p.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Parameter.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParameter_MarshalJSON(t *testing.T) {
	type fields struct {
		Refable           Refable
		CommonValidations CommonValidations
		SimpleSchema      SimpleSchema
		VendorExtensible  VendorExtensible
		ParamProps        ParamProps
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
			p := Parameter{
				Refable:           tt.fields.Refable,
				CommonValidations: tt.fields.CommonValidations,
				SimpleSchema:      tt.fields.SimpleSchema,
				VendorExtensible:  tt.fields.VendorExtensible,
				ParamProps:        tt.fields.ParamProps,
			}
			got, err := p.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parameter.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parameter.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
