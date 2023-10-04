package swag

import (
	"go/ast"
	"reflect"
	"testing"

	"github.com/go-openapi/spec"
)

func Test_newTagBaseFieldParser(t *testing.T) {
	type args struct {
		p     *Parser
		field *ast.Field
	}
	tests := []struct {
		name string
		args args
		want FieldParser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTagBaseFieldParser(tt.args.p, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTagBaseFieldParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagBaseFieldParser_ShouldSkip(t *testing.T) {
	type fields struct {
		p     *Parser
		field *ast.Field
		tag   reflect.StructTag
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
			ps := &tagBaseFieldParser{
				p:     tt.fields.p,
				field: tt.fields.field,
				tag:   tt.fields.tag,
			}
			if got := ps.ShouldSkip(); got != tt.want {
				t.Errorf("tagBaseFieldParser.ShouldSkip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagBaseFieldParser_FieldName(t *testing.T) {
	type fields struct {
		p     *Parser
		field *ast.Field
		tag   reflect.StructTag
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &tagBaseFieldParser{
				p:     tt.fields.p,
				field: tt.fields.field,
				tag:   tt.fields.tag,
			}
			got, err := ps.FieldName()
			if (err != nil) != tt.wantErr {
				t.Errorf("tagBaseFieldParser.FieldName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tagBaseFieldParser.FieldName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagBaseFieldParser_FormName(t *testing.T) {
	type fields struct {
		p     *Parser
		field *ast.Field
		tag   reflect.StructTag
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
			ps := &tagBaseFieldParser{
				p:     tt.fields.p,
				field: tt.fields.field,
				tag:   tt.fields.tag,
			}
			if got := ps.FormName(); got != tt.want {
				t.Errorf("tagBaseFieldParser.FormName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toSnakeCase(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toSnakeCase(tt.args.in); got != tt.want {
				t.Errorf("toSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toLowerCamelCase(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCamelCase(tt.args.in); got != tt.want {
				t.Errorf("toLowerCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagBaseFieldParser_CustomSchema(t *testing.T) {
	type fields struct {
		p     *Parser
		field *ast.Field
		tag   reflect.StructTag
	}
	tests := []struct {
		name    string
		fields  fields
		want    *spec.Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &tagBaseFieldParser{
				p:     tt.fields.p,
				field: tt.fields.field,
				tag:   tt.fields.tag,
			}
			got, err := ps.CustomSchema()
			if (err != nil) != tt.wantErr {
				t.Errorf("tagBaseFieldParser.CustomSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tagBaseFieldParser.CustomSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitNotWrapped(t *testing.T) {
	type args struct {
		s   string
		sep rune
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitNotWrapped(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitNotWrapped() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagBaseFieldParser_ComplementSchema(t *testing.T) {
	type fields struct {
		p     *Parser
		field *ast.Field
		tag   reflect.StructTag
	}
	type args struct {
		schema *spec.Schema
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
			ps := &tagBaseFieldParser{
				p:     tt.fields.p,
				field: tt.fields.field,
				tag:   tt.fields.tag,
			}
			if err := ps.ComplementSchema(tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("tagBaseFieldParser.ComplementSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_tagBaseFieldParser_complementSchema(t *testing.T) {
	type fields struct {
		p     *Parser
		field *ast.Field
		tag   reflect.StructTag
	}
	type args struct {
		schema *spec.Schema
		types  []string
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
			ps := &tagBaseFieldParser{
				p:     tt.fields.p,
				field: tt.fields.field,
				tag:   tt.fields.tag,
			}
			if err := ps.complementSchema(tt.args.schema, tt.args.types); (err != nil) != tt.wantErr {
				t.Errorf("tagBaseFieldParser.complementSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getFloatTag(t *testing.T) {
	type args struct {
		structTag reflect.StructTag
		tagName   string
	}
	tests := []struct {
		name    string
		args    args
		want    *float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFloatTag(tt.args.structTag, tt.args.tagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFloatTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getFloatTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIntTag(t *testing.T) {
	type args struct {
		structTag reflect.StructTag
		tagName   string
	}
	tests := []struct {
		name    string
		args    args
		want    *int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIntTag(tt.args.structTag, tt.args.tagName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIntTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getIntTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tagBaseFieldParser_IsRequired(t *testing.T) {
	type fields struct {
		p     *Parser
		field *ast.Field
		tag   reflect.StructTag
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := &tagBaseFieldParser{
				p:     tt.fields.p,
				field: tt.fields.field,
				tag:   tt.fields.tag,
			}
			got, err := ps.IsRequired()
			if (err != nil) != tt.wantErr {
				t.Errorf("tagBaseFieldParser.IsRequired() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("tagBaseFieldParser.IsRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseValidTags(t *testing.T) {
	type args struct {
		validTag string
		sf       *structField
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseValidTags(tt.args.validTag, tt.args.sf)
		})
	}
}

func Test_parseEnumTags(t *testing.T) {
	type args struct {
		enumTag string
		field   *structField
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
			if err := parseEnumTags(tt.args.enumTag, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("parseEnumTags() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_structField_setOneOf(t *testing.T) {
	type fields struct {
		schemaType   string
		arrayType    string
		formatType   string
		maximum      *float64
		minimum      *float64
		multipleOf   *float64
		maxLength    *int64
		minLength    *int64
		maxItems     *int64
		minItems     *int64
		exampleValue interface{}
		enums        []interface{}
		enumVarNames []interface{}
		unique       bool
	}
	type args struct {
		valValue string
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
			sf := &structField{
				schemaType:   tt.fields.schemaType,
				arrayType:    tt.fields.arrayType,
				formatType:   tt.fields.formatType,
				maximum:      tt.fields.maximum,
				minimum:      tt.fields.minimum,
				multipleOf:   tt.fields.multipleOf,
				maxLength:    tt.fields.maxLength,
				minLength:    tt.fields.minLength,
				maxItems:     tt.fields.maxItems,
				minItems:     tt.fields.minItems,
				exampleValue: tt.fields.exampleValue,
				enums:        tt.fields.enums,
				enumVarNames: tt.fields.enumVarNames,
				unique:       tt.fields.unique,
			}
			sf.setOneOf(tt.args.valValue)
		})
	}
}

func Test_structField_setMin(t *testing.T) {
	type fields struct {
		schemaType   string
		arrayType    string
		formatType   string
		maximum      *float64
		minimum      *float64
		multipleOf   *float64
		maxLength    *int64
		minLength    *int64
		maxItems     *int64
		minItems     *int64
		exampleValue interface{}
		enums        []interface{}
		enumVarNames []interface{}
		unique       bool
	}
	type args struct {
		valValue string
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
			sf := &structField{
				schemaType:   tt.fields.schemaType,
				arrayType:    tt.fields.arrayType,
				formatType:   tt.fields.formatType,
				maximum:      tt.fields.maximum,
				minimum:      tt.fields.minimum,
				multipleOf:   tt.fields.multipleOf,
				maxLength:    tt.fields.maxLength,
				minLength:    tt.fields.minLength,
				maxItems:     tt.fields.maxItems,
				minItems:     tt.fields.minItems,
				exampleValue: tt.fields.exampleValue,
				enums:        tt.fields.enums,
				enumVarNames: tt.fields.enumVarNames,
				unique:       tt.fields.unique,
			}
			sf.setMin(tt.args.valValue)
		})
	}
}

func Test_structField_setMax(t *testing.T) {
	type fields struct {
		schemaType   string
		arrayType    string
		formatType   string
		maximum      *float64
		minimum      *float64
		multipleOf   *float64
		maxLength    *int64
		minLength    *int64
		maxItems     *int64
		minItems     *int64
		exampleValue interface{}
		enums        []interface{}
		enumVarNames []interface{}
		unique       bool
	}
	type args struct {
		valValue string
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
			sf := &structField{
				schemaType:   tt.fields.schemaType,
				arrayType:    tt.fields.arrayType,
				formatType:   tt.fields.formatType,
				maximum:      tt.fields.maximum,
				minimum:      tt.fields.minimum,
				multipleOf:   tt.fields.multipleOf,
				maxLength:    tt.fields.maxLength,
				minLength:    tt.fields.minLength,
				maxItems:     tt.fields.maxItems,
				minItems:     tt.fields.minItems,
				exampleValue: tt.fields.exampleValue,
				enums:        tt.fields.enums,
				enumVarNames: tt.fields.enumVarNames,
				unique:       tt.fields.unique,
			}
			sf.setMax(tt.args.valValue)
		})
	}
}

func Test_parseOneOfParam2(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseOneOfParam2(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseOneOfParam2() = %v, want %v", got, tt.want)
			}
		})
	}
}
