package swag

import (
	"go/ast"
	"testing"
)

func TestTypeSpecDef_Name(t *testing.T) {
	type fields struct {
		File       *ast.File
		TypeSpec   *ast.TypeSpec
		Enums      []EnumValue
		PkgPath    string
		ParentSpec ast.Decl
		NotUnique  bool
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
			tr := &TypeSpecDef{
				File:       tt.fields.File,
				TypeSpec:   tt.fields.TypeSpec,
				Enums:      tt.fields.Enums,
				PkgPath:    tt.fields.PkgPath,
				ParentSpec: tt.fields.ParentSpec,
				NotUnique:  tt.fields.NotUnique,
			}
			if got := tr.Name(); got != tt.want {
				t.Errorf("TypeSpecDef.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeSpecDef_TypeName(t *testing.T) {
	type fields struct {
		File       *ast.File
		TypeSpec   *ast.TypeSpec
		Enums      []EnumValue
		PkgPath    string
		ParentSpec ast.Decl
		NotUnique  bool
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
			tr := &TypeSpecDef{
				File:       tt.fields.File,
				TypeSpec:   tt.fields.TypeSpec,
				Enums:      tt.fields.Enums,
				PkgPath:    tt.fields.PkgPath,
				ParentSpec: tt.fields.ParentSpec,
				NotUnique:  tt.fields.NotUnique,
			}
			if got := tr.TypeName(); got != tt.want {
				t.Errorf("TypeSpecDef.TypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeSpecDef_FullPath(t *testing.T) {
	type fields struct {
		File       *ast.File
		TypeSpec   *ast.TypeSpec
		Enums      []EnumValue
		PkgPath    string
		ParentSpec ast.Decl
		NotUnique  bool
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
			tr := &TypeSpecDef{
				File:       tt.fields.File,
				TypeSpec:   tt.fields.TypeSpec,
				Enums:      tt.fields.Enums,
				PkgPath:    tt.fields.PkgPath,
				ParentSpec: tt.fields.ParentSpec,
				NotUnique:  tt.fields.NotUnique,
			}
			if got := tr.FullPath(); got != tt.want {
				t.Errorf("TypeSpecDef.FullPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
