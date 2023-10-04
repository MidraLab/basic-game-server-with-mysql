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
	"io"
	"reflect"
	"testing"
)

func Test_newParser(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want *parser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newParser(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newParserFromReader(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want *parser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newParserFromReader(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newParserFromReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_init(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			p.init()
		})
	}
}

func Test_parser_destroy(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			p.destroy()
		})
	}
}

func Test_parser_expect(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	type args struct {
		e yaml_event_type_t
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
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			p.expect(tt.args.e)
		})
	}
}

func Test_parser_peek(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	tests := []struct {
		name   string
		fields fields
		want   yaml_event_type_t
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			if got := p.peek(); got != tt.want {
				t.Errorf("parser.peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_fail(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			p.fail()
		})
	}
}

func Test_parser_anchor(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	type args struct {
		n      *Node
		anchor []byte
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
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			p.anchor(tt.args.n, tt.args.anchor)
		})
	}
}

func Test_parser_parse(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			if got := p.parse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser.parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_node(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	type args struct {
		kind       Kind
		defaultTag string
		tag        string
		value      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			if got := p.node(tt.args.kind, tt.args.defaultTag, tt.args.tag, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser.node() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_parseChild(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	type args struct {
		parent *Node
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			if got := p.parseChild(tt.args.parent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser.parseChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_document(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			if got := p.document(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser.document() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_alias(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			if got := p.alias(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser.alias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_scalar(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			if got := p.scalar(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser.scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_sequence(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			if got := p.sequence(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser.sequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_mapping(t *testing.T) {
	type fields struct {
		parser   yaml_parser_t
		event    yaml_event_t
		doc      *Node
		anchors  map[string]*Node
		doneInit bool
		textless bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *Node
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				parser:   tt.fields.parser,
				event:    tt.fields.event,
				doc:      tt.fields.doc,
				anchors:  tt.fields.anchors,
				doneInit: tt.fields.doneInit,
				textless: tt.fields.textless,
			}
			if got := p.mapping(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser.mapping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newDecoder(t *testing.T) {
	tests := []struct {
		name string
		want *decoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDecoder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decoder_terror(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		n   *Node
		tag string
		out reflect.Value
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
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			d.terror(tt.args.n, tt.args.tag, tt.args.out)
		})
	}
}

func Test_decoder_callUnmarshaler(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		n *Node
		u Unmarshaler
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantGood bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			if gotGood := d.callUnmarshaler(tt.args.n, tt.args.u); gotGood != tt.wantGood {
				t.Errorf("decoder.callUnmarshaler() = %v, want %v", gotGood, tt.wantGood)
			}
		})
	}
}

func Test_decoder_callObsoleteUnmarshaler(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		n *Node
		u obsoleteUnmarshaler
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantGood bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			if gotGood := d.callObsoleteUnmarshaler(tt.args.n, tt.args.u); gotGood != tt.wantGood {
				t.Errorf("decoder.callObsoleteUnmarshaler() = %v, want %v", gotGood, tt.wantGood)
			}
		})
	}
}

func Test_decoder_prepare(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		n   *Node
		out reflect.Value
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		wantNewout      reflect.Value
		wantUnmarshaled bool
		wantGood        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			gotNewout, gotUnmarshaled, gotGood := d.prepare(tt.args.n, tt.args.out)
			if !reflect.DeepEqual(gotNewout, tt.wantNewout) {
				t.Errorf("decoder.prepare() gotNewout = %v, want %v", gotNewout, tt.wantNewout)
			}
			if gotUnmarshaled != tt.wantUnmarshaled {
				t.Errorf("decoder.prepare() gotUnmarshaled = %v, want %v", gotUnmarshaled, tt.wantUnmarshaled)
			}
			if gotGood != tt.wantGood {
				t.Errorf("decoder.prepare() gotGood = %v, want %v", gotGood, tt.wantGood)
			}
		})
	}
}

func Test_decoder_fieldByIndex(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		n     *Node
		v     reflect.Value
		index []int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantField reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			if gotField := d.fieldByIndex(tt.args.n, tt.args.v, tt.args.index); !reflect.DeepEqual(gotField, tt.wantField) {
				t.Errorf("decoder.fieldByIndex() = %v, want %v", gotField, tt.wantField)
			}
		})
	}
}

func Test_allowedAliasRatio(t *testing.T) {
	type args struct {
		decodeCount int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allowedAliasRatio(tt.args.decodeCount); got != tt.want {
				t.Errorf("allowedAliasRatio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decoder_unmarshal(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		n   *Node
		out reflect.Value
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantGood bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			if gotGood := d.unmarshal(tt.args.n, tt.args.out); gotGood != tt.wantGood {
				t.Errorf("decoder.unmarshal() = %v, want %v", gotGood, tt.wantGood)
			}
		})
	}
}

func Test_decoder_document(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		n   *Node
		out reflect.Value
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantGood bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			if gotGood := d.document(tt.args.n, tt.args.out); gotGood != tt.wantGood {
				t.Errorf("decoder.document() = %v, want %v", gotGood, tt.wantGood)
			}
		})
	}
}

func Test_decoder_alias(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		n   *Node
		out reflect.Value
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantGood bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			if gotGood := d.alias(tt.args.n, tt.args.out); gotGood != tt.wantGood {
				t.Errorf("decoder.alias() = %v, want %v", gotGood, tt.wantGood)
			}
		})
	}
}

func Test_resetMap(t *testing.T) {
	type args struct {
		out reflect.Value
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetMap(tt.args.out)
		})
	}
}

func Test_decoder_null(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		out reflect.Value
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
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			if got := d.null(tt.args.out); got != tt.want {
				t.Errorf("decoder.null() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decoder_scalar(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		n   *Node
		out reflect.Value
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
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			if got := d.scalar(tt.args.n, tt.args.out); got != tt.want {
				t.Errorf("decoder.scalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_settableValueOf(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := settableValueOf(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("settableValueOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decoder_sequence(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		n   *Node
		out reflect.Value
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantGood bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			if gotGood := d.sequence(tt.args.n, tt.args.out); gotGood != tt.wantGood {
				t.Errorf("decoder.sequence() = %v, want %v", gotGood, tt.wantGood)
			}
		})
	}
}

func Test_decoder_mapping(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		n   *Node
		out reflect.Value
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantGood bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			if gotGood := d.mapping(tt.args.n, tt.args.out); gotGood != tt.wantGood {
				t.Errorf("decoder.mapping() = %v, want %v", gotGood, tt.wantGood)
			}
		})
	}
}

func Test_isStringMap(t *testing.T) {
	type args struct {
		n *Node
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
			if got := isStringMap(tt.args.n); got != tt.want {
				t.Errorf("isStringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decoder_mappingStruct(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		n   *Node
		out reflect.Value
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantGood bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			if gotGood := d.mappingStruct(tt.args.n, tt.args.out); gotGood != tt.wantGood {
				t.Errorf("decoder.mappingStruct() = %v, want %v", gotGood, tt.wantGood)
			}
		})
	}
}

func Test_failWantMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			failWantMap()
		})
	}
}

func Test_decoder_merge(t *testing.T) {
	type fields struct {
		doc            *Node
		aliases        map[*Node]bool
		terrors        []string
		stringMapType  reflect.Type
		generalMapType reflect.Type
		knownFields    bool
		uniqueKeys     bool
		decodeCount    int
		aliasCount     int
		aliasDepth     int
		mergedFields   map[interface{}]bool
	}
	type args struct {
		parent *Node
		merge  *Node
		out    reflect.Value
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
			d := &decoder{
				doc:            tt.fields.doc,
				aliases:        tt.fields.aliases,
				terrors:        tt.fields.terrors,
				stringMapType:  tt.fields.stringMapType,
				generalMapType: tt.fields.generalMapType,
				knownFields:    tt.fields.knownFields,
				uniqueKeys:     tt.fields.uniqueKeys,
				decodeCount:    tt.fields.decodeCount,
				aliasCount:     tt.fields.aliasCount,
				aliasDepth:     tt.fields.aliasDepth,
				mergedFields:   tt.fields.mergedFields,
			}
			d.merge(tt.args.parent, tt.args.merge, tt.args.out)
		})
	}
}

func Test_isMerge(t *testing.T) {
	type args struct {
		n *Node
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
			if got := isMerge(tt.args.n); got != tt.want {
				t.Errorf("isMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}
