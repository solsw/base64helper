package base64helper

import (
	"encoding/base64"
	"testing"
)

func TestDecodeStringToString(t *testing.T) {
	type args struct {
		enc *base64.Encoding
		s   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "1e",
			args:    args{enc: base64.StdEncoding, s: "qwerty"},
			wantErr: true,
		},
		{name: "2e",
			args:    args{enc: base64.StdEncoding, s: "//79"},
			wantErr: true,
		},
		{name: "1",
			args: args{enc: base64.StdEncoding, s: "MQ=="},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeStringToString(tt.args.enc, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeStringToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecodeStringToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
