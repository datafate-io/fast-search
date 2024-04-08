package encoder

import (
	"context"
	"fast-search/configs"
	"time"

	pb "fast-search/internal/infrastructure/encoder/textencoderpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TextEncoderClient struct {
	config *configs.Configs
}

func NewTextEncoderClient(config *configs.Configs) *TextEncoderClient {
	return &TextEncoderClient{
		config: config,
	}
}

func (c *TextEncoderClient) EncodeText(text string) ([]float32, error) {
	conn, err := grpc.Dial(c.config.GRPCConection, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewTextEncoderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.EncodeText(ctx, &pb.TextRequest{Text: text})
	if err != nil {
		return nil, err
	}
	return r.Vector, nil
}
