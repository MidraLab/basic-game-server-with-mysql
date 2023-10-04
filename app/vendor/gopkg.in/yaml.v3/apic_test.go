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
	"bytes"
	"io"
	"testing"
)

func Test_yaml_insert_token(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		pos    int
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
			yaml_insert_token(tt.args.parser, tt.args.pos, tt.args.token)
		})
	}
}

func Test_yaml_parser_initialize(t *testing.T) {
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
			if got := yaml_parser_initialize(tt.args.parser); got != tt.want {
				t.Errorf("yaml_parser_initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_delete(t *testing.T) {
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
			yaml_parser_delete(tt.args.parser)
		})
	}
}

func Test_yaml_string_read_handler(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		buffer []byte
	}
	tests := []struct {
		name    string
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := yaml_string_read_handler(tt.args.parser, tt.args.buffer)
			if (err != nil) != tt.wantErr {
				t.Errorf("yaml_string_read_handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("yaml_string_read_handler() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_yaml_reader_read_handler(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		buffer []byte
	}
	tests := []struct {
		name    string
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := yaml_reader_read_handler(tt.args.parser, tt.args.buffer)
			if (err != nil) != tt.wantErr {
				t.Errorf("yaml_reader_read_handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("yaml_reader_read_handler() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func Test_yaml_parser_set_input_string(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		input  []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_parser_set_input_string(tt.args.parser, tt.args.input)
		})
	}
}

func Test_yaml_parser_set_input_reader(t *testing.T) {
	type args struct {
		parser *yaml_parser_t
		r      io.Reader
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_parser_set_input_reader(tt.args.parser, tt.args.r)
		})
	}
}

func Test_yaml_parser_set_encoding(t *testing.T) {
	type args struct {
		parser   *yaml_parser_t
		encoding yaml_encoding_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_parser_set_encoding(tt.args.parser, tt.args.encoding)
		})
	}
}

func Test_yaml_emitter_initialize(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_emitter_initialize(tt.args.emitter)
		})
	}
}

func Test_yaml_emitter_delete(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_emitter_delete(tt.args.emitter)
		})
	}
}

func Test_yaml_string_write_handler(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		buffer  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := yaml_string_write_handler(tt.args.emitter, tt.args.buffer); (err != nil) != tt.wantErr {
				t.Errorf("yaml_string_write_handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_yaml_writer_write_handler(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		buffer  []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := yaml_writer_write_handler(tt.args.emitter, tt.args.buffer); (err != nil) != tt.wantErr {
				t.Errorf("yaml_writer_write_handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_yaml_emitter_set_output_string(t *testing.T) {
	type args struct {
		emitter       *yaml_emitter_t
		output_buffer *[]byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_emitter_set_output_string(tt.args.emitter, tt.args.output_buffer)
		})
	}
}

func Test_yaml_emitter_set_output_writer(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			yaml_emitter_set_output_writer(tt.args.emitter, w)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("yaml_emitter_set_output_writer() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_yaml_emitter_set_encoding(t *testing.T) {
	type args struct {
		emitter  *yaml_emitter_t
		encoding yaml_encoding_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_emitter_set_encoding(tt.args.emitter, tt.args.encoding)
		})
	}
}

func Test_yaml_emitter_set_canonical(t *testing.T) {
	type args struct {
		emitter   *yaml_emitter_t
		canonical bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_emitter_set_canonical(tt.args.emitter, tt.args.canonical)
		})
	}
}

func Test_yaml_emitter_set_indent(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		indent  int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_emitter_set_indent(tt.args.emitter, tt.args.indent)
		})
	}
}

func Test_yaml_emitter_set_width(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		width   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_emitter_set_width(tt.args.emitter, tt.args.width)
		})
	}
}

func Test_yaml_emitter_set_unicode(t *testing.T) {
	type args struct {
		emitter *yaml_emitter_t
		unicode bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_emitter_set_unicode(tt.args.emitter, tt.args.unicode)
		})
	}
}

func Test_yaml_emitter_set_break(t *testing.T) {
	type args struct {
		emitter    *yaml_emitter_t
		line_break yaml_break_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_emitter_set_break(tt.args.emitter, tt.args.line_break)
		})
	}
}

func Test_yaml_stream_start_event_initialize(t *testing.T) {
	type args struct {
		event    *yaml_event_t
		encoding yaml_encoding_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_stream_start_event_initialize(tt.args.event, tt.args.encoding)
		})
	}
}

func Test_yaml_stream_end_event_initialize(t *testing.T) {
	type args struct {
		event *yaml_event_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_stream_end_event_initialize(tt.args.event)
		})
	}
}

func Test_yaml_document_start_event_initialize(t *testing.T) {
	type args struct {
		event             *yaml_event_t
		version_directive *yaml_version_directive_t
		tag_directives    []yaml_tag_directive_t
		implicit          bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_document_start_event_initialize(tt.args.event, tt.args.version_directive, tt.args.tag_directives, tt.args.implicit)
		})
	}
}

func Test_yaml_document_end_event_initialize(t *testing.T) {
	type args struct {
		event    *yaml_event_t
		implicit bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_document_end_event_initialize(tt.args.event, tt.args.implicit)
		})
	}
}

func Test_yaml_alias_event_initialize(t *testing.T) {
	type args struct {
		event  *yaml_event_t
		anchor []byte
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
			if got := yaml_alias_event_initialize(tt.args.event, tt.args.anchor); got != tt.want {
				t.Errorf("yaml_alias_event_initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_scalar_event_initialize(t *testing.T) {
	type args struct {
		event           *yaml_event_t
		anchor          []byte
		tag             []byte
		value           []byte
		plain_implicit  bool
		quoted_implicit bool
		style           yaml_scalar_style_t
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
			if got := yaml_scalar_event_initialize(tt.args.event, tt.args.anchor, tt.args.tag, tt.args.value, tt.args.plain_implicit, tt.args.quoted_implicit, tt.args.style); got != tt.want {
				t.Errorf("yaml_scalar_event_initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_sequence_start_event_initialize(t *testing.T) {
	type args struct {
		event    *yaml_event_t
		anchor   []byte
		tag      []byte
		implicit bool
		style    yaml_sequence_style_t
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
			if got := yaml_sequence_start_event_initialize(tt.args.event, tt.args.anchor, tt.args.tag, tt.args.implicit, tt.args.style); got != tt.want {
				t.Errorf("yaml_sequence_start_event_initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_sequence_end_event_initialize(t *testing.T) {
	type args struct {
		event *yaml_event_t
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
			if got := yaml_sequence_end_event_initialize(tt.args.event); got != tt.want {
				t.Errorf("yaml_sequence_end_event_initialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_mapping_start_event_initialize(t *testing.T) {
	type args struct {
		event    *yaml_event_t
		anchor   []byte
		tag      []byte
		implicit bool
		style    yaml_mapping_style_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_mapping_start_event_initialize(tt.args.event, tt.args.anchor, tt.args.tag, tt.args.implicit, tt.args.style)
		})
	}
}

func Test_yaml_mapping_end_event_initialize(t *testing.T) {
	type args struct {
		event *yaml_event_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_mapping_end_event_initialize(tt.args.event)
		})
	}
}

func Test_yaml_event_delete(t *testing.T) {
	type args struct {
		event *yaml_event_t
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			yaml_event_delete(tt.args.event)
		})
	}
}
