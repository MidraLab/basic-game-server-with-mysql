package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/uptrace/bun/schema"
)

func Test_newStructTableModel(t *testing.T) {
	type args struct {
		db    *DB
		dest  interface{}
		table *schema.Table
	}
	tests := []struct {
		name string
		args args
		want *structTableModel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newStructTableModel(tt.args.db, tt.args.dest, tt.args.table); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStructTableModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newStructTableModelValue(t *testing.T) {
	type args struct {
		db   *DB
		dest interface{}
		v    reflect.Value
	}
	tests := []struct {
		name string
		args args
		want *structTableModel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newStructTableModelValue(tt.args.db, tt.args.dest, tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStructTableModelValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_Value(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if got := m.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structTableModel.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_Table(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *schema.Table
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if got := m.Table(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structTableModel.Table() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_Relation(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	tests := []struct {
		name   string
		fields fields
		want   *schema.Relation
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if got := m.Relation(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structTableModel.Relation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_initStruct(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if err := m.initStruct(); (err != nil) != tt.wantErr {
				t.Errorf("structTableModel.initStruct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_structTableModel_mountJoins(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			m.mountJoins()
		})
	}
}

func Test_structTableModel_BeforeAppendModel(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if err := m.BeforeAppendModel(tt.args.ctx, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("structTableModel.BeforeAppendModel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_structTableModel_BeforeScanRow(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if err := m.BeforeScanRow(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("structTableModel.BeforeScanRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_structTableModel_AfterScanRow(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if err := m.AfterScanRow(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("structTableModel.AfterScanRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_structTableModel_getJoin(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *relationJoin
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if got := m.getJoin(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structTableModel.getJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_getJoins(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	tests := []struct {
		name   string
		fields fields
		want   []relationJoin
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if got := m.getJoins(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structTableModel.getJoins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_addJoin(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	type args struct {
		j relationJoin
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *relationJoin
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if got := m.addJoin(tt.args.j); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structTableModel.addJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_join(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *relationJoin
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if got := m.join(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structTableModel.join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel__join(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	type args struct {
		bind reflect.Value
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *relationJoin
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if got := m._join(tt.args.bind, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structTableModel._join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_rootValue(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	tests := []struct {
		name   string
		fields fields
		want   reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if got := m.rootValue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structTableModel.rootValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_parentIndex(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if got := m.parentIndex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structTableModel.parentIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_mount(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	type args struct {
		host reflect.Value
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			m.mount(tt.args.host)
		})
	}
}

func Test_structTableModel_updateSoftDeleteField(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	type args struct {
		tm time.Time
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if err := m.updateSoftDeleteField(tt.args.tm); (err != nil) != tt.wantErr {
				t.Errorf("structTableModel.updateSoftDeleteField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_structTableModel_ScanRows(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			got, err := m.ScanRows(tt.args.ctx, tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("structTableModel.ScanRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("structTableModel.ScanRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_ScanRow(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	type args struct {
		ctx  context.Context
		rows *sql.Rows
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if err := m.ScanRow(tt.args.ctx, tt.args.rows); (err != nil) != tt.wantErr {
				t.Errorf("structTableModel.ScanRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_structTableModel_scanRow(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	type args struct {
		ctx  context.Context
		rows *sql.Rows
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if err := m.scanRow(tt.args.ctx, tt.args.rows, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("structTableModel.scanRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_structTableModel_Scan(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if err := m.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("structTableModel.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_structTableModel_ScanColumn(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	type args struct {
		column string
		src    interface{}
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if err := m.ScanColumn(tt.args.column, tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("structTableModel.ScanColumn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_structTableModel_scanColumn(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
	}
	type args struct {
		column string
		src    interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			got, err := m.scanColumn(tt.args.column, tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("structTableModel.scanColumn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("structTableModel.scanColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_isNil(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			if got := m.isNil(); got != tt.want {
				t.Errorf("structTableModel.isNil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_structTableModel_AppendNamedArg(t *testing.T) {
	type fields struct {
		db            *DB
		table         *schema.Table
		rel           *schema.Relation
		joins         []relationJoin
		dest          interface{}
		root          reflect.Value
		index         []int
		strct         reflect.Value
		structInited  bool
		structInitErr error
		columns       []string
		scanIndex     int
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
			m := &structTableModel{
				db:            tt.fields.db,
				table:         tt.fields.table,
				rel:           tt.fields.rel,
				joins:         tt.fields.joins,
				dest:          tt.fields.dest,
				root:          tt.fields.root,
				index:         tt.fields.index,
				strct:         tt.fields.strct,
				structInited:  tt.fields.structInited,
				structInitErr: tt.fields.structInitErr,
				columns:       tt.fields.columns,
				scanIndex:     tt.fields.scanIndex,
			}
			got, got1 := m.AppendNamedArg(tt.args.fmter, tt.args.b, tt.args.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("structTableModel.AppendNamedArg() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("structTableModel.AppendNamedArg() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_unquote(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unquote(tt.args.s); got != tt.want {
				t.Errorf("unquote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitColumn(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitColumn(tt.args.s)
			if got != tt.want {
				t.Errorf("splitColumn() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("splitColumn() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
