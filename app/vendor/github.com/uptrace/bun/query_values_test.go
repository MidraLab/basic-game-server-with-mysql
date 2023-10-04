package bun

import (
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestNewValuesQuery(t *testing.T) {
	type args struct {
		db    *DB
		model interface{}
	}
	tests := []struct {
		name string
		args args
		want *ValuesQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewValuesQuery(tt.args.db, tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewValuesQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValuesQuery_Conn(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		customValueQuery customValueQuery
		withOrder        bool
	}
	type args struct {
		db IConn
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
			q := &ValuesQuery{
				baseQuery:        tt.fields.baseQuery,
				customValueQuery: tt.fields.customValueQuery,
				withOrder:        tt.fields.withOrder,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValuesQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValuesQuery_Err(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		customValueQuery customValueQuery
		withOrder        bool
	}
	type args struct {
		err error
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
			q := &ValuesQuery{
				baseQuery:        tt.fields.baseQuery,
				customValueQuery: tt.fields.customValueQuery,
				withOrder:        tt.fields.withOrder,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValuesQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValuesQuery_Column(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		customValueQuery customValueQuery
		withOrder        bool
	}
	type args struct {
		columns []string
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
			q := &ValuesQuery{
				baseQuery:        tt.fields.baseQuery,
				customValueQuery: tt.fields.customValueQuery,
				withOrder:        tt.fields.withOrder,
			}
			if got := q.Column(tt.args.columns...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValuesQuery.Column() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValuesQuery_Value(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		customValueQuery customValueQuery
		withOrder        bool
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
		want   *ValuesQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &ValuesQuery{
				baseQuery:        tt.fields.baseQuery,
				customValueQuery: tt.fields.customValueQuery,
				withOrder:        tt.fields.withOrder,
			}
			if got := q.Value(tt.args.column, tt.args.expr, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValuesQuery.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValuesQuery_WithOrder(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		customValueQuery customValueQuery
		withOrder        bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *ValuesQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &ValuesQuery{
				baseQuery:        tt.fields.baseQuery,
				customValueQuery: tt.fields.customValueQuery,
				withOrder:        tt.fields.withOrder,
			}
			if got := q.WithOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValuesQuery.WithOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValuesQuery_AppendNamedArg(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		customValueQuery customValueQuery
		withOrder        bool
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
		name  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &ValuesQuery{
				baseQuery:        tt.fields.baseQuery,
				customValueQuery: tt.fields.customValueQuery,
				withOrder:        tt.fields.withOrder,
			}
			got, got1 := q.AppendNamedArg(tt.args.fmter, tt.args.b, tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValuesQuery.AppendNamedArg() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ValuesQuery.AppendNamedArg() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestValuesQuery_AppendColumns(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		customValueQuery customValueQuery
		withOrder        bool
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
			q := &ValuesQuery{
				baseQuery:        tt.fields.baseQuery,
				customValueQuery: tt.fields.customValueQuery,
				withOrder:        tt.fields.withOrder,
			}
			got, err := q.AppendColumns(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValuesQuery.AppendColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValuesQuery.AppendColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValuesQuery_Operation(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		customValueQuery customValueQuery
		withOrder        bool
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
			q := &ValuesQuery{
				baseQuery:        tt.fields.baseQuery,
				customValueQuery: tt.fields.customValueQuery,
				withOrder:        tt.fields.withOrder,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("ValuesQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValuesQuery_AppendQuery(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		customValueQuery customValueQuery
		withOrder        bool
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
			q := &ValuesQuery{
				baseQuery:        tt.fields.baseQuery,
				customValueQuery: tt.fields.customValueQuery,
				withOrder:        tt.fields.withOrder,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValuesQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValuesQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValuesQuery_appendQuery(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		customValueQuery customValueQuery
		withOrder        bool
	}
	type args struct {
		fmter  schema.Formatter
		b      []byte
		fields []*schema.Field
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
			q := &ValuesQuery{
				baseQuery:        tt.fields.baseQuery,
				customValueQuery: tt.fields.customValueQuery,
				withOrder:        tt.fields.withOrder,
			}
			got, err := q.appendQuery(tt.args.fmter, tt.args.b, tt.args.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValuesQuery.appendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValuesQuery.appendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValuesQuery_appendValues(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		customValueQuery customValueQuery
		withOrder        bool
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
			q := &ValuesQuery{
				baseQuery:        tt.fields.baseQuery,
				customValueQuery: tt.fields.customValueQuery,
				withOrder:        tt.fields.withOrder,
			}
			got, err := q.appendValues(tt.args.fmter, tt.args.b, tt.args.fields, tt.args.strct)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValuesQuery.appendValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValuesQuery.appendValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
