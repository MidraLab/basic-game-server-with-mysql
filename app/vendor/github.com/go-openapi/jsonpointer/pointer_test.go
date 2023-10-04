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
// repository-name  jsonpointer
// repository-desc  An implementation of JSON Pointer - Go language
//
// description    Main and unique file.
//
// created        25-02-2013

package jsonpointer

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/go-openapi/swag"
)

func TestNew(t *testing.T) {
	type args struct {
		jsonPointerString string
	}
	tests := []struct {
		name    string
		args    args
		want    Pointer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.jsonPointerString)
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

func TestPointer_parse(t *testing.T) {
	type fields struct {
		referenceTokens []string
	}
	type args struct {
		jsonPointerString string
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
			p := &Pointer{
				referenceTokens: tt.fields.referenceTokens,
			}
			if err := p.parse(tt.args.jsonPointerString); (err != nil) != tt.wantErr {
				t.Errorf("Pointer.parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPointer_Get(t *testing.T) {
	type fields struct {
		referenceTokens []string
	}
	type args struct {
		document any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		want1   reflect.Kind
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pointer{
				referenceTokens: tt.fields.referenceTokens,
			}
			got, got1, err := p.Get(tt.args.document)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pointer.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pointer.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Pointer.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPointer_Set(t *testing.T) {
	type fields struct {
		referenceTokens []string
	}
	type args struct {
		document any
		value    any
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pointer{
				referenceTokens: tt.fields.referenceTokens,
			}
			got, err := p.Set(tt.args.document, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pointer.Set() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pointer.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetForToken(t *testing.T) {
	type args struct {
		document     any
		decodedToken string
	}
	tests := []struct {
		name    string
		args    args
		want    any
		want1   reflect.Kind
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GetForToken(tt.args.document, tt.args.decodedToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetForToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetForToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetForToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSetForToken(t *testing.T) {
	type args struct {
		document     any
		decodedToken string
		value        any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetForToken(tt.args.document, tt.args.decodedToken, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetForToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetForToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSingleImpl(t *testing.T) {
	type args struct {
		node         any
		decodedToken string
		nameProvider *swag.NameProvider
	}
	tests := []struct {
		name    string
		args    args
		want    any
		want1   reflect.Kind
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := getSingleImpl(tt.args.node, tt.args.decodedToken, tt.args.nameProvider)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSingleImpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSingleImpl() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getSingleImpl() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_setSingleImpl(t *testing.T) {
	type args struct {
		node         any
		data         any
		decodedToken string
		nameProvider *swag.NameProvider
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
			if err := setSingleImpl(tt.args.node, tt.args.data, tt.args.decodedToken, tt.args.nameProvider); (err != nil) != tt.wantErr {
				t.Errorf("setSingleImpl() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPointer_get(t *testing.T) {
	type fields struct {
		referenceTokens []string
	}
	type args struct {
		node         any
		nameProvider *swag.NameProvider
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		want1   reflect.Kind
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pointer{
				referenceTokens: tt.fields.referenceTokens,
			}
			got, got1, err := p.get(tt.args.node, tt.args.nameProvider)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pointer.get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pointer.get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Pointer.get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPointer_set(t *testing.T) {
	type fields struct {
		referenceTokens []string
	}
	type args struct {
		node         any
		data         any
		nameProvider *swag.NameProvider
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
			p := &Pointer{
				referenceTokens: tt.fields.referenceTokens,
			}
			if err := p.set(tt.args.node, tt.args.data, tt.args.nameProvider); (err != nil) != tt.wantErr {
				t.Errorf("Pointer.set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPointer_DecodedTokens(t *testing.T) {
	type fields struct {
		referenceTokens []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pointer{
				referenceTokens: tt.fields.referenceTokens,
			}
			if got := p.DecodedTokens(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pointer.DecodedTokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointer_IsEmpty(t *testing.T) {
	type fields struct {
		referenceTokens []string
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
			p := &Pointer{
				referenceTokens: tt.fields.referenceTokens,
			}
			if got := p.IsEmpty(); got != tt.want {
				t.Errorf("Pointer.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointer_String(t *testing.T) {
	type fields struct {
		referenceTokens []string
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
			p := &Pointer{
				referenceTokens: tt.fields.referenceTokens,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Pointer.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPointer_Offset(t *testing.T) {
	type fields struct {
		referenceTokens []string
	}
	type args struct {
		document string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pointer{
				referenceTokens: tt.fields.referenceTokens,
			}
			got, err := p.Offset(tt.args.document)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pointer.Offset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pointer.Offset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_offsetSingleObject(t *testing.T) {
	type args struct {
		dec          *json.Decoder
		decodedToken string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := offsetSingleObject(tt.args.dec, tt.args.decodedToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("offsetSingleObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("offsetSingleObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_offsetSingleArray(t *testing.T) {
	type args struct {
		dec          *json.Decoder
		decodedToken string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := offsetSingleArray(tt.args.dec, tt.args.decodedToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("offsetSingleArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("offsetSingleArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_drainSingle(t *testing.T) {
	type args struct {
		dec *json.Decoder
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
			if err := drainSingle(tt.args.dec); (err != nil) != tt.wantErr {
				t.Errorf("drainSingle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUnescape(t *testing.T) {
	type args struct {
		token string
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
			if got := Unescape(tt.args.token); got != tt.want {
				t.Errorf("Unescape() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEscape(t *testing.T) {
	type args struct {
		token string
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
			if got := Escape(tt.args.token); got != tt.want {
				t.Errorf("Escape() = %v, want %v", got, tt.want)
			}
		})
	}
}
