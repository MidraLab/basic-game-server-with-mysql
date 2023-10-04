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
)

func Test_keyList_Len(t *testing.T) {
	tests := []struct {
		name string
		l    keyList
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Len(); got != tt.want {
				t.Errorf("keyList.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keyList_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		l    keyList
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_keyList_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		l    keyList
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("keyList.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keyFloat(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name   string
		args   args
		wantF  float64
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF, gotOk := keyFloat(tt.args.v)
			if gotF != tt.wantF {
				t.Errorf("keyFloat() gotF = %v, want %v", gotF, tt.wantF)
			}
			if gotOk != tt.wantOk {
				t.Errorf("keyFloat() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_numLess(t *testing.T) {
	type args struct {
		a reflect.Value
		b reflect.Value
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
			if got := numLess(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("numLess() = %v, want %v", got, tt.want)
			}
		})
	}
}
