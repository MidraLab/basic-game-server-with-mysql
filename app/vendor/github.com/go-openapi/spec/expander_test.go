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

func Test_optionsOrDefault(t *testing.T) {
	type args struct {
		opts *ExpandOptions
	}
	tests := []struct {
		name string
		args args
		want *ExpandOptions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := optionsOrDefault(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("optionsOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpandSpec(t *testing.T) {
	type args struct {
		spec    *Swagger
		options *ExpandOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExpandSpec(tt.args.spec, tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("ExpandSpec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_baseForRoot(t *testing.T) {
	type args struct {
		root  interface{}
		cache ResolutionCache
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := baseForRoot(tt.args.root, tt.args.cache); got != tt.want {
				t.Errorf("baseForRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExpandSchema(t *testing.T) {
	type args struct {
		schema *Schema
		root   interface{}
		cache  ResolutionCache
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExpandSchema(tt.args.schema, tt.args.root, tt.args.cache); (err != nil) != tt.wantErr {
				t.Errorf("ExpandSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExpandSchemaWithBasePath(t *testing.T) {
	type args struct {
		schema *Schema
		cache  ResolutionCache
		opts   *ExpandOptions
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExpandSchemaWithBasePath(tt.args.schema, tt.args.cache, tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("ExpandSchemaWithBasePath() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_expandItems(t *testing.T) {
	type args struct {
		target     Schema
		parentRefs []string
		resolver   *schemaLoader
		basePath   string
	}
	tests := []struct {
		name    string
		args    args
		want    *Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := expandItems(tt.args.target, tt.args.parentRefs, tt.args.resolver, tt.args.basePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("expandItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expandItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expandSchema(t *testing.T) {
	type args struct {
		target     Schema
		parentRefs []string
		resolver   *schemaLoader
		basePath   string
	}
	tests := []struct {
		name    string
		args    args
		want    *Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := expandSchema(tt.args.target, tt.args.parentRefs, tt.args.resolver, tt.args.basePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("expandSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expandSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expandSchemaRef(t *testing.T) {
	type args struct {
		target     Schema
		parentRefs []string
		resolver   *schemaLoader
		basePath   string
	}
	tests := []struct {
		name    string
		args    args
		want    *Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := expandSchemaRef(tt.args.target, tt.args.parentRefs, tt.args.resolver, tt.args.basePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("expandSchemaRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("expandSchemaRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expandPathItem(t *testing.T) {
	type args struct {
		pathItem *PathItem
		resolver *schemaLoader
		basePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := expandPathItem(tt.args.pathItem, tt.args.resolver, tt.args.basePath); (err != nil) != tt.wantErr {
				t.Errorf("expandPathItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_expandOperation(t *testing.T) {
	type args struct {
		op       *Operation
		resolver *schemaLoader
		basePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := expandOperation(tt.args.op, tt.args.resolver, tt.args.basePath); (err != nil) != tt.wantErr {
				t.Errorf("expandOperation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExpandResponseWithRoot(t *testing.T) {
	type args struct {
		response *Response
		root     interface{}
		cache    ResolutionCache
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExpandResponseWithRoot(tt.args.response, tt.args.root, tt.args.cache); (err != nil) != tt.wantErr {
				t.Errorf("ExpandResponseWithRoot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExpandResponse(t *testing.T) {
	type args struct {
		response *Response
		basePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExpandResponse(tt.args.response, tt.args.basePath); (err != nil) != tt.wantErr {
				t.Errorf("ExpandResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExpandParameterWithRoot(t *testing.T) {
	type args struct {
		parameter *Parameter
		root      interface{}
		cache     ResolutionCache
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExpandParameterWithRoot(tt.args.parameter, tt.args.root, tt.args.cache); (err != nil) != tt.wantErr {
				t.Errorf("ExpandParameterWithRoot() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestExpandParameter(t *testing.T) {
	type args struct {
		parameter *Parameter
		basePath  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ExpandParameter(tt.args.parameter, tt.args.basePath); (err != nil) != tt.wantErr {
				t.Errorf("ExpandParameter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getRefAndSchema(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *Ref
		want1   *Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getRefAndSchema(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRefAndSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRefAndSchema() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getRefAndSchema() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_expandParameterOrResponse(t *testing.T) {
	type args struct {
		input    interface{}
		resolver *schemaLoader
		basePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := expandParameterOrResponse(tt.args.input, tt.args.resolver, tt.args.basePath); (err != nil) != tt.wantErr {
				t.Errorf("expandParameterOrResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
