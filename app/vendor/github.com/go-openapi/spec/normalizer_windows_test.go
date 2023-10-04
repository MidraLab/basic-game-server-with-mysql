// -build windows

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

func Test_absPath(t *testing.T) {
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
			if got := absPath(tt.args.in); got != tt.want {
				t.Errorf("absPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repairURI(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name  string
		args  args
		want  *url.URL
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := repairURI(tt.args.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repairURI() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("repairURI() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_fixWindowsURI(t *testing.T) {
	type args struct {
		u  *url.URL
		in string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fixWindowsURI(tt.args.u, tt.args.in)
		})
	}
}
