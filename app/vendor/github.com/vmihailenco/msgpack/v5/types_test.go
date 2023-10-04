package msgpack

import (
	"reflect"
	"sync"
	"testing"
)

func TestRegister(t *testing.T) {
	type args struct {
		value interface{}
		enc   encoderFunc
		dec   decoderFunc
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register(tt.args.value, tt.args.enc, tt.args.dec)
		})
	}
}

func Test_newStructCache(t *testing.T) {
	tests := []struct {
		name string
		want *structCache
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newStructCache(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStructCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structCache_Fields(t *testing.T) {
	type fields struct {
		m sync.Map
	}
	type args struct {
		typ reflect.Type
		tag string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structCache{
				m: tt.fields.m,
			}
			if got := m.Fields(tt.args.typ, tt.args.tag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structCache.Fields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_field_Omit(t *testing.T) {
	type fields struct {
		name      string
		index     []int
		omitEmpty bool
		encoder   encoderFunc
		decoder   decoderFunc
	}
	type args struct {
		strct  reflect.Value
		forced bool
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
			f := &field{
				name:      tt.fields.name,
				index:     tt.fields.index,
				omitEmpty: tt.fields.omitEmpty,
				encoder:   tt.fields.encoder,
				decoder:   tt.fields.decoder,
			}
			if got := f.Omit(tt.args.strct, tt.args.forced); got != tt.want {
				t.Errorf("field.Omit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_field_EncodeValue(t *testing.T) {
	type fields struct {
		name      string
		index     []int
		omitEmpty bool
		encoder   encoderFunc
		decoder   decoderFunc
	}
	type args struct {
		e     *Encoder
		strct reflect.Value
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
			f := &field{
				name:      tt.fields.name,
				index:     tt.fields.index,
				omitEmpty: tt.fields.omitEmpty,
				encoder:   tt.fields.encoder,
				decoder:   tt.fields.decoder,
			}
			if err := f.EncodeValue(tt.args.e, tt.args.strct); (err != nil) != tt.wantErr {
				t.Errorf("field.EncodeValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_field_DecodeValue(t *testing.T) {
	type fields struct {
		name      string
		index     []int
		omitEmpty bool
		encoder   encoderFunc
		decoder   decoderFunc
	}
	type args struct {
		d     *Decoder
		strct reflect.Value
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
			f := &field{
				name:      tt.fields.name,
				index:     tt.fields.index,
				omitEmpty: tt.fields.omitEmpty,
				encoder:   tt.fields.encoder,
				decoder:   tt.fields.decoder,
			}
			if err := f.DecodeValue(tt.args.d, tt.args.strct); (err != nil) != tt.wantErr {
				t.Errorf("field.DecodeValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newFields(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name string
		args args
		want *fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newFields(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fields_Add(t *testing.T) {
	type fields struct {
		Type         reflect.Type
		Map          map[string]*field
		List         []*field
		AsArray      bool
		hasOmitEmpty bool
	}
	type args struct {
		field *field
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
			fs := &fields{
				Type:         tt.fields.Type,
				Map:          tt.fields.Map,
				List:         tt.fields.List,
				AsArray:      tt.fields.AsArray,
				hasOmitEmpty: tt.fields.hasOmitEmpty,
			}
			fs.Add(tt.args.field)
		})
	}
}

func Test_fields_warnIfFieldExists(t *testing.T) {
	type fields struct {
		Type         reflect.Type
		Map          map[string]*field
		List         []*field
		AsArray      bool
		hasOmitEmpty bool
	}
	type args struct {
		name string
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
			fs := &fields{
				Type:         tt.fields.Type,
				Map:          tt.fields.Map,
				List:         tt.fields.List,
				AsArray:      tt.fields.AsArray,
				hasOmitEmpty: tt.fields.hasOmitEmpty,
			}
			fs.warnIfFieldExists(tt.args.name)
		})
	}
}

func Test_fields_OmitEmpty(t *testing.T) {
	type fields struct {
		Type         reflect.Type
		Map          map[string]*field
		List         []*field
		AsArray      bool
		hasOmitEmpty bool
	}
	type args struct {
		strct  reflect.Value
		forced bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*field
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &fields{
				Type:         tt.fields.Type,
				Map:          tt.fields.Map,
				List:         tt.fields.List,
				AsArray:      tt.fields.AsArray,
				hasOmitEmpty: tt.fields.hasOmitEmpty,
			}
			if got := fs.OmitEmpty(tt.args.strct, tt.args.forced); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fields.OmitEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFields(t *testing.T) {
	type args struct {
		typ         reflect.Type
		fallbackTag string
	}
	tests := []struct {
		name string
		args args
		want *fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFields(tt.args.typ, tt.args.fallbackTag); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inlineFields(t *testing.T) {
	type args struct {
		fs  *fields
		typ reflect.Type
		f   *field
		tag string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inlineFields(tt.args.fs, tt.args.typ, tt.args.f, tt.args.tag)
		})
	}
}

func Test_shouldInline(t *testing.T) {
	type args struct {
		fs  *fields
		typ reflect.Type
		f   *field
		tag string
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
			if got := shouldInline(tt.args.fs, tt.args.typ, tt.args.f, tt.args.tag); got != tt.want {
				t.Errorf("shouldInline() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEmptyValue(t *testing.T) {
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
			if got := isEmptyValue(tt.args.v); got != tt.want {
				t.Errorf("isEmptyValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fieldByIndex(t *testing.T) {
	type args struct {
		v     reflect.Value
		index []int
	}
	tests := []struct {
		name   string
		args   args
		want   reflect.Value
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOk := fieldByIndex(tt.args.v, tt.args.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldByIndex() got = %v, want %v", got, tt.want)
			}
			if gotOk != tt.wantOk {
				t.Errorf("fieldByIndex() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_fieldByIndexAlloc(t *testing.T) {
	type args struct {
		v     reflect.Value
		index []int
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
			if got := fieldByIndexAlloc(tt.args.v, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldByIndexAlloc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indirectNil(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name  string
		args  args
		want  reflect.Value
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := indirectNil(tt.args.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indirectNil() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("indirectNil() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
