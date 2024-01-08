package internal

import (
	"github.com/PNYwise/config-service/internal/grpc/handlers"
	"github.com/PNYwise/config-service/internal/service"
	config_service "github.com/PNYwise/config-service/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func InitGrpc(srv *grpc.Server, conf *viper.Viper) {
	configGrpcInit(srv, conf)
}

func configGrpcInit(srv *grpc.Server, conf *viper.Viper) {
	configservice := service.NewConfigServie(conf)
	configHandler := handlers.NewConfigHandler(configservice)

	config_service.RegisterConfigServer(srv, configHandler)
}
