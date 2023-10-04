package swag

import (
	"go/ast"
	"reflect"
	"regexp"
	"testing"

	"github.com/go-openapi/spec"
)

func TestNewOperation(t *testing.T) {
	type args struct {
		parser  *Parser
		options []func(*Operation)
	}
	tests := []struct {
		name string
		args args
		want *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOperation(tt.args.parser, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetCodeExampleFilesDirectory(t *testing.T) {
	type args struct {
		directoryPath string
	}
	tests := []struct {
		name string
		args args
		want func(*Operation)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetCodeExampleFilesDirectory(tt.args.directoryPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCodeExampleFilesDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_ParseComment(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		comment string
		astFile *ast.File
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.ParseComment(tt.args.comment, tt.args.astFile); (err != nil) != tt.wantErr {
				t.Errorf("Operation.ParseComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperation_ParseCodeSample(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		attribute     string
		in1           string
		lineRemainder string
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.ParseCodeSample(tt.args.attribute, tt.args.in1, tt.args.lineRemainder); (err != nil) != tt.wantErr {
				t.Errorf("Operation.ParseCodeSample() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperation_ParseDescriptionComment(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		lineRemainder string
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			operation.ParseDescriptionComment(tt.args.lineRemainder)
		})
	}
}

func TestOperation_ParseMetadata(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		attribute      string
		lowerAttribute string
		lineRemainder  string
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.ParseMetadata(tt.args.attribute, tt.args.lowerAttribute, tt.args.lineRemainder); (err != nil) != tt.wantErr {
				t.Errorf("Operation.ParseMetadata() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_findInSlice(t *testing.T) {
	type args struct {
		arr    []string
		target string
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
			if got := findInSlice(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("findInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_ParseParamComment(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		commentLine string
		astFile     *ast.File
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.ParseParamComment(tt.args.commentLine, tt.args.astFile); (err != nil) != tt.wantErr {
				t.Errorf("Operation.ParseParamComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperation_parseParamAttribute(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		comment    string
		objectType string
		schemaType string
		paramType  string
		param      *spec.Parameter
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.parseParamAttribute(tt.args.comment, tt.args.objectType, tt.args.schemaType, tt.args.paramType, tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("Operation.parseParamAttribute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_findAttr(t *testing.T) {
	type args struct {
		re          *regexp.Regexp
		commentLine string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findAttr(tt.args.re, tt.args.commentLine)
			if (err != nil) != tt.wantErr {
				t.Errorf("findAttr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findAttr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setStringParam(t *testing.T) {
	type args struct {
		param       *spec.Parameter
		name        string
		schemaType  string
		attr        string
		commentLine string
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
			if err := setStringParam(tt.args.param, tt.args.name, tt.args.schemaType, tt.args.attr, tt.args.commentLine); (err != nil) != tt.wantErr {
				t.Errorf("setStringParam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setNumberParam(t *testing.T) {
	type args struct {
		param       *spec.Parameter
		name        string
		schemaType  string
		attr        string
		commentLine string
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
			if err := setNumberParam(tt.args.param, tt.args.name, tt.args.schemaType, tt.args.attr, tt.args.commentLine); (err != nil) != tt.wantErr {
				t.Errorf("setNumberParam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setEnumParam(t *testing.T) {
	type args struct {
		param      *spec.Parameter
		attr       string
		objectType string
		schemaType string
		paramType  string
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
			if err := setEnumParam(tt.args.param, tt.args.attr, tt.args.objectType, tt.args.schemaType, tt.args.paramType); (err != nil) != tt.wantErr {
				t.Errorf("setEnumParam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setExtensionParam(t *testing.T) {
	type args struct {
		attr string
	}
	tests := []struct {
		name string
		args args
		want spec.Extensions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setExtensionParam(tt.args.attr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setExtensionParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setCollectionFormatParam(t *testing.T) {
	type args struct {
		param       *spec.Parameter
		name        string
		schemaType  string
		attr        string
		commentLine string
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
			if err := setCollectionFormatParam(tt.args.param, tt.args.name, tt.args.schemaType, tt.args.attr, tt.args.commentLine); (err != nil) != tt.wantErr {
				t.Errorf("setCollectionFormatParam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setDefault(t *testing.T) {
	type args struct {
		param      *spec.Parameter
		schemaType string
		value      string
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
			if err := setDefault(tt.args.param, tt.args.schemaType, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("setDefault() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setSchemaExample(t *testing.T) {
	type args struct {
		param      *spec.Parameter
		schemaType string
		value      string
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
			if err := setSchemaExample(tt.args.param, tt.args.schemaType, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("setSchemaExample() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setExample(t *testing.T) {
	type args struct {
		param      *spec.Parameter
		schemaType string
		value      string
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
			if err := setExample(tt.args.param, tt.args.schemaType, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("setExample() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_defineType(t *testing.T) {
	type args struct {
		schemaType string
		value      string
	}
	tests := []struct {
		name    string
		args    args
		wantV   interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, err := defineType(tt.args.schemaType, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("defineType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("defineType() = %v, want %v", gotV, tt.wantV)
			}
		})
	}
}

func TestOperation_ParseTagsComment(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		commentLine string
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			operation.ParseTagsComment(tt.args.commentLine)
		})
	}
}

func TestOperation_ParseAcceptComment(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		commentLine string
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.ParseAcceptComment(tt.args.commentLine); (err != nil) != tt.wantErr {
				t.Errorf("Operation.ParseAcceptComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperation_ParseProduceComment(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		commentLine string
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.ParseProduceComment(tt.args.commentLine); (err != nil) != tt.wantErr {
				t.Errorf("Operation.ParseProduceComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_parseMimeTypeList(t *testing.T) {
	type args struct {
		mimeTypeList string
		typeList     *[]string
		format       string
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
			if err := parseMimeTypeList(tt.args.mimeTypeList, tt.args.typeList, tt.args.format); (err != nil) != tt.wantErr {
				t.Errorf("parseMimeTypeList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperation_ParseRouterComment(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		commentLine string
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.ParseRouterComment(tt.args.commentLine); (err != nil) != tt.wantErr {
				t.Errorf("Operation.ParseRouterComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperation_ParseSecurityComment(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		commentLine string
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.ParseSecurityComment(tt.args.commentLine); (err != nil) != tt.wantErr {
				t.Errorf("Operation.ParseSecurityComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_findTypeDef(t *testing.T) {
	type args struct {
		importPath string
		typeName   string
	}
	tests := []struct {
		name    string
		args    args
		want    *ast.TypeSpec
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findTypeDef(tt.args.importPath, tt.args.typeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("findTypeDef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findTypeDef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_parseObjectSchema(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		refType string
		astFile *ast.File
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *spec.Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			got, err := operation.parseObjectSchema(tt.args.refType, tt.args.astFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Operation.parseObjectSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.parseObjectSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseObjectSchema(t *testing.T) {
	type args struct {
		parser  *Parser
		refType string
		astFile *ast.File
	}
	tests := []struct {
		name    string
		args    args
		want    *spec.Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseObjectSchema(tt.args.parser, tt.args.refType, tt.args.astFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseObjectSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseObjectSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFields(t *testing.T) {
	type args struct {
		s string
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
			if got := parseFields(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseCombinedObjectSchema(t *testing.T) {
	type args struct {
		parser  *Parser
		refType string
		astFile *ast.File
	}
	tests := []struct {
		name    string
		args    args
		want    *spec.Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCombinedObjectSchema(tt.args.parser, tt.args.refType, tt.args.astFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCombinedObjectSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCombinedObjectSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_parseAPIObjectSchema(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		commentLine string
		schemaType  string
		refType     string
		astFile     *ast.File
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *spec.Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			got, err := operation.parseAPIObjectSchema(tt.args.commentLine, tt.args.schemaType, tt.args.refType, tt.args.astFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("Operation.parseAPIObjectSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.parseAPIObjectSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_ParseResponseComment(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		commentLine string
		astFile     *ast.File
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.ParseResponseComment(tt.args.commentLine, tt.args.astFile); (err != nil) != tt.wantErr {
				t.Errorf("Operation.ParseResponseComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newHeaderSpec(t *testing.T) {
	type args struct {
		schemaType  string
		description string
	}
	tests := []struct {
		name string
		args args
		want spec.Header
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newHeaderSpec(tt.args.schemaType, tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newHeaderSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_ParseResponseHeaderComment(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		commentLine string
		in1         *ast.File
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.ParseResponseHeaderComment(tt.args.commentLine, tt.args.in1); (err != nil) != tt.wantErr {
				t.Errorf("Operation.ParseResponseHeaderComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperation_ParseEmptyResponseComment(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		commentLine string
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.ParseEmptyResponseComment(tt.args.commentLine); (err != nil) != tt.wantErr {
				t.Errorf("Operation.ParseEmptyResponseComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperation_ParseEmptyResponseOnly(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		commentLine string
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if err := operation.ParseEmptyResponseOnly(tt.args.commentLine); (err != nil) != tt.wantErr {
				t.Errorf("Operation.ParseEmptyResponseOnly() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperation_DefaultResponse(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	tests := []struct {
		name   string
		fields fields
		want   *spec.Response
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			if got := operation.DefaultResponse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.DefaultResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_AddResponse(t *testing.T) {
	type fields struct {
		parser              *Parser
		codeExampleFilesDir string
		Operation           spec.Operation
		RouterProperties    []RouteProperties
	}
	type args struct {
		code     int
		response *spec.Response
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
			operation := &Operation{
				parser:              tt.fields.parser,
				codeExampleFilesDir: tt.fields.codeExampleFilesDir,
				Operation:           tt.fields.Operation,
				RouterProperties:    tt.fields.RouterProperties,
			}
			operation.AddResponse(tt.args.code, tt.args.response)
		})
	}
}

func Test_createParameter(t *testing.T) {
	type args struct {
		paramType        string
		description      string
		paramName        string
		objectType       string
		schemaType       string
		required         bool
		enums            []interface{}
		collectionFormat string
	}
	tests := []struct {
		name string
		args args
		want spec.Parameter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createParameter(tt.args.paramType, tt.args.description, tt.args.paramName, tt.args.objectType, tt.args.schemaType, tt.args.required, tt.args.enums, tt.args.collectionFormat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCodeExampleForSummary(t *testing.T) {
	type args struct {
		summaryName string
		dirPath     string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCodeExampleForSummary(tt.args.summaryName, tt.args.dirPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCodeExampleForSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCodeExampleForSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}
