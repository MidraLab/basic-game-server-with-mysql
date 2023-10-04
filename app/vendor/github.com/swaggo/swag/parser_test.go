package swag

import (
	"go/ast"
	"os"
	"reflect"
	"testing"

	"github.com/KyleBanks/depth"
	"github.com/go-openapi/spec"
)

func TestNew(t *testing.T) {
	type args struct {
		options []func(*Parser)
	}
	tests := []struct {
		name string
		args args
		want *Parser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetParseDependency(t *testing.T) {
	type args struct {
		parseDependency int
	}
	tests := []struct {
		name string
		args args
		want func(*Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetParseDependency(tt.args.parseDependency); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetParseDependency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetMarkdownFileDirectory(t *testing.T) {
	type args struct {
		directoryPath string
	}
	tests := []struct {
		name string
		args args
		want func(*Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetMarkdownFileDirectory(tt.args.directoryPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetMarkdownFileDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetCodeExamplesDirectory(t *testing.T) {
	type args struct {
		directoryPath string
	}
	tests := []struct {
		name string
		args args
		want func(*Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetCodeExamplesDirectory(tt.args.directoryPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCodeExamplesDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetExcludedDirsAndFiles(t *testing.T) {
	type args struct {
		excludes string
	}
	tests := []struct {
		name string
		args args
		want func(*Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetExcludedDirsAndFiles(tt.args.excludes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetExcludedDirsAndFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetPackagePrefix(t *testing.T) {
	type args struct {
		packagePrefix string
	}
	tests := []struct {
		name string
		args args
		want func(*Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetPackagePrefix(tt.args.packagePrefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetPackagePrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetTags(t *testing.T) {
	type args struct {
		include string
	}
	tests := []struct {
		name string
		args args
		want func(*Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetTags(tt.args.include); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetParseExtension(t *testing.T) {
	type args struct {
		parseExtension string
	}
	tests := []struct {
		name string
		args args
		want func(*Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetParseExtension(tt.args.parseExtension); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetParseExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetStrict(t *testing.T) {
	type args struct {
		strict bool
	}
	tests := []struct {
		name string
		args args
		want func(*Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetStrict(tt.args.strict); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStrict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetDebugger(t *testing.T) {
	type args struct {
		logger Debugger
	}
	tests := []struct {
		name string
		args args
		want func(parser *Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetDebugger(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDebugger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetFieldParserFactory(t *testing.T) {
	type args struct {
		factory FieldParserFactory
	}
	tests := []struct {
		name string
		args args
		want func(parser *Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetFieldParserFactory(tt.args.factory); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFieldParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetOverrides(t *testing.T) {
	type args struct {
		overrides map[string]string
	}
	tests := []struct {
		name string
		args args
		want func(parser *Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetOverrides(tt.args.overrides); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetOverrides() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetCollectionFormat(t *testing.T) {
	type args struct {
		collectionFormat string
	}
	tests := []struct {
		name string
		args args
		want func(*Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetCollectionFormat(tt.args.collectionFormat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCollectionFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseUsingGoList(t *testing.T) {
	type args struct {
		enabled bool
	}
	tests := []struct {
		name string
		args args
		want func(parser *Parser)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseUsingGoList(tt.args.enabled); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseUsingGoList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ParseAPI(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		searchDir   string
		mainAPIFile string
		parseDepth  int
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if err := parser.ParseAPI(tt.args.searchDir, tt.args.mainAPIFile, tt.args.parseDepth); (err != nil) != tt.wantErr {
				t.Errorf("Parser.ParseAPI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_skipPackageByPrefix(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		pkgpath string
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if got := parser.skipPackageByPrefix(tt.args.pkgpath); got != tt.want {
				t.Errorf("Parser.skipPackageByPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ParseAPIMultiSearchDir(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		searchDirs  []string
		mainAPIFile string
		parseDepth  int
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if err := parser.ParseAPIMultiSearchDir(tt.args.searchDirs, tt.args.mainAPIFile, tt.args.parseDepth); (err != nil) != tt.wantErr {
				t.Errorf("Parser.ParseAPIMultiSearchDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getPkgName(t *testing.T) {
	type args struct {
		searchDir string
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
			got, err := getPkgName(tt.args.searchDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPkgName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getPkgName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ParseGeneralAPIInfo(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		mainAPIFile string
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if err := parser.ParseGeneralAPIInfo(tt.args.mainAPIFile); (err != nil) != tt.wantErr {
				t.Errorf("Parser.ParseGeneralAPIInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_parseGeneralAPIInfo(t *testing.T) {
	type args struct {
		parser   *Parser
		comments []string
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
			if err := parseGeneralAPIInfo(tt.args.parser, tt.args.comments); (err != nil) != tt.wantErr {
				t.Errorf("parseGeneralAPIInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_setSwaggerInfo(t *testing.T) {
	type args struct {
		swagger   *spec.Swagger
		attribute string
		value     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setSwaggerInfo(tt.args.swagger, tt.args.attribute, tt.args.value)
		})
	}
}

func Test_parseSecAttributes(t *testing.T) {
	type args struct {
		context string
		lines   []string
		index   *int
	}
	tests := []struct {
		name    string
		args    args
		want    *spec.SecurityScheme
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSecAttributes(tt.args.context, tt.args.lines, tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSecAttributes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSecAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseSecurity(t *testing.T) {
	type args struct {
		commentLine string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseSecurity(tt.args.commentLine); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSecurity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initIfEmpty(t *testing.T) {
	type args struct {
		license *spec.License
	}
	tests := []struct {
		name string
		args args
		want *spec.License
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initIfEmpty(tt.args.license); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initIfEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ParseAcceptComment(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if err := parser.ParseAcceptComment(tt.args.commentLine); (err != nil) != tt.wantErr {
				t.Errorf("Parser.ParseAcceptComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_ParseProduceComment(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if err := parser.ParseProduceComment(tt.args.commentLine); (err != nil) != tt.wantErr {
				t.Errorf("Parser.ParseProduceComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isGeneralAPIComment(t *testing.T) {
	type args struct {
		comments []string
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
			if got := isGeneralAPIComment(tt.args.comments); got != tt.want {
				t.Errorf("isGeneralAPIComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMarkdownForTag(t *testing.T) {
	type args struct {
		tagName string
		dirPath string
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
			got, err := getMarkdownForTag(tt.args.tagName, tt.args.dirPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("getMarkdownForTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMarkdownForTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isExistsScope(t *testing.T) {
	type args struct {
		scope string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isExistsScope(tt.args.scope)
			if (err != nil) != tt.wantErr {
				t.Errorf("isExistsScope() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isExistsScope() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTagsFromComment(t *testing.T) {
	type args struct {
		comment string
	}
	tests := []struct {
		name     string
		args     args
		wantTags []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTags := getTagsFromComment(tt.args.comment); !reflect.DeepEqual(gotTags, tt.wantTags) {
				t.Errorf("getTagsFromComment() = %v, want %v", gotTags, tt.wantTags)
			}
		})
	}
}

func TestParser_matchTags(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		comments []*ast.Comment
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantMatch bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if gotMatch := parser.matchTags(tt.args.comments); gotMatch != tt.wantMatch {
				t.Errorf("Parser.matchTags() = %v, want %v", gotMatch, tt.wantMatch)
			}
		})
	}
}

func Test_matchExtension(t *testing.T) {
	type args struct {
		extensionToMatch string
		comments         []*ast.Comment
	}
	tests := []struct {
		name      string
		args      args
		wantMatch bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMatch := matchExtension(tt.args.extensionToMatch, tt.args.comments); gotMatch != tt.wantMatch {
				t.Errorf("matchExtension() = %v, want %v", gotMatch, tt.wantMatch)
			}
		})
	}
}

func TestParser_ParseRouterAPIInfo(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		fileInfo *AstFileInfo
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if err := parser.ParseRouterAPIInfo(tt.args.fileInfo); (err != nil) != tt.wantErr {
				t.Errorf("Parser.ParseRouterAPIInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_refRouteMethodOp(t *testing.T) {
	type args struct {
		item   *spec.PathItem
		method string
	}
	tests := []struct {
		name   string
		args   args
		wantOp **spec.Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOp := refRouteMethodOp(tt.args.item, tt.args.method); !reflect.DeepEqual(gotOp, tt.wantOp) {
				t.Errorf("refRouteMethodOp() = %v, want %v", gotOp, tt.wantOp)
			}
		})
	}
}

func Test_processRouterOperation(t *testing.T) {
	type args struct {
		parser    *Parser
		operation *Operation
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
			if err := processRouterOperation(tt.args.parser, tt.args.operation); (err != nil) != tt.wantErr {
				t.Errorf("processRouterOperation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_convertFromSpecificToPrimitive(t *testing.T) {
	type args struct {
		typeName string
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
			got, err := convertFromSpecificToPrimitive(tt.args.typeName)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertFromSpecificToPrimitive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertFromSpecificToPrimitive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_getTypeSchema(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		typeName string
		file     *ast.File
		ref      bool
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			got, err := parser.getTypeSchema(tt.args.typeName, tt.args.file, tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.getTypeSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.getTypeSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_getRefTypeSchema(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		typeSpecDef *TypeSpecDef
		schema      *Schema
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *spec.Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if got := parser.getRefTypeSchema(tt.args.typeSpecDef, tt.args.schema); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.getRefTypeSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_isInStructStack(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		typeSpecDef *TypeSpecDef
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if got := parser.isInStructStack(tt.args.typeSpecDef); got != tt.want {
				t.Errorf("Parser.isInStructStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_ParseDefinition(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		typeSpecDef *TypeSpecDef
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			got, err := parser.ParseDefinition(tt.args.typeSpecDef)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.ParseDefinition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.ParseDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fullTypeName(t *testing.T) {
	type args struct {
		parts []string
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
			if got := fullTypeName(tt.args.parts...); got != tt.want {
				t.Errorf("fullTypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fillDefinitionDescription(t *testing.T) {
	type args struct {
		definition  *spec.Schema
		file        *ast.File
		typeSpecDef *TypeSpecDef
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fillDefinitionDescription(tt.args.definition, tt.args.file, tt.args.typeSpecDef)
		})
	}
}

func Test_extractDeclarationDescription(t *testing.T) {
	type args struct {
		commentGroups []*ast.CommentGroup
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
			if got := extractDeclarationDescription(tt.args.commentGroups...); got != tt.want {
				t.Errorf("extractDeclarationDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_parseTypeExpr(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		file     *ast.File
		typeExpr ast.Expr
		ref      bool
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			got, err := parser.parseTypeExpr(tt.args.file, tt.args.typeExpr, tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.parseTypeExpr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.parseTypeExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_parseStruct(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		file   *ast.File
		fields *ast.FieldList
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			got, err := parser.parseStruct(tt.args.file, tt.args.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.parseStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.parseStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_parseStructField(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		file  *ast.File
		field *ast.Field
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]spec.Schema
		want1   []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			got, got1, err := parser.parseStructField(tt.args.file, tt.args.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.parseStructField() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.parseStructField() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Parser.parseStructField() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getFieldType(t *testing.T) {
	type args struct {
		file                 *ast.File
		field                ast.Expr
		genericParamTypeDefs map[string]*genericTypeSpec
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
			got, err := getFieldType(tt.args.file, tt.args.field, tt.args.genericParamTypeDefs)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFieldType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getFieldType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_getUnderlyingSchema(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		schema *spec.Schema
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *spec.Schema
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if got := parser.getUnderlyingSchema(tt.args.schema); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.getUnderlyingSchema() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetSchemaTypePath(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		schema *spec.Schema
		depth  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if got := parser.GetSchemaTypePath(tt.args.schema, tt.args.depth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.GetSchemaTypePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceLastTag(t *testing.T) {
	type args struct {
		slice   []spec.Tag
		element spec.Tag
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replaceLastTag(tt.args.slice, tt.args.element)
		})
	}
}

func Test_defineTypeOfExample(t *testing.T) {
	type args struct {
		schemaType   string
		arrayType    string
		exampleValue string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := defineTypeOfExample(tt.args.schemaType, tt.args.arrayType, tt.args.exampleValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("defineTypeOfExample() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("defineTypeOfExample() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_getAllGoFileInfo(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		packageDir string
		searchDir  string
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if err := parser.getAllGoFileInfo(tt.args.packageDir, tt.args.searchDir); (err != nil) != tt.wantErr {
				t.Errorf("Parser.getAllGoFileInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_getAllGoFileInfoFromDeps(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		pkg       *depth.Pkg
		parseFlag ParseFlag
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if err := parser.getAllGoFileInfoFromDeps(tt.args.pkg, tt.args.parseFlag); (err != nil) != tt.wantErr {
				t.Errorf("Parser.getAllGoFileInfoFromDeps() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_parseFile(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		packageDir string
		path       string
		src        interface{}
		flag       ParseFlag
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if err := parser.parseFile(tt.args.packageDir, tt.args.path, tt.args.src, tt.args.flag); (err != nil) != tt.wantErr {
				t.Errorf("Parser.parseFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_checkOperationIDUniqueness(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if err := parser.checkOperationIDUniqueness(); (err != nil) != tt.wantErr {
				t.Errorf("Parser.checkOperationIDUniqueness() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_Skip(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		path string
		f    os.FileInfo
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if err := parser.Skip(tt.args.path, tt.args.f); (err != nil) != tt.wantErr {
				t.Errorf("Parser.Skip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_walkWith(t *testing.T) {
	type args struct {
		excludes    map[string]struct{}
		parseVendor bool
	}
	tests := []struct {
		name string
		args args
		want func(path string, fileInfo os.FileInfo) error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := walkWith(tt.args.excludes, tt.args.parseVendor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("walkWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_GetSwagger(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *spec.Swagger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			if got := parser.GetSwagger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.GetSwagger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_addTestType(t *testing.T) {
	type fields struct {
		swagger                 *spec.Swagger
		packages                *PackagesDefinitions
		parsedSchemas           map[*TypeSpecDef]*Schema
		outputSchemas           map[*TypeSpecDef]*Schema
		PropNamingStrategy      string
		ParseVendor             bool
		ParseDependency         ParseFlag
		ParseInternal           bool
		Strict                  bool
		RequiredByDefault       bool
		structStack             []*TypeSpecDef
		markdownFileDir         string
		codeExampleFilesDir     string
		collectionFormatInQuery string
		excludes                map[string]struct{}
		packagePrefix           []string
		parseExtension          string
		debug                   Debugger
		fieldParserFactory      FieldParserFactory
		Overrides               map[string]string
		parseGoList             bool
		tags                    map[string]struct{}
	}
	type args struct {
		typename string
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
			parser := &Parser{
				swagger:                 tt.fields.swagger,
				packages:                tt.fields.packages,
				parsedSchemas:           tt.fields.parsedSchemas,
				outputSchemas:           tt.fields.outputSchemas,
				PropNamingStrategy:      tt.fields.PropNamingStrategy,
				ParseVendor:             tt.fields.ParseVendor,
				ParseDependency:         tt.fields.ParseDependency,
				ParseInternal:           tt.fields.ParseInternal,
				Strict:                  tt.fields.Strict,
				RequiredByDefault:       tt.fields.RequiredByDefault,
				structStack:             tt.fields.structStack,
				markdownFileDir:         tt.fields.markdownFileDir,
				codeExampleFilesDir:     tt.fields.codeExampleFilesDir,
				collectionFormatInQuery: tt.fields.collectionFormatInQuery,
				excludes:                tt.fields.excludes,
				packagePrefix:           tt.fields.packagePrefix,
				parseExtension:          tt.fields.parseExtension,
				debug:                   tt.fields.debug,
				fieldParserFactory:      tt.fields.fieldParserFactory,
				Overrides:               tt.fields.Overrides,
				parseGoList:             tt.fields.parseGoList,
				tags:                    tt.fields.tags,
			}
			parser.addTestType(tt.args.typename)
		})
	}
}
