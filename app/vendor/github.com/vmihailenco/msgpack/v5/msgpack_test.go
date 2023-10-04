package msgpack

import "testing"

func TestRawMessage_EncodeMsgpack(t *testing.T) {
	type args struct {
		enc *Encoder
	}
	tests := []struct {
		name    string
		m       RawMessage
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.EncodeMsgpack(tt.args.enc); (err != nil) != tt.wantErr {
				t.Errorf("RawMessage.EncodeMsgpack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRawMessage_DecodeMsgpack(t *testing.T) {
	type args struct {
		dec *Decoder
	}
	tests := []struct {
		name    string
		m       *RawMessage
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.DecodeMsgpack(tt.args.dec); (err != nil) != tt.wantErr {
				t.Errorf("RawMessage.DecodeMsgpack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unexpectedCodeError_Error(t *testing.T) {
	type fields struct {
		code byte
		hint string
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
			err := unexpectedCodeError{
				code: tt.fields.code,
				hint: tt.fields.hint,
			}
			if got := err.Error(); got != tt.want {
				t.Errorf("unexpectedCodeError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
