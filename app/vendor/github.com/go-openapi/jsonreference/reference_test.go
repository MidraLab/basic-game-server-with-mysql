// Copyright 2013 sigu-399 ( https://github.com/sigu-399 )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author       sigu-399
// author-github  https://github.com/sigu-399
// author-mail    sigu.399@gmail.com
//
// repository-name  jsonreference
// repository-desc  An implementation of JSON Reference - Go language
//
// description    Main and unique file.
//
// created        26-02-2013

package jsonreference

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/go-openapi/jsonpointer"
)

func TestNew(t *testing.T) {
	type args struct {
		jsonReferenceString string
	}
	tests := []struct {
		name    string
		args    args
		want    Ref
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.jsonReferenceString)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustCreateRef(t *testing.T) {
	type args struct {
		ref string
	}
	tests := []struct {
		name string
		args args
		want Ref
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustCreateRef(tt.args.ref); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustCreateRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_GetURL(t *testing.T) {
	type fields struct {
		referenceURL     *url.URL
		referencePointer jsonpointer.Pointer
		HasFullURL       bool
		HasURLPathOnly   bool
		HasFragmentOnly  bool
		HasFileScheme    bool
		HasFullFilePath  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *url.URL
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ref{
				referenceURL:     tt.fields.referenceURL,
				referencePointer: tt.fields.referencePointer,
				HasFullURL:       tt.fields.HasFullURL,
				HasURLPathOnly:   tt.fields.HasURLPathOnly,
				HasFragmentOnly:  tt.fields.HasFragmentOnly,
				HasFileScheme:    tt.fields.HasFileScheme,
				HasFullFilePath:  tt.fields.HasFullFilePath,
			}
			if got := r.GetURL(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ref.GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_GetPointer(t *testing.T) {
	type fields struct {
		referenceURL     *url.URL
		referencePointer jsonpointer.Pointer
		HasFullURL       bool
		HasURLPathOnly   bool
		HasFragmentOnly  bool
		HasFileScheme    bool
		HasFullFilePath  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *jsonpointer.Pointer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ref{
				referenceURL:     tt.fields.referenceURL,
				referencePointer: tt.fields.referencePointer,
				HasFullURL:       tt.fields.HasFullURL,
				HasURLPathOnly:   tt.fields.HasURLPathOnly,
				HasFragmentOnly:  tt.fields.HasFragmentOnly,
				HasFileScheme:    tt.fields.HasFileScheme,
				HasFullFilePath:  tt.fields.HasFullFilePath,
			}
			if got := r.GetPointer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ref.GetPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_String(t *testing.T) {
	type fields struct {
		referenceURL     *url.URL
		referencePointer jsonpointer.Pointer
		HasFullURL       bool
		HasURLPathOnly   bool
		HasFragmentOnly  bool
		HasFileScheme    bool
		HasFullFilePath  bool
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
			r := &Ref{
				referenceURL:     tt.fields.referenceURL,
				referencePointer: tt.fields.referencePointer,
				HasFullURL:       tt.fields.HasFullURL,
				HasURLPathOnly:   tt.fields.HasURLPathOnly,
				HasFragmentOnly:  tt.fields.HasFragmentOnly,
				HasFileScheme:    tt.fields.HasFileScheme,
				HasFullFilePath:  tt.fields.HasFullFilePath,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("Ref.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_IsRoot(t *testing.T) {
	type fields struct {
		referenceURL     *url.URL
		referencePointer jsonpointer.Pointer
		HasFullURL       bool
		HasURLPathOnly   bool
		HasFragmentOnly  bool
		HasFileScheme    bool
		HasFullFilePath  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ref{
				referenceURL:     tt.fields.referenceURL,
				referencePointer: tt.fields.referencePointer,
				HasFullURL:       tt.fields.HasFullURL,
				HasURLPathOnly:   tt.fields.HasURLPathOnly,
				HasFragmentOnly:  tt.fields.HasFragmentOnly,
				HasFileScheme:    tt.fields.HasFileScheme,
				HasFullFilePath:  tt.fields.HasFullFilePath,
			}
			if got := r.IsRoot(); got != tt.want {
				t.Errorf("Ref.IsRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_IsCanonical(t *testing.T) {
	type fields struct {
		referenceURL     *url.URL
		referencePointer jsonpointer.Pointer
		HasFullURL       bool
		HasURLPathOnly   bool
		HasFragmentOnly  bool
		HasFileScheme    bool
		HasFullFilePath  bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ref{
				referenceURL:     tt.fields.referenceURL,
				referencePointer: tt.fields.referencePointer,
				HasFullURL:       tt.fields.HasFullURL,
				HasURLPathOnly:   tt.fields.HasURLPathOnly,
				HasFragmentOnly:  tt.fields.HasFragmentOnly,
				HasFileScheme:    tt.fields.HasFileScheme,
				HasFullFilePath:  tt.fields.HasFullFilePath,
			}
			if got := r.IsCanonical(); got != tt.want {
				t.Errorf("Ref.IsCanonical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_parse(t *testing.T) {
	type fields struct {
		referenceURL     *url.URL
		referencePointer jsonpointer.Pointer
		HasFullURL       bool
		HasURLPathOnly   bool
		HasFragmentOnly  bool
		HasFileScheme    bool
		HasFullFilePath  bool
	}
	type args struct {
		jsonReferenceString string
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
			r := &Ref{
				referenceURL:     tt.fields.referenceURL,
				referencePointer: tt.fields.referencePointer,
				HasFullURL:       tt.fields.HasFullURL,
				HasURLPathOnly:   tt.fields.HasURLPathOnly,
				HasFragmentOnly:  tt.fields.HasFragmentOnly,
				HasFileScheme:    tt.fields.HasFileScheme,
				HasFullFilePath:  tt.fields.HasFullFilePath,
			}
			if err := r.parse(tt.args.jsonReferenceString); (err != nil) != tt.wantErr {
				t.Errorf("Ref.parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRef_Inherits(t *testing.T) {
	type fields struct {
		referenceURL     *url.URL
		referencePointer jsonpointer.Pointer
		HasFullURL       bool
		HasURLPathOnly   bool
		HasFragmentOnly  bool
		HasFileScheme    bool
		HasFullFilePath  bool
	}
	type args struct {
		child Ref
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Ref
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ref{
				referenceURL:     tt.fields.referenceURL,
				referencePointer: tt.fields.referencePointer,
				HasFullURL:       tt.fields.HasFullURL,
				HasURLPathOnly:   tt.fields.HasURLPathOnly,
				HasFragmentOnly:  tt.fields.HasFragmentOnly,
				HasFileScheme:    tt.fields.HasFileScheme,
				HasFullFilePath:  tt.fields.HasFullFilePath,
			}
			got, err := r.Inherits(tt.args.child)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ref.Inherits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ref.Inherits() = %v, want %v", got, tt.want)
			}
		})
	}
}
