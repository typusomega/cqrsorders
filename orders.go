package main

import (
	"github.com/typusomega/cqrsorders/pkg/api"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
	"github.com/typusomega/cqrsorders/pkg/config"
	spec "github.com/typusomega/cqrsorders/pkg/spec/v1"
)

var mainLog = logrus.WithField("name", "main")

func main() {
	cfg := config.Get()

	address := fmt.Sprintf(":%d", cfg.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		mainLog.Fatalf("could not create listener on address: %s", address)
	}

	server := api.New()

	grpcServer := grpc.NewServer()
	spec.RegisterOrdersServer(grpcServer, server)

	mainLog.Infof("started grpc server on '%v'", address)
	if err := grpcServer.Serve(listener); err != nil {
		mainLog.Fatalf("failed to serve grpc: '%v'", err)
	}
}
