//
// Copyright (c) 2011-2019 Canonical Ltd
// Copyright (c) 2006-2010 Kirill Simonov
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package yaml

import "testing"

func Test_yaml_parser_set_reader_error(t *testing.T) {
	type args struct {
		parser  *yaml_parser_t
		problem string
		offset  int
		value   int
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
			if got := yaml_parser_set_reader_error(tt.args.parser, tt.args.problem, tt.args.offset, tt.args.value); got != tt.want {
				t.Errorf("yaml_parser_set_reader_error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_determine_encoding(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
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
			if got := yaml_parser_determine_encoding(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_determine_encoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_update_raw_buffer(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
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
			if got := yaml_parser_update_raw_buffer(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_update_raw_buffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_update_buffer(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		length int
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
			if got := yaml_parser_update_buffer(tt.args.parser, tt.args.length); got != tt.want {
				t.Errorf("yaml_parser_update_buffer() = %v, want %v", got, tt.want)
			}
		})
	}
}
