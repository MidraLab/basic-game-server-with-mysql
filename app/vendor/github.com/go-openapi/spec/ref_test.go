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

	"github.com/go-openapi/jsonreference"
)

func TestRefable_MarshalJSON(t *testing.T) {
	type fields struct {
		Ref Ref
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
			r := Refable{
				Ref: tt.fields.Ref,
			}
			got, err := r.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Refable.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Refable.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRefable_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Ref Ref
	}
	type args struct {
		d []byte
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
			r := &Refable{
				Ref: tt.fields.Ref,
			}
			if err := r.UnmarshalJSON(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("Refable.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRef_RemoteURI(t *testing.T) {
	type fields struct {
		Ref jsonreference.Ref
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
				Ref: tt.fields.Ref,
			}
			if got := r.RemoteURI(); got != tt.want {
				t.Errorf("Ref.RemoteURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_IsValidURI(t *testing.T) {
	type fields struct {
		Ref jsonreference.Ref
	}
	type args struct {
		basepaths []string
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
			r := &Ref{
				Ref: tt.fields.Ref,
			}
			if got := r.IsValidURI(tt.args.basepaths...); got != tt.want {
				t.Errorf("Ref.IsValidURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_Inherits(t *testing.T) {
	type fields struct {
		Ref jsonreference.Ref
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
				Ref: tt.fields.Ref,
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

func TestNewRef(t *testing.T) {
	type args struct {
		refURI string
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
			got, err := NewRef(tt.args.refURI)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustCreateRef(t *testing.T) {
	type args struct {
		refURI string
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
			if got := MustCreateRef(tt.args.refURI); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MustCreateRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_MarshalJSON(t *testing.T) {
	type fields struct {
		Ref jsonreference.Ref
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
			r := Ref{
				Ref: tt.fields.Ref,
			}
			got, err := r.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Ref.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ref.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Ref jsonreference.Ref
	}
	type args struct {
		d []byte
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
				Ref: tt.fields.Ref,
			}
			if err := r.UnmarshalJSON(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("Ref.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRef_GobEncode(t *testing.T) {
	type fields struct {
		Ref jsonreference.Ref
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
			r := Ref{
				Ref: tt.fields.Ref,
			}
			got, err := r.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Ref.GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ref.GobEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRef_GobDecode(t *testing.T) {
	type fields struct {
		Ref jsonreference.Ref
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
			r := &Ref{
				Ref: tt.fields.Ref,
			}
			if err := r.GobDecode(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Ref.GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRef_fromMap(t *testing.T) {
	type fields struct {
		Ref jsonreference.Ref
	}
	type args struct {
		v map[string]interface{}
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
				Ref: tt.fields.Ref,
			}
			if err := r.fromMap(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Ref.fromMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
