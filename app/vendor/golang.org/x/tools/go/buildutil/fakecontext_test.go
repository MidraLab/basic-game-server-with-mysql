// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buildutil

import (
	"go/build"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestFakeContext(t *testing.T) {
	type args struct {
		pkgs map[string]map[string]string
	}
	tests := []struct {
		name string
		args args
		want *build.Context
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FakeContext(tt.args.pkgs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byName_Len(t *testing.T) {
	tests := []struct {
		name string
		s    byName
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("byName.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_byName_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    byName
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Swap(tt.args.i, tt.args.j)
		})
	}
}

func Test_byName_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		s    byName
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("byName.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeFileInfo_Name(t *testing.T) {
	tests := []struct {
		name string
		fi   fakeFileInfo
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fi.Name(); got != tt.want {
				t.Errorf("fakeFileInfo.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeFileInfo_Sys(t *testing.T) {
	tests := []struct {
		name string
		f    fakeFileInfo
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Sys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fakeFileInfo.Sys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeFileInfo_ModTime(t *testing.T) {
	tests := []struct {
		name string
		f    fakeFileInfo
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ModTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fakeFileInfo.ModTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeFileInfo_IsDir(t *testing.T) {
	tests := []struct {
		name string
		f    fakeFileInfo
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.IsDir(); got != tt.want {
				t.Errorf("fakeFileInfo.IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeFileInfo_Size(t *testing.T) {
	tests := []struct {
		name string
		f    fakeFileInfo
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Size(); got != tt.want {
				t.Errorf("fakeFileInfo.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeFileInfo_Mode(t *testing.T) {
	tests := []struct {
		name string
		f    fakeFileInfo
		want os.FileMode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Mode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fakeFileInfo.Mode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeDirInfo_Name(t *testing.T) {
	tests := []struct {
		name string
		fd   fakeDirInfo
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fd.Name(); got != tt.want {
				t.Errorf("fakeDirInfo.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeDirInfo_Sys(t *testing.T) {
	tests := []struct {
		name string
		f    fakeDirInfo
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Sys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fakeDirInfo.Sys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeDirInfo_ModTime(t *testing.T) {
	tests := []struct {
		name string
		f    fakeDirInfo
		want time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.ModTime(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fakeDirInfo.ModTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeDirInfo_IsDir(t *testing.T) {
	tests := []struct {
		name string
		f    fakeDirInfo
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.IsDir(); got != tt.want {
				t.Errorf("fakeDirInfo.IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeDirInfo_Size(t *testing.T) {
	tests := []struct {
		name string
		f    fakeDirInfo
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Size(); got != tt.want {
				t.Errorf("fakeDirInfo.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fakeDirInfo_Mode(t *testing.T) {
	tests := []struct {
		name string
		f    fakeDirInfo
		want os.FileMode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Mode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fakeDirInfo.Mode() = %v, want %v", got, tt.want)
			}
		})
	}
}
