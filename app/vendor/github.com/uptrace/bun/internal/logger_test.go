package internal

import (
	"log"
	"testing"
)

func Test_logger_Printf(t *testing.T) {
	type fields struct {
		log *log.Logger
	}
	type args struct {
		format string
		v      []interface{}
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
			l := &logger{
				log: tt.fields.log,
			}
			l.Printf(tt.args.format, tt.args.v...)
		})
	}
}
