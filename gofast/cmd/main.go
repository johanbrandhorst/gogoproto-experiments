// Copyright 2017 Johan Brandhorst. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"context"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	pbGofast "github.com/johanbrandhorst/gogoproto-experiments/gofast/server"
	"github.com/johanbrandhorst/gogoproto-experiments/gofast/timestamp"
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

type goFastSrv struct{}

func (g goFastSrv) Empty(_ context.Context, _ *empty.Empty) (*pbGofast.TestReply, error) {
	return &pbGofast.TestReply{
		RequestTime: &timestamp.Timestamp{
			Seconds: time.Now().Unix(),
			Nanos:   int32(time.Now().UnixNano()),
		},
	}, nil
}

func (g goFastSrv) TestMethod(_ context.Context, _ *pbGofast.TestRequest) (*pbGofast.TestReply, error) {
	return &pbGofast.TestReply{
		RequestTime: &timestamp.Timestamp{
			Seconds: time.Now().Unix(),
			Nanos:   int32(time.Now().UnixNano()),
		},
	}, nil
}

func main() {
	gs := grpc.NewServer()
	pbGofast.RegisterTestServiceServer(gs, &goFastSrv{})
	reflection.Register(gs)

	addr := "localhost:10000"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		logger.WithError(err).Fatal("Failed to start listener")
	}

	logger.Info("Serving on https://" + addr)
	logger.Fatal(gs.Serve(l))
}
