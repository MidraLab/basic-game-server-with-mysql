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

func TestResponse_JSONLookup(t *testing.T) {
	type fields struct {
		Refable          Refable
		ResponseProps    ResponseProps
		VendorExtensible VendorExtensible
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
			r := Response{
				Refable:          tt.fields.Refable,
				ResponseProps:    tt.fields.ResponseProps,
				VendorExtensible: tt.fields.VendorExtensible,
			}
			got, err := r.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Response.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Refable          Refable
		ResponseProps    ResponseProps
		VendorExtensible VendorExtensible
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
			r := &Response{
				Refable:          tt.fields.Refable,
				ResponseProps:    tt.fields.ResponseProps,
				VendorExtensible: tt.fields.VendorExtensible,
			}
			if err := r.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Response.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponse_MarshalJSON(t *testing.T) {
	type fields struct {
		Refable          Refable
		ResponseProps    ResponseProps
		VendorExtensible VendorExtensible
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
			r := Response{
				Refable:          tt.fields.Refable,
				ResponseProps:    tt.fields.ResponseProps,
				VendorExtensible: tt.fields.VendorExtensible,
			}
			got, err := r.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Response.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewResponse(t *testing.T) {
	tests := []struct {
		name string
		want *Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponseRef(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want *Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResponseRef(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResponseRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_WithDescription(t *testing.T) {
	type fields struct {
		Refable          Refable
		ResponseProps    ResponseProps
		VendorExtensible VendorExtensible
	}
	type args struct {
		description string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Refable:          tt.fields.Refable,
				ResponseProps:    tt.fields.ResponseProps,
				VendorExtensible: tt.fields.VendorExtensible,
			}
			if got := r.WithDescription(tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.WithDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_WithSchema(t *testing.T) {
	type fields struct {
		Refable          Refable
		ResponseProps    ResponseProps
		VendorExtensible VendorExtensible
	}
	type args struct {
		schema *Schema
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Refable:          tt.fields.Refable,
				ResponseProps:    tt.fields.ResponseProps,
				VendorExtensible: tt.fields.VendorExtensible,
			}
			if got := r.WithSchema(tt.args.schema); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.WithSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_AddHeader(t *testing.T) {
	type fields struct {
		Refable          Refable
		ResponseProps    ResponseProps
		VendorExtensible VendorExtensible
	}
	type args struct {
		name   string
		header *Header
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Refable:          tt.fields.Refable,
				ResponseProps:    tt.fields.ResponseProps,
				VendorExtensible: tt.fields.VendorExtensible,
			}
			if got := r.AddHeader(tt.args.name, tt.args.header); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.AddHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_RemoveHeader(t *testing.T) {
	type fields struct {
		Refable          Refable
		ResponseProps    ResponseProps
		VendorExtensible VendorExtensible
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Refable:          tt.fields.Refable,
				ResponseProps:    tt.fields.ResponseProps,
				VendorExtensible: tt.fields.VendorExtensible,
			}
			if got := r.RemoveHeader(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.RemoveHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponse_AddExample(t *testing.T) {
	type fields struct {
		Refable          Refable
		ResponseProps    ResponseProps
		VendorExtensible VendorExtensible
	}
	type args struct {
		mediaType string
		example   interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Response{
				Refable:          tt.fields.Refable,
				ResponseProps:    tt.fields.ResponseProps,
				VendorExtensible: tt.fields.VendorExtensible,
			}
			if got := r.AddExample(tt.args.mediaType, tt.args.example); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Response.AddExample() = %v, want %v", got, tt.want)
			}
		})
	}
}
