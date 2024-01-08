package main

import (
	"log"
	"net"
	"time"

	"github.com/PNYwise/config-service/internal"
	"github.com/PNYwise/config-service/internal/configs"
	"google.golang.org/grpc"
)

func main() {
	time.Local = time.UTC
	srv := grpc.NewServer()
	conf := configs.New()

	internal.InitGrpc(srv, conf)

	port := conf.GetString("config-service.config.app.port")
	log.Println("Starting RPC server at", ":"+port)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", ":"+port, err)
	}
	log.Println("server Started at", ":"+port)
	log.Fatal(srv.Serve(l))
}
