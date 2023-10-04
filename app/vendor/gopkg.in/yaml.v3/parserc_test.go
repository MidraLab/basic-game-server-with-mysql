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

func Test_peek_token(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
	}
	tests := []struct {
		name string
		args args
		want *yaml_token_t
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peek_token(tt.args.parser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("peek_token() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_unfold_comments(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		token  *yaml_token_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_parser_unfold_comments(tt.args.parser, tt.args.token)
		})
	}
}

func Test_skip_token(t *testing.T) {
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
			skip_token(tt.args.parser)
		})
	}
}

func Test_yaml_parser_parse(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
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
			if got := yaml_parser_parse(tt.args.parser, tt.args.event); got != tt.want {
				t.Errorf("yaml_parser_parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_set_parser_error(t *testing.T) {
	type args struct {
		parser       *yaml_parser_t
		problem      string
		problem_mark yaml_mark_t
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
			if got := yaml_parser_set_parser_error(tt.args.parser, tt.args.problem, tt.args.problem_mark); got != tt.want {
				t.Errorf("yaml_parser_set_parser_error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_set_parser_error_context(t *testing.T) {
	type args struct {
		parser       *yaml_parser_t
		context      string
		context_mark yaml_mark_t
		problem      string
		problem_mark yaml_mark_t
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
			if got := yaml_parser_set_parser_error_context(tt.args.parser, tt.args.context, tt.args.context_mark, tt.args.problem, tt.args.problem_mark); got != tt.want {
				t.Errorf("yaml_parser_set_parser_error_context() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_state_machine(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
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
			if got := yaml_parser_state_machine(tt.args.parser, tt.args.event); got != tt.want {
				t.Errorf("yaml_parser_state_machine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_stream_start(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
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
			if got := yaml_parser_parse_stream_start(tt.args.parser, tt.args.event); got != tt.want {
				t.Errorf("yaml_parser_parse_stream_start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_document_start(t *testing.T) {
	type args struct {
		parser   *yaml_parser_t
		event    *yaml_event_t
		implicit bool
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
			if got := yaml_parser_parse_document_start(tt.args.parser, tt.args.event, tt.args.implicit); got != tt.want {
				t.Errorf("yaml_parser_parse_document_start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_document_content(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
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
			if got := yaml_parser_parse_document_content(tt.args.parser, tt.args.event); got != tt.want {
				t.Errorf("yaml_parser_parse_document_content() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_document_end(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
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
			if got := yaml_parser_parse_document_end(tt.args.parser, tt.args.event); got != tt.want {
				t.Errorf("yaml_parser_parse_document_end() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_set_event_comments(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_parser_set_event_comments(tt.args.parser, tt.args.event)
		})
	}
}

func Test_yaml_parser_parse_node(t *testing.T) {
	type args struct {
		parser              *yaml_parser_t
		event               *yaml_event_t
		block               bool
		indentless_sequence bool
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
			if got := yaml_parser_parse_node(tt.args.parser, tt.args.event, tt.args.block, tt.args.indentless_sequence); got != tt.want {
				t.Errorf("yaml_parser_parse_node() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_block_sequence_entry(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
		first  bool
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
			if got := yaml_parser_parse_block_sequence_entry(tt.args.parser, tt.args.event, tt.args.first); got != tt.want {
				t.Errorf("yaml_parser_parse_block_sequence_entry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_indentless_sequence_entry(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
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
			if got := yaml_parser_parse_indentless_sequence_entry(tt.args.parser, tt.args.event); got != tt.want {
				t.Errorf("yaml_parser_parse_indentless_sequence_entry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_split_stem_comment(t *testing.T) {
	type args struct {
		parser   *yaml_parser_t
		stem_len int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_parser_split_stem_comment(tt.args.parser, tt.args.stem_len)
		})
	}
}

func Test_yaml_parser_parse_block_mapping_key(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
		first  bool
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
			if got := yaml_parser_parse_block_mapping_key(tt.args.parser, tt.args.event, tt.args.first); got != tt.want {
				t.Errorf("yaml_parser_parse_block_mapping_key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_block_mapping_value(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
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
			if got := yaml_parser_parse_block_mapping_value(tt.args.parser, tt.args.event); got != tt.want {
				t.Errorf("yaml_parser_parse_block_mapping_value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_flow_sequence_entry(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
		first  bool
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
			if got := yaml_parser_parse_flow_sequence_entry(tt.args.parser, tt.args.event, tt.args.first); got != tt.want {
				t.Errorf("yaml_parser_parse_flow_sequence_entry() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_flow_sequence_entry_mapping_key(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
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
			if got := yaml_parser_parse_flow_sequence_entry_mapping_key(tt.args.parser, tt.args.event); got != tt.want {
				t.Errorf("yaml_parser_parse_flow_sequence_entry_mapping_key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_flow_sequence_entry_mapping_value(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
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
			if got := yaml_parser_parse_flow_sequence_entry_mapping_value(tt.args.parser, tt.args.event); got != tt.want {
				t.Errorf("yaml_parser_parse_flow_sequence_entry_mapping_value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_flow_sequence_entry_mapping_end(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
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
			if got := yaml_parser_parse_flow_sequence_entry_mapping_end(tt.args.parser, tt.args.event); got != tt.want {
				t.Errorf("yaml_parser_parse_flow_sequence_entry_mapping_end() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_flow_mapping_key(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
		first  bool
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
			if got := yaml_parser_parse_flow_mapping_key(tt.args.parser, tt.args.event, tt.args.first); got != tt.want {
				t.Errorf("yaml_parser_parse_flow_mapping_key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_parse_flow_mapping_value(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
		empty  bool
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
			if got := yaml_parser_parse_flow_mapping_value(tt.args.parser, tt.args.event, tt.args.empty); got != tt.want {
				t.Errorf("yaml_parser_parse_flow_mapping_value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_process_empty_scalar(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		event  *yaml_event_t
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
			if got := yaml_parser_process_empty_scalar(tt.args.parser, tt.args.event, tt.args.mark); got != tt.want {
				t.Errorf("yaml_parser_process_empty_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_process_directives(t *testing.T) {
	type args struct {
		parser                *yaml_parser_t
		version_directive_ref **yaml_version_directive_t
		tag_directives_ref    *[]yaml_tag_directive_t
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
			if got := yaml_parser_process_directives(tt.args.parser, tt.args.version_directive_ref, tt.args.tag_directives_ref); got != tt.want {
				t.Errorf("yaml_parser_process_directives() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_append_tag_directive(t *testing.T) {
	type args struct {
		parser           *yaml_parser_t
		value            yaml_tag_directive_t
		allow_duplicates bool
		mark             yaml_mark_t
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
			if got := yaml_parser_append_tag_directive(tt.args.parser, tt.args.value, tt.args.allow_duplicates, tt.args.mark); got != tt.want {
				t.Errorf("yaml_parser_append_tag_directive() = %v, want %v", got, tt.want)
			}
		})
	}
}
