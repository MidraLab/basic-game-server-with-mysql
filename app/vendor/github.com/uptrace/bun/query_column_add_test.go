package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestNewAddColumnQuery(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name string
		args args
		want *AddColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAddColumnQuery(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAddColumnQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddColumnQuery_Conn(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		ifNotExists bool
	}
	type args struct {
		db IConn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AddColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &AddColumnQuery{
				baseQuery:   tt.fields.baseQuery,
				ifNotExists: tt.fields.ifNotExists,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddColumnQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddColumnQuery_Model(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		ifNotExists bool
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AddColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &AddColumnQuery{
				baseQuery:   tt.fields.baseQuery,
				ifNotExists: tt.fields.ifNotExists,
			}
			if got := q.Model(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddColumnQuery.Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddColumnQuery_Err(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		ifNotExists bool
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AddColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &AddColumnQuery{
				baseQuery:   tt.fields.baseQuery,
				ifNotExists: tt.fields.ifNotExists,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddColumnQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddColumnQuery_Apply(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		ifNotExists bool
	}
	type args struct {
		fn func(*AddColumnQuery) *AddColumnQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AddColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &AddColumnQuery{
				baseQuery:   tt.fields.baseQuery,
				ifNotExists: tt.fields.ifNotExists,
			}
			if got := q.Apply(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddColumnQuery.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddColumnQuery_Table(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		ifNotExists bool
	}
	type args struct {
		tables []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AddColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &AddColumnQuery{
				baseQuery:   tt.fields.baseQuery,
				ifNotExists: tt.fields.ifNotExists,
			}
			if got := q.Table(tt.args.tables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddColumnQuery.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddColumnQuery_TableExpr(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		ifNotExists bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AddColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &AddColumnQuery{
				baseQuery:   tt.fields.baseQuery,
				ifNotExists: tt.fields.ifNotExists,
			}
			if got := q.TableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddColumnQuery.TableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddColumnQuery_ModelTableExpr(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		ifNotExists bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AddColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &AddColumnQuery{
				baseQuery:   tt.fields.baseQuery,
				ifNotExists: tt.fields.ifNotExists,
			}
			if got := q.ModelTableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddColumnQuery.ModelTableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddColumnQuery_ColumnExpr(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		ifNotExists bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AddColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &AddColumnQuery{
				baseQuery:   tt.fields.baseQuery,
				ifNotExists: tt.fields.ifNotExists,
			}
			if got := q.ColumnExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddColumnQuery.ColumnExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddColumnQuery_IfNotExists(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		ifNotExists bool
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
			q := &AddColumnQuery{
				baseQuery:   tt.fields.baseQuery,
				ifNotExists: tt.fields.ifNotExists,
			}
			if got := q.IfNotExists(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddColumnQuery.IfNotExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddColumnQuery_Operation(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		ifNotExists bool
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
			q := &AddColumnQuery{
				baseQuery:   tt.fields.baseQuery,
				ifNotExists: tt.fields.ifNotExists,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("AddColumnQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddColumnQuery_AppendQuery(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		ifNotExists bool
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
			q := &AddColumnQuery{
				baseQuery:   tt.fields.baseQuery,
				ifNotExists: tt.fields.ifNotExists,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddColumnQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddColumnQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddColumnQuery_Exec(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		ifNotExists bool
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
			q := &AddColumnQuery{
				baseQuery:   tt.fields.baseQuery,
				ifNotExists: tt.fields.ifNotExists,
			}
			got, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddColumnQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddColumnQuery.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
