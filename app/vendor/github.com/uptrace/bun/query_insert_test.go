package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestNewInsertQuery(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name string
		args args
		want *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewInsertQuery(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewInsertQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Conn(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		db IConn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Model(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Model(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Err(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Apply(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		fn func(*InsertQuery) *InsertQuery
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Apply(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_With(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		name  string
		query schema.QueryAppender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.With(tt.args.name, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.With() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_WithRecursive(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		name  string
		query schema.QueryAppender
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.WithRecursive(tt.args.name, tt.args.query); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.WithRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Table(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		tables []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Table(tt.args.tables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_TableExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.TableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.TableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_ModelTableExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.ModelTableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.ModelTableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Column(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Column(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Column() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_ColumnExpr(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.ColumnExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.ColumnExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_ExcludeColumn(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		columns []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.ExcludeColumn(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.ExcludeColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Value(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		column string
		expr   string
		args   []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Value(tt.args.column, tt.args.expr, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Where(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Where(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Where() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_WhereOr(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.WhereOr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.WhereOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Returning(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Returning(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Returning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Ignore(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Ignore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Ignore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Replace(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Replace(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Operation(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("InsertQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_AppendQuery(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_appendColumnsValues(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		fmter      schema.Formatter
		b          []byte
		skipOutput bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			got, err := q.appendColumnsValues(tt.args.fmter, tt.args.b, tt.args.skipOutput)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertQuery.appendColumnsValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.appendColumnsValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_appendStructValues(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		fmter  schema.Formatter
		b      []byte
		fields []*schema.Field
		strct  reflect.Value
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			got, err := q.appendStructValues(tt.args.fmter, tt.args.b, tt.args.fields, tt.args.strct)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertQuery.appendStructValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.appendStructValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_appendSliceValues(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		fmter  schema.Formatter
		b      []byte
		fields []*schema.Field
		slice  reflect.Value
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			got, err := q.appendSliceValues(tt.args.fmter, tt.args.b, tt.args.fields, tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertQuery.appendSliceValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.appendSliceValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_getFields(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*schema.Field
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			got, err := q.getFields()
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertQuery.getFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.getFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_appendFields(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		fmter  schema.Formatter
		b      []byte
		fields []*schema.Field
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.appendFields(tt.args.fmter, tt.args.b, tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.appendFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_On(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		s    string
		args []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.On(tt.args.s, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.On() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Set(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *InsertQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.Set(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_appendOn(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			got, err := q.appendOn(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertQuery.appendOn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.appendOn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_onConflictDoUpdate(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.onConflictDoUpdate(); got != tt.want {
				t.Errorf("InsertQuery.onConflictDoUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_onDuplicateKeyUpdate(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.onDuplicateKeyUpdate(); got != tt.want {
				t.Errorf("InsertQuery.onDuplicateKeyUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_appendSetExcluded(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		b      []byte
		fields []*schema.Field
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.appendSetExcluded(tt.args.b, tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.appendSetExcluded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_appendSetValues(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		b      []byte
		fields []*schema.Field
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.appendSetValues(tt.args.b, tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.appendSetValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_Scan(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if err := q.Scan(tt.args.ctx, tt.args.dest...); (err != nil) != tt.wantErr {
				t.Errorf("InsertQuery.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsertQuery_Exec(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			got, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_scanOrExec(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			got, err := q.scanOrExec(tt.args.ctx, tt.args.dest, tt.args.hasDest)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertQuery.scanOrExec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertQuery.scanOrExec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertQuery_beforeInsertHook(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if err := q.beforeInsertHook(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("InsertQuery.beforeInsertHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsertQuery_afterInsertHook(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if err := q.afterInsertHook(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("InsertQuery.afterInsertHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsertQuery_tryLastInsertID(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
	}
	type args struct {
		res  sql.Result
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if err := q.tryLastInsertID(tt.args.res, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("InsertQuery.tryLastInsertID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsertQuery_String(t *testing.T) {
	type fields struct {
		whereBaseQuery   whereBaseQuery
		returningQuery   returningQuery
		customValueQuery customValueQuery
		on               schema.QueryWithArgs
		setQuery         setQuery
		ignore           bool
		replace          bool
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
			q := &InsertQuery{
				whereBaseQuery:   tt.fields.whereBaseQuery,
				returningQuery:   tt.fields.returningQuery,
				customValueQuery: tt.fields.customValueQuery,
				on:               tt.fields.on,
				setQuery:         tt.fields.setQuery,
				ignore:           tt.fields.ignore,
				replace:          tt.fields.replace,
			}
			if got := q.String(); got != tt.want {
				t.Errorf("InsertQuery.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
