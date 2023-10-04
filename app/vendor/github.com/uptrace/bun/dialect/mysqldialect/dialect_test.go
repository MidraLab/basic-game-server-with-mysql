package mysqldialect

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/uptrace/bun/dialect"
	"github.com/uptrace/bun/dialect/feature"
	"github.com/uptrace/bun/schema"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Dialect
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_Init(t *testing.T) {
	type fields struct {
		BaseDialect schema.BaseDialect
		tables      *schema.Tables
		features    feature.Feature
	}
	type args struct {
		db *sql.DB
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
			d := &Dialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			d.Init(tt.args.db)
		})
	}
}

func Test_cleanupVersion(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanupVersion(tt.args.s); got != tt.want {
				t.Errorf("cleanupVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_Name(t *testing.T) {
	type fields struct {
		BaseDialect schema.BaseDialect
		tables      *schema.Tables
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
			d := &Dialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.Name(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dialect.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_Features(t *testing.T) {
	type fields struct {
		BaseDialect schema.BaseDialect
		tables      *schema.Tables
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
			d := &Dialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.Features(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dialect.Features() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_Tables(t *testing.T) {
	type fields struct {
		BaseDialect schema.BaseDialect
		tables      *schema.Tables
		features    feature.Feature
	}
	tests := []struct {
		name   string
		fields fields
		want   *schema.Tables
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.Tables(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dialect.Tables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_OnTable(t *testing.T) {
	type fields struct {
		BaseDialect schema.BaseDialect
		tables      *schema.Tables
		features    feature.Feature
	}
	type args struct {
		table *schema.Table
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
			d := &Dialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			d.OnTable(tt.args.table)
		})
	}
}

func TestDialect_IdentQuote(t *testing.T) {
	type fields struct {
		BaseDialect schema.BaseDialect
		tables      *schema.Tables
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
			d := &Dialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.IdentQuote(); got != tt.want {
				t.Errorf("Dialect.IdentQuote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_AppendTime(t *testing.T) {
	type fields struct {
		BaseDialect schema.BaseDialect
		tables      *schema.Tables
		features    feature.Feature
	}
	type args struct {
		b  []byte
		tm time.Time
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
			d := &Dialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.AppendTime(tt.args.b, tt.args.tm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dialect.AppendTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_AppendString(t *testing.T) {
	type fields struct {
		BaseDialect schema.BaseDialect
		tables      *schema.Tables
		features    feature.Feature
	}
	type args struct {
		b []byte
		s string
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
			d := &Dialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.AppendString(tt.args.b, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dialect.AppendString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_AppendBytes(t *testing.T) {
	type fields struct {
		BaseDialect schema.BaseDialect
		tables      *schema.Tables
		features    feature.Feature
	}
	type args struct {
		b  []byte
		bs []byte
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
			d := &Dialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.AppendBytes(tt.args.b, tt.args.bs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dialect.AppendBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_AppendJSON(t *testing.T) {
	type fields struct {
		BaseDialect schema.BaseDialect
		tables      *schema.Tables
		features    feature.Feature
	}
	type args struct {
		b     []byte
		jsonb []byte
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
			d := &Dialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.AppendJSON(tt.args.b, tt.args.jsonb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dialect.AppendJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDialect_DefaultVarcharLen(t *testing.T) {
	type fields struct {
		BaseDialect schema.BaseDialect
		tables      *schema.Tables
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
			d := &Dialect{
				BaseDialect: tt.fields.BaseDialect,
				tables:      tt.fields.tables,
				features:    tt.fields.features,
			}
			if got := d.DefaultVarcharLen(); got != tt.want {
				t.Errorf("Dialect.DefaultVarcharLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlType(t *testing.T) {
	type args struct {
		field *schema.Field
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sqlType(tt.args.field); got != tt.want {
				t.Errorf("sqlType() = %v, want %v", got, tt.want)
			}
		})
	}
}
