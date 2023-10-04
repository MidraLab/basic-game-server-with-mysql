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

func TestSwagger_JSONLookup(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		SwaggerProps     SwaggerProps
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
			s := Swagger{
				VendorExtensible: tt.fields.VendorExtensible,
				SwaggerProps:     tt.fields.SwaggerProps,
			}
			got, err := s.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Swagger.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Swagger.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwagger_MarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		SwaggerProps     SwaggerProps
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
			s := Swagger{
				VendorExtensible: tt.fields.VendorExtensible,
				SwaggerProps:     tt.fields.SwaggerProps,
			}
			got, err := s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Swagger.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Swagger.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwagger_UnmarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		SwaggerProps     SwaggerProps
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
			s := &Swagger{
				VendorExtensible: tt.fields.VendorExtensible,
				SwaggerProps:     tt.fields.SwaggerProps,
			}
			if err := s.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Swagger.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSwagger_GobEncode(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		SwaggerProps     SwaggerProps
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
			s := Swagger{
				VendorExtensible: tt.fields.VendorExtensible,
				SwaggerProps:     tt.fields.SwaggerProps,
			}
			got, err := s.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Swagger.GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Swagger.GobEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwagger_GobDecode(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		SwaggerProps     SwaggerProps
	}
	type args struct {
		b []byte
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
			s := &Swagger{
				VendorExtensible: tt.fields.VendorExtensible,
				SwaggerProps:     tt.fields.SwaggerProps,
			}
			if err := s.GobDecode(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Swagger.GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSwaggerProps_GobEncode(t *testing.T) {
	type fields struct {
		ID                  string
		Consumes            []string
		Produces            []string
		Schemes             []string
		Swagger             string
		Info                *Info
		Host                string
		BasePath            string
		Paths               *Paths
		Definitions         Definitions
		Parameters          map[string]Parameter
		Responses           map[string]Response
		SecurityDefinitions SecurityDefinitions
		Security            []map[string][]string
		Tags                []Tag
		ExternalDocs        *ExternalDocumentation
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
			o := SwaggerProps{
				ID:                  tt.fields.ID,
				Consumes:            tt.fields.Consumes,
				Produces:            tt.fields.Produces,
				Schemes:             tt.fields.Schemes,
				Swagger:             tt.fields.Swagger,
				Info:                tt.fields.Info,
				Host:                tt.fields.Host,
				BasePath:            tt.fields.BasePath,
				Paths:               tt.fields.Paths,
				Definitions:         tt.fields.Definitions,
				Parameters:          tt.fields.Parameters,
				Responses:           tt.fields.Responses,
				SecurityDefinitions: tt.fields.SecurityDefinitions,
				Security:            tt.fields.Security,
				Tags:                tt.fields.Tags,
				ExternalDocs:        tt.fields.ExternalDocs,
			}
			got, err := o.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("SwaggerProps.GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwaggerProps.GobEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwaggerProps_GobDecode(t *testing.T) {
	type fields struct {
		ID                  string
		Consumes            []string
		Produces            []string
		Schemes             []string
		Swagger             string
		Info                *Info
		Host                string
		BasePath            string
		Paths               *Paths
		Definitions         Definitions
		Parameters          map[string]Parameter
		Responses           map[string]Response
		SecurityDefinitions SecurityDefinitions
		Security            []map[string][]string
		Tags                []Tag
		ExternalDocs        *ExternalDocumentation
	}
	type args struct {
		b []byte
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
			o := &SwaggerProps{
				ID:                  tt.fields.ID,
				Consumes:            tt.fields.Consumes,
				Produces:            tt.fields.Produces,
				Schemes:             tt.fields.Schemes,
				Swagger:             tt.fields.Swagger,
				Info:                tt.fields.Info,
				Host:                tt.fields.Host,
				BasePath:            tt.fields.BasePath,
				Paths:               tt.fields.Paths,
				Definitions:         tt.fields.Definitions,
				Parameters:          tt.fields.Parameters,
				Responses:           tt.fields.Responses,
				SecurityDefinitions: tt.fields.SecurityDefinitions,
				Security:            tt.fields.Security,
				Tags:                tt.fields.Tags,
				ExternalDocs:        tt.fields.ExternalDocs,
			}
			if err := o.GobDecode(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("SwaggerProps.GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSchemaOrBool_JSONLookup(t *testing.T) {
	type fields struct {
		Allows bool
		Schema *Schema
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
			s := SchemaOrBool{
				Allows: tt.fields.Allows,
				Schema: tt.fields.Schema,
			}
			got, err := s.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("SchemaOrBool.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SchemaOrBool.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaOrBool_MarshalJSON(t *testing.T) {
	type fields struct {
		Allows bool
		Schema *Schema
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
			s := SchemaOrBool{
				Allows: tt.fields.Allows,
				Schema: tt.fields.Schema,
			}
			got, err := s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("SchemaOrBool.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SchemaOrBool.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaOrBool_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Allows bool
		Schema *Schema
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
			s := &SchemaOrBool{
				Allows: tt.fields.Allows,
				Schema: tt.fields.Schema,
			}
			if err := s.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SchemaOrBool.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSchemaOrStringArray_JSONLookup(t *testing.T) {
	type fields struct {
		Schema   *Schema
		Property []string
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
			s := SchemaOrStringArray{
				Schema:   tt.fields.Schema,
				Property: tt.fields.Property,
			}
			got, err := s.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("SchemaOrStringArray.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SchemaOrStringArray.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaOrStringArray_MarshalJSON(t *testing.T) {
	type fields struct {
		Schema   *Schema
		Property []string
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
			s := SchemaOrStringArray{
				Schema:   tt.fields.Schema,
				Property: tt.fields.Property,
			}
			got, err := s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("SchemaOrStringArray.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SchemaOrStringArray.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaOrStringArray_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Schema   *Schema
		Property []string
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
			s := &SchemaOrStringArray{
				Schema:   tt.fields.Schema,
				Property: tt.fields.Property,
			}
			if err := s.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SchemaOrStringArray.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringOrArray_Contains(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		s    StringOrArray
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.value); got != tt.want {
				t.Errorf("StringOrArray.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaOrArray_JSONLookup(t *testing.T) {
	type fields struct {
		Schema  *Schema
		Schemas []Schema
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
			s := SchemaOrArray{
				Schema:  tt.fields.Schema,
				Schemas: tt.fields.Schemas,
			}
			got, err := s.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("SchemaOrArray.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SchemaOrArray.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringOrArray_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		s       *StringOrArray
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("StringOrArray.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStringOrArray_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		s       StringOrArray
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("StringOrArray.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringOrArray.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaOrArray_Len(t *testing.T) {
	type fields struct {
		Schema  *Schema
		Schemas []Schema
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SchemaOrArray{
				Schema:  tt.fields.Schema,
				Schemas: tt.fields.Schemas,
			}
			if got := s.Len(); got != tt.want {
				t.Errorf("SchemaOrArray.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaOrArray_ContainsType(t *testing.T) {
	type fields struct {
		Schema  *Schema
		Schemas []Schema
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SchemaOrArray{
				Schema:  tt.fields.Schema,
				Schemas: tt.fields.Schemas,
			}
			if got := s.ContainsType(tt.args.name); got != tt.want {
				t.Errorf("SchemaOrArray.ContainsType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaOrArray_MarshalJSON(t *testing.T) {
	type fields struct {
		Schema  *Schema
		Schemas []Schema
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
			s := SchemaOrArray{
				Schema:  tt.fields.Schema,
				Schemas: tt.fields.Schemas,
			}
			got, err := s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("SchemaOrArray.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SchemaOrArray.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaOrArray_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Schema  *Schema
		Schemas []Schema
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
			s := &SchemaOrArray{
				Schema:  tt.fields.Schema,
				Schemas: tt.fields.Schemas,
			}
			if err := s.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SchemaOrArray.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
