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

// Package yaml implements YAML support for the Go language.
//
// Source code and other details for the project are available at GitHub:
//
//   https://github.com/go-yaml/yaml
//
package yaml

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type args struct {
		in  []byte
		out interface{}
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
			if err := Unmarshal(tt.args.in, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewDecoder(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want *Decoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDecoder(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecoder_KnownFields(t *testing.T) {
	type fields struct {
		parser      *parser
		knownFields bool
	}
	type args struct {
		enable bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec := &Decoder{
				parser:      tt.fields.parser,
				knownFields: tt.fields.knownFields,
			}
			dec.KnownFields(tt.args.enable)
		})
	}
}

func TestDecoder_Decode(t *testing.T) {
	type fields struct {
		parser      *parser
		knownFields bool
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dec := &Decoder{
				parser:      tt.fields.parser,
				knownFields: tt.fields.knownFields,
			}
			if err := dec.Decode(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Decoder.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNode_Decode(t *testing.T) {
	type fields struct {
		Kind        Kind
		Style       Style
		Tag         string
		Value       string
		Anchor      string
		Alias       *Node
		Content     []*Node
		HeadComment string
		LineComment string
		FootComment string
		Line        int
		Column      int
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Kind:        tt.fields.Kind,
				Style:       tt.fields.Style,
				Tag:         tt.fields.Tag,
				Value:       tt.fields.Value,
				Anchor:      tt.fields.Anchor,
				Alias:       tt.fields.Alias,
				Content:     tt.fields.Content,
				HeadComment: tt.fields.HeadComment,
				LineComment: tt.fields.LineComment,
				FootComment: tt.fields.FootComment,
				Line:        tt.fields.Line,
				Column:      tt.fields.Column,
			}
			if err := n.Decode(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Node.Decode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unmarshal(t *testing.T) {
	type args struct {
		in     []byte
		out    interface{}
		strict bool
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
			if err := unmarshal(tt.args.in, tt.args.out, tt.args.strict); (err != nil) != tt.wantErr {
				t.Errorf("unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMarshal(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := Marshal(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Marshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("Marshal() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestNewEncoder(t *testing.T) {
	tests := []struct {
		name  string
		want  *Encoder
		wantW string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if got := NewEncoder(w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEncoder() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("NewEncoder() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestEncoder_Encode(t *testing.T) {
	type fields struct {
		encoder *encoder
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				encoder: tt.fields.encoder,
			}
			if err := e.Encode(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.Encode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNode_Encode(t *testing.T) {
	type fields struct {
		Kind        Kind
		Style       Style
		Tag         string
		Value       string
		Anchor      string
		Alias       *Node
		Content     []*Node
		HeadComment string
		LineComment string
		FootComment string
		Line        int
		Column      int
	}
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Kind:        tt.fields.Kind,
				Style:       tt.fields.Style,
				Tag:         tt.fields.Tag,
				Value:       tt.fields.Value,
				Anchor:      tt.fields.Anchor,
				Alias:       tt.fields.Alias,
				Content:     tt.fields.Content,
				HeadComment: tt.fields.HeadComment,
				LineComment: tt.fields.LineComment,
				FootComment: tt.fields.FootComment,
				Line:        tt.fields.Line,
				Column:      tt.fields.Column,
			}
			if err := n.Encode(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("Node.Encode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEncoder_SetIndent(t *testing.T) {
	type fields struct {
		encoder *encoder
	}
	type args struct {
		spaces int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				encoder: tt.fields.encoder,
			}
			e.SetIndent(tt.args.spaces)
		})
	}
}

func TestEncoder_Close(t *testing.T) {
	type fields struct {
		encoder *encoder
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoder{
				encoder: tt.fields.encoder,
			}
			if err := e.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Encoder.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_handleErr(t *testing.T) {
	type args struct {
		err *error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handleErr(tt.args.err)
		})
	}
}

func Test_fail(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fail(tt.args.err)
		})
	}
}

func Test_failf(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			failf(tt.args.format, tt.args.args...)
		})
	}
}

func TestTypeError_Error(t *testing.T) {
	type fields struct {
		Errors []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &TypeError{
				Errors: tt.fields.Errors,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("TypeError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_IsZero(t *testing.T) {
	type fields struct {
		Kind        Kind
		Style       Style
		Tag         string
		Value       string
		Anchor      string
		Alias       *Node
		Content     []*Node
		HeadComment string
		LineComment string
		FootComment string
		Line        int
		Column      int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Kind:        tt.fields.Kind,
				Style:       tt.fields.Style,
				Tag:         tt.fields.Tag,
				Value:       tt.fields.Value,
				Anchor:      tt.fields.Anchor,
				Alias:       tt.fields.Alias,
				Content:     tt.fields.Content,
				HeadComment: tt.fields.HeadComment,
				LineComment: tt.fields.LineComment,
				FootComment: tt.fields.FootComment,
				Line:        tt.fields.Line,
				Column:      tt.fields.Column,
			}
			if got := n.IsZero(); got != tt.want {
				t.Errorf("Node.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_LongTag(t *testing.T) {
	type fields struct {
		Kind        Kind
		Style       Style
		Tag         string
		Value       string
		Anchor      string
		Alias       *Node
		Content     []*Node
		HeadComment string
		LineComment string
		FootComment string
		Line        int
		Column      int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Kind:        tt.fields.Kind,
				Style:       tt.fields.Style,
				Tag:         tt.fields.Tag,
				Value:       tt.fields.Value,
				Anchor:      tt.fields.Anchor,
				Alias:       tt.fields.Alias,
				Content:     tt.fields.Content,
				HeadComment: tt.fields.HeadComment,
				LineComment: tt.fields.LineComment,
				FootComment: tt.fields.FootComment,
				Line:        tt.fields.Line,
				Column:      tt.fields.Column,
			}
			if got := n.LongTag(); got != tt.want {
				t.Errorf("Node.LongTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_ShortTag(t *testing.T) {
	type fields struct {
		Kind        Kind
		Style       Style
		Tag         string
		Value       string
		Anchor      string
		Alias       *Node
		Content     []*Node
		HeadComment string
		LineComment string
		FootComment string
		Line        int
		Column      int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Kind:        tt.fields.Kind,
				Style:       tt.fields.Style,
				Tag:         tt.fields.Tag,
				Value:       tt.fields.Value,
				Anchor:      tt.fields.Anchor,
				Alias:       tt.fields.Alias,
				Content:     tt.fields.Content,
				HeadComment: tt.fields.HeadComment,
				LineComment: tt.fields.LineComment,
				FootComment: tt.fields.FootComment,
				Line:        tt.fields.Line,
				Column:      tt.fields.Column,
			}
			if got := n.ShortTag(); got != tt.want {
				t.Errorf("Node.ShortTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_indicatedString(t *testing.T) {
	type fields struct {
		Kind        Kind
		Style       Style
		Tag         string
		Value       string
		Anchor      string
		Alias       *Node
		Content     []*Node
		HeadComment string
		LineComment string
		FootComment string
		Line        int
		Column      int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Kind:        tt.fields.Kind,
				Style:       tt.fields.Style,
				Tag:         tt.fields.Tag,
				Value:       tt.fields.Value,
				Anchor:      tt.fields.Anchor,
				Alias:       tt.fields.Alias,
				Content:     tt.fields.Content,
				HeadComment: tt.fields.HeadComment,
				LineComment: tt.fields.LineComment,
				FootComment: tt.fields.FootComment,
				Line:        tt.fields.Line,
				Column:      tt.fields.Column,
			}
			if got := n.indicatedString(); got != tt.want {
				t.Errorf("Node.indicatedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_SetString(t *testing.T) {
	type fields struct {
		Kind        Kind
		Style       Style
		Tag         string
		Value       string
		Anchor      string
		Alias       *Node
		Content     []*Node
		HeadComment string
		LineComment string
		FootComment string
		Line        int
		Column      int
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Node{
				Kind:        tt.fields.Kind,
				Style:       tt.fields.Style,
				Tag:         tt.fields.Tag,
				Value:       tt.fields.Value,
				Anchor:      tt.fields.Anchor,
				Alias:       tt.fields.Alias,
				Content:     tt.fields.Content,
				HeadComment: tt.fields.HeadComment,
				LineComment: tt.fields.LineComment,
				FootComment: tt.fields.FootComment,
				Line:        tt.fields.Line,
				Column:      tt.fields.Column,
			}
			n.SetString(tt.args.s)
		})
	}
}

func Test_getStructInfo(t *testing.T) {
	type args struct {
		st reflect.Type
	}
	tests := []struct {
		name    string
		args    args
		want    *structInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getStructInfo(tt.args.st)
			if (err != nil) != tt.wantErr {
				t.Errorf("getStructInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStructInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isZero(t *testing.T) {
	type args struct {
		v reflect.Value
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
			if got := isZero(tt.args.v); got != tt.want {
				t.Errorf("isZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
