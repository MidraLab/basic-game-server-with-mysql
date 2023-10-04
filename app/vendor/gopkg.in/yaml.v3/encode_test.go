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
	"bytes"
	"reflect"
	"testing"
)

func Test_newEncoder(t *testing.T) {
	tests := []struct {
		name string
		want *encoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newEncoder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newEncoderWithWriter(t *testing.T) {
	tests := []struct {
		name  string
		want  *encoder
		wantW string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if got := newEncoderWithWriter(w); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newEncoderWithWriter() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("newEncoderWithWriter() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_encoder_init(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.init()
		})
	}
}

func Test_encoder_finish(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.finish()
		})
	}
}

func Test_encoder_destroy(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.destroy()
		})
	}
}

func Test_encoder_emit(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.emit()
		})
	}
}

func Test_encoder_must(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		ok bool
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.must(tt.args.ok)
		})
	}
}

func Test_encoder_marshalDoc(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		tag string
		in  reflect.Value
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.marshalDoc(tt.args.tag, tt.args.in)
		})
	}
}

func Test_encoder_marshal(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		tag string
		in  reflect.Value
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.marshal(tt.args.tag, tt.args.in)
		})
	}
}

func Test_encoder_mapv(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		tag string
		in  reflect.Value
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.mapv(tt.args.tag, tt.args.in)
		})
	}
}

func Test_encoder_fieldByIndex(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			if gotField := e.fieldByIndex(tt.args.v, tt.args.index); !reflect.DeepEqual(gotField, tt.wantField) {
				t.Errorf("encoder.fieldByIndex() = %v, want %v", gotField, tt.wantField)
			}
		})
	}
}

func Test_encoder_structv(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		tag string
		in  reflect.Value
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.structv(tt.args.tag, tt.args.in)
		})
	}
}

func Test_encoder_mappingv(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		tag string
		f   func()
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.mappingv(tt.args.tag, tt.args.f)
		})
	}
}

func Test_encoder_slicev(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		tag string
		in  reflect.Value
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.slicev(tt.args.tag, tt.args.in)
		})
	}
}

func Test_isBase60Float(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := isBase60Float(tt.args.s); gotResult != tt.wantResult {
				t.Errorf("isBase60Float() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_isOldBool(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := isOldBool(tt.args.s); gotResult != tt.wantResult {
				t.Errorf("isOldBool() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_encoder_stringv(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		tag string
		in  reflect.Value
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.stringv(tt.args.tag, tt.args.in)
		})
	}
}

func Test_encoder_boolv(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		tag string
		in  reflect.Value
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.boolv(tt.args.tag, tt.args.in)
		})
	}
}

func Test_encoder_intv(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		tag string
		in  reflect.Value
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.intv(tt.args.tag, tt.args.in)
		})
	}
}

func Test_encoder_uintv(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		tag string
		in  reflect.Value
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.uintv(tt.args.tag, tt.args.in)
		})
	}
}

func Test_encoder_timev(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		tag string
		in  reflect.Value
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.timev(tt.args.tag, tt.args.in)
		})
	}
}

func Test_encoder_floatv(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		tag string
		in  reflect.Value
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.floatv(tt.args.tag, tt.args.in)
		})
	}
}

func Test_encoder_nilv(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.nilv()
		})
	}
}

func Test_encoder_emitScalar(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		value  string
		anchor string
		tag    string
		style  yaml_scalar_style_t
		head   []byte
		line   []byte
		foot   []byte
		tail   []byte
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.emitScalar(tt.args.value, tt.args.anchor, tt.args.tag, tt.args.style, tt.args.head, tt.args.line, tt.args.foot, tt.args.tail)
		})
	}
}

func Test_encoder_nodev(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		in reflect.Value
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.nodev(tt.args.in)
		})
	}
}

func Test_encoder_node(t *testing.T) {
	type fields struct {
		emitter  yaml_emitter_t
		event    yaml_event_t
		out      []byte
		flow     bool
		indent   int
		doneInit bool
	}
	type args struct {
		node *Node
		tail string
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
			e := &encoder{
				emitter:  tt.fields.emitter,
				event:    tt.fields.event,
				out:      tt.fields.out,
				flow:     tt.fields.flow,
				indent:   tt.fields.indent,
				doneInit: tt.fields.doneInit,
			}
			e.node(tt.args.node, tt.args.tail)
		})
	}
}
