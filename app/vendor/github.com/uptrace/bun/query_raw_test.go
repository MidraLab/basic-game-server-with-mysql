package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestDB_Raw(t *testing.T) {
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
			if got := db.Raw(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.Raw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRawQuery(t *testing.T) {
	type args struct {
		db    *DB
		query string
		args  []interface{}
	}
	tests := []struct {
		name string
		args args
		want *RawQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRawQuery(tt.args.db, tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRawQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawQuery_Conn(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
		query     string
		args      []interface{}
	}
	type args struct {
		db IConn
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
			q := &RawQuery{
				baseQuery: tt.fields.baseQuery,
				query:     tt.fields.query,
				args:      tt.fields.args,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RawQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawQuery_Err(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
		query     string
		args      []interface{}
	}
	type args struct {
		err error
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
			q := &RawQuery{
				baseQuery: tt.fields.baseQuery,
				query:     tt.fields.query,
				args:      tt.fields.args,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RawQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawQuery_Exec(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
		query     string
		args      []interface{}
	}
	type args struct {
		ctx  context.Context
		dest []interface{}
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
			q := &RawQuery{
				baseQuery: tt.fields.baseQuery,
				query:     tt.fields.query,
				args:      tt.fields.args,
			}
			got, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RawQuery.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawQuery_Scan(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
		query     string
		args      []interface{}
	}
	type args struct {
		ctx  context.Context
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
			q := &RawQuery{
				baseQuery: tt.fields.baseQuery,
				query:     tt.fields.query,
				args:      tt.fields.args,
			}
			if err := q.Scan(tt.args.ctx, tt.args.dest...); (err != nil) != tt.wantErr {
				t.Errorf("RawQuery.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRawQuery_scanOrExec(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
		query     string
		args      []interface{}
	}
	type args struct {
		ctx     context.Context
		dest    []interface{}
		hasDest bool
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
			q := &RawQuery{
				baseQuery: tt.fields.baseQuery,
				query:     tt.fields.query,
				args:      tt.fields.args,
			}
			got, err := q.scanOrExec(tt.args.ctx, tt.args.dest, tt.args.hasDest)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawQuery.scanOrExec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RawQuery.scanOrExec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawQuery_AppendQuery(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
		query     string
		args      []interface{}
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &RawQuery{
				baseQuery: tt.fields.baseQuery,
				query:     tt.fields.query,
				args:      tt.fields.args,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("RawQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RawQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawQuery_Operation(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
		query     string
		args      []interface{}
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
			q := &RawQuery{
				baseQuery: tt.fields.baseQuery,
				query:     tt.fields.query,
				args:      tt.fields.args,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("RawQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}
