package bun

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
)

func Test_newScanModel(t *testing.T) {
	type args struct {
		db   *DB
		dest []interface{}
	}
	tests := []struct {
		name string
		args args
		want *scanModel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newScanModel(tt.args.db, tt.args.dest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newScanModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanModel_Value(t *testing.T) {
	type fields struct {
		db        *DB
		dest      []interface{}
		scanIndex int
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
			m := &scanModel{
				db:        tt.fields.db,
				dest:      tt.fields.dest,
				scanIndex: tt.fields.scanIndex,
			}
			if got := m.Value(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scanModel.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanModel_ScanRows(t *testing.T) {
	type fields struct {
		db        *DB
		dest      []interface{}
		scanIndex int
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
			m := &scanModel{
				db:        tt.fields.db,
				dest:      tt.fields.dest,
				scanIndex: tt.fields.scanIndex,
			}
			got, err := m.ScanRows(tt.args.ctx, tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("scanModel.ScanRows() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("scanModel.ScanRows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanModel_ScanRow(t *testing.T) {
	type fields struct {
		db        *DB
		dest      []interface{}
		scanIndex int
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
			m := &scanModel{
				db:        tt.fields.db,
				dest:      tt.fields.dest,
				scanIndex: tt.fields.scanIndex,
			}
			if err := m.ScanRow(tt.args.ctx, tt.args.rows); (err != nil) != tt.wantErr {
				t.Errorf("scanModel.ScanRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_scanModel_Scan(t *testing.T) {
	type fields struct {
		db        *DB
		dest      []interface{}
		scanIndex int
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
			m := &scanModel{
				db:        tt.fields.db,
				dest:      tt.fields.dest,
				scanIndex: tt.fields.scanIndex,
			}
			if err := m.Scan(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("scanModel.Scan() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
