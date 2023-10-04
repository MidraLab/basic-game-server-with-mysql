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

package swag

import (
	"reflect"
	"testing"
)

func TestJoinByFormat(t *testing.T) {
	type args struct {
		data   []string
		format string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinByFormat(tt.args.data, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JoinByFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitByFormat(t *testing.T) {
	type args struct {
		data   string
		format string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitByFormat(tt.args.data, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitByFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byInitialism_Len(t *testing.T) {
	tests := []struct {
		name string
		s    byInitialism
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("byInitialism.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byInitialism_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    byInitialism
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_byInitialism_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    byInitialism
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("byInitialism.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trim(t *testing.T) {
	type args struct {
		str string
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
			if got := trim(tt.args.str); got != tt.want {
				t.Errorf("trim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_upper(t *testing.T) {
	type args struct {
		str string
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
			if got := upper(tt.args.str); got != tt.want {
				t.Errorf("upper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lower(t *testing.T) {
	type args struct {
		str string
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
			if got := lower(tt.args.str); got != tt.want {
				t.Errorf("lower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamelize(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name          string
		args          args
		wantCamelized string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCamelized := Camelize(tt.args.word); gotCamelized != tt.wantCamelized {
				t.Errorf("Camelize() = %v, want %v", gotCamelized, tt.wantCamelized)
			}
		})
	}
}

func TestToFileName(t *testing.T) {
	type args struct {
		name string
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
			if got := ToFileName(tt.args.name); got != tt.want {
				t.Errorf("ToFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToCommandName(t *testing.T) {
	type args struct {
		name string
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
			if got := ToCommandName(tt.args.name); got != tt.want {
				t.Errorf("ToCommandName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToHumanNameLower(t *testing.T) {
	type args struct {
		name string
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
			if got := ToHumanNameLower(tt.args.name); got != tt.want {
				t.Errorf("ToHumanNameLower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToHumanNameTitle(t *testing.T) {
	type args struct {
		name string
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
			if got := ToHumanNameTitle(tt.args.name); got != tt.want {
				t.Errorf("ToHumanNameTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToJSONName(t *testing.T) {
	type args struct {
		name string
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
			if got := ToJSONName(tt.args.name); got != tt.want {
				t.Errorf("ToJSONName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToVarName(t *testing.T) {
	type args struct {
		name string
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
			if got := ToVarName(tt.args.name); got != tt.want {
				t.Errorf("ToVarName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToGoName(t *testing.T) {
	type args struct {
		name string
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
			if got := ToGoName(tt.args.name); got != tt.want {
				t.Errorf("ToGoName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsStrings(t *testing.T) {
	type args struct {
		coll []string
		item string
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
			if got := ContainsStrings(tt.args.coll, tt.args.item); got != tt.want {
				t.Errorf("ContainsStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsStringsCI(t *testing.T) {
	type args struct {
		coll []string
		item string
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
			if got := ContainsStringsCI(tt.args.coll, tt.args.item); got != tt.want {
				t.Errorf("ContainsStringsCI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsZero(t *testing.T) {
	type args struct {
		data interface{}
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
			if got := IsZero(tt.args.data); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddInitialisms(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddInitialisms(tt.args.words...)
		})
	}
}
