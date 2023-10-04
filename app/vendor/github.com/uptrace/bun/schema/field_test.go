package schema

import (
	"reflect"
	"testing"

	"github.com/uptrace/bun/internal/tagparser"
)

func TestField_String(t *testing.T) {
	type fields struct {
		StructField        reflect.StructField
		IsPtr              bool
		Tag                tagparser.Tag
		IndirectType       reflect.Type
		Index              []int
		Name               string
		SQLName            Safe
		GoName             string
		DiscoveredSQLType  string
		UserSQLType        string
		CreateTableSQLType string
		SQLDefault         string
		OnDelete           string
		OnUpdate           string
		IsPK               bool
		NotNull            bool
		NullZero           bool
		AutoIncrement      bool
		Identity           bool
		Append             AppenderFunc
		Scan               ScannerFunc
		IsZero             IsZeroerFunc
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
			f := &Field{
				StructField:        tt.fields.StructField,
				IsPtr:              tt.fields.IsPtr,
				Tag:                tt.fields.Tag,
				IndirectType:       tt.fields.IndirectType,
				Index:              tt.fields.Index,
				Name:               tt.fields.Name,
				SQLName:            tt.fields.SQLName,
				GoName:             tt.fields.GoName,
				DiscoveredSQLType:  tt.fields.DiscoveredSQLType,
				UserSQLType:        tt.fields.UserSQLType,
				CreateTableSQLType: tt.fields.CreateTableSQLType,
				SQLDefault:         tt.fields.SQLDefault,
				OnDelete:           tt.fields.OnDelete,
				OnUpdate:           tt.fields.OnUpdate,
				IsPK:               tt.fields.IsPK,
				NotNull:            tt.fields.NotNull,
				NullZero:           tt.fields.NullZero,
				AutoIncrement:      tt.fields.AutoIncrement,
				Identity:           tt.fields.Identity,
				Append:             tt.fields.Append,
				Scan:               tt.fields.Scan,
				IsZero:             tt.fields.IsZero,
			}
			if got := f.String(); got != tt.want {
				t.Errorf("Field.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_Clone(t *testing.T) {
	type fields struct {
		StructField        reflect.StructField
		IsPtr              bool
		Tag                tagparser.Tag
		IndirectType       reflect.Type
		Index              []int
		Name               string
		SQLName            Safe
		GoName             string
		DiscoveredSQLType  string
		UserSQLType        string
		CreateTableSQLType string
		SQLDefault         string
		OnDelete           string
		OnUpdate           string
		IsPK               bool
		NotNull            bool
		NullZero           bool
		AutoIncrement      bool
		Identity           bool
		Append             AppenderFunc
		Scan               ScannerFunc
		IsZero             IsZeroerFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   *Field
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Field{
				StructField:        tt.fields.StructField,
				IsPtr:              tt.fields.IsPtr,
				Tag:                tt.fields.Tag,
				IndirectType:       tt.fields.IndirectType,
				Index:              tt.fields.Index,
				Name:               tt.fields.Name,
				SQLName:            tt.fields.SQLName,
				GoName:             tt.fields.GoName,
				DiscoveredSQLType:  tt.fields.DiscoveredSQLType,
				UserSQLType:        tt.fields.UserSQLType,
				CreateTableSQLType: tt.fields.CreateTableSQLType,
				SQLDefault:         tt.fields.SQLDefault,
				OnDelete:           tt.fields.OnDelete,
				OnUpdate:           tt.fields.OnUpdate,
				IsPK:               tt.fields.IsPK,
				NotNull:            tt.fields.NotNull,
				NullZero:           tt.fields.NullZero,
				AutoIncrement:      tt.fields.AutoIncrement,
				Identity:           tt.fields.Identity,
				Append:             tt.fields.Append,
				Scan:               tt.fields.Scan,
				IsZero:             tt.fields.IsZero,
			}
			if got := f.Clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Field.Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_Value(t *testing.T) {
	type fields struct {
		StructField        reflect.StructField
		IsPtr              bool
		Tag                tagparser.Tag
		IndirectType       reflect.Type
		Index              []int
		Name               string
		SQLName            Safe
		GoName             string
		DiscoveredSQLType  string
		UserSQLType        string
		CreateTableSQLType string
		SQLDefault         string
		OnDelete           string
		OnUpdate           string
		IsPK               bool
		NotNull            bool
		NullZero           bool
		AutoIncrement      bool
		Identity           bool
		Append             AppenderFunc
		Scan               ScannerFunc
		IsZero             IsZeroerFunc
	}
	type args struct {
		strct reflect.Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Field{
				StructField:        tt.fields.StructField,
				IsPtr:              tt.fields.IsPtr,
				Tag:                tt.fields.Tag,
				IndirectType:       tt.fields.IndirectType,
				Index:              tt.fields.Index,
				Name:               tt.fields.Name,
				SQLName:            tt.fields.SQLName,
				GoName:             tt.fields.GoName,
				DiscoveredSQLType:  tt.fields.DiscoveredSQLType,
				UserSQLType:        tt.fields.UserSQLType,
				CreateTableSQLType: tt.fields.CreateTableSQLType,
				SQLDefault:         tt.fields.SQLDefault,
				OnDelete:           tt.fields.OnDelete,
				OnUpdate:           tt.fields.OnUpdate,
				IsPK:               tt.fields.IsPK,
				NotNull:            tt.fields.NotNull,
				NullZero:           tt.fields.NullZero,
				AutoIncrement:      tt.fields.AutoIncrement,
				Identity:           tt.fields.Identity,
				Append:             tt.fields.Append,
				Scan:               tt.fields.Scan,
				IsZero:             tt.fields.IsZero,
			}
			if got := f.Value(tt.args.strct); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Field.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_HasNilValue(t *testing.T) {
	type fields struct {
		StructField        reflect.StructField
		IsPtr              bool
		Tag                tagparser.Tag
		IndirectType       reflect.Type
		Index              []int
		Name               string
		SQLName            Safe
		GoName             string
		DiscoveredSQLType  string
		UserSQLType        string
		CreateTableSQLType string
		SQLDefault         string
		OnDelete           string
		OnUpdate           string
		IsPK               bool
		NotNull            bool
		NullZero           bool
		AutoIncrement      bool
		Identity           bool
		Append             AppenderFunc
		Scan               ScannerFunc
		IsZero             IsZeroerFunc
	}
	type args struct {
		v reflect.Value
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
			f := &Field{
				StructField:        tt.fields.StructField,
				IsPtr:              tt.fields.IsPtr,
				Tag:                tt.fields.Tag,
				IndirectType:       tt.fields.IndirectType,
				Index:              tt.fields.Index,
				Name:               tt.fields.Name,
				SQLName:            tt.fields.SQLName,
				GoName:             tt.fields.GoName,
				DiscoveredSQLType:  tt.fields.DiscoveredSQLType,
				UserSQLType:        tt.fields.UserSQLType,
				CreateTableSQLType: tt.fields.CreateTableSQLType,
				SQLDefault:         tt.fields.SQLDefault,
				OnDelete:           tt.fields.OnDelete,
				OnUpdate:           tt.fields.OnUpdate,
				IsPK:               tt.fields.IsPK,
				NotNull:            tt.fields.NotNull,
				NullZero:           tt.fields.NullZero,
				AutoIncrement:      tt.fields.AutoIncrement,
				Identity:           tt.fields.Identity,
				Append:             tt.fields.Append,
				Scan:               tt.fields.Scan,
				IsZero:             tt.fields.IsZero,
			}
			if got := f.HasNilValue(tt.args.v); got != tt.want {
				t.Errorf("Field.HasNilValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_HasZeroValue(t *testing.T) {
	type fields struct {
		StructField        reflect.StructField
		IsPtr              bool
		Tag                tagparser.Tag
		IndirectType       reflect.Type
		Index              []int
		Name               string
		SQLName            Safe
		GoName             string
		DiscoveredSQLType  string
		UserSQLType        string
		CreateTableSQLType string
		SQLDefault         string
		OnDelete           string
		OnUpdate           string
		IsPK               bool
		NotNull            bool
		NullZero           bool
		AutoIncrement      bool
		Identity           bool
		Append             AppenderFunc
		Scan               ScannerFunc
		IsZero             IsZeroerFunc
	}
	type args struct {
		v reflect.Value
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
			f := &Field{
				StructField:        tt.fields.StructField,
				IsPtr:              tt.fields.IsPtr,
				Tag:                tt.fields.Tag,
				IndirectType:       tt.fields.IndirectType,
				Index:              tt.fields.Index,
				Name:               tt.fields.Name,
				SQLName:            tt.fields.SQLName,
				GoName:             tt.fields.GoName,
				DiscoveredSQLType:  tt.fields.DiscoveredSQLType,
				UserSQLType:        tt.fields.UserSQLType,
				CreateTableSQLType: tt.fields.CreateTableSQLType,
				SQLDefault:         tt.fields.SQLDefault,
				OnDelete:           tt.fields.OnDelete,
				OnUpdate:           tt.fields.OnUpdate,
				IsPK:               tt.fields.IsPK,
				NotNull:            tt.fields.NotNull,
				NullZero:           tt.fields.NullZero,
				AutoIncrement:      tt.fields.AutoIncrement,
				Identity:           tt.fields.Identity,
				Append:             tt.fields.Append,
				Scan:               tt.fields.Scan,
				IsZero:             tt.fields.IsZero,
			}
			if got := f.HasZeroValue(tt.args.v); got != tt.want {
				t.Errorf("Field.HasZeroValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_AppendValue(t *testing.T) {
	type fields struct {
		StructField        reflect.StructField
		IsPtr              bool
		Tag                tagparser.Tag
		IndirectType       reflect.Type
		Index              []int
		Name               string
		SQLName            Safe
		GoName             string
		DiscoveredSQLType  string
		UserSQLType        string
		CreateTableSQLType string
		SQLDefault         string
		OnDelete           string
		OnUpdate           string
		IsPK               bool
		NotNull            bool
		NullZero           bool
		AutoIncrement      bool
		Identity           bool
		Append             AppenderFunc
		Scan               ScannerFunc
		IsZero             IsZeroerFunc
	}
	type args struct {
		fmter Formatter
		b     []byte
		strct reflect.Value
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Field{
				StructField:        tt.fields.StructField,
				IsPtr:              tt.fields.IsPtr,
				Tag:                tt.fields.Tag,
				IndirectType:       tt.fields.IndirectType,
				Index:              tt.fields.Index,
				Name:               tt.fields.Name,
				SQLName:            tt.fields.SQLName,
				GoName:             tt.fields.GoName,
				DiscoveredSQLType:  tt.fields.DiscoveredSQLType,
				UserSQLType:        tt.fields.UserSQLType,
				CreateTableSQLType: tt.fields.CreateTableSQLType,
				SQLDefault:         tt.fields.SQLDefault,
				OnDelete:           tt.fields.OnDelete,
				OnUpdate:           tt.fields.OnUpdate,
				IsPK:               tt.fields.IsPK,
				NotNull:            tt.fields.NotNull,
				NullZero:           tt.fields.NullZero,
				AutoIncrement:      tt.fields.AutoIncrement,
				Identity:           tt.fields.Identity,
				Append:             tt.fields.Append,
				Scan:               tt.fields.Scan,
				IsZero:             tt.fields.IsZero,
			}
			if got := f.AppendValue(tt.args.fmter, tt.args.b, tt.args.strct); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Field.AppendValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_ScanWithCheck(t *testing.T) {
	type fields struct {
		StructField        reflect.StructField
		IsPtr              bool
		Tag                tagparser.Tag
		IndirectType       reflect.Type
		Index              []int
		Name               string
		SQLName            Safe
		GoName             string
		DiscoveredSQLType  string
		UserSQLType        string
		CreateTableSQLType string
		SQLDefault         string
		OnDelete           string
		OnUpdate           string
		IsPK               bool
		NotNull            bool
		NullZero           bool
		AutoIncrement      bool
		Identity           bool
		Append             AppenderFunc
		Scan               ScannerFunc
		IsZero             IsZeroerFunc
	}
	type args struct {
		fv  reflect.Value
		src interface{}
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
			f := &Field{
				StructField:        tt.fields.StructField,
				IsPtr:              tt.fields.IsPtr,
				Tag:                tt.fields.Tag,
				IndirectType:       tt.fields.IndirectType,
				Index:              tt.fields.Index,
				Name:               tt.fields.Name,
				SQLName:            tt.fields.SQLName,
				GoName:             tt.fields.GoName,
				DiscoveredSQLType:  tt.fields.DiscoveredSQLType,
				UserSQLType:        tt.fields.UserSQLType,
				CreateTableSQLType: tt.fields.CreateTableSQLType,
				SQLDefault:         tt.fields.SQLDefault,
				OnDelete:           tt.fields.OnDelete,
				OnUpdate:           tt.fields.OnUpdate,
				IsPK:               tt.fields.IsPK,
				NotNull:            tt.fields.NotNull,
				NullZero:           tt.fields.NullZero,
				AutoIncrement:      tt.fields.AutoIncrement,
				Identity:           tt.fields.Identity,
				Append:             tt.fields.Append,
				Scan:               tt.fields.Scan,
				IsZero:             tt.fields.IsZero,
			}
			if err := f.ScanWithCheck(tt.args.fv, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Field.ScanWithCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestField_ScanValue(t *testing.T) {
	type fields struct {
		StructField        reflect.StructField
		IsPtr              bool
		Tag                tagparser.Tag
		IndirectType       reflect.Type
		Index              []int
		Name               string
		SQLName            Safe
		GoName             string
		DiscoveredSQLType  string
		UserSQLType        string
		CreateTableSQLType string
		SQLDefault         string
		OnDelete           string
		OnUpdate           string
		IsPK               bool
		NotNull            bool
		NullZero           bool
		AutoIncrement      bool
		Identity           bool
		Append             AppenderFunc
		Scan               ScannerFunc
		IsZero             IsZeroerFunc
	}
	type args struct {
		strct reflect.Value
		src   interface{}
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
			f := &Field{
				StructField:        tt.fields.StructField,
				IsPtr:              tt.fields.IsPtr,
				Tag:                tt.fields.Tag,
				IndirectType:       tt.fields.IndirectType,
				Index:              tt.fields.Index,
				Name:               tt.fields.Name,
				SQLName:            tt.fields.SQLName,
				GoName:             tt.fields.GoName,
				DiscoveredSQLType:  tt.fields.DiscoveredSQLType,
				UserSQLType:        tt.fields.UserSQLType,
				CreateTableSQLType: tt.fields.CreateTableSQLType,
				SQLDefault:         tt.fields.SQLDefault,
				OnDelete:           tt.fields.OnDelete,
				OnUpdate:           tt.fields.OnUpdate,
				IsPK:               tt.fields.IsPK,
				NotNull:            tt.fields.NotNull,
				NullZero:           tt.fields.NullZero,
				AutoIncrement:      tt.fields.AutoIncrement,
				Identity:           tt.fields.Identity,
				Append:             tt.fields.Append,
				Scan:               tt.fields.Scan,
				IsZero:             tt.fields.IsZero,
			}
			if err := f.ScanValue(tt.args.strct, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("Field.ScanValue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestField_SkipUpdate(t *testing.T) {
	type fields struct {
		StructField        reflect.StructField
		IsPtr              bool
		Tag                tagparser.Tag
		IndirectType       reflect.Type
		Index              []int
		Name               string
		SQLName            Safe
		GoName             string
		DiscoveredSQLType  string
		UserSQLType        string
		CreateTableSQLType string
		SQLDefault         string
		OnDelete           string
		OnUpdate           string
		IsPK               bool
		NotNull            bool
		NullZero           bool
		AutoIncrement      bool
		Identity           bool
		Append             AppenderFunc
		Scan               ScannerFunc
		IsZero             IsZeroerFunc
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
			f := &Field{
				StructField:        tt.fields.StructField,
				IsPtr:              tt.fields.IsPtr,
				Tag:                tt.fields.Tag,
				IndirectType:       tt.fields.IndirectType,
				Index:              tt.fields.Index,
				Name:               tt.fields.Name,
				SQLName:            tt.fields.SQLName,
				GoName:             tt.fields.GoName,
				DiscoveredSQLType:  tt.fields.DiscoveredSQLType,
				UserSQLType:        tt.fields.UserSQLType,
				CreateTableSQLType: tt.fields.CreateTableSQLType,
				SQLDefault:         tt.fields.SQLDefault,
				OnDelete:           tt.fields.OnDelete,
				OnUpdate:           tt.fields.OnUpdate,
				IsPK:               tt.fields.IsPK,
				NotNull:            tt.fields.NotNull,
				NullZero:           tt.fields.NullZero,
				AutoIncrement:      tt.fields.AutoIncrement,
				Identity:           tt.fields.Identity,
				Append:             tt.fields.Append,
				Scan:               tt.fields.Scan,
				IsZero:             tt.fields.IsZero,
			}
			if got := f.SkipUpdate(); got != tt.want {
				t.Errorf("Field.SkipUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_indexEqual(t *testing.T) {
	type args struct {
		ind1 []int
		ind2 []int
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
			if got := indexEqual(tt.args.ind1, tt.args.ind2); got != tt.want {
				t.Errorf("indexEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
