package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestNewDropColumnQuery(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name string
		args args
		want *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDropColumnQuery(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDropColumnQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropColumnQuery_Conn(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
	}
	type args struct {
		db IConn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropColumnQuery{
				baseQuery: tt.fields.baseQuery,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropColumnQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropColumnQuery_Model(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropColumnQuery{
				baseQuery: tt.fields.baseQuery,
			}
			if got := q.Model(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropColumnQuery.Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropColumnQuery_Err(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropColumnQuery{
				baseQuery: tt.fields.baseQuery,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropColumnQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropColumnQuery_Apply(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
	}
	type args struct {
		fn func(*DropColumnQuery) *DropColumnQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropColumnQuery{
				baseQuery: tt.fields.baseQuery,
			}
			if got := q.Apply(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropColumnQuery.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropColumnQuery_Table(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
	}
	type args struct {
		tables []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropColumnQuery{
				baseQuery: tt.fields.baseQuery,
			}
			if got := q.Table(tt.args.tables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropColumnQuery.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropColumnQuery_TableExpr(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropColumnQuery{
				baseQuery: tt.fields.baseQuery,
			}
			if got := q.TableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropColumnQuery.TableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropColumnQuery_ModelTableExpr(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropColumnQuery{
				baseQuery: tt.fields.baseQuery,
			}
			if got := q.ModelTableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropColumnQuery.ModelTableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropColumnQuery_Column(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
	}
	type args struct {
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropColumnQuery{
				baseQuery: tt.fields.baseQuery,
			}
			if got := q.Column(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropColumnQuery.Column() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropColumnQuery_ColumnExpr(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropColumnQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropColumnQuery{
				baseQuery: tt.fields.baseQuery,
			}
			if got := q.ColumnExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropColumnQuery.ColumnExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropColumnQuery_Operation(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
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
			q := &DropColumnQuery{
				baseQuery: tt.fields.baseQuery,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("DropColumnQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropColumnQuery_AppendQuery(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
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
			q := &DropColumnQuery{
				baseQuery: tt.fields.baseQuery,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("DropColumnQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropColumnQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropColumnQuery_Exec(t *testing.T) {
	type fields struct {
		baseQuery baseQuery
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
			q := &DropColumnQuery{
				baseQuery: tt.fields.baseQuery,
			}
			got, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DropColumnQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropColumnQuery.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
