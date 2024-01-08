package handlers

import (
	"context"

	"github.com/PNYwise/config-service/internal/domain"
	config_service "github.com/PNYwise/config-service/proto"
	"github.com/golang/protobuf/ptypes/empty"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type configHandler struct {
	config_service.UnimplementedConfigServer
	configServie domain.IConfigService
}

func NewConfigHandler(
	configServie domain.IConfigService,
) *configHandler {
	return &configHandler{
		configServie: configServie,
	}
}

func (c *configHandler) Get(ctx context.Context, _ *empty.Empty) (*structpb.Value, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.DataLoss, "UnaryEcho: failed to get metadata")
	}

	idValues, idExists := md["id"]
	tokenValues, tokenExists := md["token"]

	if !idExists || len(idValues) == 0 || !tokenExists || len(tokenValues) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID or token metadata missing or empty")
	}

	configs, err := c.configServie.Get(&domain.ConfigRequest{
		ID:    idValues[0],
		Token: tokenValues[0],
	})

	if err != nil {
		return nil, status.Errorf(codes.DataLoss, err.Error())
	}
	strings := &structpb.Value{
		Kind: &structpb.Value_StringValue{StringValue: configs.(string)},
	}
	return strings, nil
}
func (c *configHandler) Set(context.Context, *structpb.Value) (*structpb.Value, error) {
	return nil, nil
}
func (c *configHandler) Update(context.Context, *structpb.Value) (*structpb.Value, error) {
	return nil, nil
}
func (c *configHandler) Delete(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, nil
}
