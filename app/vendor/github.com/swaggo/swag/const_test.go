package swag

import (
	"go/ast"
	"go/token"
	"reflect"
	"testing"
)

func TestEvaluateEscapedChar(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvaluateEscapedChar(tt.args.text); got != tt.want {
				t.Errorf("EvaluateEscapedChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvaluateEscapedString(t *testing.T) {
	type args struct {
		text string
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
			if got := EvaluateEscapedString(tt.args.text); got != tt.want {
				t.Errorf("EvaluateEscapedString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvaluateDataConversion(t *testing.T) {
	type args struct {
		x        interface{}
		typeName string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvaluateDataConversion(tt.args.x, tt.args.typeName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvaluateDataConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvaluateUnary(t *testing.T) {
	type args struct {
		x        interface{}
		operator token.Token
		xtype    ast.Expr
	}
	tests := []struct {
		name  string
		args  args
		want  interface{}
		want1 ast.Expr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := EvaluateUnary(tt.args.x, tt.args.operator, tt.args.xtype)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvaluateUnary() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("EvaluateUnary() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestEvaluateBinary(t *testing.T) {
	type args struct {
		x        interface{}
		y        interface{}
		operator token.Token
		xtype    ast.Expr
		ytype    ast.Expr
	}
	tests := []struct {
		name  string
		args  args
		want  interface{}
		want1 ast.Expr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := EvaluateBinary(tt.args.x, tt.args.y, tt.args.operator, tt.args.xtype, tt.args.ytype)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EvaluateBinary() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("EvaluateBinary() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
