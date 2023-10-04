// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package semver implements comparison of semantic version strings.
// In this package, semantic version strings must begin with a leading "v",
// as in "v1.0.0".
//
// The general form of a semantic version string accepted by this package is
//
//	vMAJOR[.MINOR[.PATCH[-PRERELEASE][+BUILD]]]
//
// where square brackets indicate optional parts of the syntax;
// MAJOR, MINOR, and PATCH are decimal integers without extra leading zeros;
// PRERELEASE and BUILD are each a series of non-empty dot-separated identifiers
// using only alphanumeric characters and hyphens; and
// all-numeric PRERELEASE identifiers must not have leading zeros.
//
// This package follows Semantic Versioning 2.0.0 (see semver.org)
// with two exceptions. First, it requires the "v" prefix. Second, it recognizes
// vMAJOR and vMAJOR.MINOR (with no prerelease or build suffixes)
// as shorthands for vMAJOR.0.0 and vMAJOR.MINOR.0.
package semver

import (
	"reflect"
	"testing"
)

func TestIsValid(t *testing.T) {
	type args struct {
		v string
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
			if got := IsValid(tt.args.v); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCanonical(t *testing.T) {
	type args struct {
		v string
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
			if got := Canonical(tt.args.v); got != tt.want {
				t.Errorf("Canonical() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMajor(t *testing.T) {
	type args struct {
		v string
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
			if got := Major(tt.args.v); got != tt.want {
				t.Errorf("Major() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMajorMinor(t *testing.T) {
	type args struct {
		v string
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
			if got := MajorMinor(tt.args.v); got != tt.want {
				t.Errorf("MajorMinor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrerelease(t *testing.T) {
	type args struct {
		v string
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
			if got := Prerelease(tt.args.v); got != tt.want {
				t.Errorf("Prerelease() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuild(t *testing.T) {
	type args struct {
		v string
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
			if got := Build(tt.args.v); got != tt.want {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	type args struct {
		v string
		w string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.v, tt.args.w); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		v string
		w string
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
			if got := Max(tt.args.v, tt.args.w); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByVersion_Len(t *testing.T) {
	tests := []struct {
		name string
		vs   ByVersion
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vs.Len(); got != tt.want {
				t.Errorf("ByVersion.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByVersion_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		vs   ByVersion
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.vs.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestByVersion_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		vs   ByVersion
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vs.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("ByVersion.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort(tt.args.list)
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name   string
		args   args
		wantP  parsed
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, gotOk := parse(tt.args.v)
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("parse() gotP = %v, want %v", gotP, tt.wantP)
			}
			if gotOk != tt.wantOk {
				t.Errorf("parse() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_parseInt(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name     string
		args     args
		wantT    string
		wantRest string
		wantOk   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotT, gotRest, gotOk := parseInt(tt.args.v)
			if gotT != tt.wantT {
				t.Errorf("parseInt() gotT = %v, want %v", gotT, tt.wantT)
			}
			if gotRest != tt.wantRest {
				t.Errorf("parseInt() gotRest = %v, want %v", gotRest, tt.wantRest)
			}
			if gotOk != tt.wantOk {
				t.Errorf("parseInt() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_parsePrerelease(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name     string
		args     args
		wantT    string
		wantRest string
		wantOk   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotT, gotRest, gotOk := parsePrerelease(tt.args.v)
			if gotT != tt.wantT {
				t.Errorf("parsePrerelease() gotT = %v, want %v", gotT, tt.wantT)
			}
			if gotRest != tt.wantRest {
				t.Errorf("parsePrerelease() gotRest = %v, want %v", gotRest, tt.wantRest)
			}
			if gotOk != tt.wantOk {
				t.Errorf("parsePrerelease() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_parseBuild(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name     string
		args     args
		wantT    string
		wantRest string
		wantOk   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotT, gotRest, gotOk := parseBuild(tt.args.v)
			if gotT != tt.wantT {
				t.Errorf("parseBuild() gotT = %v, want %v", gotT, tt.wantT)
			}
			if gotRest != tt.wantRest {
				t.Errorf("parseBuild() gotRest = %v, want %v", gotRest, tt.wantRest)
			}
			if gotOk != tt.wantOk {
				t.Errorf("parseBuild() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func Test_isIdentChar(t *testing.T) {
	type args struct {
		c byte
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
			if got := isIdentChar(tt.args.c); got != tt.want {
				t.Errorf("isIdentChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBadNum(t *testing.T) {
	type args struct {
		v string
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
			if got := isBadNum(tt.args.v); got != tt.want {
				t.Errorf("isBadNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNum(t *testing.T) {
	type args struct {
		v string
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
			if got := isNum(tt.args.v); got != tt.want {
				t.Errorf("isNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compareInt(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareInt(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("compareInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_comparePrerelease(t *testing.T) {
	type args struct {
		x string
		y string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comparePrerelease(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("comparePrerelease() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextIdent(t *testing.T) {
	type args struct {
		x string
	}
	tests := []struct {
		name     string
		args     args
		wantDx   string
		wantRest string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDx, gotRest := nextIdent(tt.args.x)
			if gotDx != tt.wantDx {
				t.Errorf("nextIdent() gotDx = %v, want %v", gotDx, tt.wantDx)
			}
			if gotRest != tt.wantRest {
				t.Errorf("nextIdent() gotRest = %v, want %v", gotRest, tt.wantRest)
			}
		})
	}
}
