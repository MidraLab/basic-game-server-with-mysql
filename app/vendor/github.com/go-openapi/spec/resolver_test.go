package spec

import (
	"reflect"
	"testing"
)

func Test_resolveAnyWithBase(t *testing.T) {
	type args struct {
		root    interface{}
		ref     *Ref
		result  interface{}
		options *ExpandOptions
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
			if err := resolveAnyWithBase(tt.args.root, tt.args.ref, tt.args.result, tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("resolveAnyWithBase() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestResolveRefWithBase(t *testing.T) {
	type args struct {
		root    interface{}
		ref     *Ref
		options *ExpandOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveRefWithBase(tt.args.root, tt.args.ref, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveRefWithBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveRefWithBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveRef(t *testing.T) {
	type args struct {
		root interface{}
		ref  *Ref
	}
	tests := []struct {
		name    string
		args    args
		want    *Schema
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveRef(tt.args.root, tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveRef() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveRef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveParameterWithBase(t *testing.T) {
	type args struct {
		root    interface{}
		ref     Ref
		options *ExpandOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *Parameter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveParameterWithBase(tt.args.root, tt.args.ref, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveParameterWithBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveParameterWithBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveParameter(t *testing.T) {
	type args struct {
		root interface{}
		ref  Ref
	}
	tests := []struct {
		name    string
		args    args
		want    *Parameter
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveParameter(tt.args.root, tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveParameter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveResponseWithBase(t *testing.T) {
	type args struct {
		root    interface{}
		ref     Ref
		options *ExpandOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveResponseWithBase(tt.args.root, tt.args.ref, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveResponseWithBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveResponseWithBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveResponse(t *testing.T) {
	type args struct {
		root interface{}
		ref  Ref
	}
	tests := []struct {
		name    string
		args    args
		want    *Response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveResponse(tt.args.root, tt.args.ref)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolvePathItemWithBase(t *testing.T) {
	type args struct {
		root    interface{}
		ref     Ref
		options *ExpandOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *PathItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolvePathItemWithBase(tt.args.root, tt.args.ref, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolvePathItemWithBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolvePathItemWithBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolvePathItem(t *testing.T) {
	type args struct {
		root    interface{}
		ref     Ref
		options *ExpandOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *PathItem
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolvePathItem(tt.args.root, tt.args.ref, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolvePathItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolvePathItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveItemsWithBase(t *testing.T) {
	type args struct {
		root    interface{}
		ref     Ref
		options *ExpandOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *Items
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveItemsWithBase(tt.args.root, tt.args.ref, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveItemsWithBase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveItemsWithBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveItems(t *testing.T) {
	type args struct {
		root    interface{}
		ref     Ref
		options *ExpandOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *Items
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ResolveItems(tt.args.root, tt.args.ref, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResolveItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
