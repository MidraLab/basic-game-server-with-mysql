// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package loader

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"go/types"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestPackageInfo_String(t *testing.T) {
	type fields struct {
		Pkg                   *types.Package
		Importable            bool
		TransitivelyErrorFree bool
		Files                 []*ast.File
		Errors                []error
		Info                  types.Info
		dir                   string
		checker               *types.Checker
		errorFunc             func(error)
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
			info := &PackageInfo{
				Pkg:                   tt.fields.Pkg,
				Importable:            tt.fields.Importable,
				TransitivelyErrorFree: tt.fields.TransitivelyErrorFree,
				Files:                 tt.fields.Files,
				Errors:                tt.fields.Errors,
				Info:                  tt.fields.Info,
				dir:                   tt.fields.dir,
				checker:               tt.fields.checker,
				errorFunc:             tt.fields.errorFunc,
			}
			if got := info.String(); got != tt.want {
				t.Errorf("PackageInfo.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackageInfo_appendError(t *testing.T) {
	type fields struct {
		Pkg                   *types.Package
		Importable            bool
		TransitivelyErrorFree bool
		Files                 []*ast.File
		Errors                []error
		Info                  types.Info
		dir                   string
		checker               *types.Checker
		errorFunc             func(error)
	}
	type args struct {
		err error
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
			info := &PackageInfo{
				Pkg:                   tt.fields.Pkg,
				Importable:            tt.fields.Importable,
				TransitivelyErrorFree: tt.fields.TransitivelyErrorFree,
				Files:                 tt.fields.Files,
				Errors:                tt.fields.Errors,
				Info:                  tt.fields.Info,
				dir:                   tt.fields.dir,
				checker:               tt.fields.checker,
				errorFunc:             tt.fields.errorFunc,
			}
			info.appendError(tt.args.err)
		})
	}
}

func TestConfig_fset(t *testing.T) {
	type fields struct {
		Fset                *token.FileSet
		ParserMode          parser.Mode
		TypeChecker         types.Config
		TypeCheckFuncBodies func(path string) bool
		Build               *build.Context
		Cwd                 string
		DisplayPath         func(path string) string
		AllowErrors         bool
		CreatePkgs          []PkgSpec
		ImportPkgs          map[string]bool
		FindPackage         func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
		AfterTypeCheck      func(info *PackageInfo, files []*ast.File)
	}
	tests := []struct {
		name   string
		fields fields
		want   *token.FileSet
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Fset:                tt.fields.Fset,
				ParserMode:          tt.fields.ParserMode,
				TypeChecker:         tt.fields.TypeChecker,
				TypeCheckFuncBodies: tt.fields.TypeCheckFuncBodies,
				Build:               tt.fields.Build,
				Cwd:                 tt.fields.Cwd,
				DisplayPath:         tt.fields.DisplayPath,
				AllowErrors:         tt.fields.AllowErrors,
				CreatePkgs:          tt.fields.CreatePkgs,
				ImportPkgs:          tt.fields.ImportPkgs,
				FindPackage:         tt.fields.FindPackage,
				AfterTypeCheck:      tt.fields.AfterTypeCheck,
			}
			if got := conf.fset(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.fset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_ParseFile(t *testing.T) {
	type fields struct {
		Fset                *token.FileSet
		ParserMode          parser.Mode
		TypeChecker         types.Config
		TypeCheckFuncBodies func(path string) bool
		Build               *build.Context
		Cwd                 string
		DisplayPath         func(path string) string
		AllowErrors         bool
		CreatePkgs          []PkgSpec
		ImportPkgs          map[string]bool
		FindPackage         func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
		AfterTypeCheck      func(info *PackageInfo, files []*ast.File)
	}
	type args struct {
		filename string
		src      interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *ast.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Fset:                tt.fields.Fset,
				ParserMode:          tt.fields.ParserMode,
				TypeChecker:         tt.fields.TypeChecker,
				TypeCheckFuncBodies: tt.fields.TypeCheckFuncBodies,
				Build:               tt.fields.Build,
				Cwd:                 tt.fields.Cwd,
				DisplayPath:         tt.fields.DisplayPath,
				AllowErrors:         tt.fields.AllowErrors,
				CreatePkgs:          tt.fields.CreatePkgs,
				ImportPkgs:          tt.fields.ImportPkgs,
				FindPackage:         tt.fields.FindPackage,
				AfterTypeCheck:      tt.fields.AfterTypeCheck,
			}
			got, err := conf.ParseFile(tt.args.filename, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.ParseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.ParseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_FromArgs(t *testing.T) {
	type fields struct {
		Fset                *token.FileSet
		ParserMode          parser.Mode
		TypeChecker         types.Config
		TypeCheckFuncBodies func(path string) bool
		Build               *build.Context
		Cwd                 string
		DisplayPath         func(path string) string
		AllowErrors         bool
		CreatePkgs          []PkgSpec
		ImportPkgs          map[string]bool
		FindPackage         func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
		AfterTypeCheck      func(info *PackageInfo, files []*ast.File)
	}
	type args struct {
		args  []string
		xtest bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Fset:                tt.fields.Fset,
				ParserMode:          tt.fields.ParserMode,
				TypeChecker:         tt.fields.TypeChecker,
				TypeCheckFuncBodies: tt.fields.TypeCheckFuncBodies,
				Build:               tt.fields.Build,
				Cwd:                 tt.fields.Cwd,
				DisplayPath:         tt.fields.DisplayPath,
				AllowErrors:         tt.fields.AllowErrors,
				CreatePkgs:          tt.fields.CreatePkgs,
				ImportPkgs:          tt.fields.ImportPkgs,
				FindPackage:         tt.fields.FindPackage,
				AfterTypeCheck:      tt.fields.AfterTypeCheck,
			}
			got, err := conf.FromArgs(tt.args.args, tt.args.xtest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.FromArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.FromArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_CreateFromFilenames(t *testing.T) {
	type fields struct {
		Fset                *token.FileSet
		ParserMode          parser.Mode
		TypeChecker         types.Config
		TypeCheckFuncBodies func(path string) bool
		Build               *build.Context
		Cwd                 string
		DisplayPath         func(path string) string
		AllowErrors         bool
		CreatePkgs          []PkgSpec
		ImportPkgs          map[string]bool
		FindPackage         func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
		AfterTypeCheck      func(info *PackageInfo, files []*ast.File)
	}
	type args struct {
		path      string
		filenames []string
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
			conf := &Config{
				Fset:                tt.fields.Fset,
				ParserMode:          tt.fields.ParserMode,
				TypeChecker:         tt.fields.TypeChecker,
				TypeCheckFuncBodies: tt.fields.TypeCheckFuncBodies,
				Build:               tt.fields.Build,
				Cwd:                 tt.fields.Cwd,
				DisplayPath:         tt.fields.DisplayPath,
				AllowErrors:         tt.fields.AllowErrors,
				CreatePkgs:          tt.fields.CreatePkgs,
				ImportPkgs:          tt.fields.ImportPkgs,
				FindPackage:         tt.fields.FindPackage,
				AfterTypeCheck:      tt.fields.AfterTypeCheck,
			}
			conf.CreateFromFilenames(tt.args.path, tt.args.filenames...)
		})
	}
}

func TestConfig_CreateFromFiles(t *testing.T) {
	type fields struct {
		Fset                *token.FileSet
		ParserMode          parser.Mode
		TypeChecker         types.Config
		TypeCheckFuncBodies func(path string) bool
		Build               *build.Context
		Cwd                 string
		DisplayPath         func(path string) string
		AllowErrors         bool
		CreatePkgs          []PkgSpec
		ImportPkgs          map[string]bool
		FindPackage         func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
		AfterTypeCheck      func(info *PackageInfo, files []*ast.File)
	}
	type args struct {
		path  string
		files []*ast.File
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
			conf := &Config{
				Fset:                tt.fields.Fset,
				ParserMode:          tt.fields.ParserMode,
				TypeChecker:         tt.fields.TypeChecker,
				TypeCheckFuncBodies: tt.fields.TypeCheckFuncBodies,
				Build:               tt.fields.Build,
				Cwd:                 tt.fields.Cwd,
				DisplayPath:         tt.fields.DisplayPath,
				AllowErrors:         tt.fields.AllowErrors,
				CreatePkgs:          tt.fields.CreatePkgs,
				ImportPkgs:          tt.fields.ImportPkgs,
				FindPackage:         tt.fields.FindPackage,
				AfterTypeCheck:      tt.fields.AfterTypeCheck,
			}
			conf.CreateFromFiles(tt.args.path, tt.args.files...)
		})
	}
}

func TestConfig_ImportWithTests(t *testing.T) {
	type fields struct {
		Fset                *token.FileSet
		ParserMode          parser.Mode
		TypeChecker         types.Config
		TypeCheckFuncBodies func(path string) bool
		Build               *build.Context
		Cwd                 string
		DisplayPath         func(path string) string
		AllowErrors         bool
		CreatePkgs          []PkgSpec
		ImportPkgs          map[string]bool
		FindPackage         func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
		AfterTypeCheck      func(info *PackageInfo, files []*ast.File)
	}
	type args struct {
		path string
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
			conf := &Config{
				Fset:                tt.fields.Fset,
				ParserMode:          tt.fields.ParserMode,
				TypeChecker:         tt.fields.TypeChecker,
				TypeCheckFuncBodies: tt.fields.TypeCheckFuncBodies,
				Build:               tt.fields.Build,
				Cwd:                 tt.fields.Cwd,
				DisplayPath:         tt.fields.DisplayPath,
				AllowErrors:         tt.fields.AllowErrors,
				CreatePkgs:          tt.fields.CreatePkgs,
				ImportPkgs:          tt.fields.ImportPkgs,
				FindPackage:         tt.fields.FindPackage,
				AfterTypeCheck:      tt.fields.AfterTypeCheck,
			}
			conf.ImportWithTests(tt.args.path)
		})
	}
}

func TestConfig_Import(t *testing.T) {
	type fields struct {
		Fset                *token.FileSet
		ParserMode          parser.Mode
		TypeChecker         types.Config
		TypeCheckFuncBodies func(path string) bool
		Build               *build.Context
		Cwd                 string
		DisplayPath         func(path string) string
		AllowErrors         bool
		CreatePkgs          []PkgSpec
		ImportPkgs          map[string]bool
		FindPackage         func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
		AfterTypeCheck      func(info *PackageInfo, files []*ast.File)
	}
	type args struct {
		path string
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
			conf := &Config{
				Fset:                tt.fields.Fset,
				ParserMode:          tt.fields.ParserMode,
				TypeChecker:         tt.fields.TypeChecker,
				TypeCheckFuncBodies: tt.fields.TypeCheckFuncBodies,
				Build:               tt.fields.Build,
				Cwd:                 tt.fields.Cwd,
				DisplayPath:         tt.fields.DisplayPath,
				AllowErrors:         tt.fields.AllowErrors,
				CreatePkgs:          tt.fields.CreatePkgs,
				ImportPkgs:          tt.fields.ImportPkgs,
				FindPackage:         tt.fields.FindPackage,
				AfterTypeCheck:      tt.fields.AfterTypeCheck,
			}
			conf.Import(tt.args.path)
		})
	}
}

func TestConfig_addImport(t *testing.T) {
	type fields struct {
		Fset                *token.FileSet
		ParserMode          parser.Mode
		TypeChecker         types.Config
		TypeCheckFuncBodies func(path string) bool
		Build               *build.Context
		Cwd                 string
		DisplayPath         func(path string) string
		AllowErrors         bool
		CreatePkgs          []PkgSpec
		ImportPkgs          map[string]bool
		FindPackage         func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
		AfterTypeCheck      func(info *PackageInfo, files []*ast.File)
	}
	type args struct {
		path  string
		tests bool
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
			conf := &Config{
				Fset:                tt.fields.Fset,
				ParserMode:          tt.fields.ParserMode,
				TypeChecker:         tt.fields.TypeChecker,
				TypeCheckFuncBodies: tt.fields.TypeCheckFuncBodies,
				Build:               tt.fields.Build,
				Cwd:                 tt.fields.Cwd,
				DisplayPath:         tt.fields.DisplayPath,
				AllowErrors:         tt.fields.AllowErrors,
				CreatePkgs:          tt.fields.CreatePkgs,
				ImportPkgs:          tt.fields.ImportPkgs,
				FindPackage:         tt.fields.FindPackage,
				AfterTypeCheck:      tt.fields.AfterTypeCheck,
			}
			conf.addImport(tt.args.path, tt.args.tests)
		})
	}
}

func TestProgram_PathEnclosingInterval(t *testing.T) {
	type fields struct {
		Fset        *token.FileSet
		Created     []*PackageInfo
		Imported    map[string]*PackageInfo
		AllPackages map[*types.Package]*PackageInfo
		importMap   map[string]*types.Package
	}
	type args struct {
		start token.Pos
		end   token.Pos
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantPkg   *PackageInfo
		wantPath  []ast.Node
		wantExact bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prog := &Program{
				Fset:        tt.fields.Fset,
				Created:     tt.fields.Created,
				Imported:    tt.fields.Imported,
				AllPackages: tt.fields.AllPackages,
				importMap:   tt.fields.importMap,
			}
			gotPkg, gotPath, gotExact := prog.PathEnclosingInterval(tt.args.start, tt.args.end)
			if !reflect.DeepEqual(gotPkg, tt.wantPkg) {
				t.Errorf("Program.PathEnclosingInterval() gotPkg = %v, want %v", gotPkg, tt.wantPkg)
			}
			if !reflect.DeepEqual(gotPath, tt.wantPath) {
				t.Errorf("Program.PathEnclosingInterval() gotPath = %v, want %v", gotPath, tt.wantPath)
			}
			if gotExact != tt.wantExact {
				t.Errorf("Program.PathEnclosingInterval() gotExact = %v, want %v", gotExact, tt.wantExact)
			}
		})
	}
}

func TestProgram_InitialPackages(t *testing.T) {
	type fields struct {
		Fset        *token.FileSet
		Created     []*PackageInfo
		Imported    map[string]*PackageInfo
		AllPackages map[*types.Package]*PackageInfo
		importMap   map[string]*types.Package
	}
	tests := []struct {
		name   string
		fields fields
		want   []*PackageInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prog := &Program{
				Fset:        tt.fields.Fset,
				Created:     tt.fields.Created,
				Imported:    tt.fields.Imported,
				AllPackages: tt.fields.AllPackages,
				importMap:   tt.fields.importMap,
			}
			if got := prog.InitialPackages(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Program.InitialPackages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProgram_Package(t *testing.T) {
	type fields struct {
		Fset        *token.FileSet
		Created     []*PackageInfo
		Imported    map[string]*PackageInfo
		AllPackages map[*types.Package]*PackageInfo
		importMap   map[string]*types.Package
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PackageInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prog := &Program{
				Fset:        tt.fields.Fset,
				Created:     tt.fields.Created,
				Imported:    tt.fields.Imported,
				AllPackages: tt.fields.AllPackages,
				importMap:   tt.fields.importMap,
			}
			if got := prog.Package(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Program.Package() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_importInfo_awaitCompletion(t *testing.T) {
	type fields struct {
		path     string
		info     *PackageInfo
		complete chan struct{}
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ii := &importInfo{
				path:     tt.fields.path,
				info:     tt.fields.info,
				complete: tt.fields.complete,
			}
			ii.awaitCompletion()
		})
	}
}

func Test_importInfo_Complete(t *testing.T) {
	type fields struct {
		path     string
		info     *PackageInfo
		complete chan struct{}
	}
	type args struct {
		info *PackageInfo
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
			ii := &importInfo{
				path:     tt.fields.path,
				info:     tt.fields.info,
				complete: tt.fields.complete,
			}
			ii.Complete(tt.args.info)
		})
	}
}

func TestConfig_Load(t *testing.T) {
	type fields struct {
		Fset                *token.FileSet
		ParserMode          parser.Mode
		TypeChecker         types.Config
		TypeCheckFuncBodies func(path string) bool
		Build               *build.Context
		Cwd                 string
		DisplayPath         func(path string) string
		AllowErrors         bool
		CreatePkgs          []PkgSpec
		ImportPkgs          map[string]bool
		FindPackage         func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
		AfterTypeCheck      func(info *PackageInfo, files []*ast.File)
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Program
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Fset:                tt.fields.Fset,
				ParserMode:          tt.fields.ParserMode,
				TypeChecker:         tt.fields.TypeChecker,
				TypeCheckFuncBodies: tt.fields.TypeCheckFuncBodies,
				Build:               tt.fields.Build,
				Cwd:                 tt.fields.Cwd,
				DisplayPath:         tt.fields.DisplayPath,
				AllowErrors:         tt.fields.AllowErrors,
				CreatePkgs:          tt.fields.CreatePkgs,
				ImportPkgs:          tt.fields.ImportPkgs,
				FindPackage:         tt.fields.FindPackage,
				AfterTypeCheck:      tt.fields.AfterTypeCheck,
			}
			got, err := conf.Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byImportPath_Len(t *testing.T) {
	tests := []struct {
		name string
		b    byImportPath
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Len(); got != tt.want {
				t.Errorf("byImportPath.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byImportPath_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		b    byImportPath
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("byImportPath.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byImportPath_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		b    byImportPath
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_markErrorFreePackages(t *testing.T) {
	type args struct {
		allPackages map[*types.Package]*PackageInfo
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			markErrorFreePackages(tt.args.allPackages)
		})
	}
}

func TestConfig_build(t *testing.T) {
	type fields struct {
		Fset                *token.FileSet
		ParserMode          parser.Mode
		TypeChecker         types.Config
		TypeCheckFuncBodies func(path string) bool
		Build               *build.Context
		Cwd                 string
		DisplayPath         func(path string) string
		AllowErrors         bool
		CreatePkgs          []PkgSpec
		ImportPkgs          map[string]bool
		FindPackage         func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
		AfterTypeCheck      func(info *PackageInfo, files []*ast.File)
	}
	tests := []struct {
		name   string
		fields fields
		want   *build.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Fset:                tt.fields.Fset,
				ParserMode:          tt.fields.ParserMode,
				TypeChecker:         tt.fields.TypeChecker,
				TypeCheckFuncBodies: tt.fields.TypeCheckFuncBodies,
				Build:               tt.fields.Build,
				Cwd:                 tt.fields.Cwd,
				DisplayPath:         tt.fields.DisplayPath,
				AllowErrors:         tt.fields.AllowErrors,
				CreatePkgs:          tt.fields.CreatePkgs,
				ImportPkgs:          tt.fields.ImportPkgs,
				FindPackage:         tt.fields.FindPackage,
				AfterTypeCheck:      tt.fields.AfterTypeCheck,
			}
			if got := conf.build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_parsePackageFiles(t *testing.T) {
	type fields struct {
		Fset                *token.FileSet
		ParserMode          parser.Mode
		TypeChecker         types.Config
		TypeCheckFuncBodies func(path string) bool
		Build               *build.Context
		Cwd                 string
		DisplayPath         func(path string) string
		AllowErrors         bool
		CreatePkgs          []PkgSpec
		ImportPkgs          map[string]bool
		FindPackage         func(ctxt *build.Context, importPath, fromDir string, mode build.ImportMode) (*build.Package, error)
		AfterTypeCheck      func(info *PackageInfo, files []*ast.File)
	}
	type args struct {
		bp    *build.Package
		which rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []*ast.File
		want1  []error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Config{
				Fset:                tt.fields.Fset,
				ParserMode:          tt.fields.ParserMode,
				TypeChecker:         tt.fields.TypeChecker,
				TypeCheckFuncBodies: tt.fields.TypeCheckFuncBodies,
				Build:               tt.fields.Build,
				Cwd:                 tt.fields.Cwd,
				DisplayPath:         tt.fields.DisplayPath,
				AllowErrors:         tt.fields.AllowErrors,
				CreatePkgs:          tt.fields.CreatePkgs,
				ImportPkgs:          tt.fields.ImportPkgs,
				FindPackage:         tt.fields.FindPackage,
				AfterTypeCheck:      tt.fields.AfterTypeCheck,
			}
			got, got1 := conf.parsePackageFiles(tt.args.bp, tt.args.which)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.parsePackageFiles() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Config.parsePackageFiles() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_importer_doImport(t *testing.T) {
	type fields struct {
		conf       *Config
		start      time.Time
		progMu     sync.Mutex
		prog       *Program
		findpkgMu  sync.Mutex
		findpkg    map[findpkgKey]*findpkgValue
		importedMu sync.Mutex
		imported   map[string]*importInfo
		graphMu    sync.Mutex
		graph      map[string]map[string]bool
	}
	type args struct {
		from *PackageInfo
		to   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Package
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			imp := &importer{
				conf:       tt.fields.conf,
				start:      tt.fields.start,
				progMu:     tt.fields.progMu,
				prog:       tt.fields.prog,
				findpkgMu:  tt.fields.findpkgMu,
				findpkg:    tt.fields.findpkg,
				importedMu: tt.fields.importedMu,
				imported:   tt.fields.imported,
				graphMu:    tt.fields.graphMu,
				graph:      tt.fields.graph,
			}
			got, err := imp.doImport(tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("importer.doImport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importer.doImport() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_importer_findPackage(t *testing.T) {
	type fields struct {
		conf       *Config
		start      time.Time
		progMu     sync.Mutex
		prog       *Program
		findpkgMu  sync.Mutex
		findpkg    map[findpkgKey]*findpkgValue
		importedMu sync.Mutex
		imported   map[string]*importInfo
		graphMu    sync.Mutex
		graph      map[string]map[string]bool
	}
	type args struct {
		importPath string
		fromDir    string
		mode       build.ImportMode
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *build.Package
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			imp := &importer{
				conf:       tt.fields.conf,
				start:      tt.fields.start,
				progMu:     tt.fields.progMu,
				prog:       tt.fields.prog,
				findpkgMu:  tt.fields.findpkgMu,
				findpkg:    tt.fields.findpkg,
				importedMu: tt.fields.importedMu,
				imported:   tt.fields.imported,
				graphMu:    tt.fields.graphMu,
				graph:      tt.fields.graph,
			}
			got, err := imp.findPackage(tt.args.importPath, tt.args.fromDir, tt.args.mode)
			if (err != nil) != tt.wantErr {
				t.Errorf("importer.findPackage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importer.findPackage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_importer_importAll(t *testing.T) {
	type fields struct {
		conf       *Config
		start      time.Time
		progMu     sync.Mutex
		prog       *Program
		findpkgMu  sync.Mutex
		findpkg    map[findpkgKey]*findpkgValue
		importedMu sync.Mutex
		imported   map[string]*importInfo
		graphMu    sync.Mutex
		graph      map[string]map[string]bool
	}
	type args struct {
		fromPath string
		fromDir  string
		imports  map[string]bool
		mode     build.ImportMode
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantInfos  []*PackageInfo
		wantErrors []importError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			imp := &importer{
				conf:       tt.fields.conf,
				start:      tt.fields.start,
				progMu:     tt.fields.progMu,
				prog:       tt.fields.prog,
				findpkgMu:  tt.fields.findpkgMu,
				findpkg:    tt.fields.findpkg,
				importedMu: tt.fields.importedMu,
				imported:   tt.fields.imported,
				graphMu:    tt.fields.graphMu,
				graph:      tt.fields.graph,
			}
			gotInfos, gotErrors := imp.importAll(tt.args.fromPath, tt.args.fromDir, tt.args.imports, tt.args.mode)
			if !reflect.DeepEqual(gotInfos, tt.wantInfos) {
				t.Errorf("importer.importAll() gotInfos = %v, want %v", gotInfos, tt.wantInfos)
			}
			if !reflect.DeepEqual(gotErrors, tt.wantErrors) {
				t.Errorf("importer.importAll() gotErrors = %v, want %v", gotErrors, tt.wantErrors)
			}
		})
	}
}

func Test_importer_findPath(t *testing.T) {
	type fields struct {
		conf       *Config
		start      time.Time
		progMu     sync.Mutex
		prog       *Program
		findpkgMu  sync.Mutex
		findpkg    map[findpkgKey]*findpkgValue
		importedMu sync.Mutex
		imported   map[string]*importInfo
		graphMu    sync.Mutex
		graph      map[string]map[string]bool
	}
	type args struct {
		from string
		to   string
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
			imp := &importer{
				conf:       tt.fields.conf,
				start:      tt.fields.start,
				progMu:     tt.fields.progMu,
				prog:       tt.fields.prog,
				findpkgMu:  tt.fields.findpkgMu,
				findpkg:    tt.fields.findpkg,
				importedMu: tt.fields.importedMu,
				imported:   tt.fields.imported,
				graphMu:    tt.fields.graphMu,
				graph:      tt.fields.graph,
			}
			if got := imp.findPath(tt.args.from, tt.args.to); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importer.findPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_importer_startLoad(t *testing.T) {
	type fields struct {
		conf       *Config
		start      time.Time
		progMu     sync.Mutex
		prog       *Program
		findpkgMu  sync.Mutex
		findpkg    map[findpkgKey]*findpkgValue
		importedMu sync.Mutex
		imported   map[string]*importInfo
		graphMu    sync.Mutex
		graph      map[string]map[string]bool
	}
	type args struct {
		bp *build.Package
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *importInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			imp := &importer{
				conf:       tt.fields.conf,
				start:      tt.fields.start,
				progMu:     tt.fields.progMu,
				prog:       tt.fields.prog,
				findpkgMu:  tt.fields.findpkgMu,
				findpkg:    tt.fields.findpkg,
				importedMu: tt.fields.importedMu,
				imported:   tt.fields.imported,
				graphMu:    tt.fields.graphMu,
				graph:      tt.fields.graph,
			}
			if got := imp.startLoad(tt.args.bp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importer.startLoad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_importer_load(t *testing.T) {
	type fields struct {
		conf       *Config
		start      time.Time
		progMu     sync.Mutex
		prog       *Program
		findpkgMu  sync.Mutex
		findpkg    map[findpkgKey]*findpkgValue
		importedMu sync.Mutex
		imported   map[string]*importInfo
		graphMu    sync.Mutex
		graph      map[string]map[string]bool
	}
	type args struct {
		bp *build.Package
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PackageInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			imp := &importer{
				conf:       tt.fields.conf,
				start:      tt.fields.start,
				progMu:     tt.fields.progMu,
				prog:       tt.fields.prog,
				findpkgMu:  tt.fields.findpkgMu,
				findpkg:    tt.fields.findpkg,
				importedMu: tt.fields.importedMu,
				imported:   tt.fields.imported,
				graphMu:    tt.fields.graphMu,
				graph:      tt.fields.graph,
			}
			if got := imp.load(tt.args.bp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importer.load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_importer_addFiles(t *testing.T) {
	type fields struct {
		conf       *Config
		start      time.Time
		progMu     sync.Mutex
		prog       *Program
		findpkgMu  sync.Mutex
		findpkg    map[findpkgKey]*findpkgValue
		importedMu sync.Mutex
		imported   map[string]*importInfo
		graphMu    sync.Mutex
		graph      map[string]map[string]bool
	}
	type args struct {
		info       *PackageInfo
		files      []*ast.File
		cycleCheck bool
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
			imp := &importer{
				conf:       tt.fields.conf,
				start:      tt.fields.start,
				progMu:     tt.fields.progMu,
				prog:       tt.fields.prog,
				findpkgMu:  tt.fields.findpkgMu,
				findpkg:    tt.fields.findpkg,
				importedMu: tt.fields.importedMu,
				imported:   tt.fields.imported,
				graphMu:    tt.fields.graphMu,
				graph:      tt.fields.graph,
			}
			imp.addFiles(tt.args.info, tt.args.files, tt.args.cycleCheck)
		})
	}
}

func Test_importer_newPackageInfo(t *testing.T) {
	type fields struct {
		conf       *Config
		start      time.Time
		progMu     sync.Mutex
		prog       *Program
		findpkgMu  sync.Mutex
		findpkg    map[findpkgKey]*findpkgValue
		importedMu sync.Mutex
		imported   map[string]*importInfo
		graphMu    sync.Mutex
		graph      map[string]map[string]bool
	}
	type args struct {
		path string
		dir  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PackageInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			imp := &importer{
				conf:       tt.fields.conf,
				start:      tt.fields.start,
				progMu:     tt.fields.progMu,
				prog:       tt.fields.prog,
				findpkgMu:  tt.fields.findpkgMu,
				findpkg:    tt.fields.findpkg,
				importedMu: tt.fields.importedMu,
				imported:   tt.fields.imported,
				graphMu:    tt.fields.graphMu,
				graph:      tt.fields.graph,
			}
			if got := imp.newPackageInfo(tt.args.path, tt.args.dir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("importer.newPackageInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_closure_Import(t *testing.T) {
	type fields struct {
		imp  *importer
		info *PackageInfo
	}
	type args struct {
		to string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *types.Package
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := closure{
				imp:  tt.fields.imp,
				info: tt.fields.info,
			}
			got, err := c.Import(tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("closure.Import() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("closure.Import() = %v, want %v", got, tt.want)
			}
		})
	}
}
