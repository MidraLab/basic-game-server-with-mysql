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

func TestResponses_JSONLookup(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		ResponsesProps   ResponsesProps
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
			r := Responses{
				VendorExtensible: tt.fields.VendorExtensible,
				ResponsesProps:   tt.fields.ResponsesProps,
			}
			got, err := r.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Responses.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Responses.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponses_UnmarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		ResponsesProps   ResponsesProps
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
			r := &Responses{
				VendorExtensible: tt.fields.VendorExtensible,
				ResponsesProps:   tt.fields.ResponsesProps,
			}
			if err := r.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Responses.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResponses_MarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		ResponsesProps   ResponsesProps
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
			r := Responses{
				VendorExtensible: tt.fields.VendorExtensible,
				ResponsesProps:   tt.fields.ResponsesProps,
			}
			got, err := r.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Responses.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Responses.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponsesProps_MarshalJSON(t *testing.T) {
	type fields struct {
		Default             *Response
		StatusCodeResponses map[int]Response
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
			r := ResponsesProps{
				Default:             tt.fields.Default,
				StatusCodeResponses: tt.fields.StatusCodeResponses,
			}
			got, err := r.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResponsesProps.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResponsesProps.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResponsesProps_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Default             *Response
		StatusCodeResponses map[int]Response
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
			r := &ResponsesProps{
				Default:             tt.fields.Default,
				StatusCodeResponses: tt.fields.StatusCodeResponses,
			}
			if err := r.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ResponsesProps.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
