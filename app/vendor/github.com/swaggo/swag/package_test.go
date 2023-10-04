package swag

import (
	"go/ast"
	"reflect"
	"testing"
)

func TestNewPackageDefinitions(t *testing.T) {
	type args struct {
		name    string
		pkgPath string
	}
	tests := []struct {
		name string
		args args
		want *PackageDefinitions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPackageDefinitions(tt.args.name, tt.args.pkgPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPackageDefinitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackageDefinitions_AddFile(t *testing.T) {
	type fields struct {
		Files           map[string]*ast.File
		TypeDefinitions map[string]*TypeSpecDef
		ConstTable      map[string]*ConstVariable
		OrderedConst    []*ConstVariable
		Name            string
		Path            string
	}
	type args struct {
		pkgPath string
		file    *ast.File
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PackageDefinitions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg := &PackageDefinitions{
				Files:           tt.fields.Files,
				TypeDefinitions: tt.fields.TypeDefinitions,
				ConstTable:      tt.fields.ConstTable,
				OrderedConst:    tt.fields.OrderedConst,
				Name:            tt.fields.Name,
				Path:            tt.fields.Path,
			}
			if got := pkg.AddFile(tt.args.pkgPath, tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackageDefinitions.AddFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackageDefinitions_AddTypeSpec(t *testing.T) {
	type fields struct {
		Files           map[string]*ast.File
		TypeDefinitions map[string]*TypeSpecDef
		ConstTable      map[string]*ConstVariable
		OrderedConst    []*ConstVariable
		Name            string
		Path            string
	}
	type args struct {
		name     string
		typeSpec *TypeSpecDef
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PackageDefinitions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg := &PackageDefinitions{
				Files:           tt.fields.Files,
				TypeDefinitions: tt.fields.TypeDefinitions,
				ConstTable:      tt.fields.ConstTable,
				OrderedConst:    tt.fields.OrderedConst,
				Name:            tt.fields.Name,
				Path:            tt.fields.Path,
			}
			if got := pkg.AddTypeSpec(tt.args.name, tt.args.typeSpec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackageDefinitions.AddTypeSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackageDefinitions_AddConst(t *testing.T) {
	type fields struct {
		Files           map[string]*ast.File
		TypeDefinitions map[string]*TypeSpecDef
		ConstTable      map[string]*ConstVariable
		OrderedConst    []*ConstVariable
		Name            string
		Path            string
	}
	type args struct {
		astFile   *ast.File
		valueSpec *ast.ValueSpec
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PackageDefinitions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg := &PackageDefinitions{
				Files:           tt.fields.Files,
				TypeDefinitions: tt.fields.TypeDefinitions,
				ConstTable:      tt.fields.ConstTable,
				OrderedConst:    tt.fields.OrderedConst,
				Name:            tt.fields.Name,
				Path:            tt.fields.Path,
			}
			if got := pkg.AddConst(tt.args.astFile, tt.args.valueSpec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackageDefinitions.AddConst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackageDefinitions_evaluateConstValue(t *testing.T) {
	type fields struct {
		Files           map[string]*ast.File
		TypeDefinitions map[string]*TypeSpecDef
		ConstTable      map[string]*ConstVariable
		OrderedConst    []*ConstVariable
		Name            string
		Path            string
	}
	type args struct {
		file            *ast.File
		iota            int
		expr            ast.Expr
		globalEvaluator ConstVariableGlobalEvaluator
		recursiveStack  map[string]struct{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
		want1  ast.Expr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pkg := &PackageDefinitions{
				Files:           tt.fields.Files,
				TypeDefinitions: tt.fields.TypeDefinitions,
				ConstTable:      tt.fields.ConstTable,
				OrderedConst:    tt.fields.OrderedConst,
				Name:            tt.fields.Name,
				Path:            tt.fields.Path,
			}
			got, got1 := pkg.evaluateConstValue(tt.args.file, tt.args.iota, tt.args.expr, tt.args.globalEvaluator, tt.args.recursiveStack)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackageDefinitions.evaluateConstValue() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PackageDefinitions.evaluateConstValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
