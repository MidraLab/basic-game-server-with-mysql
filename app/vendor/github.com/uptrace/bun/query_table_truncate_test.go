package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestNewTruncateTableQuery(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name string
		args args
		want *TruncateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTruncateTableQuery(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTruncateTableQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateTableQuery_Conn(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		cascadeQuery     cascadeQuery
		continueIdentity bool
	}
	type args struct {
		db IConn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TruncateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &TruncateTableQuery{
				baseQuery:        tt.fields.baseQuery,
				cascadeQuery:     tt.fields.cascadeQuery,
				continueIdentity: tt.fields.continueIdentity,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruncateTableQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateTableQuery_Model(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		cascadeQuery     cascadeQuery
		continueIdentity bool
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TruncateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &TruncateTableQuery{
				baseQuery:        tt.fields.baseQuery,
				cascadeQuery:     tt.fields.cascadeQuery,
				continueIdentity: tt.fields.continueIdentity,
			}
			if got := q.Model(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruncateTableQuery.Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateTableQuery_Err(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		cascadeQuery     cascadeQuery
		continueIdentity bool
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TruncateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &TruncateTableQuery{
				baseQuery:        tt.fields.baseQuery,
				cascadeQuery:     tt.fields.cascadeQuery,
				continueIdentity: tt.fields.continueIdentity,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruncateTableQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateTableQuery_Table(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		cascadeQuery     cascadeQuery
		continueIdentity bool
	}
	type args struct {
		tables []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TruncateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &TruncateTableQuery{
				baseQuery:        tt.fields.baseQuery,
				cascadeQuery:     tt.fields.cascadeQuery,
				continueIdentity: tt.fields.continueIdentity,
			}
			if got := q.Table(tt.args.tables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruncateTableQuery.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateTableQuery_TableExpr(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		cascadeQuery     cascadeQuery
		continueIdentity bool
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *TruncateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &TruncateTableQuery{
				baseQuery:        tt.fields.baseQuery,
				cascadeQuery:     tt.fields.cascadeQuery,
				continueIdentity: tt.fields.continueIdentity,
			}
			if got := q.TableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruncateTableQuery.TableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateTableQuery_ContinueIdentity(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		cascadeQuery     cascadeQuery
		continueIdentity bool
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
			q := &TruncateTableQuery{
				baseQuery:        tt.fields.baseQuery,
				cascadeQuery:     tt.fields.cascadeQuery,
				continueIdentity: tt.fields.continueIdentity,
			}
			if got := q.ContinueIdentity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruncateTableQuery.ContinueIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateTableQuery_Cascade(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		cascadeQuery     cascadeQuery
		continueIdentity bool
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
			q := &TruncateTableQuery{
				baseQuery:        tt.fields.baseQuery,
				cascadeQuery:     tt.fields.cascadeQuery,
				continueIdentity: tt.fields.continueIdentity,
			}
			if got := q.Cascade(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruncateTableQuery.Cascade() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateTableQuery_Restrict(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		cascadeQuery     cascadeQuery
		continueIdentity bool
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
			q := &TruncateTableQuery{
				baseQuery:        tt.fields.baseQuery,
				cascadeQuery:     tt.fields.cascadeQuery,
				continueIdentity: tt.fields.continueIdentity,
			}
			if got := q.Restrict(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruncateTableQuery.Restrict() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateTableQuery_Operation(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		cascadeQuery     cascadeQuery
		continueIdentity bool
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
			q := &TruncateTableQuery{
				baseQuery:        tt.fields.baseQuery,
				cascadeQuery:     tt.fields.cascadeQuery,
				continueIdentity: tt.fields.continueIdentity,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("TruncateTableQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateTableQuery_AppendQuery(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		cascadeQuery     cascadeQuery
		continueIdentity bool
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
			q := &TruncateTableQuery{
				baseQuery:        tt.fields.baseQuery,
				cascadeQuery:     tt.fields.cascadeQuery,
				continueIdentity: tt.fields.continueIdentity,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("TruncateTableQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruncateTableQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncateTableQuery_Exec(t *testing.T) {
	type fields struct {
		baseQuery        baseQuery
		cascadeQuery     cascadeQuery
		continueIdentity bool
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
			q := &TruncateTableQuery{
				baseQuery:        tt.fields.baseQuery,
				cascadeQuery:     tt.fields.cascadeQuery,
				continueIdentity: tt.fields.continueIdentity,
			}
			got, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("TruncateTableQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TruncateTableQuery.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
