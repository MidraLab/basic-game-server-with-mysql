package swag

import (
	"context"
	"go/build"
	"reflect"
	"testing"
)

func Test_listPackages(t *testing.T) {
	type args struct {
		ctx  context.Context
		dir  string
		env  []string
		args []string
	}
	tests := []struct {
		name     string
		args     args
		wantPkgs []*build.Package
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPkgs, err := listPackages(tt.args.ctx, tt.args.dir, tt.args.env, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("listPackages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPkgs, tt.wantPkgs) {
				t.Errorf("listPackages() = %v, want %v", gotPkgs, tt.wantPkgs)
			}
		})
	}
}

func TestParser_getAllGoFileInfoFromDepsByList(t *testing.T) {
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
		pkg       *build.Package
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
			if err := parser.getAllGoFileInfoFromDepsByList(tt.args.pkg, tt.args.parseFlag); (err != nil) != tt.wantErr {
				t.Errorf("Parser.getAllGoFileInfoFromDepsByList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
