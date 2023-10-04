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

import (
	"reflect"
	"testing"
)

func Test_cache(t *testing.T) {
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
			if got := cache(tt.args.parser, tt.args.length); got != tt.want {
				t.Errorf("cache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_skip(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skip(tt.args.parser)
		})
	}
}

func Test_skip_line(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			skip_line(tt.args.parser)
		})
	}
}

func Test_read(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		s      []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := read(tt.args.parser, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_read_line(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		s      []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := read_line(tt.args.parser, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("read_line() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		token  *yaml_token_t
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
			if got := yaml_parser_scan(tt.args.parser, tt.args.token); got != tt.want {
				t.Errorf("yaml_parser_scan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_set_scanner_error(t *testing.T) {
	type args struct {
		parser       *yaml_parser_t
		context      string
		context_mark yaml_mark_t
		problem      string
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
			if got := yaml_parser_set_scanner_error(tt.args.parser, tt.args.context, tt.args.context_mark, tt.args.problem); got != tt.want {
				t.Errorf("yaml_parser_set_scanner_error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_set_scanner_tag_error(t *testing.T) {
	type args struct {
		parser       *yaml_parser_t
		directive    bool
		context_mark yaml_mark_t
		problem      string
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
			if got := yaml_parser_set_scanner_tag_error(tt.args.parser, tt.args.directive, tt.args.context_mark, tt.args.problem); got != tt.want {
				t.Errorf("yaml_parser_set_scanner_tag_error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trace(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want func()
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trace(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_more_tokens(t *testing.T) {
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
			if got := yaml_parser_fetch_more_tokens(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_fetch_more_tokens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_next_token(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
	}
	tests := []struct {
		name   string
		args   args
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOk := yaml_parser_fetch_next_token(tt.args.parser); gotOk != tt.wantOk {
				t.Errorf("yaml_parser_fetch_next_token() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_yaml_simple_key_is_valid(t *testing.T) {
	type args struct {
		parser     *yaml_parser_t
		simple_key *yaml_simple_key_t
	}
	tests := []struct {
		name      string
		args      args
		wantValid bool
		wantOk    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValid, gotOk := yaml_simple_key_is_valid(tt.args.parser, tt.args.simple_key)
			if gotValid != tt.wantValid {
				t.Errorf("yaml_simple_key_is_valid() gotValid = %v, want %v", gotValid, tt.wantValid)
			}
			if gotOk != tt.wantOk {
				t.Errorf("yaml_simple_key_is_valid() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_yaml_parser_save_simple_key(t *testing.T) {
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
			if got := yaml_parser_save_simple_key(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_save_simple_key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_remove_simple_key(t *testing.T) {
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
			if got := yaml_parser_remove_simple_key(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_remove_simple_key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_increase_flow_level(t *testing.T) {
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
			if got := yaml_parser_increase_flow_level(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_increase_flow_level() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_decrease_flow_level(t *testing.T) {
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
			if got := yaml_parser_decrease_flow_level(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_decrease_flow_level() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_roll_indent(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		column int
		number int
		typ    yaml_token_type_t
		mark   yaml_mark_t
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
			if got := yaml_parser_roll_indent(tt.args.parser, tt.args.column, tt.args.number, tt.args.typ, tt.args.mark); got != tt.want {
				t.Errorf("yaml_parser_roll_indent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_unroll_indent(t *testing.T) {
	type args struct {
		parser    *yaml_parser_t
		column    int
		scan_mark yaml_mark_t
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
			if got := yaml_parser_unroll_indent(tt.args.parser, tt.args.column, tt.args.scan_mark); got != tt.want {
				t.Errorf("yaml_parser_unroll_indent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_stream_start(t *testing.T) {
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
			if got := yaml_parser_fetch_stream_start(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_fetch_stream_start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_stream_end(t *testing.T) {
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
			if got := yaml_parser_fetch_stream_end(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_fetch_stream_end() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_directive(t *testing.T) {
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
			if got := yaml_parser_fetch_directive(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_fetch_directive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_document_indicator(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		typ    yaml_token_type_t
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
			if got := yaml_parser_fetch_document_indicator(tt.args.parser, tt.args.typ); got != tt.want {
				t.Errorf("yaml_parser_fetch_document_indicator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_flow_collection_start(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		typ    yaml_token_type_t
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
			if got := yaml_parser_fetch_flow_collection_start(tt.args.parser, tt.args.typ); got != tt.want {
				t.Errorf("yaml_parser_fetch_flow_collection_start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_flow_collection_end(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		typ    yaml_token_type_t
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
			if got := yaml_parser_fetch_flow_collection_end(tt.args.parser, tt.args.typ); got != tt.want {
				t.Errorf("yaml_parser_fetch_flow_collection_end() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_flow_entry(t *testing.T) {
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
			if got := yaml_parser_fetch_flow_entry(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_fetch_flow_entry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_block_entry(t *testing.T) {
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
			if got := yaml_parser_fetch_block_entry(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_fetch_block_entry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_key(t *testing.T) {
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
			if got := yaml_parser_fetch_key(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_fetch_key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_value(t *testing.T) {
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
			if got := yaml_parser_fetch_value(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_fetch_value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_anchor(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		typ    yaml_token_type_t
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
			if got := yaml_parser_fetch_anchor(tt.args.parser, tt.args.typ); got != tt.want {
				t.Errorf("yaml_parser_fetch_anchor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_tag(t *testing.T) {
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
			if got := yaml_parser_fetch_tag(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_fetch_tag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_block_scalar(t *testing.T) {
	type args struct {
		parser  *yaml_parser_t
		literal bool
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
			if got := yaml_parser_fetch_block_scalar(tt.args.parser, tt.args.literal); got != tt.want {
				t.Errorf("yaml_parser_fetch_block_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_flow_scalar(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		single bool
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
			if got := yaml_parser_fetch_flow_scalar(tt.args.parser, tt.args.single); got != tt.want {
				t.Errorf("yaml_parser_fetch_flow_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_fetch_plain_scalar(t *testing.T) {
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
			if got := yaml_parser_fetch_plain_scalar(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_fetch_plain_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_to_next_token(t *testing.T) {
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
			if got := yaml_parser_scan_to_next_token(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_scan_to_next_token() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_directive(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		token  *yaml_token_t
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
			if got := yaml_parser_scan_directive(tt.args.parser, tt.args.token); got != tt.want {
				t.Errorf("yaml_parser_scan_directive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_directive_name(t *testing.T) {
	type args struct {
		parser     *yaml_parser_t
		start_mark yaml_mark_t
		name       *[]byte
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
			if got := yaml_parser_scan_directive_name(tt.args.parser, tt.args.start_mark, tt.args.name); got != tt.want {
				t.Errorf("yaml_parser_scan_directive_name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_version_directive_value(t *testing.T) {
	type args struct {
		parser     *yaml_parser_t
		start_mark yaml_mark_t
		major      *int8
		minor      *int8
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
			if got := yaml_parser_scan_version_directive_value(tt.args.parser, tt.args.start_mark, tt.args.major, tt.args.minor); got != tt.want {
				t.Errorf("yaml_parser_scan_version_directive_value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_version_directive_number(t *testing.T) {
	type args struct {
		parser     *yaml_parser_t
		start_mark yaml_mark_t
		number     *int8
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
			if got := yaml_parser_scan_version_directive_number(tt.args.parser, tt.args.start_mark, tt.args.number); got != tt.want {
				t.Errorf("yaml_parser_scan_version_directive_number() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_tag_directive_value(t *testing.T) {
	type args struct {
		parser     *yaml_parser_t
		start_mark yaml_mark_t
		handle     *[]byte
		prefix     *[]byte
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
			if got := yaml_parser_scan_tag_directive_value(tt.args.parser, tt.args.start_mark, tt.args.handle, tt.args.prefix); got != tt.want {
				t.Errorf("yaml_parser_scan_tag_directive_value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_anchor(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		token  *yaml_token_t
		typ    yaml_token_type_t
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
			if got := yaml_parser_scan_anchor(tt.args.parser, tt.args.token, tt.args.typ); got != tt.want {
				t.Errorf("yaml_parser_scan_anchor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_tag(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		token  *yaml_token_t
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
			if got := yaml_parser_scan_tag(tt.args.parser, tt.args.token); got != tt.want {
				t.Errorf("yaml_parser_scan_tag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_tag_handle(t *testing.T) {
	type args struct {
		parser     *yaml_parser_t
		directive  bool
		start_mark yaml_mark_t
		handle     *[]byte
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
			if got := yaml_parser_scan_tag_handle(tt.args.parser, tt.args.directive, tt.args.start_mark, tt.args.handle); got != tt.want {
				t.Errorf("yaml_parser_scan_tag_handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_tag_uri(t *testing.T) {
	type args struct {
		parser     *yaml_parser_t
		directive  bool
		head       []byte
		start_mark yaml_mark_t
		uri        *[]byte
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
			if got := yaml_parser_scan_tag_uri(tt.args.parser, tt.args.directive, tt.args.head, tt.args.start_mark, tt.args.uri); got != tt.want {
				t.Errorf("yaml_parser_scan_tag_uri() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_uri_escapes(t *testing.T) {
	type args struct {
		parser     *yaml_parser_t
		directive  bool
		start_mark yaml_mark_t
		s          *[]byte
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
			if got := yaml_parser_scan_uri_escapes(tt.args.parser, tt.args.directive, tt.args.start_mark, tt.args.s); got != tt.want {
				t.Errorf("yaml_parser_scan_uri_escapes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_block_scalar(t *testing.T) {
	type args struct {
		parser  *yaml_parser_t
		token   *yaml_token_t
		literal bool
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
			if got := yaml_parser_scan_block_scalar(tt.args.parser, tt.args.token, tt.args.literal); got != tt.want {
				t.Errorf("yaml_parser_scan_block_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_block_scalar_breaks(t *testing.T) {
	type args struct {
		parser     *yaml_parser_t
		indent     *int
		breaks     *[]byte
		start_mark yaml_mark_t
		end_mark   *yaml_mark_t
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
			if got := yaml_parser_scan_block_scalar_breaks(tt.args.parser, tt.args.indent, tt.args.breaks, tt.args.start_mark, tt.args.end_mark); got != tt.want {
				t.Errorf("yaml_parser_scan_block_scalar_breaks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_flow_scalar(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		token  *yaml_token_t
		single bool
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
			if got := yaml_parser_scan_flow_scalar(tt.args.parser, tt.args.token, tt.args.single); got != tt.want {
				t.Errorf("yaml_parser_scan_flow_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_plain_scalar(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		token  *yaml_token_t
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
			if got := yaml_parser_scan_plain_scalar(tt.args.parser, tt.args.token); got != tt.want {
				t.Errorf("yaml_parser_scan_plain_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_line_comment(t *testing.T) {
	type args struct {
		parser     *yaml_parser_t
		token_mark yaml_mark_t
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
			if got := yaml_parser_scan_line_comment(tt.args.parser, tt.args.token_mark); got != tt.want {
				t.Errorf("yaml_parser_scan_line_comment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_scan_comments(t *testing.T) {
	type args struct {
		parser    *yaml_parser_t
		scan_mark yaml_mark_t
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
			if got := yaml_parser_scan_comments(tt.args.parser, tt.args.scan_mark); got != tt.want {
				t.Errorf("yaml_parser_scan_comments() = %v, want %v", got, tt.want)
			}
		})
	}
}
