package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/dialect/feature"
	"github.com/uptrace/bun/internal"
	"github.com/uptrace/bun/schema"
)

func TestWithDiscardUnknownColumns(t *testing.T) {
	tests := []struct {
		name string
		want DBOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithDiscardUnknownColumns(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithDiscardUnknownColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDB(t *testing.T) {
	type args struct {
		sqldb   *sql.DB
		dialect schema.Dialect
		opts    []DBOption
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDB(tt.args.sqldb, tt.args.dialect, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_String(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
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
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.String(); got != tt.want {
				t.Errorf("DB.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_DBStats(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   DBStats
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.DBStats(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.DBStats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewValues(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ValuesQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewValues(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewMerge(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewMerge(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewSelect(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewSelect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewInsert(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewInsert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewUpdate(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewUpdate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewDelete(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewDelete(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewRaw(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RawQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewRaw(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewCreateTable(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewCreateTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewCreateTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewDropTable(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewDropTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewDropTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewCreateIndex(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewCreateIndex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewCreateIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewDropIndex(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewDropIndex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewDropIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewTruncateTable(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *TruncateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewTruncateTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewTruncateTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewAddColumn(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *AddColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewAddColumn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewAddColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_NewDropColumn(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.NewDropColumn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.NewDropColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_ResetModel(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx    context.Context
		models []interface{}
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
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if err := db.ResetModel(tt.args.ctx, tt.args.models...); (err != nil) != tt.wantErr {
				t.Errorf("DB.ResetModel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Dialect(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   schema.Dialect
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.Dialect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.Dialect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_ScanRows(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx  context.Context
		rows *sql.Rows
		dest []interface{}
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
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if err := db.ScanRows(tt.args.ctx, tt.args.rows, tt.args.dest...); (err != nil) != tt.wantErr {
				t.Errorf("DB.ScanRows() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_ScanRow(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx  context.Context
		rows *sql.Rows
		dest []interface{}
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
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if err := db.ScanRow(tt.args.ctx, tt.args.rows, tt.args.dest...); (err != nil) != tt.wantErr {
				t.Errorf("DB.ScanRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_AddQueryHook(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		hook QueryHook
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
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			db.AddQueryHook(tt.args.hook)
		})
	}
}

func TestDB_Table(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *schema.Table
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.Table(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_RegisterModel(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		models []interface{}
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
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			db.RegisterModel(tt.args.models...)
		})
	}
}

func TestDB_clone(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.clone(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_WithNamedArg(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		name  string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.WithNamedArg(tt.args.name, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.WithNamedArg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Formatter(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   schema.Formatter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.Formatter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.Formatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_UpdateFQN(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		alias  string
		column string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Ident
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.UpdateFQN(tt.args.alias, tt.args.column); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.UpdateFQN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_HasFeature(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		feat feature.Feature
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
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.HasFeature(tt.args.feat); got != tt.want {
				t.Errorf("DB.HasFeature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Exec(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    sql.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			got, err := db.Exec(tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_ExecContext(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    sql.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			got, err := db.ExecContext(tt.args.ctx, tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.ExecContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.ExecContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Query(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *sql.Rows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			got, err := db.Query(tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_QueryContext(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *sql.Rows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			got, err := db.QueryContext(tt.args.ctx, tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.QueryContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.QueryContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_QueryRow(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *sql.Row
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.QueryRow(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.QueryRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_QueryRowContext(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *sql.Row
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.QueryRowContext(tt.args.ctx, tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.QueryRowContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_format(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.format(tt.args.query, tt.args.args); got != tt.want {
				t.Errorf("DB.format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Conn(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Conn
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			got, err := db.Conn(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.Conn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_ExecContext(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    sql.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			got, err := c.ExecContext(tt.args.ctx, tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn.ExecContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.ExecContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_QueryContext(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *sql.Rows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			got, err := c.QueryContext(tt.args.ctx, tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn.QueryContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.QueryContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_QueryRowContext(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *sql.Row
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.QueryRowContext(tt.args.ctx, tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.QueryRowContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_Dialect(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   schema.Dialect
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.Dialect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.Dialect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewValues(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ValuesQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewValues(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewMerge(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewMerge(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewSelect(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewSelect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewInsert(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewInsert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewUpdate(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewUpdate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewDelete(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewDelete(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewRaw(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RawQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewRaw(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewCreateTable(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewCreateTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewCreateTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewDropTable(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewDropTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewDropTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewCreateIndex(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewCreateIndex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewCreateIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewDropIndex(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewDropIndex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewDropIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewTruncateTable(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   *TruncateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewTruncateTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewTruncateTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewAddColumn(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   *AddColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewAddColumn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewAddColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_NewDropColumn(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if got := c.NewDropColumn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.NewDropColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConn_RunInTx(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	type args struct {
		ctx  context.Context
		opts *sql.TxOptions
		fn   func(ctx context.Context, tx Tx) error
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
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			if err := c.RunInTx(tt.args.ctx, tt.args.opts, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("Conn.RunInTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConn_BeginTx(t *testing.T) {
	type fields struct {
		db   *DB
		Conn *sql.Conn
	}
	type args struct {
		ctx  context.Context
		opts *sql.TxOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Tx
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Conn{
				db:   tt.fields.db,
				Conn: tt.fields.Conn,
			}
			got, err := c.BeginTx(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Conn.BeginTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conn.BeginTx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_Prepare(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Stmt
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			got, err := db.Prepare(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.Prepare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.Prepare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_PrepareContext(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Stmt
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			got, err := db.PrepareContext(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.PrepareContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.PrepareContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_RunInTx(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx  context.Context
		opts *sql.TxOptions
		fn   func(ctx context.Context, tx Tx) error
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
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if err := db.RunInTx(tt.args.ctx, tt.args.opts, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("DB.RunInTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_Begin(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name    string
		fields  fields
		want    Tx
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			got, err := db.Begin()
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.Begin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.Begin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_BeginTx(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	type args struct {
		ctx  context.Context
		opts *sql.TxOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Tx
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			got, err := db.BeginTx(tt.args.ctx, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("DB.BeginTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.BeginTx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_Commit(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
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
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if err := tx.Commit(); (err != nil) != tt.wantErr {
				t.Errorf("Tx.Commit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTx_commitTX(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
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
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if err := tx.commitTX(); (err != nil) != tt.wantErr {
				t.Errorf("Tx.commitTX() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTx_commitSP(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
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
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if err := tx.commitSP(); (err != nil) != tt.wantErr {
				t.Errorf("Tx.commitSP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTx_Rollback(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
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
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if err := tx.Rollback(); (err != nil) != tt.wantErr {
				t.Errorf("Tx.Rollback() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTx_rollbackTX(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
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
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if err := tx.rollbackTX(); (err != nil) != tt.wantErr {
				t.Errorf("Tx.rollbackTX() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTx_rollbackSP(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
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
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if err := tx.rollbackSP(); (err != nil) != tt.wantErr {
				t.Errorf("Tx.rollbackSP() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTx_Exec(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    sql.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			got, err := tx.Exec(tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tx.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_ExecContext(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    sql.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			got, err := tx.ExecContext(tt.args.ctx, tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tx.ExecContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.ExecContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_Query(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *sql.Rows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			got, err := tx.Query(tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tx.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_QueryContext(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *sql.Rows
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			got, err := tx.QueryContext(tt.args.ctx, tt.args.query, tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tx.QueryContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.QueryContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_QueryRow(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *sql.Row
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.QueryRow(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.QueryRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_QueryRowContext(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	type args struct {
		ctx   context.Context
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *sql.Row
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.QueryRowContext(tt.args.ctx, tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.QueryRowContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_Begin(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name    string
		fields  fields
		want    Tx
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			got, err := tx.Begin()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tx.Begin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.Begin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_BeginTx(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	type args struct {
		ctx context.Context
		in1 *sql.TxOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Tx
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			got, err := tx.BeginTx(tt.args.ctx, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tx.BeginTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.BeginTx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_RunInTx(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	type args struct {
		ctx context.Context
		in1 *sql.TxOptions
		fn  func(ctx context.Context, tx Tx) error
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
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if err := tx.RunInTx(tt.args.ctx, tt.args.in1, tt.args.fn); (err != nil) != tt.wantErr {
				t.Errorf("Tx.RunInTx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTx_Dialect(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   schema.Dialect
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.Dialect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.Dialect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewValues(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *ValuesQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewValues(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewMerge(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   *MergeQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewMerge(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewMerge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewSelect(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   *SelectQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewSelect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewInsert(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewInsert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewUpdate(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   *UpdateQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewUpdate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewDelete(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   *DeleteQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewDelete(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewRaw(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RawQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewRaw(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewCreateTable(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewCreateTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewCreateTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewDropTable(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewDropTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewDropTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewCreateIndex(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   *CreateIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewCreateIndex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewCreateIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewDropIndex(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropIndexQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewDropIndex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewDropIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewTruncateTable(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   *TruncateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewTruncateTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewTruncateTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewAddColumn(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   *AddColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewAddColumn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewAddColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTx_NewDropColumn(t *testing.T) {
	type fields struct {
		ctx  context.Context
		db   *DB
		name string
		Tx   *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := Tx{
				ctx:  tt.fields.ctx,
				db:   tt.fields.db,
				name: tt.fields.name,
				Tx:   tt.fields.Tx,
			}
			if got := tx.NewDropColumn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tx.NewDropColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDB_makeQueryBytes(t *testing.T) {
	type fields struct {
		DB         *sql.DB
		dialect    schema.Dialect
		features   feature.Feature
		queryHooks []QueryHook
		fmter      schema.Formatter
		flags      internal.Flag
		stats      DBStats
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				DB:         tt.fields.DB,
				dialect:    tt.fields.dialect,
				features:   tt.fields.features,
				queryHooks: tt.fields.queryHooks,
				fmter:      tt.fields.fmter,
				flags:      tt.fields.flags,
				stats:      tt.fields.stats,
			}
			if got := db.makeQueryBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.makeQueryBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
