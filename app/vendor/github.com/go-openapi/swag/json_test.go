// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swag

import (
	"reflect"
	"sync"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WriteJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WriteJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadJSON(t *testing.T) {
	type args struct {
		data  []byte
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ReadJSON(tt.args.data, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("ReadJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDynamicJSONToStruct(t *testing.T) {
	type args struct {
		data   interface{}
		target interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DynamicJSONToStruct(tt.args.data, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("DynamicJSONToStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConcatJSON(t *testing.T) {
	type args struct {
		blobs [][]byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatJSON(tt.args.blobs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConcatJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToDynamicJSON(t *testing.T) {
	type args struct {
		data interface{}
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
			if got := ToDynamicJSON(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDynamicJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromDynamicJSON(t *testing.T) {
	type args struct {
		data   interface{}
		target interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FromDynamicJSON(tt.args.data, tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("FromDynamicJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewNameProvider(t *testing.T) {
	tests := []struct {
		name string
		want *NameProvider
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNameProvider(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNameProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildnameIndex(t *testing.T) {
	type args struct {
		tpe        reflect.Type
		idx        map[string]string
		reverseIdx map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildnameIndex(tt.args.tpe, tt.args.idx, tt.args.reverseIdx)
		})
	}
}

func Test_newNameIndex(t *testing.T) {
	type args struct {
		tpe reflect.Type
	}
	tests := []struct {
		name string
		args args
		want nameIndex
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newNameIndex(tt.args.tpe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNameIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNameProvider_GetJSONNames(t *testing.T) {
	type fields struct {
		lock  *sync.Mutex
		index map[reflect.Type]nameIndex
	}
	type args struct {
		subject interface{}
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
			n := &NameProvider{
				lock:  tt.fields.lock,
				index: tt.fields.index,
			}
			if got := n.GetJSONNames(tt.args.subject); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NameProvider.GetJSONNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNameProvider_GetJSONName(t *testing.T) {
	type fields struct {
		lock  *sync.Mutex
		index map[reflect.Type]nameIndex
	}
	type args struct {
		subject interface{}
		name    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NameProvider{
				lock:  tt.fields.lock,
				index: tt.fields.index,
			}
			got, got1 := n.GetJSONName(tt.args.subject, tt.args.name)
			if got != tt.want {
				t.Errorf("NameProvider.GetJSONName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NameProvider.GetJSONName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNameProvider_GetJSONNameForType(t *testing.T) {
	type fields struct {
		lock  *sync.Mutex
		index map[reflect.Type]nameIndex
	}
	type args struct {
		tpe  reflect.Type
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NameProvider{
				lock:  tt.fields.lock,
				index: tt.fields.index,
			}
			got, got1 := n.GetJSONNameForType(tt.args.tpe, tt.args.name)
			if got != tt.want {
				t.Errorf("NameProvider.GetJSONNameForType() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NameProvider.GetJSONNameForType() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNameProvider_makeNameIndex(t *testing.T) {
	type fields struct {
		lock  *sync.Mutex
		index map[reflect.Type]nameIndex
	}
	type args struct {
		tpe reflect.Type
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   nameIndex
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NameProvider{
				lock:  tt.fields.lock,
				index: tt.fields.index,
			}
			if got := n.makeNameIndex(tt.args.tpe); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NameProvider.makeNameIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNameProvider_GetGoName(t *testing.T) {
	type fields struct {
		lock  *sync.Mutex
		index map[reflect.Type]nameIndex
	}
	type args struct {
		subject interface{}
		name    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NameProvider{
				lock:  tt.fields.lock,
				index: tt.fields.index,
			}
			got, got1 := n.GetGoName(tt.args.subject, tt.args.name)
			if got != tt.want {
				t.Errorf("NameProvider.GetGoName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NameProvider.GetGoName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNameProvider_GetGoNameForType(t *testing.T) {
	type fields struct {
		lock  *sync.Mutex
		index map[reflect.Type]nameIndex
	}
	type args struct {
		tpe  reflect.Type
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NameProvider{
				lock:  tt.fields.lock,
				index: tt.fields.index,
			}
			got, got1 := n.GetGoNameForType(tt.args.tpe, tt.args.name)
			if got != tt.want {
				t.Errorf("NameProvider.GetGoNameForType() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NameProvider.GetGoNameForType() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
