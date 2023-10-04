package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func Test_newMapModel(t *testing.T) {
	type args struct {
		db   *DB
		dest *map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want *mapModel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newMapModel(tt.args.db, tt.args.dest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMapModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapModel_Value(t *testing.T) {
	type fields struct {
		db           *DB
		dest         *map[string]interface{}
		m            map[string]interface{}
		rows         *sql.Rows
		columns      []string
		_columnTypes []*sql.ColumnType
		scanIndex    int
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mapModel{
				db:           tt.fields.db,
				dest:         tt.fields.dest,
				m:            tt.fields.m,
				rows:         tt.fields.rows,
				columns:      tt.fields.columns,
				_columnTypes: tt.fields._columnTypes,
				scanIndex:    tt.fields.scanIndex,
			}
			if got := m.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapModel.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapModel_ScanRows(t *testing.T) {
	type fields struct {
		db           *DB
		dest         *map[string]interface{}
		m            map[string]interface{}
		rows         *sql.Rows
		columns      []string
		_columnTypes []*sql.ColumnType
		scanIndex    int
	}
	type args struct {
		ctx  context.Context
		rows *sql.Rows
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mapModel{
				db:           tt.fields.db,
				dest:         tt.fields.dest,
				m:            tt.fields.m,
				rows:         tt.fields.rows,
				columns:      tt.fields.columns,
				_columnTypes: tt.fields._columnTypes,
				scanIndex:    tt.fields.scanIndex,
			}
			got, err := m.ScanRows(tt.args.ctx, tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapModel.ScanRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("mapModel.ScanRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapModel_Scan(t *testing.T) {
	type fields struct {
		db           *DB
		dest         *map[string]interface{}
		m            map[string]interface{}
		rows         *sql.Rows
		columns      []string
		_columnTypes []*sql.ColumnType
		scanIndex    int
	}
	type args struct {
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
			m := &mapModel{
				db:           tt.fields.db,
				dest:         tt.fields.dest,
				m:            tt.fields.m,
				rows:         tt.fields.rows,
				columns:      tt.fields.columns,
				_columnTypes: tt.fields._columnTypes,
				scanIndex:    tt.fields.scanIndex,
			}
			if err := m.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("mapModel.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mapModel_columnTypes(t *testing.T) {
	type fields struct {
		db           *DB
		dest         *map[string]interface{}
		m            map[string]interface{}
		rows         *sql.Rows
		columns      []string
		_columnTypes []*sql.ColumnType
		scanIndex    int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*sql.ColumnType
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mapModel{
				db:           tt.fields.db,
				dest:         tt.fields.dest,
				m:            tt.fields.m,
				rows:         tt.fields.rows,
				columns:      tt.fields.columns,
				_columnTypes: tt.fields._columnTypes,
				scanIndex:    tt.fields.scanIndex,
			}
			got, err := m.columnTypes()
			if (err != nil) != tt.wantErr {
				t.Errorf("mapModel.columnTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapModel.columnTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapModel_scanRaw(t *testing.T) {
	type fields struct {
		db           *DB
		dest         *map[string]interface{}
		m            map[string]interface{}
		rows         *sql.Rows
		columns      []string
		_columnTypes []*sql.ColumnType
		scanIndex    int
	}
	type args struct {
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
			m := &mapModel{
				db:           tt.fields.db,
				dest:         tt.fields.dest,
				m:            tt.fields.m,
				rows:         tt.fields.rows,
				columns:      tt.fields.columns,
				_columnTypes: tt.fields._columnTypes,
				scanIndex:    tt.fields.scanIndex,
			}
			if err := m.scanRaw(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("mapModel.scanRaw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mapModel_appendColumnsValues(t *testing.T) {
	type fields struct {
		db           *DB
		dest         *map[string]interface{}
		m            map[string]interface{}
		rows         *sql.Rows
		columns      []string
		_columnTypes []*sql.ColumnType
		scanIndex    int
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
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
			m := &mapModel{
				db:           tt.fields.db,
				dest:         tt.fields.dest,
				m:            tt.fields.m,
				rows:         tt.fields.rows,
				columns:      tt.fields.columns,
				_columnTypes: tt.fields._columnTypes,
				scanIndex:    tt.fields.scanIndex,
			}
			if got := m.appendColumnsValues(tt.args.fmter, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapModel.appendColumnsValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapModel_appendSet(t *testing.T) {
	type fields struct {
		db           *DB
		dest         *map[string]interface{}
		m            map[string]interface{}
		rows         *sql.Rows
		columns      []string
		_columnTypes []*sql.ColumnType
		scanIndex    int
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
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
			m := &mapModel{
				db:           tt.fields.db,
				dest:         tt.fields.dest,
				m:            tt.fields.m,
				rows:         tt.fields.rows,
				columns:      tt.fields.columns,
				_columnTypes: tt.fields._columnTypes,
				scanIndex:    tt.fields.scanIndex,
			}
			if got := m.appendSet(tt.args.fmter, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapModel.appendSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeDest(t *testing.T) {
	type args struct {
		v interface{}
		n int
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeDest(tt.args.v, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeDest() = %v, want %v", got, tt.want)
			}
		})
	}
}
