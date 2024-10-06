package main

import (
	transmitter "Transmitter/pkg/transmitter"
	"context"
	"log/slog"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger := setupLogger()

	var connection *grpc.ClientConn
	connection, err := grpc.Dial(":8888", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logger.Error("Failed to connect to grpc server %s", err)
		return
	}
	defer connection.Close()

	logger.Info("Created gRPC client connection")

	transmitterClient := transmitter.NewTransmitterClient(connection)
	clientRequest := &transmitter.TransmitRequest{}

	response, err := transmitterClient.Transmit(context.Background(), clientRequest)
	if err != nil {
		logger.Error("Error while transmitting: ", err)
		return
	}

	logger.Info("Response from server: ", response)
}

func setupLogger() *slog.Logger {
	var log *slog.Logger

	log = slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	return log
}
