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

func TestMustLoadJSONSchemaDraft04(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustLoadJSONSchemaDraft04(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustLoadJSONSchemaDraft04() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONSchemaDraft04(t *testing.T) {
	tests := []struct {
		name    string
		want    *Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JSONSchemaDraft04()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONSchemaDraft04() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONSchemaDraft04() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustLoadSwagger20Schema(t *testing.T) {
	tests := []struct {
		name string
		want *Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustLoadSwagger20Schema(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustLoadSwagger20Schema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSwagger20Schema(t *testing.T) {
	tests := []struct {
		name    string
		want    *Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Swagger20Schema()
			if (err != nil) != tt.wantErr {
				t.Errorf("Swagger20Schema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Swagger20Schema() = %v, want %v", got, tt.want)
			}
		})
	}
}
