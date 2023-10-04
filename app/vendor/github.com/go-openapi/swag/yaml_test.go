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
	"encoding/json"
	"reflect"
	"testing"

	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
	yaml "gopkg.in/yaml.v3"
)

func TestYAMLMatcher(t *testing.T) {
	type args struct {
		path string
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
			if got := YAMLMatcher(tt.args.path); got != tt.want {
				t.Errorf("YAMLMatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYAMLToJSON(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    json.RawMessage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := YAMLToJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("YAMLToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YAMLToJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToYAMLDoc(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BytesToYAMLDoc(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("BytesToYAMLDoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToYAMLDoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yamlNode(t *testing.T) {
	type args struct {
		root *yaml.Node
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := yamlNode(tt.args.root)
			if (err != nil) != tt.wantErr {
				t.Errorf("yamlNode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("yamlNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yamlDocument(t *testing.T) {
	type args struct {
		node *yaml.Node
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := yamlDocument(tt.args.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("yamlDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("yamlDocument() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yamlMapping(t *testing.T) {
	type args struct {
		node *yaml.Node
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := yamlMapping(tt.args.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("yamlMapping() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("yamlMapping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yamlSequence(t *testing.T) {
	type args struct {
		node *yaml.Node
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := yamlSequence(tt.args.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("yamlSequence() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("yamlSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yamlScalar(t *testing.T) {
	type args struct {
		node *yaml.Node
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := yamlScalar(tt.args.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("yamlScalar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("yamlScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_yamlStringScalarC(t *testing.T) {
	type args struct {
		node *yaml.Node
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := yamlStringScalarC(tt.args.node)
			if (err != nil) != tt.wantErr {
				t.Errorf("yamlStringScalarC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("yamlStringScalarC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONMapSlice_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		s       JSONMapSlice
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONMapSlice.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONMapSlice.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONMapSlice_MarshalEasyJSON(t *testing.T) {
	type args struct {
		w *jwriter.Writer
	}
	tests := []struct {
		name string
		s    JSONMapSlice
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.MarshalEasyJSON(tt.args.w)
		})
	}
}

func TestJSONMapSlice_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		s       *JSONMapSlice
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("JSONMapSlice.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJSONMapSlice_UnmarshalEasyJSON(t *testing.T) {
	type args struct {
		in *jlexer.Lexer
	}
	tests := []struct {
		name string
		s    *JSONMapSlice
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.UnmarshalEasyJSON(tt.args.in)
		})
	}
}

func TestJSONMapSlice_MarshalYAML(t *testing.T) {
	tests := []struct {
		name    string
		s       JSONMapSlice
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.MarshalYAML()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONMapSlice.MarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONMapSlice.MarshalYAML() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_json2yaml(t *testing.T) {
	type args struct {
		item interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *yaml.Node
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json2yaml(tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("json2yaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("json2yaml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONMapItem_MarshalJSON(t *testing.T) {
	type fields struct {
		Key   string
		Value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := JSONMapItem{
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			got, err := s.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONMapItem.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONMapItem.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSONMapItem_MarshalEasyJSON(t *testing.T) {
	type fields struct {
		Key   string
		Value interface{}
	}
	type args struct {
		w *jwriter.Writer
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
			s := JSONMapItem{
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			s.MarshalEasyJSON(tt.args.w)
		})
	}
}

func TestJSONMapItem_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Key   string
		Value interface{}
	}
	type args struct {
		data []byte
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
			s := &JSONMapItem{
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if err := s.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("JSONMapItem.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJSONMapItem_UnmarshalEasyJSON(t *testing.T) {
	type fields struct {
		Key   string
		Value interface{}
	}
	type args struct {
		in *jlexer.Lexer
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
			s := &JSONMapItem{
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			s.UnmarshalEasyJSON(tt.args.in)
		})
	}
}

func Test_transformData(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantOut interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := transformData(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("transformData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("transformData() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestYAMLDoc(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    json.RawMessage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := YAMLDoc(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("YAMLDoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YAMLDoc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestYAMLData(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := YAMLData(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("YAMLData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("YAMLData() = %v, want %v", got, tt.want)
			}
		})
	}
}
