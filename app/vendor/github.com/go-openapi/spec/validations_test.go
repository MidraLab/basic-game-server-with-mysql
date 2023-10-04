package spec

import (
	"reflect"
	"testing"
)

func TestCommonValidations_SetValidations(t *testing.T) {
	type fields struct {
		Maximum          *float64
		ExclusiveMaximum bool
		Minimum          *float64
		ExclusiveMinimum bool
		MaxLength        *int64
		MinLength        *int64
		Pattern          string
		MaxItems         *int64
		MinItems         *int64
		UniqueItems      bool
		MultipleOf       *float64
		Enum             []interface{}
	}
	type args struct {
		val SchemaValidations
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
			v := &CommonValidations{
				Maximum:          tt.fields.Maximum,
				ExclusiveMaximum: tt.fields.ExclusiveMaximum,
				Minimum:          tt.fields.Minimum,
				ExclusiveMinimum: tt.fields.ExclusiveMinimum,
				MaxLength:        tt.fields.MaxLength,
				MinLength:        tt.fields.MinLength,
				Pattern:          tt.fields.Pattern,
				MaxItems:         tt.fields.MaxItems,
				MinItems:         tt.fields.MinItems,
				UniqueItems:      tt.fields.UniqueItems,
				MultipleOf:       tt.fields.MultipleOf,
				Enum:             tt.fields.Enum,
			}
			v.SetValidations(tt.args.val)
		})
	}
}

func Test_clearedValidations_apply(t *testing.T) {
	type args struct {
		cbs []func(string, interface{})
	}
	tests := []struct {
		name string
		c    clearedValidations
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.apply(tt.args.cbs)
		})
	}
}

func TestCommonValidations_ClearNumberValidations(t *testing.T) {
	type fields struct {
		Maximum          *float64
		ExclusiveMaximum bool
		Minimum          *float64
		ExclusiveMinimum bool
		MaxLength        *int64
		MinLength        *int64
		Pattern          string
		MaxItems         *int64
		MinItems         *int64
		UniqueItems      bool
		MultipleOf       *float64
		Enum             []interface{}
	}
	type args struct {
		cbs []func(string, interface{})
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
			v := &CommonValidations{
				Maximum:          tt.fields.Maximum,
				ExclusiveMaximum: tt.fields.ExclusiveMaximum,
				Minimum:          tt.fields.Minimum,
				ExclusiveMinimum: tt.fields.ExclusiveMinimum,
				MaxLength:        tt.fields.MaxLength,
				MinLength:        tt.fields.MinLength,
				Pattern:          tt.fields.Pattern,
				MaxItems:         tt.fields.MaxItems,
				MinItems:         tt.fields.MinItems,
				UniqueItems:      tt.fields.UniqueItems,
				MultipleOf:       tt.fields.MultipleOf,
				Enum:             tt.fields.Enum,
			}
			v.ClearNumberValidations(tt.args.cbs...)
		})
	}
}

func TestCommonValidations_ClearStringValidations(t *testing.T) {
	type fields struct {
		Maximum          *float64
		ExclusiveMaximum bool
		Minimum          *float64
		ExclusiveMinimum bool
		MaxLength        *int64
		MinLength        *int64
		Pattern          string
		MaxItems         *int64
		MinItems         *int64
		UniqueItems      bool
		MultipleOf       *float64
		Enum             []interface{}
	}
	type args struct {
		cbs []func(string, interface{})
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
			v := &CommonValidations{
				Maximum:          tt.fields.Maximum,
				ExclusiveMaximum: tt.fields.ExclusiveMaximum,
				Minimum:          tt.fields.Minimum,
				ExclusiveMinimum: tt.fields.ExclusiveMinimum,
				MaxLength:        tt.fields.MaxLength,
				MinLength:        tt.fields.MinLength,
				Pattern:          tt.fields.Pattern,
				MaxItems:         tt.fields.MaxItems,
				MinItems:         tt.fields.MinItems,
				UniqueItems:      tt.fields.UniqueItems,
				MultipleOf:       tt.fields.MultipleOf,
				Enum:             tt.fields.Enum,
			}
			v.ClearStringValidations(tt.args.cbs...)
		})
	}
}

func TestCommonValidations_ClearArrayValidations(t *testing.T) {
	type fields struct {
		Maximum          *float64
		ExclusiveMaximum bool
		Minimum          *float64
		ExclusiveMinimum bool
		MaxLength        *int64
		MinLength        *int64
		Pattern          string
		MaxItems         *int64
		MinItems         *int64
		UniqueItems      bool
		MultipleOf       *float64
		Enum             []interface{}
	}
	type args struct {
		cbs []func(string, interface{})
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
			v := &CommonValidations{
				Maximum:          tt.fields.Maximum,
				ExclusiveMaximum: tt.fields.ExclusiveMaximum,
				Minimum:          tt.fields.Minimum,
				ExclusiveMinimum: tt.fields.ExclusiveMinimum,
				MaxLength:        tt.fields.MaxLength,
				MinLength:        tt.fields.MinLength,
				Pattern:          tt.fields.Pattern,
				MaxItems:         tt.fields.MaxItems,
				MinItems:         tt.fields.MinItems,
				UniqueItems:      tt.fields.UniqueItems,
				MultipleOf:       tt.fields.MultipleOf,
				Enum:             tt.fields.Enum,
			}
			v.ClearArrayValidations(tt.args.cbs...)
		})
	}
}

func TestCommonValidations_Validations(t *testing.T) {
	type fields struct {
		Maximum          *float64
		ExclusiveMaximum bool
		Minimum          *float64
		ExclusiveMinimum bool
		MaxLength        *int64
		MinLength        *int64
		Pattern          string
		MaxItems         *int64
		MinItems         *int64
		UniqueItems      bool
		MultipleOf       *float64
		Enum             []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   SchemaValidations
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := CommonValidations{
				Maximum:          tt.fields.Maximum,
				ExclusiveMaximum: tt.fields.ExclusiveMaximum,
				Minimum:          tt.fields.Minimum,
				ExclusiveMinimum: tt.fields.ExclusiveMinimum,
				MaxLength:        tt.fields.MaxLength,
				MinLength:        tt.fields.MinLength,
				Pattern:          tt.fields.Pattern,
				MaxItems:         tt.fields.MaxItems,
				MinItems:         tt.fields.MinItems,
				UniqueItems:      tt.fields.UniqueItems,
				MultipleOf:       tt.fields.MultipleOf,
				Enum:             tt.fields.Enum,
			}
			if got := v.Validations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommonValidations.Validations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommonValidations_HasNumberValidations(t *testing.T) {
	type fields struct {
		Maximum          *float64
		ExclusiveMaximum bool
		Minimum          *float64
		ExclusiveMinimum bool
		MaxLength        *int64
		MinLength        *int64
		Pattern          string
		MaxItems         *int64
		MinItems         *int64
		UniqueItems      bool
		MultipleOf       *float64
		Enum             []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := CommonValidations{
				Maximum:          tt.fields.Maximum,
				ExclusiveMaximum: tt.fields.ExclusiveMaximum,
				Minimum:          tt.fields.Minimum,
				ExclusiveMinimum: tt.fields.ExclusiveMinimum,
				MaxLength:        tt.fields.MaxLength,
				MinLength:        tt.fields.MinLength,
				Pattern:          tt.fields.Pattern,
				MaxItems:         tt.fields.MaxItems,
				MinItems:         tt.fields.MinItems,
				UniqueItems:      tt.fields.UniqueItems,
				MultipleOf:       tt.fields.MultipleOf,
				Enum:             tt.fields.Enum,
			}
			if got := v.HasNumberValidations(); got != tt.want {
				t.Errorf("CommonValidations.HasNumberValidations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommonValidations_HasStringValidations(t *testing.T) {
	type fields struct {
		Maximum          *float64
		ExclusiveMaximum bool
		Minimum          *float64
		ExclusiveMinimum bool
		MaxLength        *int64
		MinLength        *int64
		Pattern          string
		MaxItems         *int64
		MinItems         *int64
		UniqueItems      bool
		MultipleOf       *float64
		Enum             []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := CommonValidations{
				Maximum:          tt.fields.Maximum,
				ExclusiveMaximum: tt.fields.ExclusiveMaximum,
				Minimum:          tt.fields.Minimum,
				ExclusiveMinimum: tt.fields.ExclusiveMinimum,
				MaxLength:        tt.fields.MaxLength,
				MinLength:        tt.fields.MinLength,
				Pattern:          tt.fields.Pattern,
				MaxItems:         tt.fields.MaxItems,
				MinItems:         tt.fields.MinItems,
				UniqueItems:      tt.fields.UniqueItems,
				MultipleOf:       tt.fields.MultipleOf,
				Enum:             tt.fields.Enum,
			}
			if got := v.HasStringValidations(); got != tt.want {
				t.Errorf("CommonValidations.HasStringValidations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommonValidations_HasArrayValidations(t *testing.T) {
	type fields struct {
		Maximum          *float64
		ExclusiveMaximum bool
		Minimum          *float64
		ExclusiveMinimum bool
		MaxLength        *int64
		MinLength        *int64
		Pattern          string
		MaxItems         *int64
		MinItems         *int64
		UniqueItems      bool
		MultipleOf       *float64
		Enum             []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := CommonValidations{
				Maximum:          tt.fields.Maximum,
				ExclusiveMaximum: tt.fields.ExclusiveMaximum,
				Minimum:          tt.fields.Minimum,
				ExclusiveMinimum: tt.fields.ExclusiveMinimum,
				MaxLength:        tt.fields.MaxLength,
				MinLength:        tt.fields.MinLength,
				Pattern:          tt.fields.Pattern,
				MaxItems:         tt.fields.MaxItems,
				MinItems:         tt.fields.MinItems,
				UniqueItems:      tt.fields.UniqueItems,
				MultipleOf:       tt.fields.MultipleOf,
				Enum:             tt.fields.Enum,
			}
			if got := v.HasArrayValidations(); got != tt.want {
				t.Errorf("CommonValidations.HasArrayValidations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommonValidations_HasEnum(t *testing.T) {
	type fields struct {
		Maximum          *float64
		ExclusiveMaximum bool
		Minimum          *float64
		ExclusiveMinimum bool
		MaxLength        *int64
		MinLength        *int64
		Pattern          string
		MaxItems         *int64
		MinItems         *int64
		UniqueItems      bool
		MultipleOf       *float64
		Enum             []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := CommonValidations{
				Maximum:          tt.fields.Maximum,
				ExclusiveMaximum: tt.fields.ExclusiveMaximum,
				Minimum:          tt.fields.Minimum,
				ExclusiveMinimum: tt.fields.ExclusiveMinimum,
				MaxLength:        tt.fields.MaxLength,
				MinLength:        tt.fields.MinLength,
				Pattern:          tt.fields.Pattern,
				MaxItems:         tt.fields.MaxItems,
				MinItems:         tt.fields.MinItems,
				UniqueItems:      tt.fields.UniqueItems,
				MultipleOf:       tt.fields.MultipleOf,
				Enum:             tt.fields.Enum,
			}
			if got := v.HasEnum(); got != tt.want {
				t.Errorf("CommonValidations.HasEnum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaValidations_HasObjectValidations(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		PatternProperties SchemaProperties
		MaxProperties     *int64
		MinProperties     *int64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := SchemaValidations{
				CommonValidations: tt.fields.CommonValidations,
				PatternProperties: tt.fields.PatternProperties,
				MaxProperties:     tt.fields.MaxProperties,
				MinProperties:     tt.fields.MinProperties,
			}
			if got := v.HasObjectValidations(); got != tt.want {
				t.Errorf("SchemaValidations.HasObjectValidations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaValidations_SetValidations(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		PatternProperties SchemaProperties
		MaxProperties     *int64
		MinProperties     *int64
	}
	type args struct {
		val SchemaValidations
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
			v := &SchemaValidations{
				CommonValidations: tt.fields.CommonValidations,
				PatternProperties: tt.fields.PatternProperties,
				MaxProperties:     tt.fields.MaxProperties,
				MinProperties:     tt.fields.MinProperties,
			}
			v.SetValidations(tt.args.val)
		})
	}
}

func TestSchemaValidations_Validations(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		PatternProperties SchemaProperties
		MaxProperties     *int64
		MinProperties     *int64
	}
	tests := []struct {
		name   string
		fields fields
		want   SchemaValidations
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := SchemaValidations{
				CommonValidations: tt.fields.CommonValidations,
				PatternProperties: tt.fields.PatternProperties,
				MaxProperties:     tt.fields.MaxProperties,
				MinProperties:     tt.fields.MinProperties,
			}
			if got := v.Validations(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SchemaValidations.Validations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSchemaValidations_ClearObjectValidations(t *testing.T) {
	type fields struct {
		CommonValidations CommonValidations
		PatternProperties SchemaProperties
		MaxProperties     *int64
		MinProperties     *int64
	}
	type args struct {
		cbs []func(string, interface{})
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
			v := &SchemaValidations{
				CommonValidations: tt.fields.CommonValidations,
				PatternProperties: tt.fields.PatternProperties,
				MaxProperties:     tt.fields.MaxProperties,
				MinProperties:     tt.fields.MinProperties,
			}
			v.ClearObjectValidations(tt.args.cbs...)
		})
	}
}
