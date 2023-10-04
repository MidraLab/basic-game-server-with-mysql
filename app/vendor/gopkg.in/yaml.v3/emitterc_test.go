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

func Test_flush(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := flush(tt.args.emitter); got != tt.want {
				t.Errorf("flush() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_put(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		value   byte
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
			if got := put(tt.args.emitter, tt.args.value); got != tt.want {
				t.Errorf("put() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_put_break(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := put_break(tt.args.emitter); got != tt.want {
				t.Errorf("put_break() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_write(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		s       []byte
		i       *int
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
			if got := write(tt.args.emitter, tt.args.s, tt.args.i); got != tt.want {
				t.Errorf("write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_write_all(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		s       []byte
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
			if got := write_all(tt.args.emitter, tt.args.s); got != tt.want {
				t.Errorf("write_all() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_write_break(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		s       []byte
		i       *int
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
			if got := write_break(tt.args.emitter, tt.args.s, tt.args.i); got != tt.want {
				t.Errorf("write_break() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_set_emitter_error(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		problem string
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
			if got := yaml_emitter_set_emitter_error(tt.args.emitter, tt.args.problem); got != tt.want {
				t.Errorf("yaml_emitter_set_emitter_error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
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
			if got := yaml_emitter_emit(tt.args.emitter, tt.args.event); got != tt.want {
				t.Errorf("yaml_emitter_emit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_need_more_events(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_need_more_events(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_need_more_events() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_append_tag_directive(t *testing.T) {
	type args struct {
		emitter          *yaml_emitter_t
		value            *yaml_tag_directive_t
		allow_duplicates bool
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
			if got := yaml_emitter_append_tag_directive(tt.args.emitter, tt.args.value, tt.args.allow_duplicates); got != tt.want {
				t.Errorf("yaml_emitter_append_tag_directive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_increase_indent(t *testing.T) {
	type args struct {
		emitter    *yaml_emitter_t
		flow       bool
		indentless bool
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
			if got := yaml_emitter_increase_indent(tt.args.emitter, tt.args.flow, tt.args.indentless); got != tt.want {
				t.Errorf("yaml_emitter_increase_indent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_state_machine(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
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
			if got := yaml_emitter_state_machine(tt.args.emitter, tt.args.event); got != tt.want {
				t.Errorf("yaml_emitter_state_machine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_stream_start(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
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
			if got := yaml_emitter_emit_stream_start(tt.args.emitter, tt.args.event); got != tt.want {
				t.Errorf("yaml_emitter_emit_stream_start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_document_start(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
		first   bool
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
			if got := yaml_emitter_emit_document_start(tt.args.emitter, tt.args.event, tt.args.first); got != tt.want {
				t.Errorf("yaml_emitter_emit_document_start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_document_content(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
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
			if got := yaml_emitter_emit_document_content(tt.args.emitter, tt.args.event); got != tt.want {
				t.Errorf("yaml_emitter_emit_document_content() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_document_end(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
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
			if got := yaml_emitter_emit_document_end(tt.args.emitter, tt.args.event); got != tt.want {
				t.Errorf("yaml_emitter_emit_document_end() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_flow_sequence_item(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
		first   bool
		trail   bool
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
			if got := yaml_emitter_emit_flow_sequence_item(tt.args.emitter, tt.args.event, tt.args.first, tt.args.trail); got != tt.want {
				t.Errorf("yaml_emitter_emit_flow_sequence_item() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_flow_mapping_key(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
		first   bool
		trail   bool
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
			if got := yaml_emitter_emit_flow_mapping_key(tt.args.emitter, tt.args.event, tt.args.first, tt.args.trail); got != tt.want {
				t.Errorf("yaml_emitter_emit_flow_mapping_key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_flow_mapping_value(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
		simple  bool
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
			if got := yaml_emitter_emit_flow_mapping_value(tt.args.emitter, tt.args.event, tt.args.simple); got != tt.want {
				t.Errorf("yaml_emitter_emit_flow_mapping_value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_block_sequence_item(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
		first   bool
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
			if got := yaml_emitter_emit_block_sequence_item(tt.args.emitter, tt.args.event, tt.args.first); got != tt.want {
				t.Errorf("yaml_emitter_emit_block_sequence_item() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_block_mapping_key(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
		first   bool
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
			if got := yaml_emitter_emit_block_mapping_key(tt.args.emitter, tt.args.event, tt.args.first); got != tt.want {
				t.Errorf("yaml_emitter_emit_block_mapping_key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_block_mapping_value(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
		simple  bool
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
			if got := yaml_emitter_emit_block_mapping_value(tt.args.emitter, tt.args.event, tt.args.simple); got != tt.want {
				t.Errorf("yaml_emitter_emit_block_mapping_value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_silent_nil_event(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
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
			if got := yaml_emitter_silent_nil_event(tt.args.emitter, tt.args.event); got != tt.want {
				t.Errorf("yaml_emitter_silent_nil_event() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_node(t *testing.T) {
	type args struct {
		emitter    *yaml_emitter_t
		event      *yaml_event_t
		root       bool
		sequence   bool
		mapping    bool
		simple_key bool
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
			if got := yaml_emitter_emit_node(tt.args.emitter, tt.args.event, tt.args.root, tt.args.sequence, tt.args.mapping, tt.args.simple_key); got != tt.want {
				t.Errorf("yaml_emitter_emit_node() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_alias(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
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
			if got := yaml_emitter_emit_alias(tt.args.emitter, tt.args.event); got != tt.want {
				t.Errorf("yaml_emitter_emit_alias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_scalar(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
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
			if got := yaml_emitter_emit_scalar(tt.args.emitter, tt.args.event); got != tt.want {
				t.Errorf("yaml_emitter_emit_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_sequence_start(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
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
			if got := yaml_emitter_emit_sequence_start(tt.args.emitter, tt.args.event); got != tt.want {
				t.Errorf("yaml_emitter_emit_sequence_start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_emit_mapping_start(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
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
			if got := yaml_emitter_emit_mapping_start(tt.args.emitter, tt.args.event); got != tt.want {
				t.Errorf("yaml_emitter_emit_mapping_start() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_check_empty_document(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_check_empty_document(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_check_empty_document() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_check_empty_sequence(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_check_empty_sequence(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_check_empty_sequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_check_empty_mapping(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_check_empty_mapping(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_check_empty_mapping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_check_simple_key(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_check_simple_key(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_check_simple_key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_select_scalar_style(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
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
			if got := yaml_emitter_select_scalar_style(tt.args.emitter, tt.args.event); got != tt.want {
				t.Errorf("yaml_emitter_select_scalar_style() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_process_anchor(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_process_anchor(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_process_anchor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_process_tag(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_process_tag(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_process_tag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_process_scalar(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_process_scalar(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_process_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_process_head_comment(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_process_head_comment(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_process_head_comment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_process_line_comment(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_process_line_comment(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_process_line_comment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_process_foot_comment(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_process_foot_comment(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_process_foot_comment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_analyze_version_directive(t *testing.T) {
	type args struct {
		emitter           *yaml_emitter_t
		version_directive *yaml_version_directive_t
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
			if got := yaml_emitter_analyze_version_directive(tt.args.emitter, tt.args.version_directive); got != tt.want {
				t.Errorf("yaml_emitter_analyze_version_directive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_analyze_tag_directive(t *testing.T) {
	type args struct {
		emitter       *yaml_emitter_t
		tag_directive *yaml_tag_directive_t
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
			if got := yaml_emitter_analyze_tag_directive(tt.args.emitter, tt.args.tag_directive); got != tt.want {
				t.Errorf("yaml_emitter_analyze_tag_directive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_analyze_anchor(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		anchor  []byte
		alias   bool
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
			if got := yaml_emitter_analyze_anchor(tt.args.emitter, tt.args.anchor, tt.args.alias); got != tt.want {
				t.Errorf("yaml_emitter_analyze_anchor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_analyze_tag(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		tag     []byte
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
			if got := yaml_emitter_analyze_tag(tt.args.emitter, tt.args.tag); got != tt.want {
				t.Errorf("yaml_emitter_analyze_tag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_analyze_scalar(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		value   []byte
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
			if got := yaml_emitter_analyze_scalar(tt.args.emitter, tt.args.value); got != tt.want {
				t.Errorf("yaml_emitter_analyze_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_analyze_event(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		event   *yaml_event_t
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
			if got := yaml_emitter_analyze_event(tt.args.emitter, tt.args.event); got != tt.want {
				t.Errorf("yaml_emitter_analyze_event() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_bom(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_write_bom(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_write_bom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_indent(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
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
			if got := yaml_emitter_write_indent(tt.args.emitter); got != tt.want {
				t.Errorf("yaml_emitter_write_indent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_indicator(t *testing.T) {
	type args struct {
		emitter         *yaml_emitter_t
		indicator       []byte
		need_whitespace bool
		is_whitespace   bool
		is_indention    bool
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
			if got := yaml_emitter_write_indicator(tt.args.emitter, tt.args.indicator, tt.args.need_whitespace, tt.args.is_whitespace, tt.args.is_indention); got != tt.want {
				t.Errorf("yaml_emitter_write_indicator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_anchor(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		value   []byte
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
			if got := yaml_emitter_write_anchor(tt.args.emitter, tt.args.value); got != tt.want {
				t.Errorf("yaml_emitter_write_anchor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_tag_handle(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		value   []byte
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
			if got := yaml_emitter_write_tag_handle(tt.args.emitter, tt.args.value); got != tt.want {
				t.Errorf("yaml_emitter_write_tag_handle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_tag_content(t *testing.T) {
	type args struct {
		emitter         *yaml_emitter_t
		value           []byte
		need_whitespace bool
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
			if got := yaml_emitter_write_tag_content(tt.args.emitter, tt.args.value, tt.args.need_whitespace); got != tt.want {
				t.Errorf("yaml_emitter_write_tag_content() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_plain_scalar(t *testing.T) {
	type args struct {
		emitter      *yaml_emitter_t
		value        []byte
		allow_breaks bool
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
			if got := yaml_emitter_write_plain_scalar(tt.args.emitter, tt.args.value, tt.args.allow_breaks); got != tt.want {
				t.Errorf("yaml_emitter_write_plain_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_single_quoted_scalar(t *testing.T) {
	type args struct {
		emitter      *yaml_emitter_t
		value        []byte
		allow_breaks bool
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
			if got := yaml_emitter_write_single_quoted_scalar(tt.args.emitter, tt.args.value, tt.args.allow_breaks); got != tt.want {
				t.Errorf("yaml_emitter_write_single_quoted_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_double_quoted_scalar(t *testing.T) {
	type args struct {
		emitter      *yaml_emitter_t
		value        []byte
		allow_breaks bool
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
			if got := yaml_emitter_write_double_quoted_scalar(tt.args.emitter, tt.args.value, tt.args.allow_breaks); got != tt.want {
				t.Errorf("yaml_emitter_write_double_quoted_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_block_scalar_hints(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		value   []byte
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
			if got := yaml_emitter_write_block_scalar_hints(tt.args.emitter, tt.args.value); got != tt.want {
				t.Errorf("yaml_emitter_write_block_scalar_hints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_literal_scalar(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		value   []byte
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
			if got := yaml_emitter_write_literal_scalar(tt.args.emitter, tt.args.value); got != tt.want {
				t.Errorf("yaml_emitter_write_literal_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_folded_scalar(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		value   []byte
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
			if got := yaml_emitter_write_folded_scalar(tt.args.emitter, tt.args.value); got != tt.want {
				t.Errorf("yaml_emitter_write_folded_scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_emitter_write_comment(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		comment []byte
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
			if got := yaml_emitter_write_comment(tt.args.emitter, tt.args.comment); got != tt.want {
				t.Errorf("yaml_emitter_write_comment() = %v, want %v", got, tt.want)
			}
		})
	}
}
