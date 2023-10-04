package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
)

func Test_newSliceModel(t *testing.T) {
	type args struct {
		db     *DB
		dest   []interface{}
		values []reflect.Value
	}
	tests := []struct {
		name string
		args args
		want *sliceModel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSliceModel(tt.args.db, tt.args.dest, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSliceModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceModel_Value(t *testing.T) {
	type fields struct {
		dest      []interface{}
		values    []reflect.Value
		scanIndex int
		info      []sliceInfo
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
			m := &sliceModel{
				dest:      tt.fields.dest,
				values:    tt.fields.values,
				scanIndex: tt.fields.scanIndex,
				info:      tt.fields.info,
			}
			if got := m.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceModel.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceModel_ScanRows(t *testing.T) {
	type fields struct {
		dest      []interface{}
		values    []reflect.Value
		scanIndex int
		info      []sliceInfo
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
			m := &sliceModel{
				dest:      tt.fields.dest,
				values:    tt.fields.values,
				scanIndex: tt.fields.scanIndex,
				info:      tt.fields.info,
			}
			got, err := m.ScanRows(tt.args.ctx, tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("sliceModel.ScanRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sliceModel.ScanRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceModel_Scan(t *testing.T) {
	type fields struct {
		dest      []interface{}
		values    []reflect.Value
		scanIndex int
		info      []sliceInfo
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
			m := &sliceModel{
				dest:      tt.fields.dest,
				values:    tt.fields.values,
				scanIndex: tt.fields.scanIndex,
				info:      tt.fields.info,
			}
			if err := m.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("sliceModel.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
