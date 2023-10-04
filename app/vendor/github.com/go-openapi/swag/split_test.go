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

func Test_split(t *testing.T) {
	type args struct {
		str string
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
			if got := split(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitter_split(t *testing.T) {
	type fields struct {
		postSplitInitialismCheck bool
		initialisms              []string
	}
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []nameLexem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &splitter{
				postSplitInitialismCheck: tt.fields.postSplitInitialismCheck,
				initialisms:              tt.fields.initialisms,
			}
			if got := s.split(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitter.split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newSplitter(t *testing.T) {
	type args struct {
		options []splitterOption
	}
	tests := []struct {
		name string
		args args
		want *splitter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSplitter(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSplitter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_withPostSplitInitialismCheck(t *testing.T) {
	type args struct {
		s *splitter
	}
	tests := []struct {
		name string
		args args
		want *splitter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := withPostSplitInitialismCheck(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("withPostSplitInitialismCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitter_toNameLexems(t *testing.T) {
	type fields struct {
		postSplitInitialismCheck bool
		initialisms              []string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []nameLexem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &splitter{
				postSplitInitialismCheck: tt.fields.postSplitInitialismCheck,
				initialisms:              tt.fields.initialisms,
			}
			if got := s.toNameLexems(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitter.toNameLexems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitter_gatherInitialismMatches(t *testing.T) {
	type fields struct {
		postSplitInitialismCheck bool
		initialisms              []string
	}
	type args struct {
		nameRunes []rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   initialismMatches
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &splitter{
				postSplitInitialismCheck: tt.fields.postSplitInitialismCheck,
				initialisms:              tt.fields.initialisms,
			}
			if got := s.gatherInitialismMatches(tt.args.nameRunes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitter.gatherInitialismMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitter_mapMatchesToNameLexems(t *testing.T) {
	type fields struct {
		postSplitInitialismCheck bool
		initialisms              []string
	}
	type args struct {
		nameRunes []rune
		matches   initialismMatches
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []nameLexem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &splitter{
				postSplitInitialismCheck: tt.fields.postSplitInitialismCheck,
				initialisms:              tt.fields.initialisms,
			}
			if got := s.mapMatchesToNameLexems(tt.args.nameRunes, tt.args.matches); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitter.mapMatchesToNameLexems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitter_initialismRuneEqual(t *testing.T) {
	type fields struct {
		postSplitInitialismCheck bool
		initialisms              []string
	}
	type args struct {
		a rune
		b rune
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
			s := &splitter{
				postSplitInitialismCheck: tt.fields.postSplitInitialismCheck,
				initialisms:              tt.fields.initialisms,
			}
			if got := s.initialismRuneEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("splitter.initialismRuneEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitter_breakInitialism(t *testing.T) {
	type fields struct {
		postSplitInitialismCheck bool
		initialisms              []string
	}
	type args struct {
		original string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   nameLexem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &splitter{
				postSplitInitialismCheck: tt.fields.postSplitInitialismCheck,
				initialisms:              tt.fields.initialisms,
			}
			if got := s.breakInitialism(tt.args.original); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitter.breakInitialism() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitter_breakCasualString(t *testing.T) {
	type fields struct {
		postSplitInitialismCheck bool
		initialisms              []string
	}
	type args struct {
		str []rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []nameLexem
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &splitter{
				postSplitInitialismCheck: tt.fields.postSplitInitialismCheck,
				initialisms:              tt.fields.initialisms,
			}
			if got := s.breakCasualString(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitter.breakCasualString() = %v, want %v", got, tt.want)
			}
		})
	}
}
