package tagparser

import (
	"reflect"
	"testing"
)

func TestTag_IsZero(t *testing.T) {
	type fields struct {
		Name    string
		Options map[string][]string
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
			tr := Tag{
				Name:    tt.fields.Name,
				Options: tt.fields.Options,
			}
			if got := tr.IsZero(); got != tt.want {
				t.Errorf("Tag.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTag_HasOption(t *testing.T) {
	type fields struct {
		Name    string
		Options map[string][]string
	}
	type args struct {
		name string
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
			tr := Tag{
				Name:    tt.fields.Name,
				Options: tt.fields.Options,
			}
			if got := tr.HasOption(tt.args.name); got != tt.want {
				t.Errorf("Tag.HasOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTag_Option(t *testing.T) {
	type fields struct {
		Name    string
		Options map[string][]string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tag{
				Name:    tt.fields.Name,
				Options: tt.fields.Options,
			}
			got, got1 := tr.Option(tt.args.name)
			if got != tt.want {
				t.Errorf("Tag.Option() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Tag.Option() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Tag
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_setName(t *testing.T) {
	type fields struct {
		s        string
		i        int
		tag      Tag
		seenName bool
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
			p := &parser{
				s:        tt.fields.s,
				i:        tt.fields.i,
				tag:      tt.fields.tag,
				seenName: tt.fields.seenName,
			}
			p.setName(tt.args.name)
		})
	}
}

func Test_parser_addOption(t *testing.T) {
	type fields struct {
		s        string
		i        int
		tag      Tag
		seenName bool
	}
	type args struct {
		key   string
		value string
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
				s:        tt.fields.s,
				i:        tt.fields.i,
				tag:      tt.fields.tag,
				seenName: tt.fields.seenName,
			}
			p.addOption(tt.args.key, tt.args.value)
		})
	}
}

func Test_parser_parse(t *testing.T) {
	type fields struct {
		s        string
		i        int
		tag      Tag
		seenName bool
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
				s:        tt.fields.s,
				i:        tt.fields.i,
				tag:      tt.fields.tag,
				seenName: tt.fields.seenName,
			}
			p.parse()
		})
	}
}

func Test_parser_parseKeyValue(t *testing.T) {
	type fields struct {
		s        string
		i        int
		tag      Tag
		seenName bool
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
				s:        tt.fields.s,
				i:        tt.fields.i,
				tag:      tt.fields.tag,
				seenName: tt.fields.seenName,
			}
			p.parseKeyValue()
		})
	}
}

func Test_parser_parseValue(t *testing.T) {
	type fields struct {
		s        string
		i        int
		tag      Tag
		seenName bool
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
			p := &parser{
				s:        tt.fields.s,
				i:        tt.fields.i,
				tag:      tt.fields.tag,
				seenName: tt.fields.seenName,
			}
			if got := p.parseValue(); got != tt.want {
				t.Errorf("parser.parseValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_parseQuotedValue(t *testing.T) {
	type fields struct {
		s        string
		i        int
		tag      Tag
		seenName bool
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
			p := &parser{
				s:        tt.fields.s,
				i:        tt.fields.i,
				tag:      tt.fields.tag,
				seenName: tt.fields.seenName,
			}
			if got := p.parseQuotedValue(); got != tt.want {
				t.Errorf("parser.parseQuotedValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_skipPairs(t *testing.T) {
	type fields struct {
		s        string
		i        int
		tag      Tag
		seenName bool
	}
	type args struct {
		start byte
		end   byte
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
				s:        tt.fields.s,
				i:        tt.fields.i,
				tag:      tt.fields.tag,
				seenName: tt.fields.seenName,
			}
			p.skipPairs(tt.args.start, tt.args.end)
		})
	}
}

func Test_parser_valid(t *testing.T) {
	type fields struct {
		s        string
		i        int
		tag      Tag
		seenName bool
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
			p := &parser{
				s:        tt.fields.s,
				i:        tt.fields.i,
				tag:      tt.fields.tag,
				seenName: tt.fields.seenName,
			}
			if got := p.valid(); got != tt.want {
				t.Errorf("parser.valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_read(t *testing.T) {
	type fields struct {
		s        string
		i        int
		tag      Tag
		seenName bool
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				s:        tt.fields.s,
				i:        tt.fields.i,
				tag:      tt.fields.tag,
				seenName: tt.fields.seenName,
			}
			if got := p.read(); got != tt.want {
				t.Errorf("parser.read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parser_peek(t *testing.T) {
	type fields struct {
		s        string
		i        int
		tag      Tag
		seenName bool
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				s:        tt.fields.s,
				i:        tt.fields.i,
				tag:      tt.fields.tag,
				seenName: tt.fields.seenName,
			}
			if got := p.peek(); got != tt.want {
				t.Errorf("parser.peek() = %v, want %v", got, tt.want)
			}
		})
	}
}
