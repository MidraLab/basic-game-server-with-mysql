package spec

import (
	"reflect"
	"testing"
)

func TestOrderSchemaItems_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		items   OrderSchemaItems
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.items.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderSchemaItems.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderSchemaItems.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderSchemaItems_Len(t *testing.T) {
	tests := []struct {
		name  string
		items OrderSchemaItems
		want  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.items.Len(); got != tt.want {
				t.Errorf("OrderSchemaItems.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderSchemaItems_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name  string
		items OrderSchemaItems
		args  args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.items.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestOrderSchemaItems_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		items   OrderSchemaItems
		args    args
		wantRet bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := tt.items.Less(tt.args.i, tt.args.j); gotRet != tt.wantRet {
				t.Errorf("OrderSchemaItems.Less() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestSchemaProperties_ToOrderedSchemaItems(t *testing.T) {
	tests := []struct {
		name       string
		properties SchemaProperties
		want       OrderSchemaItems
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.properties.ToOrderedSchemaItems(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SchemaProperties.ToOrderedSchemaItems() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaProperties_MarshalJSON(t *testing.T) {
	tests := []struct {
		name       string
		properties SchemaProperties
		want       []byte
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.properties.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("SchemaProperties.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SchemaProperties.MarshalJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
