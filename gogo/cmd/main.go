// Copyright 2017 Johan Brandhorst. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"context"
	"net"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	pbGoGo "github.com/johanbrandhorst/gogoproto-experiments/gogo/server"
)

var logger *logrus.Logger

func init() {
	logger = logrus.StandardLogger()
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})
	// Should only be done from init functions
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(logger.Out, logger.Out, logger.Out))
}

type gogoSrv struct{}

func (g gogoSrv) Empty(_ context.Context, _ *types.Empty) (*pbGoGo.TestReply, error) {
	t := time.Now()
	return &pbGoGo.TestReply{
		RequestTime: &t,
	}, nil
}

func (g gogoSrv) TestMethod(_ context.Context, _ *pbGoGo.TestRequest) (*pbGoGo.TestReply, error) {
	t := time.Now()
	return &pbGoGo.TestReply{
		RequestTime: &t,
	}, nil
}

func main() {
	gs := grpc.NewServer()
	pbGoGo.RegisterTestServiceServer(gs, &gogoSrv{})
	reflection.Register(gs)

	addr := "localhost:10000"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		logger.WithError(err).Fatal("Failed to start listener")
	}

	logger.Info("Serving on https://" + addr)
	logger.Fatal(gs.Serve(l))
}
