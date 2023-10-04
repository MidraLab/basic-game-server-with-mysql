package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func TestNewCreateTableQuery(t *testing.T) {
	type args struct {
		db *DB
	}
	tests := []struct {
		name string
		args args
		want *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCreateTableQuery(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateTableQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_Conn(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		db IConn
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.Conn(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.Conn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_Model(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		model interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.Model(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.Model() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_Err(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.Err(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.Err() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_Table(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		tables []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.Table(tt.args.tables...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_TableExpr(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.TableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.TableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_ModelTableExpr(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.ModelTableExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.ModelTableExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_ColumnExpr(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.ColumnExpr(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.ColumnExpr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_Temp(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.Temp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.Temp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_IfNotExists(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.IfNotExists(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.IfNotExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_Varchar(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.Varchar(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.Varchar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_ForeignKey(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.ForeignKey(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.ForeignKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_PartitionBy(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.PartitionBy(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.PartitionBy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_TableSpace(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		tablespace string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateTableQuery
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.TableSpace(tt.args.tablespace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.TableSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_WithForeignKeys(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.WithForeignKeys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.WithForeignKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_Operation(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.Operation(); got != tt.want {
				t.Errorf("CreateTableQuery.Operation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_AppendQuery(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			got, err := q.AppendQuery(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTableQuery.AppendQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.AppendQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_appendSQLType(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		b     []byte
		field *schema.Field
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.appendSQLType(tt.args.b, tt.args.field); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.appendSQLType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_appendUniqueConstraints(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.appendUniqueConstraints(tt.args.fmter, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.appendUniqueConstraints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_appendUniqueConstraint(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		fmter  schema.Formatter
		b      []byte
		name   string
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.appendUniqueConstraint(tt.args.fmter, tt.args.b, tt.args.name, tt.args.fields...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.appendUniqueConstraint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_appendFKConstraints(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			got, err := q.appendFKConstraints(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTableQuery.appendFKConstraints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.appendFKConstraints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_appendPKConstraint(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
	}
	type args struct {
		b   []byte
		pks []*schema.Field
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if got := q.appendPKConstraint(tt.args.b, tt.args.pks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.appendPKConstraint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_Exec(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			got, err := q.Exec(tt.args.ctx, tt.args.dest...)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTableQuery.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTableQuery.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTableQuery_beforeCreateTableHook(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if err := q.beforeCreateTableHook(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CreateTableQuery.beforeCreateTableHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateTableQuery_afterCreateTableHook(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		temp        bool
		ifNotExists bool
		varchar     int
		fks         []schema.QueryWithArgs
		partitionBy schema.QueryWithArgs
		tablespace  schema.QueryWithArgs
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
			q := &CreateTableQuery{
				baseQuery:   tt.fields.baseQuery,
				temp:        tt.fields.temp,
				ifNotExists: tt.fields.ifNotExists,
				varchar:     tt.fields.varchar,
				fks:         tt.fields.fks,
				partitionBy: tt.fields.partitionBy,
				tablespace:  tt.fields.tablespace,
			}
			if err := q.afterCreateTableHook(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CreateTableQuery.afterCreateTableHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
