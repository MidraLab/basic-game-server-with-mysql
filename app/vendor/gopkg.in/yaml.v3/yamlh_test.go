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

func Test_yaml_token_type_t_String(t *testing.T) {
	tests := []struct {
		name string
		tt   yaml_token_type_t
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tt.String(); got != tt.want {
				t.Errorf("yaml_token_type_t.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_event_type_t_String(t *testing.T) {
	tests := []struct {
		name string
		e    yaml_event_type_t
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("yaml_event_type_t.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_event_t_scalar_style(t *testing.T) {
	type fields struct {
		typ               yaml_event_type_t
		start_mark        yaml_mark_t
		end_mark          yaml_mark_t
		encoding          yaml_encoding_t
		version_directive *yaml_version_directive_t
		tag_directives    []yaml_tag_directive_t
		head_comment      []byte
		line_comment      []byte
		foot_comment      []byte
		tail_comment      []byte
		anchor            []byte
		tag               []byte
		value             []byte
		implicit          bool
		quoted_implicit   bool
		style             yaml_style_t
	}
	tests := []struct {
		name   string
		fields fields
		want   yaml_scalar_style_t
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &yaml_event_t{
				typ:               tt.fields.typ,
				start_mark:        tt.fields.start_mark,
				end_mark:          tt.fields.end_mark,
				encoding:          tt.fields.encoding,
				version_directive: tt.fields.version_directive,
				tag_directives:    tt.fields.tag_directives,
				head_comment:      tt.fields.head_comment,
				line_comment:      tt.fields.line_comment,
				foot_comment:      tt.fields.foot_comment,
				tail_comment:      tt.fields.tail_comment,
				anchor:            tt.fields.anchor,
				tag:               tt.fields.tag,
				value:             tt.fields.value,
				implicit:          tt.fields.implicit,
				quoted_implicit:   tt.fields.quoted_implicit,
				style:             tt.fields.style,
			}
			if got := e.scalar_style(); got != tt.want {
				t.Errorf("yaml_event_t.scalar_style() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_event_t_sequence_style(t *testing.T) {
	type fields struct {
		typ               yaml_event_type_t
		start_mark        yaml_mark_t
		end_mark          yaml_mark_t
		encoding          yaml_encoding_t
		version_directive *yaml_version_directive_t
		tag_directives    []yaml_tag_directive_t
		head_comment      []byte
		line_comment      []byte
		foot_comment      []byte
		tail_comment      []byte
		anchor            []byte
		tag               []byte
		value             []byte
		implicit          bool
		quoted_implicit   bool
		style             yaml_style_t
	}
	tests := []struct {
		name   string
		fields fields
		want   yaml_sequence_style_t
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &yaml_event_t{
				typ:               tt.fields.typ,
				start_mark:        tt.fields.start_mark,
				end_mark:          tt.fields.end_mark,
				encoding:          tt.fields.encoding,
				version_directive: tt.fields.version_directive,
				tag_directives:    tt.fields.tag_directives,
				head_comment:      tt.fields.head_comment,
				line_comment:      tt.fields.line_comment,
				foot_comment:      tt.fields.foot_comment,
				tail_comment:      tt.fields.tail_comment,
				anchor:            tt.fields.anchor,
				tag:               tt.fields.tag,
				value:             tt.fields.value,
				implicit:          tt.fields.implicit,
				quoted_implicit:   tt.fields.quoted_implicit,
				style:             tt.fields.style,
			}
			if got := e.sequence_style(); got != tt.want {
				t.Errorf("yaml_event_t.sequence_style() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_event_t_mapping_style(t *testing.T) {
	type fields struct {
		typ               yaml_event_type_t
		start_mark        yaml_mark_t
		end_mark          yaml_mark_t
		encoding          yaml_encoding_t
		version_directive *yaml_version_directive_t
		tag_directives    []yaml_tag_directive_t
		head_comment      []byte
		line_comment      []byte
		foot_comment      []byte
		tail_comment      []byte
		anchor            []byte
		tag               []byte
		value             []byte
		implicit          bool
		quoted_implicit   bool
		style             yaml_style_t
	}
	tests := []struct {
		name   string
		fields fields
		want   yaml_mapping_style_t
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &yaml_event_t{
				typ:               tt.fields.typ,
				start_mark:        tt.fields.start_mark,
				end_mark:          tt.fields.end_mark,
				encoding:          tt.fields.encoding,
				version_directive: tt.fields.version_directive,
				tag_directives:    tt.fields.tag_directives,
				head_comment:      tt.fields.head_comment,
				line_comment:      tt.fields.line_comment,
				foot_comment:      tt.fields.foot_comment,
				tail_comment:      tt.fields.tail_comment,
				anchor:            tt.fields.anchor,
				tag:               tt.fields.tag,
				value:             tt.fields.value,
				implicit:          tt.fields.implicit,
				quoted_implicit:   tt.fields.quoted_implicit,
				style:             tt.fields.style,
			}
			if got := e.mapping_style(); got != tt.want {
				t.Errorf("yaml_event_t.mapping_style() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yaml_parser_state_t_String(t *testing.T) {
	tests := []struct {
		name string
		ps   yaml_parser_state_t
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ps.String(); got != tt.want {
				t.Errorf("yaml_parser_state_t.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
