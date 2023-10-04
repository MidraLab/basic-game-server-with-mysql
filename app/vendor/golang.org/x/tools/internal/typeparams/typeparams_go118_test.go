// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.18
// +build go1.18

package typeparams

import (
	"go/ast"
	"go/types"
	"reflect"
	"testing"
)

func TestForTypeSpec(t *testing.T) {
	type args struct {
		n *ast.TypeSpec
	}
	tests := []struct {
		name string
		args args
		want *ast.FieldList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ForTypeSpec(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForTypeSpec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForFuncType(t *testing.T) {
	type args struct {
		n *ast.FuncType
	}
	tests := []struct {
		name string
		args args
		want *ast.FieldList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ForFuncType(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForFuncType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTypeParam(t *testing.T) {
	type args struct {
		name       *types.TypeName
		constraint types.Type
	}
	tests := []struct {
		name string
		args args
		want *TypeParam
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTypeParam(tt.args.name, tt.args.constraint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTypeParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetTypeParamConstraint(t *testing.T) {
	type args struct {
		tparam     *TypeParam
		constraint types.Type
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetTypeParamConstraint(tt.args.tparam, tt.args.constraint)
		})
	}
}

func TestNewSignatureType(t *testing.T) {
	type args struct {
		recv           *types.Var
		recvTypeParams []*TypeParam
		typeParams     []*TypeParam
		params         *types.Tuple
		results        *types.Tuple
		variadic       bool
	}
	tests := []struct {
		name string
		args args
		want *types.Signature
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSignatureType(tt.args.recv, tt.args.recvTypeParams, tt.args.typeParams, tt.args.params, tt.args.results, tt.args.variadic); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSignatureType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForSignature(t *testing.T) {
	type args struct {
		sig *types.Signature
	}
	tests := []struct {
		name string
		args args
		want *TypeParamList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ForSignature(tt.args.sig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForSignature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecvTypeParams(t *testing.T) {
	type args struct {
		sig *types.Signature
	}
	tests := []struct {
		name string
		args args
		want *TypeParamList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecvTypeParams(tt.args.sig); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecvTypeParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsComparable(t *testing.T) {
	type args struct {
		iface *types.Interface
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
			if got := IsComparable(tt.args.iface); got != tt.want {
				t.Errorf("IsComparable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMethodSet(t *testing.T) {
	type args struct {
		iface *types.Interface
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
			if got := IsMethodSet(tt.args.iface); got != tt.want {
				t.Errorf("IsMethodSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsImplicit(t *testing.T) {
	type args struct {
		iface *types.Interface
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
			if got := IsImplicit(tt.args.iface); got != tt.want {
				t.Errorf("IsImplicit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarkImplicit(t *testing.T) {
	type args struct {
		iface *types.Interface
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MarkImplicit(tt.args.iface)
		})
	}
}

func TestForNamed(t *testing.T) {
	type args struct {
		named *types.Named
	}
	tests := []struct {
		name string
		args args
		want *TypeParamList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ForNamed(tt.args.named); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ForNamed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetForNamed(t *testing.T) {
	type args struct {
		n       *types.Named
		tparams []*TypeParam
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetForNamed(tt.args.n, tt.args.tparams)
		})
	}
}

func TestNamedTypeArgs(t *testing.T) {
	type args struct {
		named *types.Named
	}
	tests := []struct {
		name string
		args args
		want *TypeList
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NamedTypeArgs(tt.args.named); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NamedTypeArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNamedTypeOrigin(t *testing.T) {
	type args struct {
		named *types.Named
	}
	tests := []struct {
		name string
		args args
		want *types.Named
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NamedTypeOrigin(tt.args.named); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NamedTypeOrigin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTerm(t *testing.T) {
	type args struct {
		tilde bool
		typ   types.Type
	}
	tests := []struct {
		name string
		args args
		want *Term
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTerm(tt.args.tilde, tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTerm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUnion(t *testing.T) {
	type args struct {
		terms []*Term
	}
	tests := []struct {
		name string
		args args
		want *Union
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUnion(tt.args.terms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitInstanceInfo(t *testing.T) {
	type args struct {
		info *types.Info
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitInstanceInfo(tt.args.info)
		})
	}
}

func TestGetInstances(t *testing.T) {
	type args struct {
		info *types.Info
	}
	tests := []struct {
		name string
		args args
		want map[*ast.Ident]Instance
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstances(tt.args.info); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewContext(t *testing.T) {
	tests := []struct {
		name string
		want *Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContext(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInstantiate(t *testing.T) {
	type args struct {
		ctxt     *Context
		typ      types.Type
		targs    []types.Type
		validate bool
	}
	tests := []struct {
		name    string
		args    args
		want    types.Type
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Instantiate(tt.args.ctxt, tt.args.typ, tt.args.targs, tt.args.validate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Instantiate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Instantiate() = %v, want %v", got, tt.want)
			}
		})
	}
}
