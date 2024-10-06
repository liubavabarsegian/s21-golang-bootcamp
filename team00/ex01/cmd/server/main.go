package main

import (
	pb "Transmitter/internal/grpc/transmitter"
	transmitter "Transmitter/pkg/transmitter"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	logger := setupLogger(envLocal)

	grpcServer := grpc.NewServer()
	transmitterServer := pb.NewTransmitterServer(*logger, 100)
	transmitter.RegisterTransmitterServer(grpcServer, transmitterServer)

	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		logger.Error("Cannot start server: ", err)
		return
	}

	logger.Info("Started server")

	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Error("cannot start server: ", err)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
