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

func TestXMLObject_WithName(t *testing.T) {
	type fields struct {
		Name      string
		Namespace string
		Prefix    string
		Attribute bool
		Wrapped   bool
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *XMLObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XMLObject{
				Name:      tt.fields.Name,
				Namespace: tt.fields.Namespace,
				Prefix:    tt.fields.Prefix,
				Attribute: tt.fields.Attribute,
				Wrapped:   tt.fields.Wrapped,
			}
			if got := x.WithName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XMLObject.WithName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMLObject_WithNamespace(t *testing.T) {
	type fields struct {
		Name      string
		Namespace string
		Prefix    string
		Attribute bool
		Wrapped   bool
	}
	type args struct {
		namespace string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *XMLObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XMLObject{
				Name:      tt.fields.Name,
				Namespace: tt.fields.Namespace,
				Prefix:    tt.fields.Prefix,
				Attribute: tt.fields.Attribute,
				Wrapped:   tt.fields.Wrapped,
			}
			if got := x.WithNamespace(tt.args.namespace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XMLObject.WithNamespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMLObject_WithPrefix(t *testing.T) {
	type fields struct {
		Name      string
		Namespace string
		Prefix    string
		Attribute bool
		Wrapped   bool
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *XMLObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XMLObject{
				Name:      tt.fields.Name,
				Namespace: tt.fields.Namespace,
				Prefix:    tt.fields.Prefix,
				Attribute: tt.fields.Attribute,
				Wrapped:   tt.fields.Wrapped,
			}
			if got := x.WithPrefix(tt.args.prefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XMLObject.WithPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMLObject_AsAttribute(t *testing.T) {
	type fields struct {
		Name      string
		Namespace string
		Prefix    string
		Attribute bool
		Wrapped   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *XMLObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XMLObject{
				Name:      tt.fields.Name,
				Namespace: tt.fields.Namespace,
				Prefix:    tt.fields.Prefix,
				Attribute: tt.fields.Attribute,
				Wrapped:   tt.fields.Wrapped,
			}
			if got := x.AsAttribute(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XMLObject.AsAttribute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMLObject_AsElement(t *testing.T) {
	type fields struct {
		Name      string
		Namespace string
		Prefix    string
		Attribute bool
		Wrapped   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *XMLObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XMLObject{
				Name:      tt.fields.Name,
				Namespace: tt.fields.Namespace,
				Prefix:    tt.fields.Prefix,
				Attribute: tt.fields.Attribute,
				Wrapped:   tt.fields.Wrapped,
			}
			if got := x.AsElement(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XMLObject.AsElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMLObject_AsWrapped(t *testing.T) {
	type fields struct {
		Name      string
		Namespace string
		Prefix    string
		Attribute bool
		Wrapped   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *XMLObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XMLObject{
				Name:      tt.fields.Name,
				Namespace: tt.fields.Namespace,
				Prefix:    tt.fields.Prefix,
				Attribute: tt.fields.Attribute,
				Wrapped:   tt.fields.Wrapped,
			}
			if got := x.AsWrapped(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XMLObject.AsWrapped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXMLObject_AsUnwrapped(t *testing.T) {
	type fields struct {
		Name      string
		Namespace string
		Prefix    string
		Attribute bool
		Wrapped   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *XMLObject
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XMLObject{
				Name:      tt.fields.Name,
				Namespace: tt.fields.Namespace,
				Prefix:    tt.fields.Prefix,
				Attribute: tt.fields.Attribute,
				Wrapped:   tt.fields.Wrapped,
			}
			if got := x.AsUnwrapped(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XMLObject.AsUnwrapped() = %v, want %v", got, tt.want)
			}
		})
	}
}
