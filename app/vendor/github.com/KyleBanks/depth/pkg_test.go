package depth

import (
	"go/build"
	"testing"
)

func TestPkg_Resolve(t *testing.T) {
	type fields struct {
		Name     string
		SrcDir   string
		Internal bool
		Resolved bool
		Test     bool
		Tree     *Tree
		Parent   *Pkg
		Deps     []Pkg
		Raw      *build.Package
	}
	type args struct {
		i Importer
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
			p := &Pkg{
				Name:     tt.fields.Name,
				SrcDir:   tt.fields.SrcDir,
				Internal: tt.fields.Internal,
				Resolved: tt.fields.Resolved,
				Test:     tt.fields.Test,
				Tree:     tt.fields.Tree,
				Parent:   tt.fields.Parent,
				Deps:     tt.fields.Deps,
				Raw:      tt.fields.Raw,
			}
			p.Resolve(tt.args.i)
		})
	}
}

func TestPkg_setDeps(t *testing.T) {
	type fields struct {
		Name     string
		SrcDir   string
		Internal bool
		Resolved bool
		Test     bool
		Tree     *Tree
		Parent   *Pkg
		Deps     []Pkg
		Raw      *build.Package
	}
	type args struct {
		i       Importer
		imports []string
		srcDir  string
		unique  map[string]struct{}
		isTest  bool
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
			p := &Pkg{
				Name:     tt.fields.Name,
				SrcDir:   tt.fields.SrcDir,
				Internal: tt.fields.Internal,
				Resolved: tt.fields.Resolved,
				Test:     tt.fields.Test,
				Tree:     tt.fields.Tree,
				Parent:   tt.fields.Parent,
				Deps:     tt.fields.Deps,
				Raw:      tt.fields.Raw,
			}
			p.setDeps(tt.args.i, tt.args.imports, tt.args.srcDir, tt.args.unique, tt.args.isTest)
		})
	}
}

func TestPkg_addDep(t *testing.T) {
	type fields struct {
		Name     string
		SrcDir   string
		Internal bool
		Resolved bool
		Test     bool
		Tree     *Tree
		Parent   *Pkg
		Deps     []Pkg
		Raw      *build.Package
	}
	type args struct {
		i      Importer
		name   string
		srcDir string
		isTest bool
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
			p := &Pkg{
				Name:     tt.fields.Name,
				SrcDir:   tt.fields.SrcDir,
				Internal: tt.fields.Internal,
				Resolved: tt.fields.Resolved,
				Test:     tt.fields.Test,
				Tree:     tt.fields.Tree,
				Parent:   tt.fields.Parent,
				Deps:     tt.fields.Deps,
				Raw:      tt.fields.Raw,
			}
			p.addDep(tt.args.i, tt.args.name, tt.args.srcDir, tt.args.isTest)
		})
	}
}

func TestPkg_isParent(t *testing.T) {
	type fields struct {
		Name     string
		SrcDir   string
		Internal bool
		Resolved bool
		Test     bool
		Tree     *Tree
		Parent   *Pkg
		Deps     []Pkg
		Raw      *build.Package
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
			p := &Pkg{
				Name:     tt.fields.Name,
				SrcDir:   tt.fields.SrcDir,
				Internal: tt.fields.Internal,
				Resolved: tt.fields.Resolved,
				Test:     tt.fields.Test,
				Tree:     tt.fields.Tree,
				Parent:   tt.fields.Parent,
				Deps:     tt.fields.Deps,
				Raw:      tt.fields.Raw,
			}
			if got := p.isParent(tt.args.name); got != tt.want {
				t.Errorf("Pkg.isParent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPkg_depth(t *testing.T) {
	type fields struct {
		Name     string
		SrcDir   string
		Internal bool
		Resolved bool
		Test     bool
		Tree     *Tree
		Parent   *Pkg
		Deps     []Pkg
		Raw      *build.Package
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pkg{
				Name:     tt.fields.Name,
				SrcDir:   tt.fields.SrcDir,
				Internal: tt.fields.Internal,
				Resolved: tt.fields.Resolved,
				Test:     tt.fields.Test,
				Tree:     tt.fields.Tree,
				Parent:   tt.fields.Parent,
				Deps:     tt.fields.Deps,
				Raw:      tt.fields.Raw,
			}
			if got := p.depth(); got != tt.want {
				t.Errorf("Pkg.depth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPkg_cleanName(t *testing.T) {
	type fields struct {
		Name     string
		SrcDir   string
		Internal bool
		Resolved bool
		Test     bool
		Tree     *Tree
		Parent   *Pkg
		Deps     []Pkg
		Raw      *build.Package
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
			p := &Pkg{
				Name:     tt.fields.Name,
				SrcDir:   tt.fields.SrcDir,
				Internal: tt.fields.Internal,
				Resolved: tt.fields.Resolved,
				Test:     tt.fields.Test,
				Tree:     tt.fields.Tree,
				Parent:   tt.fields.Parent,
				Deps:     tt.fields.Deps,
				Raw:      tt.fields.Raw,
			}
			if got := p.cleanName(); got != tt.want {
				t.Errorf("Pkg.cleanName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPkg_String(t *testing.T) {
	type fields struct {
		Name     string
		SrcDir   string
		Internal bool
		Resolved bool
		Test     bool
		Tree     *Tree
		Parent   *Pkg
		Deps     []Pkg
		Raw      *build.Package
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
			p := &Pkg{
				Name:     tt.fields.Name,
				SrcDir:   tt.fields.SrcDir,
				Internal: tt.fields.Internal,
				Resolved: tt.fields.Resolved,
				Test:     tt.fields.Test,
				Tree:     tt.fields.Tree,
				Parent:   tt.fields.Parent,
				Deps:     tt.fields.Deps,
				Raw:      tt.fields.Raw,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Pkg.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byInternalAndName_Len(t *testing.T) {
	tests := []struct {
		name string
		b    byInternalAndName
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Len(); got != tt.want {
				t.Errorf("byInternalAndName.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byInternalAndName_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		b    byInternalAndName
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

func Test_byInternalAndName_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		b    byInternalAndName
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("byInternalAndName.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
