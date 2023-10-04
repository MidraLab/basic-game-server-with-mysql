//
// Copyright (c) 2011-2019 Canonical Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yaml

import (
	"reflect"
	"testing"
	"time"
)

func Test_shortTag(t *testing.T) {
	type args struct {
		tag string
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
			if got := shortTag(tt.args.tag); got != tt.want {
				t.Errorf("shortTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longTag(t *testing.T) {
	type args struct {
		tag string
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
			if got := longTag(tt.args.tag); got != tt.want {
				t.Errorf("longTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resolvableTag(t *testing.T) {
	type args struct {
		tag string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolvableTag(tt.args.tag); got != tt.want {
				t.Errorf("resolvableTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_resolve(t *testing.T) {
	type args struct {
		tag string
		in  string
	}
	tests := []struct {
		name     string
		args     args
		wantRtag string
		wantOut  interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRtag, gotOut := resolve(tt.args.tag, tt.args.in)
			if gotRtag != tt.wantRtag {
				t.Errorf("resolve() gotRtag = %v, want %v", gotRtag, tt.wantRtag)
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("resolve() gotOut = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_encodeBase64(t *testing.T) {
	type args struct {
		s string
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
			if got := encodeBase64(tt.args.s); got != tt.want {
				t.Errorf("encodeBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseTimestamp(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  time.Time
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseTimestamp(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseTimestamp() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseTimestamp() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
