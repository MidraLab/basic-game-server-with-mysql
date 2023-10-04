package schema

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/uptrace/bun/dialect"
	"github.com/uptrace/bun/dialect/feature"
)

func TestBaseDialect_AppendUint32(t *testing.T) {
	type args struct {
		b []byte
		n uint32
	}
	tests := []struct {
		name string
		b    BaseDialect
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseDialect{}
			if got := b.AppendUint32(tt.args.b, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseDialect.AppendUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseDialect_AppendUint64(t *testing.T) {
	type args struct {
		b []byte
		n uint64
	}
	tests := []struct {
		name string
		b    BaseDialect
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseDialect{}
			if got := b.AppendUint64(tt.args.b, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseDialect.AppendUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseDialect_AppendTime(t *testing.T) {
	type args struct {
		b  []byte
		tm time.Time
	}
	tests := []struct {
		name string
		b    BaseDialect
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseDialect{}
			if got := b.AppendTime(tt.args.b, tt.args.tm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseDialect.AppendTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseDialect_AppendString(t *testing.T) {
	type args struct {
		b []byte
		s string
	}
	tests := []struct {
		name string
		b    BaseDialect
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseDialect{}
			if got := b.AppendString(tt.args.b, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseDialect.AppendString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseDialect_AppendBytes(t *testing.T) {
	type args struct {
		b  []byte
		bs []byte
	}
	tests := []struct {
		name string
		b    BaseDialect
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseDialect{}
			if got := b.AppendBytes(tt.args.b, tt.args.bs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseDialect.AppendBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseDialect_AppendJSON(t *testing.T) {
	type args struct {
		b     []byte
		jsonb []byte
	}
	tests := []struct {
		name string
		b    BaseDialect
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseDialect{}
			if got := b.AppendJSON(tt.args.b, tt.args.jsonb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseDialect.AppendJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseDialect_AppendBool(t *testing.T) {
	type args struct {
		b []byte
		v bool
	}
	tests := []struct {
		name string
		b    BaseDialect
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BaseDialect{}
			if got := b.AppendBool(tt.args.b, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseDialect.AppendBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newNopDialect(t *testing.T) {
	tests := []struct {
		name string
		want *nopDialect
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newNopDialect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNopDialect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nopDialect_Init(t *testing.T) {
	type fields struct {
		BaseDialect BaseDialect
		tables      *Tables
		features    feature.Feature
	}
	type args struct {
		in0 *sql.DB
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
			d := &nopDialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			d.Init(tt.args.in0)
		})
	}
}

func Test_nopDialect_Name(t *testing.T) {
	type fields struct {
		BaseDialect BaseDialect
		tables      *Tables
		features    feature.Feature
	}
	tests := []struct {
		name   string
		fields fields
		want   dialect.Name
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &nopDialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.Name(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nopDialect.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nopDialect_Features(t *testing.T) {
	type fields struct {
		BaseDialect BaseDialect
		tables      *Tables
		features    feature.Feature
	}
	tests := []struct {
		name   string
		fields fields
		want   feature.Feature
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &nopDialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.Features(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nopDialect.Features() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nopDialect_Tables(t *testing.T) {
	type fields struct {
		BaseDialect BaseDialect
		tables      *Tables
		features    feature.Feature
	}
	tests := []struct {
		name   string
		fields fields
		want   *Tables
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &nopDialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.Tables(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nopDialect.Tables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nopDialect_OnField(t *testing.T) {
	type fields struct {
		BaseDialect BaseDialect
		tables      *Tables
		features    feature.Feature
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
			d := &nopDialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			d.OnField(tt.args.field)
		})
	}
}

func Test_nopDialect_OnTable(t *testing.T) {
	type fields struct {
		BaseDialect BaseDialect
		tables      *Tables
		features    feature.Feature
	}
	type args struct {
		table *Table
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
			d := &nopDialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			d.OnTable(tt.args.table)
		})
	}
}

func Test_nopDialect_IdentQuote(t *testing.T) {
	type fields struct {
		BaseDialect BaseDialect
		tables      *Tables
		features    feature.Feature
	}
	tests := []struct {
		name   string
		fields fields
		want   byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &nopDialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.IdentQuote(); got != tt.want {
				t.Errorf("nopDialect.IdentQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nopDialect_DefaultVarcharLen(t *testing.T) {
	type fields struct {
		BaseDialect BaseDialect
		tables      *Tables
		features    feature.Feature
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &nopDialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.DefaultVarcharLen(); got != tt.want {
				t.Errorf("nopDialect.DefaultVarcharLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
