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
	"net/url"
	"reflect"
	"testing"
)

func Test_normalizeURI(t *testing.T) {
	type args struct {
		refPath string
		base    string
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
			if got := normalizeURI(tt.args.refPath, tt.args.base); got != tt.want {
				t.Errorf("normalizeURI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_denormalizeRef(t *testing.T) {
	type args struct {
		ref                  *Ref
		originalRelativeBase string
		id                   string
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
			if got := denormalizeRef(tt.args.ref, tt.args.originalRelativeBase, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("denormalizeRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rebase(t *testing.T) {
	type args struct {
		ref      *Ref
		v        *url.URL
		notEqual bool
	}
	tests := []struct {
		name  string
		args  args
		want  Ref
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := rebase(tt.args.ref, tt.args.v, tt.args.notEqual)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rebase() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("rebase() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_normalizeRef(t *testing.T) {
	type args struct {
		ref          *Ref
		relativeBase string
	}
	tests := []struct {
		name string
		args args
		want *Ref
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeRef(tt.args.ref, tt.args.relativeBase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalizeRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalizeBase(t *testing.T) {
	type args struct {
		in string
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
			if got := normalizeBase(tt.args.in); got != tt.want {
				t.Errorf("normalizeBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
