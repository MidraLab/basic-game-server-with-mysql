package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"
)

func Test_newSliceTableModel(t *testing.T) {
	type args struct {
		db       *DB
		dest     interface{}
		slice    reflect.Value
		elemType reflect.Type
	}
	tests := []struct {
		name string
		args args
		want *sliceTableModel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSliceTableModel(tt.args.db, tt.args.dest, tt.args.slice, tt.args.elemType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSliceTableModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceTableModel_init(t *testing.T) {
	type fields struct {
		structTableModel structTableModel
		slice            reflect.Value
		sliceLen         int
		sliceOfPtr       bool
		nextElem         func() reflect.Value
	}
	type args struct {
		sliceType reflect.Type
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
			m := &sliceTableModel{
				structTableModel: tt.fields.structTableModel,
				slice:            tt.fields.slice,
				sliceLen:         tt.fields.sliceLen,
				sliceOfPtr:       tt.fields.sliceOfPtr,
				nextElem:         tt.fields.nextElem,
			}
			m.init(tt.args.sliceType)
		})
	}
}

func Test_sliceTableModel_join(t *testing.T) {
	type fields struct {
		structTableModel structTableModel
		slice            reflect.Value
		sliceLen         int
		sliceOfPtr       bool
		nextElem         func() reflect.Value
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
			m := &sliceTableModel{
				structTableModel: tt.fields.structTableModel,
				slice:            tt.fields.slice,
				sliceLen:         tt.fields.sliceLen,
				sliceOfPtr:       tt.fields.sliceOfPtr,
				nextElem:         tt.fields.nextElem,
			}
			if got := m.join(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceTableModel.join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceTableModel_ScanRows(t *testing.T) {
	type fields struct {
		structTableModel structTableModel
		slice            reflect.Value
		sliceLen         int
		sliceOfPtr       bool
		nextElem         func() reflect.Value
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
			m := &sliceTableModel{
				structTableModel: tt.fields.structTableModel,
				slice:            tt.fields.slice,
				sliceLen:         tt.fields.sliceLen,
				sliceOfPtr:       tt.fields.sliceOfPtr,
				nextElem:         tt.fields.nextElem,
			}
			got, err := m.ScanRows(tt.args.ctx, tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("sliceTableModel.ScanRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sliceTableModel.ScanRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceTableModel_BeforeAppendModel(t *testing.T) {
	type fields struct {
		structTableModel structTableModel
		slice            reflect.Value
		sliceLen         int
		sliceOfPtr       bool
		nextElem         func() reflect.Value
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
			m := &sliceTableModel{
				structTableModel: tt.fields.structTableModel,
				slice:            tt.fields.slice,
				sliceLen:         tt.fields.sliceLen,
				sliceOfPtr:       tt.fields.sliceOfPtr,
				nextElem:         tt.fields.nextElem,
			}
			if err := m.BeforeAppendModel(tt.args.ctx, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("sliceTableModel.BeforeAppendModel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sliceTableModel_updateSoftDeleteField(t *testing.T) {
	type fields struct {
		structTableModel structTableModel
		slice            reflect.Value
		sliceLen         int
		sliceOfPtr       bool
		nextElem         func() reflect.Value
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
			m := &sliceTableModel{
				structTableModel: tt.fields.structTableModel,
				slice:            tt.fields.slice,
				sliceLen:         tt.fields.sliceLen,
				sliceOfPtr:       tt.fields.sliceOfPtr,
				nextElem:         tt.fields.nextElem,
			}
			if err := m.updateSoftDeleteField(tt.args.tm); (err != nil) != tt.wantErr {
				t.Errorf("sliceTableModel.updateSoftDeleteField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
