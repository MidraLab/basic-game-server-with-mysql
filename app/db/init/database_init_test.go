package db_init

import (
	"testing"

	"github.com/uptrace/bun"
)

func TestCreateTable(t *testing.T) {
	type args struct {
		db *bun.DB
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateTable(tt.args.db)
		})
	}
}
