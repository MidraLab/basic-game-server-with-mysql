package bun

import (
	"reflect"
	"testing"

	"github.com/uptrace/bun/schema"
)

func Test_newModel(t *testing.T) {
	type args struct {
		db   *DB
		dest []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    Model
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newModel(tt.args.db, tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("newModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newSingleModel(t *testing.T) {
	type args struct {
		db   *DB
		dest interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    Model
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newSingleModel(tt.args.db, tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("newSingleModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSingleModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newModel(t *testing.T) {
	type args struct {
		db   *DB
		dest interface{}
		scan bool
	}
	tests := []struct {
		name    string
		args    args
		want    Model
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := _newModel(tt.args.db, tt.args.dest, tt.args.scan)
			if (err != nil) != tt.wantErr {
				t.Errorf("_newModel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("_newModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newTableModelIndex(t *testing.T) {
	type args struct {
		db    *DB
		table *schema.Table
		root  reflect.Value
		index []int
		rel   *schema.Relation
	}
	tests := []struct {
		name    string
		args    args
		want    TableModel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newTableModelIndex(tt.args.db, tt.args.table, tt.args.root, tt.args.index, tt.args.rel)
			if (err != nil) != tt.wantErr {
				t.Errorf("newTableModelIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTableModelIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validMap(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validMap(tt.args.typ); (err != nil) != tt.wantErr {
				t.Errorf("validMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_isSingleRowModel(t *testing.T) {
	type args struct {
		m Model
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSingleRowModel(tt.args.m); got != tt.want {
				t.Errorf("isSingleRowModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
