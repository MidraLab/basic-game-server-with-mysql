//go:build !go1.18
// +build !go1.18

package swag

import (
	"go/ast"
	"reflect"
	"testing"

	"github.com/go-openapi/spec"
)

func TestPackagesDefinitions_parametrizeGenericType(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		file            *ast.File
		original        *TypeSpecDef
		fullGenericForm string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TypeSpecDef
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkgDefs := &PackagesDefinitions{
				files:             tt.fields.files,
				packages:          tt.fields.packages,
				uniqueDefinitions: tt.fields.uniqueDefinitions,
				parseDependency:   tt.fields.parseDependency,
				debug:             tt.fields.debug,
			}
			if got := pkgDefs.parametrizeGenericType(tt.args.file, tt.args.original, tt.args.fullGenericForm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackagesDefinitions.parametrizeGenericType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGenericFieldType(t *testing.T) {
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
			got, err := getGenericFieldType(tt.args.file, tt.args.field, tt.args.genericParamTypeDefs)
			if (err != nil) != tt.wantErr {
				t.Errorf("getGenericFieldType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getGenericFieldType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_parseGenericTypeExpr(t *testing.T) {
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
			got, err := parser.parseGenericTypeExpr(tt.args.file, tt.args.typeExpr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.parseGenericTypeExpr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.parseGenericTypeExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}
