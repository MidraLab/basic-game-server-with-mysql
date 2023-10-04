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

func Test_baseQuery_DB(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.DB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.DB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_GetConn(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	tests := []struct {
		name   string
		fields fields
		want   IConn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.GetConn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.GetConn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_GetModel(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	tests := []struct {
		name   string
		fields fields
		want   Model
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.GetModel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.GetModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_GetTableName(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.GetTableName(); got != tt.want {
				t.Errorf("baseQuery.GetTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_setConn(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		db IConn
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			q.setConn(tt.args.db)
		})
	}
}

func Test_baseQuery_setModel(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		modeli interface{}
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			q.setModel(tt.args.modeli)
		})
	}
}

func Test_baseQuery_setErr(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		err error
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			q.setErr(tt.args.err)
		})
	}
}

func Test_baseQuery_getModel(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		dest []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Model
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.getModel(tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.getModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.getModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_beforeAppendModel(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		ctx   context.Context
		query Query
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if err := q.beforeAppendModel(tt.args.ctx, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.beforeAppendModel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_baseQuery_hasFeature(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		feature feature.Feature
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.hasFeature(tt.args.feature); got != tt.want {
				t.Errorf("baseQuery.hasFeature() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_checkSoftDelete(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if err := q.checkSoftDelete(); (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.checkSoftDelete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_baseQuery_whereDeleted(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			q.whereDeleted()
		})
	}
}

func Test_baseQuery_whereAllWithDeleted(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			q.whereAllWithDeleted()
		})
	}
}

func Test_baseQuery_isSoftDelete(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.isSoftDelete(); got != tt.want {
				t.Errorf("baseQuery.isSoftDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_addWith(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		name      string
		query     schema.QueryAppender
		recursive bool
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			q.addWith(tt.args.name, tt.args.query, tt.args.recursive)
		})
	}
}

func Test_baseQuery_appendWith(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.appendWith(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.appendWith() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.appendWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_appendCTE(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
		cte   withQuery
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.appendCTE(tt.args.fmter, tt.args.b, tt.args.cte)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.appendCTE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.appendCTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_appendSelectFromValues(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		fmter  schema.Formatter
		b      []byte
		cte    withQuery
		values *ValuesQuery
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.appendSelectFromValues(tt.args.fmter, tt.args.b, tt.args.cte, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.appendSelectFromValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.appendSelectFromValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_addTable(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		table schema.QueryWithArgs
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			q.addTable(tt.args.table)
		})
	}
}

func Test_baseQuery_addColumn(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		column schema.QueryWithArgs
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			q.addColumn(tt.args.column)
		})
	}
}

func Test_baseQuery_excludeColumn(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		columns []string
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			q.excludeColumn(tt.args.columns)
		})
	}
}

func Test_baseQuery__excludeColumn(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		column string
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q._excludeColumn(tt.args.column); got != tt.want {
				t.Errorf("baseQuery._excludeColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_modelHasTableName(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.modelHasTableName(); got != tt.want {
				t.Errorf("baseQuery.modelHasTableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_hasTables(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.hasTables(); got != tt.want {
				t.Errorf("baseQuery.hasTables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_appendTables(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.appendTables(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.appendTables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.appendTables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_appendTablesWithAlias(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.appendTablesWithAlias(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.appendTablesWithAlias() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.appendTablesWithAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery__appendTables(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		fmter     schema.Formatter
		b         []byte
		withAlias bool
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q._appendTables(tt.args.fmter, tt.args.b, tt.args.withAlias)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery._appendTables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery._appendTables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_appendFirstTable(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.appendFirstTable(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.appendFirstTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.appendFirstTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_appendFirstTableWithAlias(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.appendFirstTableWithAlias(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.appendFirstTableWithAlias() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.appendFirstTableWithAlias() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery__appendFirstTable(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		fmter     schema.Formatter
		b         []byte
		withAlias bool
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q._appendFirstTable(tt.args.fmter, tt.args.b, tt.args.withAlias)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery._appendFirstTable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery._appendFirstTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_hasMultiTables(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.hasMultiTables(); got != tt.want {
				t.Errorf("baseQuery.hasMultiTables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_appendOtherTables(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.appendOtherTables(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.appendOtherTables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.appendOtherTables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_appendColumns(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.appendColumns(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.appendColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.appendColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_getFields(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.getFields()
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.getFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.getFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_getDataFields(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.getDataFields()
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.getDataFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.getDataFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery__getFields(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		omitPK bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*schema.Field
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q._getFields(tt.args.omitPK)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery._getFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery._getFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_scan(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		ctx     context.Context
		iquery  Query
		query   string
		model   Model
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.scan(tt.args.ctx, tt.args.iquery, tt.args.query, tt.args.model, tt.args.hasDest)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.scan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.scan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_exec(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
	}
	type args struct {
		ctx    context.Context
		iquery Query
		query  string
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, err := q.exec(tt.args.ctx, tt.args.iquery, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("baseQuery.exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.exec() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_AppendNamedArg(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			got, got1 := q.AppendNamedArg(tt.args.fmter, tt.args.b, tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.AppendNamedArg() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("baseQuery.AppendNamedArg() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_baseQuery_Dialect(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.Dialect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.Dialect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewValues(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewValues(tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewSelect(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewSelect(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewSelect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewInsert(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewInsert(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewUpdate(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewUpdate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewDelete(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewDelete(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewRaw(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewRaw(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewCreateTable(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewCreateTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewCreateTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewDropTable(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewDropTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewDropTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewCreateIndex(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewCreateIndex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewCreateIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewDropIndex(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewDropIndex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewDropIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewTruncateTable(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewTruncateTable(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewTruncateTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewAddColumn(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewAddColumn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewAddColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_baseQuery_NewDropColumn(t *testing.T) {
	type fields struct {
		db             *DB
		conn           IConn
		model          Model
		err            error
		tableModel     TableModel
		table          *schema.Table
		with           []withQuery
		modelTableName schema.QueryWithArgs
		tables         []schema.QueryWithArgs
		columns        []schema.QueryWithArgs
		flags          internal.Flag
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
			q := &baseQuery{
				db:             tt.fields.db,
				conn:           tt.fields.conn,
				model:          tt.fields.model,
				err:            tt.fields.err,
				tableModel:     tt.fields.tableModel,
				table:          tt.fields.table,
				with:           tt.fields.with,
				modelTableName: tt.fields.modelTableName,
				tables:         tt.fields.tables,
				columns:        tt.fields.columns,
				flags:          tt.fields.flags,
			}
			if got := q.NewDropColumn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("baseQuery.NewDropColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendColumns(t *testing.T) {
	type args struct {
		b      []byte
		table  schema.Safe
		fields []*schema.Field
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := appendColumns(tt.args.b, tt.args.table, tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatterWithModel(t *testing.T) {
	type args struct {
		fmter schema.Formatter
		model schema.NamedArgAppender
	}
	tests := []struct {
		name string
		args args
		want schema.Formatter
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatterWithModel(tt.args.fmter, tt.args.model); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatterWithModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_whereBaseQuery_addWhere(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		where       []schema.QueryWithSep
		whereFields []*schema.Field
	}
	type args struct {
		where schema.QueryWithSep
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
			q := &whereBaseQuery{
				baseQuery:   tt.fields.baseQuery,
				where:       tt.fields.where,
				whereFields: tt.fields.whereFields,
			}
			q.addWhere(tt.args.where)
		})
	}
}

func Test_whereBaseQuery_addWhereGroup(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		where       []schema.QueryWithSep
		whereFields []*schema.Field
	}
	type args struct {
		sep   string
		where []schema.QueryWithSep
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
			q := &whereBaseQuery{
				baseQuery:   tt.fields.baseQuery,
				where:       tt.fields.where,
				whereFields: tt.fields.whereFields,
			}
			q.addWhereGroup(tt.args.sep, tt.args.where)
		})
	}
}

func Test_whereBaseQuery_addWhereCols(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		where       []schema.QueryWithSep
		whereFields []*schema.Field
	}
	type args struct {
		cols []string
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
			q := &whereBaseQuery{
				baseQuery:   tt.fields.baseQuery,
				where:       tt.fields.where,
				whereFields: tt.fields.whereFields,
			}
			q.addWhereCols(tt.args.cols)
		})
	}
}

func Test_whereBaseQuery_mustAppendWhere(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		where       []schema.QueryWithSep
		whereFields []*schema.Field
	}
	type args struct {
		fmter     schema.Formatter
		b         []byte
		withAlias bool
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
			q := &whereBaseQuery{
				baseQuery:   tt.fields.baseQuery,
				where:       tt.fields.where,
				whereFields: tt.fields.whereFields,
			}
			got, err := q.mustAppendWhere(tt.args.fmter, tt.args.b, tt.args.withAlias)
			if (err != nil) != tt.wantErr {
				t.Errorf("whereBaseQuery.mustAppendWhere() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("whereBaseQuery.mustAppendWhere() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_whereBaseQuery_appendWhere(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		where       []schema.QueryWithSep
		whereFields []*schema.Field
	}
	type args struct {
		fmter     schema.Formatter
		b         []byte
		withAlias bool
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
			q := &whereBaseQuery{
				baseQuery:   tt.fields.baseQuery,
				where:       tt.fields.where,
				whereFields: tt.fields.whereFields,
			}
			got, err := q.appendWhere(tt.args.fmter, tt.args.b, tt.args.withAlias)
			if (err != nil) != tt.wantErr {
				t.Errorf("whereBaseQuery.appendWhere() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("whereBaseQuery.appendWhere() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_appendWhere(t *testing.T) {
	type args struct {
		fmter schema.Formatter
		b     []byte
		where []schema.QueryWithSep
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := appendWhere(tt.args.fmter, tt.args.b, tt.args.where)
			if (err != nil) != tt.wantErr {
				t.Errorf("appendWhere() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("appendWhere() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_whereBaseQuery_appendWhereFields(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		where       []schema.QueryWithSep
		whereFields []*schema.Field
	}
	type args struct {
		fmter     schema.Formatter
		b         []byte
		fields    []*schema.Field
		withAlias bool
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
			q := &whereBaseQuery{
				baseQuery:   tt.fields.baseQuery,
				where:       tt.fields.where,
				whereFields: tt.fields.whereFields,
			}
			got, err := q.appendWhereFields(tt.args.fmter, tt.args.b, tt.args.fields, tt.args.withAlias)
			if (err != nil) != tt.wantErr {
				t.Errorf("whereBaseQuery.appendWhereFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("whereBaseQuery.appendWhereFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_whereBaseQuery_appendWhereStructFields(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		where       []schema.QueryWithSep
		whereFields []*schema.Field
	}
	type args struct {
		fmter     schema.Formatter
		b         []byte
		model     *structTableModel
		fields    []*schema.Field
		withAlias bool
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
			q := &whereBaseQuery{
				baseQuery:   tt.fields.baseQuery,
				where:       tt.fields.where,
				whereFields: tt.fields.whereFields,
			}
			got, err := q.appendWhereStructFields(tt.args.fmter, tt.args.b, tt.args.model, tt.args.fields, tt.args.withAlias)
			if (err != nil) != tt.wantErr {
				t.Errorf("whereBaseQuery.appendWhereStructFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("whereBaseQuery.appendWhereStructFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_whereBaseQuery_appendWhereSliceFields(t *testing.T) {
	type fields struct {
		baseQuery   baseQuery
		where       []schema.QueryWithSep
		whereFields []*schema.Field
	}
	type args struct {
		fmter     schema.Formatter
		b         []byte
		model     *sliceTableModel
		fields    []*schema.Field
		withAlias bool
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
			q := &whereBaseQuery{
				baseQuery:   tt.fields.baseQuery,
				where:       tt.fields.where,
				whereFields: tt.fields.whereFields,
			}
			got, err := q.appendWhereSliceFields(tt.args.fmter, tt.args.b, tt.args.model, tt.args.fields, tt.args.withAlias)
			if (err != nil) != tt.wantErr {
				t.Errorf("whereBaseQuery.appendWhereSliceFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("whereBaseQuery.appendWhereSliceFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_returningQuery_addReturning(t *testing.T) {
	type fields struct {
		returning       []schema.QueryWithArgs
		returningFields []*schema.Field
	}
	type args struct {
		ret schema.QueryWithArgs
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
			q := &returningQuery{
				returning:       tt.fields.returning,
				returningFields: tt.fields.returningFields,
			}
			q.addReturning(tt.args.ret)
		})
	}
}

func Test_returningQuery_addReturningField(t *testing.T) {
	type fields struct {
		returning       []schema.QueryWithArgs
		returningFields []*schema.Field
	}
	type args struct {
		field *schema.Field
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
			q := &returningQuery{
				returning:       tt.fields.returning,
				returningFields: tt.fields.returningFields,
			}
			q.addReturningField(tt.args.field)
		})
	}
}

func Test_returningQuery_appendReturning(t *testing.T) {
	type fields struct {
		returning       []schema.QueryWithArgs
		returningFields []*schema.Field
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
			q := &returningQuery{
				returning:       tt.fields.returning,
				returningFields: tt.fields.returningFields,
			}
			got, err := q.appendReturning(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("returningQuery.appendReturning() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("returningQuery.appendReturning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_returningQuery_appendOutput(t *testing.T) {
	type fields struct {
		returning       []schema.QueryWithArgs
		returningFields []*schema.Field
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
			q := &returningQuery{
				returning:       tt.fields.returning,
				returningFields: tt.fields.returningFields,
			}
			got, err := q.appendOutput(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("returningQuery.appendOutput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("returningQuery.appendOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_returningQuery__appendReturning(t *testing.T) {
	type fields struct {
		returning       []schema.QueryWithArgs
		returningFields []*schema.Field
	}
	type args struct {
		fmter schema.Formatter
		b     []byte
		table string
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
			q := &returningQuery{
				returning:       tt.fields.returning,
				returningFields: tt.fields.returningFields,
			}
			got, err := q._appendReturning(tt.args.fmter, tt.args.b, tt.args.table)
			if (err != nil) != tt.wantErr {
				t.Errorf("returningQuery._appendReturning() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("returningQuery._appendReturning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_returningQuery_hasReturning(t *testing.T) {
	type fields struct {
		returning       []schema.QueryWithArgs
		returningFields []*schema.Field
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
			q := &returningQuery{
				returning:       tt.fields.returning,
				returningFields: tt.fields.returningFields,
			}
			if got := q.hasReturning(); got != tt.want {
				t.Errorf("returningQuery.hasReturning() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_customValueQuery_addValue(t *testing.T) {
	type fields struct {
		modelValues map[string]schema.QueryWithArgs
		extraValues []columnValue
	}
	type args struct {
		table  *schema.Table
		column string
		value  string
		args   []interface{}
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
			q := &customValueQuery{
				modelValues: tt.fields.modelValues,
				extraValues: tt.fields.extraValues,
			}
			q.addValue(tt.args.table, tt.args.column, tt.args.value, tt.args.args)
		})
	}
}

func Test_setQuery_addSet(t *testing.T) {
	type fields struct {
		set []schema.QueryWithArgs
	}
	type args struct {
		set schema.QueryWithArgs
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
			q := &setQuery{
				set: tt.fields.set,
			}
			q.addSet(tt.args.set)
		})
	}
}

func Test_setQuery_appendSet(t *testing.T) {
	type fields struct {
		set []schema.QueryWithArgs
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
			q := setQuery{
				set: tt.fields.set,
			}
			got, err := q.appendSet(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("setQuery.appendSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setQuery.appendSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cascadeQuery_appendCascade(t *testing.T) {
	type fields struct {
		cascade  bool
		restrict bool
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
			q := cascadeQuery{
				cascade:  tt.fields.cascade,
				restrict: tt.fields.restrict,
			}
			if got := q.appendCascade(tt.args.fmter, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cascadeQuery.appendCascade() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idxHintsQuery_lazyUse(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	tests := []struct {
		name   string
		fields fields
		want   *indexHints
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			if got := ih.lazyUse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("idxHintsQuery.lazyUse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idxHintsQuery_lazyIgnore(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	tests := []struct {
		name   string
		fields fields
		want   *indexHints
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			if got := ih.lazyIgnore(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("idxHintsQuery.lazyIgnore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idxHintsQuery_lazyForce(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	tests := []struct {
		name   string
		fields fields
		want   *indexHints
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			if got := ih.lazyForce(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("idxHintsQuery.lazyForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idxHintsQuery_appendIndexes(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		hints   []schema.QueryWithArgs
		indexes []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []schema.QueryWithArgs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			if got := ih.appendIndexes(tt.args.hints, tt.args.indexes...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("idxHintsQuery.appendIndexes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idxHintsQuery_addUseIndex(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		indexes []string
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			ih.addUseIndex(tt.args.indexes...)
		})
	}
}

func Test_idxHintsQuery_addUseIndexForJoin(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		indexes []string
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			ih.addUseIndexForJoin(tt.args.indexes...)
		})
	}
}

func Test_idxHintsQuery_addUseIndexForOrderBy(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		indexes []string
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			ih.addUseIndexForOrderBy(tt.args.indexes...)
		})
	}
}

func Test_idxHintsQuery_addUseIndexForGroupBy(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		indexes []string
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			ih.addUseIndexForGroupBy(tt.args.indexes...)
		})
	}
}

func Test_idxHintsQuery_addIgnoreIndex(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		indexes []string
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			ih.addIgnoreIndex(tt.args.indexes...)
		})
	}
}

func Test_idxHintsQuery_addIgnoreIndexForJoin(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		indexes []string
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			ih.addIgnoreIndexForJoin(tt.args.indexes...)
		})
	}
}

func Test_idxHintsQuery_addIgnoreIndexForOrderBy(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		indexes []string
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			ih.addIgnoreIndexForOrderBy(tt.args.indexes...)
		})
	}
}

func Test_idxHintsQuery_addIgnoreIndexForGroupBy(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		indexes []string
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			ih.addIgnoreIndexForGroupBy(tt.args.indexes...)
		})
	}
}

func Test_idxHintsQuery_addForceIndex(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		indexes []string
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			ih.addForceIndex(tt.args.indexes...)
		})
	}
}

func Test_idxHintsQuery_addForceIndexForJoin(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		indexes []string
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			ih.addForceIndexForJoin(tt.args.indexes...)
		})
	}
}

func Test_idxHintsQuery_addForceIndexForOrderBy(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		indexes []string
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			ih.addForceIndexForOrderBy(tt.args.indexes...)
		})
	}
}

func Test_idxHintsQuery_addForceIndexForGroupBy(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		indexes []string
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			ih.addForceIndexForGroupBy(tt.args.indexes...)
		})
	}
}

func Test_idxHintsQuery_appendIndexHints(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			got, err := ih.appendIndexHints(tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("idxHintsQuery.appendIndexHints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("idxHintsQuery.appendIndexHints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idxHintsQuery_bufIndexHint(t *testing.T) {
	type fields struct {
		use    *indexHints
		ignore *indexHints
		force  *indexHints
	}
	type args struct {
		name  string
		hints []schema.QueryWithArgs
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
			ih := &idxHintsQuery{
				use:    tt.fields.use,
				ignore: tt.fields.ignore,
				force:  tt.fields.force,
			}
			got, err := ih.bufIndexHint(tt.args.name, tt.args.hints, tt.args.fmter, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("idxHintsQuery.bufIndexHint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("idxHintsQuery.bufIndexHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
