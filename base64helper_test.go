package base64helper

import (
	"encoding/base64"
	"errors"
	"testing"
)

func TestDecodeStringToString(t *testing.T) {
	type args struct {
		enc *base64.Encoding
		s   string
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantErr   bool
		wantErrIs error
	}{
		{name: "error 1",
			args: args{
				enc: base64.StdEncoding,
				s:   "qwerty",
			},
			wantErr: true,
		},
		{name: "ErrInvalidUTF8",
			args: args{
				enc: base64.StdEncoding,
				s:   "//79",
			},
			wantErr:   true,
			wantErrIs: ErrInvalidUTF8,
		},
		{name: "non-nil enc",
			args: args{enc: base64.StdEncoding, s: "MQ=="},
			want: "1",
		},
		{name: "nil enc",
			args: args{enc: nil, s: "MQ=="},
			want: "1",
		},
		{name: "empty",
			args: args{enc: base64.StdEncoding, s: ""},
			want: "",
		},
		{name: "multibyte utf8",
			args: args{enc: base64.StdEncoding, s: base64.StdEncoding.EncodeToString([]byte("привет"))},
			want: "привет",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeStringToString(tt.args.enc, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeStringToString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErrIs != nil && !errors.Is(err, tt.wantErrIs) {
				t.Errorf("DecodeStringToString() error = %v, wantErrIs %v", err, tt.wantErrIs)
			}
			if got != tt.want {
				t.Errorf("DecodeStringToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
