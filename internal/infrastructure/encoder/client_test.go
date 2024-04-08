package encoder

import (
	"fast-search/configs"
	"fmt"
	"testing"
)

func TestNewTextEncoderClient(t *testing.T) {
	type args struct {
		config *configs.Configs
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestNewTextEncoderClient",
			args: args{
				config: &configs.Configs{
					GRPCConection: "localhost:50051",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = NewTextEncoderClient(tt.args.config)
		})
	}
}

func TestTextEncoderClient_EncodeText(t *testing.T) {
	type fields struct {
		config *configs.Configs
	}
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []float32
		wantErr bool
	}{
		{
			name: "TestTextEncoderClient_EncodeText",
			fields: fields{
				config: &configs.Configs{
					GRPCConection: "localhost:50051",
				},
			},
			args: args{
				text: "test",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TextEncoderClient{
				config: tt.fields.config,
			}
			got, err := c.EncodeText(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("TextEncoderClient.EncodeText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
