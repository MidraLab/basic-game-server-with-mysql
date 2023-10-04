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

package spec

import (
	"reflect"
	"testing"
)

func TestOperationProps_MarshalJSON(t *testing.T) {
	type fields struct {
		Description  string
		Consumes     []string
		Produces     []string
		Schemes      []string
		Tags         []string
		Summary      string
		ExternalDocs *ExternalDocumentation
		ID           string
		Deprecated   bool
		Security     []map[string][]string
		Parameters   []Parameter
		Responses    *Responses
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
			op := OperationProps{
				Description:  tt.fields.Description,
				Consumes:     tt.fields.Consumes,
				Produces:     tt.fields.Produces,
				Schemes:      tt.fields.Schemes,
				Tags:         tt.fields.Tags,
				Summary:      tt.fields.Summary,
				ExternalDocs: tt.fields.ExternalDocs,
				ID:           tt.fields.ID,
				Deprecated:   tt.fields.Deprecated,
				Security:     tt.fields.Security,
				Parameters:   tt.fields.Parameters,
				Responses:    tt.fields.Responses,
			}
			got, err := op.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("OperationProps.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OperationProps.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_SuccessResponse(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	tests := []struct {
		name   string
		fields fields
		want   *Response
		want1  int
		want2  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			got, got1, got2 := o.SuccessResponse()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.SuccessResponse() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Operation.SuccessResponse() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Operation.SuccessResponse() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestOperation_JSONLookup(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			got, err := o.JSONLookup(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("Operation.JSONLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.JSONLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_UnmarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
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
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if err := o.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Operation.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperation_MarshalJSON(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
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
			o := Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			got, err := o.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Operation.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOperation(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOperation(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_WithID(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.WithID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.WithID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_WithDescription(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		description string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.WithDescription(tt.args.description); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.WithDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_WithSummary(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		summary string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.WithSummary(tt.args.summary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.WithSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_WithExternalDocs(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		description string
		url         string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.WithExternalDocs(tt.args.description, tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.WithExternalDocs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_Deprecate(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	tests := []struct {
		name   string
		fields fields
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.Deprecate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.Deprecate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_Undeprecate(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	tests := []struct {
		name   string
		fields fields
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.Undeprecate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.Undeprecate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_WithConsumes(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		mediaTypes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.WithConsumes(tt.args.mediaTypes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.WithConsumes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_WithProduces(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		mediaTypes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.WithProduces(tt.args.mediaTypes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.WithProduces() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_WithTags(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		tags []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.WithTags(tt.args.tags...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.WithTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_AddParam(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		param *Parameter
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.AddParam(tt.args.param); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.AddParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_RemoveParam(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		name string
		in   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.RemoveParam(tt.args.name, tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.RemoveParam() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_SecuredWith(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		name   string
		scopes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.SecuredWith(tt.args.name, tt.args.scopes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.SecuredWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_WithDefaultResponse(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		response *Response
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.WithDefaultResponse(tt.args.response); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.WithDefaultResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_RespondsWith(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		code     int
		response *Response
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Operation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if got := o.RespondsWith(tt.args.code, tt.args.response); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.RespondsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_GobEncode(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
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
			o := Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			got, err := o.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("Operation.GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Operation.GobEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperation_GobDecode(t *testing.T) {
	type fields struct {
		VendorExtensible VendorExtensible
		OperationProps   OperationProps
	}
	type args struct {
		b []byte
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
			o := &Operation{
				VendorExtensible: tt.fields.VendorExtensible,
				OperationProps:   tt.fields.OperationProps,
			}
			if err := o.GobDecode(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Operation.GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperationProps_GobEncode(t *testing.T) {
	type fields struct {
		Description  string
		Consumes     []string
		Produces     []string
		Schemes      []string
		Tags         []string
		Summary      string
		ExternalDocs *ExternalDocumentation
		ID           string
		Deprecated   bool
		Security     []map[string][]string
		Parameters   []Parameter
		Responses    *Responses
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
			op := OperationProps{
				Description:  tt.fields.Description,
				Consumes:     tt.fields.Consumes,
				Produces:     tt.fields.Produces,
				Schemes:      tt.fields.Schemes,
				Tags:         tt.fields.Tags,
				Summary:      tt.fields.Summary,
				ExternalDocs: tt.fields.ExternalDocs,
				ID:           tt.fields.ID,
				Deprecated:   tt.fields.Deprecated,
				Security:     tt.fields.Security,
				Parameters:   tt.fields.Parameters,
				Responses:    tt.fields.Responses,
			}
			got, err := op.GobEncode()
			if (err != nil) != tt.wantErr {
				t.Errorf("OperationProps.GobEncode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OperationProps.GobEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperationProps_GobDecode(t *testing.T) {
	type fields struct {
		Description  string
		Consumes     []string
		Produces     []string
		Schemes      []string
		Tags         []string
		Summary      string
		ExternalDocs *ExternalDocumentation
		ID           string
		Deprecated   bool
		Security     []map[string][]string
		Parameters   []Parameter
		Responses    *Responses
	}
	type args struct {
		b []byte
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
			op := &OperationProps{
				Description:  tt.fields.Description,
				Consumes:     tt.fields.Consumes,
				Produces:     tt.fields.Produces,
				Schemes:      tt.fields.Schemes,
				Tags:         tt.fields.Tags,
				Summary:      tt.fields.Summary,
				ExternalDocs: tt.fields.ExternalDocs,
				ID:           tt.fields.ID,
				Deprecated:   tt.fields.Deprecated,
				Security:     tt.fields.Security,
				Parameters:   tt.fields.Parameters,
				Responses:    tt.fields.Responses,
			}
			if err := op.GobDecode(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("OperationProps.GobDecode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
