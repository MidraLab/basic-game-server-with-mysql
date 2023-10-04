package tagparser

import (
	"reflect"
	"testing"

	"github.com/vmihailenco/tagparser/v2/internal/parser"
)

func TestTag_HasOption(t *testing.T) {
	type fields struct {
		Name    string
		Options map[string]string
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
			tr := &Tag{
				Name:    tt.fields.Name,
				Options: tt.fields.Options,
			}
			if got := tr.HasOption(tt.args.name); got != tt.want {
				t.Errorf("Tag.HasOption() = %v, want %v", got, tt.want)
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
		want *Tag
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

func Test_tagParser_setTagOption(t *testing.T) {
	type fields struct {
		Parser  *parser.Parser
		Tag     Tag
		hasName bool
		key     string
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
			p := &tagParser{
				Parser:  tt.fields.Parser,
				Tag:     tt.fields.Tag,
				hasName: tt.fields.hasName,
				key:     tt.fields.key,
			}
			p.setTagOption(tt.args.key, tt.args.value)
		})
	}
}

func Test_tagParser_parseKey(t *testing.T) {
	type fields struct {
		Parser  *parser.Parser
		Tag     Tag
		hasName bool
		key     string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &tagParser{
				Parser:  tt.fields.Parser,
				Tag:     tt.fields.Tag,
				hasName: tt.fields.hasName,
				key:     tt.fields.key,
			}
			p.parseKey()
		})
	}
}

func Test_tagParser_parseValue(t *testing.T) {
	type fields struct {
		Parser  *parser.Parser
		Tag     Tag
		hasName bool
		key     string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &tagParser{
				Parser:  tt.fields.Parser,
				Tag:     tt.fields.Tag,
				hasName: tt.fields.hasName,
				key:     tt.fields.key,
			}
			p.parseValue()
		})
	}
}

func Test_tagParser_readBrackets(t *testing.T) {
	type fields struct {
		Parser  *parser.Parser
		Tag     Tag
		hasName bool
		key     string
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &tagParser{
				Parser:  tt.fields.Parser,
				Tag:     tt.fields.Tag,
				hasName: tt.fields.hasName,
				key:     tt.fields.key,
			}
			if got := p.readBrackets(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagParser.readBrackets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagParser_parseQuotedValue(t *testing.T) {
	type fields struct {
		Parser  *parser.Parser
		Tag     Tag
		hasName bool
		key     string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &tagParser{
				Parser:  tt.fields.Parser,
				Tag:     tt.fields.Tag,
				hasName: tt.fields.hasName,
				key:     tt.fields.key,
			}
			p.parseQuotedValue()
		})
	}
}
