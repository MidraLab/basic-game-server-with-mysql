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

func TestBooleanProperty(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BooleanProperty(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BooleanProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoolProperty(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolProperty(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoolProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringProperty(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringProperty(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCharProperty(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CharProperty(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CharProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64Property(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float64Property(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64Property() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32Property(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float32Property(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32Property() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Property(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int8Property(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8Property() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Property(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int16Property(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16Property() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Property(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Property(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32Property() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Property(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Property(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64Property() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrFmtProperty(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name string
		args args
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrFmtProperty(tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrFmtProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateProperty(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateProperty(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateTimeProperty(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateTimeProperty(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DateTimeProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapProperty(t *testing.T) {
	type args struct {
		property *Schema
	}
	tests := []struct {
		name string
		args args
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapProperty(tt.args.property); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRefProperty(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RefProperty(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRefSchema(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RefSchema(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RefSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayProperty(t *testing.T) {
	type args struct {
		items *Schema
	}
	tests := []struct {
		name string
		args args
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayProperty(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComposedSchema(t *testing.T) {
	type args struct {
		schemas []Schema
	}
	tests := []struct {
		name string
		args args
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComposedSchema(tt.args.schemas...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComposedSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaURL_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		r       SchemaURL
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("SchemaURL.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SchemaURL.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaURL_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		r       *SchemaURL
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SchemaURL.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSchemaURL_fromMap(t *testing.T) {
	type args struct {
		v map[string]interface{}
	}
	tests := []struct {
		name    string
		r       *SchemaURL
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.fromMap(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("SchemaURL.fromMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSchema_JSONLookup(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
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
			s := Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			got, err := s.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Schema.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithID(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithTitle(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		title string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithTitle(tt.args.title); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithDescription(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		description string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithDescription(tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithProperties(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		schemas map[string]Schema
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithProperties(tt.args.schemas); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_SetProperty(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		name   string
		schema Schema
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.SetProperty(tt.args.name, tt.args.schema); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.SetProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithAllOf(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		schemas []Schema
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithAllOf(tt.args.schemas...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithAllOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithMaxProperties(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		max int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithMaxProperties(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithMaxProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithMinProperties(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		min int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithMinProperties(tt.args.min); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithMinProperties() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_Typed(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		tpe    string
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.Typed(tt.args.tpe, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.Typed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_AddType(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		tpe    string
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.AddType(tt.args.tpe, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.AddType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_AsNullable(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.AsNullable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.AsNullable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_CollectionOf(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		items Schema
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.CollectionOf(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.CollectionOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithDefault(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		defaultValue interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithDefault(tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithRequired(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		items []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithRequired(tt.args.items...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_AddRequired(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		items []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.AddRequired(tt.args.items...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.AddRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithMaxLength(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		max int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithMaxLength(tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithMinLength(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		min int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithMinLength(tt.args.min); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithMinLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithPattern(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		pattern string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithPattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithMultipleOf(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		number float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithMultipleOf(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithMultipleOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithMaximum(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		max       float64
		exclusive bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithMaximum(tt.args.max, tt.args.exclusive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithMinimum(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		min       float64
		exclusive bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithMinimum(tt.args.min, tt.args.exclusive); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithEnum(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithEnum(tt.args.values...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithEnum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithMaxItems(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		size int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithMaxItems(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithMaxItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithMinItems(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		size int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithMinItems(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithMinItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_UniqueValues(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.UniqueValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.UniqueValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_AllowDuplicates(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.AllowDuplicates(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.AllowDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_AddToAllOf(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		schemas []Schema
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.AddToAllOf(tt.args.schemas...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.AddToAllOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithDiscriminator(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		discriminator string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithDiscriminator(tt.args.discriminator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithDiscriminator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_AsReadOnly(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.AsReadOnly(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.AsReadOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_AsWritable(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.AsWritable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.AsWritable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithExample(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		example interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithExample(tt.args.example); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithExample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithExternalDocs(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		description string
		url         string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithExternalDocs(tt.args.description, tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithExternalDocs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithXMLName(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithXMLName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithXMLName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithXMLNamespace(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		namespace string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithXMLNamespace(tt.args.namespace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithXMLNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_WithXMLPrefix(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithXMLPrefix(tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithXMLPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_AsXMLAttribute(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.AsXMLAttribute(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.AsXMLAttribute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_AsXMLElement(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.AsXMLElement(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.AsXMLElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_AsWrappedXML(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.AsWrappedXML(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.AsWrappedXML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_AsUnwrappedXML(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.AsUnwrappedXML(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.AsUnwrappedXML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_SetValidations(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		val SchemaValidations
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
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			s.SetValidations(tt.args.val)
		})
	}
}

func TestSchema_WithValidations(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	type args struct {
		val SchemaValidations
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.WithValidations(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.WithValidations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_Validations(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   SchemaValidations
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if got := s.Validations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.Validations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_MarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
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
			s := Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			got, err := s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Schema.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schema.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchema_UnmarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible   VendorExtensible
		SchemaProps        SchemaProps
		SwaggerSchemaProps SwaggerSchemaProps
		ExtraProps         map[string]interface{}
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
			s := &Schema{
				VendorExtensible:   tt.fields.VendorExtensible,
				SchemaProps:        tt.fields.SchemaProps,
				SwaggerSchemaProps: tt.fields.SwaggerSchemaProps,
				ExtraProps:         tt.fields.ExtraProps,
			}
			if err := s.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Schema.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
