package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestNewDropTableQuery(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name string
		args args
		want *DropTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDropTableQuery(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDropTableQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_Conn(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
	}
	type args struct {
		db IConn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropTableQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_Model(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			if got := q.Model(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropTableQuery.Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_Err(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropTableQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_Table(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
	}
	type args struct {
		tables []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			if got := q.Table(tt.args.tables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropTableQuery.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_TableExpr(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			if got := q.TableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropTableQuery.TableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_ModelTableExpr(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DropTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			if got := q.ModelTableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropTableQuery.ModelTableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_IfExists(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
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
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			if got := q.IfExists(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropTableQuery.IfExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_Cascade(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
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
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			if got := q.Cascade(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropTableQuery.Cascade() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_Restrict(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
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
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			if got := q.Restrict(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropTableQuery.Restrict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_Operation(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
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
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("DropTableQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_AppendQuery(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
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
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("DropTableQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropTableQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_Exec(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
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
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			got, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("DropTableQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DropTableQuery.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropTableQuery_beforeDropTableHook(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
	}
	type args struct {
		ctx context.Context
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
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			if err := q.beforeDropTableHook(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("DropTableQuery.beforeDropTableHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDropTableQuery_afterDropTableHook(t *testing.T) {
	type fields struct {
		baseQuery    baseQuery
		cascadeQuery cascadeQuery
		ifExists     bool
	}
	type args struct {
		ctx context.Context
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
			q := &DropTableQuery{
				baseQuery:    tt.fields.baseQuery,
				cascadeQuery: tt.fields.cascadeQuery,
				ifExists:     tt.fields.ifExists,
			}
			if err := q.afterDropTableHook(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("DropTableQuery.afterDropTableHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
