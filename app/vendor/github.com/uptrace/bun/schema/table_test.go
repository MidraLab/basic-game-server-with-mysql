package schema

import (
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/uptrace/bun/internal"
)

func TestSetTableNameInflector(t *testing.T) {
	type args struct {
		fn func(string) string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetTableNameInflector(tt.args.fn)
		})
	}
}

func Test_newTable(t *testing.T) {
	type args struct {
		dialect Dialect
		typ     reflect.Type
	}
	tests := []struct {
		name string
		args args
		want *Table
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newTable(tt.args.dialect, tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_init1(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			tr.init1()
		})
	}
}

func TestTable_init2(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			tr.init2()
		})
	}
}

func TestTable_setName(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		name string
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			tr.setName(tt.args.name)
		})
	}
}

func TestTable_String(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.String(); got != tt.want {
				t.Errorf("Table.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_CheckPKs(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if err := tr.CheckPKs(); (err != nil) != tt.wantErr {
				t.Errorf("Table.CheckPKs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTable_addField(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		field *Field
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			tr.addField(tt.args.field)
		})
	}
}

func TestTable_removeField(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		field *Field
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			tr.removeField(tt.args.field)
		})
	}
}

func TestTable_fieldWithLock(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Field
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.fieldWithLock(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table.fieldWithLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_HasField(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.HasField(tt.args.name); got != tt.want {
				t.Errorf("Table.HasField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_Field(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Field
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			got, err := tr.Field(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Table.Field() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table.Field() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_fieldByGoName(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Field
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.fieldByGoName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table.fieldByGoName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_initFields(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			tr.initFields()
		})
	}
}

func TestTable_addFields(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		typ    reflect.Type
		prefix string
		index  []int
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			tr.addFields(tt.args.typ, tt.args.prefix, tt.args.index)
		})
	}
}

func TestTable_processBaseModelField(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		f reflect.StructField
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			tr.processBaseModelField(tt.args.f)
		})
	}
}

func TestTable_newField(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		f      reflect.StructField
		prefix string
		index  []int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Field
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.newField(tt.args.f, tt.args.prefix, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table.newField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_initRelations(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			tr.initRelations()
		})
	}
}

func TestTable_tryRelation(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		field *Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.tryRelation(tt.args.field); got != tt.want {
				t.Errorf("Table.tryRelation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_initRelation(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		field *Field
		rel   string
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			tr.initRelation(tt.args.field, tt.args.rel)
		})
	}
}

func TestTable_addRelation(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		rel *Relation
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			tr.addRelation(tt.args.rel)
		})
	}
}

func TestTable_belongsToRelation(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		field *Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Relation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.belongsToRelation(tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table.belongsToRelation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_hasOneRelation(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		field *Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Relation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.hasOneRelation(tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table.hasOneRelation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_hasManyRelation(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		field *Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Relation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.hasManyRelation(tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table.hasManyRelation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_m2mRelation(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		field *Field
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Relation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.m2mRelation(tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table.m2mRelation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_inlineFields(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		field *Field
		seen  map[reflect.Type]struct{}
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			tr.inlineFields(tt.args.field, tt.args.seen)
		})
	}
}

func TestTable_Dialect(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	tests := []struct {
		name   string
		fields fields
		want   Dialect
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.Dialect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table.Dialect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_HasBeforeAppendModelHook(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.HasBeforeAppendModelHook(); got != tt.want {
				t.Errorf("Table.HasBeforeAppendModelHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_HasBeforeScanHook(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.HasBeforeScanHook(); got != tt.want {
				t.Errorf("Table.HasBeforeScanHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_HasAfterScanHook(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.HasAfterScanHook(); got != tt.want {
				t.Errorf("Table.HasAfterScanHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_HasBeforeScanRowHook(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.HasBeforeScanRowHook(); got != tt.want {
				t.Errorf("Table.HasBeforeScanRowHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_HasAfterScanRowHook(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
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
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.HasAfterScanRowHook(); got != tt.want {
				t.Errorf("Table.HasAfterScanRowHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_AppendNamedArg(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		fmter Formatter
		b     []byte
		name  string
		strct reflect.Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			got, got1 := tr.AppendNamedArg(tt.args.fmter, tt.args.b, tt.args.name, tt.args.strct)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Table.AppendNamedArg() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Table.AppendNamedArg() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestTable_quoteTableName(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Safe
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.quoteTableName(tt.args.s); got != tt.want {
				t.Errorf("Table.quoteTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTable_quoteIdent(t *testing.T) {
	type fields struct {
		dialect               Dialect
		Type                  reflect.Type
		ZeroValue             reflect.Value
		ZeroIface             interface{}
		TypeName              string
		ModelName             string
		Name                  string
		SQLName               Safe
		SQLNameForSelects     Safe
		Alias                 string
		SQLAlias              Safe
		Fields                []*Field
		PKs                   []*Field
		DataFields            []*Field
		fieldsMapMu           sync.RWMutex
		FieldMap              map[string]*Field
		Relations             map[string]*Relation
		Unique                map[string][]*Field
		SoftDeleteField       *Field
		UpdateSoftDeleteField func(fv reflect.Value, tm time.Time) error
		allFields             []*Field
		flags                 internal.Flag
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Safe
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Table{
				dialect:               tt.fields.dialect,
				Type:                  tt.fields.Type,
				ZeroValue:             tt.fields.ZeroValue,
				ZeroIface:             tt.fields.ZeroIface,
				TypeName:              tt.fields.TypeName,
				ModelName:             tt.fields.ModelName,
				Name:                  tt.fields.Name,
				SQLName:               tt.fields.SQLName,
				SQLNameForSelects:     tt.fields.SQLNameForSelects,
				Alias:                 tt.fields.Alias,
				SQLAlias:              tt.fields.SQLAlias,
				Fields:                tt.fields.Fields,
				PKs:                   tt.fields.PKs,
				DataFields:            tt.fields.DataFields,
				fieldsMapMu:           tt.fields.fieldsMapMu,
				FieldMap:              tt.fields.FieldMap,
				Relations:             tt.fields.Relations,
				Unique:                tt.fields.Unique,
				SoftDeleteField:       tt.fields.SoftDeleteField,
				UpdateSoftDeleteField: tt.fields.UpdateSoftDeleteField,
				allFields:             tt.fields.allFields,
				flags:                 tt.fields.flags,
			}
			if got := tr.quoteIdent(tt.args.s); got != tt.want {
				t.Errorf("Table.quoteIdent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isKnownTableOption(t *testing.T) {
	type args struct {
		name string
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
			if got := isKnownTableOption(tt.args.name); got != tt.want {
				t.Errorf("isKnownTableOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isKnownFieldOption(t *testing.T) {
	type args struct {
		name string
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
			if got := isKnownFieldOption(tt.args.name); got != tt.want {
				t.Errorf("isKnownFieldOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isKnownFKRule(t *testing.T) {
	type args struct {
		name string
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
			if got := isKnownFKRule(tt.args.name); got != tt.want {
				t.Errorf("isKnownFKRule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeField(t *testing.T) {
	type args struct {
		fields []*Field
		field  *Field
	}
	tests := []struct {
		name string
		args args
		want []*Field
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeField(tt.args.fields, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeField() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseRelationJoin(t *testing.T) {
	type args struct {
		join []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseRelationJoin(tt.args.join)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRelationJoin() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parseRelationJoin() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_softDeleteFieldUpdater(t *testing.T) {
	type args struct {
		field *Field
	}
	tests := []struct {
		name string
		args args
		want func(fv reflect.Value, tm time.Time) error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := softDeleteFieldUpdater(tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("softDeleteFieldUpdater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_softDeleteFieldUpdaterFallback(t *testing.T) {
	type args struct {
		field *Field
	}
	tests := []struct {
		name string
		args args
		want func(fv reflect.Value, tm time.Time) error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := softDeleteFieldUpdaterFallback(tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("softDeleteFieldUpdaterFallback() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_withIndex(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := withIndex(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("withIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
