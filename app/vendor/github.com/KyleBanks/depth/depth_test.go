// Package depth provides the ability to traverse and retrieve Go source code dependencies in the form of
// internal and external packages.
//
// For example, the dependencies of the stdlib `strings` package can be resolved like so:
//
// 	import "github.com/KyleBanks/depth"
//
//	var t depth.Tree
// 	err := t.Resolve("strings")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// Output: "strings has 4 dependencies."
// 	log.Printf("%v has %v dependencies.", t.Root.Name, len(t.Root.Deps))
//
// For additional customization, simply set the appropriate flags on the `Tree` before resolving:
//
// 	import "github.com/KyleBanks/depth"
//
// 	t := depth.Tree {
//  	ResolveInternal: true,
//   	ResolveTest: true,
//   	MaxDepth: 10,
// 	}
// 	err := t.Resolve("strings")
package depth

import "testing"

func TestTree_Resolve(t *testing.T) {
	type fields struct {
		Root            *Pkg
		ResolveInternal bool
		ResolveTest     bool
		MaxDepth        int
		Importer        Importer
		importCache     map[string]struct{}
	}
	type args struct {
		name string
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
			tr := &Tree{
				Root:            tt.fields.Root,
				ResolveInternal: tt.fields.ResolveInternal,
				ResolveTest:     tt.fields.ResolveTest,
				MaxDepth:        tt.fields.MaxDepth,
				Importer:        tt.fields.Importer,
				importCache:     tt.fields.importCache,
			}
			if err := tr.Resolve(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Tree.Resolve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTree_shouldResolveInternal(t *testing.T) {
	type fields struct {
		Root            *Pkg
		ResolveInternal bool
		ResolveTest     bool
		MaxDepth        int
		Importer        Importer
		importCache     map[string]struct{}
	}
	type args struct {
		parent *Pkg
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
			tr := &Tree{
				Root:            tt.fields.Root,
				ResolveInternal: tt.fields.ResolveInternal,
				ResolveTest:     tt.fields.ResolveTest,
				MaxDepth:        tt.fields.MaxDepth,
				Importer:        tt.fields.Importer,
				importCache:     tt.fields.importCache,
			}
			if got := tr.shouldResolveInternal(tt.args.parent); got != tt.want {
				t.Errorf("Tree.shouldResolveInternal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_isAtMaxDepth(t *testing.T) {
	type fields struct {
		Root            *Pkg
		ResolveInternal bool
		ResolveTest     bool
		MaxDepth        int
		Importer        Importer
		importCache     map[string]struct{}
	}
	type args struct {
		p *Pkg
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
			tr := &Tree{
				Root:            tt.fields.Root,
				ResolveInternal: tt.fields.ResolveInternal,
				ResolveTest:     tt.fields.ResolveTest,
				MaxDepth:        tt.fields.MaxDepth,
				Importer:        tt.fields.Importer,
				importCache:     tt.fields.importCache,
			}
			if got := tr.isAtMaxDepth(tt.args.p); got != tt.want {
				t.Errorf("Tree.isAtMaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_hasSeenImport(t *testing.T) {
	type fields struct {
		Root            *Pkg
		ResolveInternal bool
		ResolveTest     bool
		MaxDepth        int
		Importer        Importer
		importCache     map[string]struct{}
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
			tr := &Tree{
				Root:            tt.fields.Root,
				ResolveInternal: tt.fields.ResolveInternal,
				ResolveTest:     tt.fields.ResolveTest,
				MaxDepth:        tt.fields.MaxDepth,
				Importer:        tt.fields.Importer,
				importCache:     tt.fields.importCache,
			}
			if got := tr.hasSeenImport(tt.args.name); got != tt.want {
				t.Errorf("Tree.hasSeenImport() = %v, want %v", got, tt.want)
			}
		})
	}
}
