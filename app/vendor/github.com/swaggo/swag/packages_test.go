package swag

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"
)

func TestNewPackagesDefinitions(t *testing.T) {
	tests := []struct {
		name string
		want *PackagesDefinitions
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPackagesDefinitions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPackagesDefinitions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackagesDefinitions_ParseFile(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
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
			pkgDefs := &PackagesDefinitions{
				files:             tt.fields.files,
				packages:          tt.fields.packages,
				uniqueDefinitions: tt.fields.uniqueDefinitions,
				parseDependency:   tt.fields.parseDependency,
				debug:             tt.fields.debug,
			}
			if err := pkgDefs.ParseFile(tt.args.packageDir, tt.args.path, tt.args.src, tt.args.flag); (err != nil) != tt.wantErr {
				t.Errorf("PackagesDefinitions.ParseFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPackagesDefinitions_collectAstFile(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		fileSet    *token.FileSet
		packageDir string
		path       string
		astFile    *ast.File
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
			pkgDefs := &PackagesDefinitions{
				files:             tt.fields.files,
				packages:          tt.fields.packages,
				uniqueDefinitions: tt.fields.uniqueDefinitions,
				parseDependency:   tt.fields.parseDependency,
				debug:             tt.fields.debug,
			}
			if err := pkgDefs.collectAstFile(tt.args.fileSet, tt.args.packageDir, tt.args.path, tt.args.astFile, tt.args.flag); (err != nil) != tt.wantErr {
				t.Errorf("PackagesDefinitions.collectAstFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPackagesDefinitions_RangeFiles(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		handle func(info *AstFileInfo) error
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
			pkgDefs := &PackagesDefinitions{
				files:             tt.fields.files,
				packages:          tt.fields.packages,
				uniqueDefinitions: tt.fields.uniqueDefinitions,
				parseDependency:   tt.fields.parseDependency,
				debug:             tt.fields.debug,
			}
			if err := pkgDefs.RangeFiles(tt.args.handle); (err != nil) != tt.wantErr {
				t.Errorf("PackagesDefinitions.RangeFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPackagesDefinitions_ParseTypes(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[*TypeSpecDef]*Schema
		wantErr bool
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
			got, err := pkgDefs.ParseTypes()
			if (err != nil) != tt.wantErr {
				t.Errorf("PackagesDefinitions.ParseTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackagesDefinitions.ParseTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackagesDefinitions_parseTypesFromFile(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		astFile       *ast.File
		packagePath   string
		parsedSchemas map[*TypeSpecDef]*Schema
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
			pkgDefs := &PackagesDefinitions{
				files:             tt.fields.files,
				packages:          tt.fields.packages,
				uniqueDefinitions: tt.fields.uniqueDefinitions,
				parseDependency:   tt.fields.parseDependency,
				debug:             tt.fields.debug,
			}
			pkgDefs.parseTypesFromFile(tt.args.astFile, tt.args.packagePath, tt.args.parsedSchemas)
		})
	}
}

func TestPackagesDefinitions_parseFunctionScopedTypesFromFile(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		astFile       *ast.File
		packagePath   string
		parsedSchemas map[*TypeSpecDef]*Schema
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
			pkgDefs := &PackagesDefinitions{
				files:             tt.fields.files,
				packages:          tt.fields.packages,
				uniqueDefinitions: tt.fields.uniqueDefinitions,
				parseDependency:   tt.fields.parseDependency,
				debug:             tt.fields.debug,
			}
			pkgDefs.parseFunctionScopedTypesFromFile(tt.args.astFile, tt.args.packagePath, tt.args.parsedSchemas)
		})
	}
}

func TestPackagesDefinitions_collectConstVariables(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		astFile            *ast.File
		packagePath        string
		generalDeclaration *ast.GenDecl
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
			pkgDefs := &PackagesDefinitions{
				files:             tt.fields.files,
				packages:          tt.fields.packages,
				uniqueDefinitions: tt.fields.uniqueDefinitions,
				parseDependency:   tt.fields.parseDependency,
				debug:             tt.fields.debug,
			}
			pkgDefs.collectConstVariables(tt.args.astFile, tt.args.packagePath, tt.args.generalDeclaration)
		})
	}
}

func TestPackagesDefinitions_evaluateAllConstVariables(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	tests := []struct {
		name   string
		fields fields
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
			pkgDefs.evaluateAllConstVariables()
		})
	}
}

func TestPackagesDefinitions_EvaluateConstValue(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		pkg            *PackageDefinitions
		cv             *ConstVariable
		recursiveStack map[string]struct{}
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
			pkgDefs := &PackagesDefinitions{
				files:             tt.fields.files,
				packages:          tt.fields.packages,
				uniqueDefinitions: tt.fields.uniqueDefinitions,
				parseDependency:   tt.fields.parseDependency,
				debug:             tt.fields.debug,
			}
			got, got1 := pkgDefs.EvaluateConstValue(tt.args.pkg, tt.args.cv, tt.args.recursiveStack)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackagesDefinitions.EvaluateConstValue() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PackagesDefinitions.EvaluateConstValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPackagesDefinitions_EvaluateConstValueByName(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		file              *ast.File
		pkgName           string
		constVariableName string
		recursiveStack    map[string]struct{}
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
			pkgDefs := &PackagesDefinitions{
				files:             tt.fields.files,
				packages:          tt.fields.packages,
				uniqueDefinitions: tt.fields.uniqueDefinitions,
				parseDependency:   tt.fields.parseDependency,
				debug:             tt.fields.debug,
			}
			got, got1 := pkgDefs.EvaluateConstValueByName(tt.args.file, tt.args.pkgName, tt.args.constVariableName, tt.args.recursiveStack)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackagesDefinitions.EvaluateConstValueByName() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PackagesDefinitions.EvaluateConstValueByName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestPackagesDefinitions_collectConstEnums(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		parsedSchemas map[*TypeSpecDef]*Schema
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
			pkgDefs := &PackagesDefinitions{
				files:             tt.fields.files,
				packages:          tt.fields.packages,
				uniqueDefinitions: tt.fields.uniqueDefinitions,
				parseDependency:   tt.fields.parseDependency,
				debug:             tt.fields.debug,
			}
			pkgDefs.collectConstEnums(tt.args.parsedSchemas)
		})
	}
}

func TestPackagesDefinitions_removeAllNotUniqueTypes(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	tests := []struct {
		name   string
		fields fields
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
			pkgDefs.removeAllNotUniqueTypes()
		})
	}
}

func TestPackagesDefinitions_findTypeSpec(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		pkgPath  string
		typeName string
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
			if got := pkgDefs.findTypeSpec(tt.args.pkgPath, tt.args.typeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackagesDefinitions.findTypeSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackagesDefinitions_loadExternalPackage(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		importPath string
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
			pkgDefs := &PackagesDefinitions{
				files:             tt.fields.files,
				packages:          tt.fields.packages,
				uniqueDefinitions: tt.fields.uniqueDefinitions,
				parseDependency:   tt.fields.parseDependency,
				debug:             tt.fields.debug,
			}
			if err := pkgDefs.loadExternalPackage(tt.args.importPath); (err != nil) != tt.wantErr {
				t.Errorf("PackagesDefinitions.loadExternalPackage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPackagesDefinitions_findPackagePathFromImports(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		pkg  string
		file *ast.File
	}
	tests := []struct {
		name                 string
		fields               fields
		args                 args
		wantMatchedPkgPaths  []string
		wantExternalPkgPaths []string
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
			gotMatchedPkgPaths, gotExternalPkgPaths := pkgDefs.findPackagePathFromImports(tt.args.pkg, tt.args.file)
			if !reflect.DeepEqual(gotMatchedPkgPaths, tt.wantMatchedPkgPaths) {
				t.Errorf("PackagesDefinitions.findPackagePathFromImports() gotMatchedPkgPaths = %v, want %v", gotMatchedPkgPaths, tt.wantMatchedPkgPaths)
			}
			if !reflect.DeepEqual(gotExternalPkgPaths, tt.wantExternalPkgPaths) {
				t.Errorf("PackagesDefinitions.findPackagePathFromImports() gotExternalPkgPaths = %v, want %v", gotExternalPkgPaths, tt.wantExternalPkgPaths)
			}
		})
	}
}

func TestPackagesDefinitions_findTypeSpecFromPackagePaths(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		matchedPkgPaths  []string
		externalPkgPaths []string
		name             string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantTypeDef *TypeSpecDef
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
			if gotTypeDef := pkgDefs.findTypeSpecFromPackagePaths(tt.args.matchedPkgPaths, tt.args.externalPkgPaths, tt.args.name); !reflect.DeepEqual(gotTypeDef, tt.wantTypeDef) {
				t.Errorf("PackagesDefinitions.findTypeSpecFromPackagePaths() = %v, want %v", gotTypeDef, tt.wantTypeDef)
			}
		})
	}
}

func TestPackagesDefinitions_FindTypeSpec(t *testing.T) {
	type fields struct {
		files             map[*ast.File]*AstFileInfo
		packages          map[string]*PackageDefinitions
		uniqueDefinitions map[string]*TypeSpecDef
		parseDependency   ParseFlag
		debug             Debugger
	}
	type args struct {
		typeName string
		file     *ast.File
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
			if got := pkgDefs.FindTypeSpec(tt.args.typeName, tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackagesDefinitions.FindTypeSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}
