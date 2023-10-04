package jlexer

import "testing"

func TestLexerError_Error(t *testing.T) {
	type fields struct {
		Reason string
		Offset int
		Data   string
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
			l := &LexerError{
				Reason: tt.fields.Reason,
				Offset: tt.fields.Offset,
				Data:   tt.fields.Data,
			}
			if got := l.Error(); got != tt.want {
				t.Errorf("LexerError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
