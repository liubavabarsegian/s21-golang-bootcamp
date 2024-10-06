package main

import (
	transmitter "Transmitter/pkg/transmitter"
	"context"
	"log"
	"log/slog"
	"math"
	"os"

	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultAddress = "localhost:8888"
	defaultK       = 2.0
)

func main() {
	logger := setupLogger()
	address, k := ParseFLags()

	// Создание gRPC клиента
	conn, err := grpc.NewClient(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Error("Failed to connect: ", err)
	}
	defer conn.Close()
	logger.Info("Created gRPC client connection")

	transmitterClient := transmitter.NewTransmitterClient(conn)
	clientRequest := &transmitter.TransmitRequest{}

	stream, err := transmitterClient.Transmit(context.Background(), clientRequest)
	if err != nil {
		logger.Error("Error while transmitting: ", err)
		return
	}

	DetectAnomalies(*logger, stream, k)

}

func setupLogger() *slog.Logger {
	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	return log
}

func ParseFLags() (address *string, k *float64) {
	address = flag.String("address", defaultAddress, "Server address")
	k = flag.Float64("k", defaultK, "Anomaly detection coefficient")
	flag.Parse()
	return
}

func DetectAnomalies(logger slog.Logger, stream grpc.ServerStreamingClient[transmitter.TransmitResponse], k *float64) {

	var sum float64    // Сумма всех частот для вычисления среднего значения
	var sumSq float64  // Сумма квадратов частот для вычисления дисперсии
	var count int      // Счетчик полученных сообщений
	var mean float64   // Среднее значение частот
	var stdDev float64 // Стандартное отклонение частот

	for {
		msg, err := stream.Recv() // Получение сообщения из потока
		if err != nil {
			log.Fatalf("Stream error: %v", err)
		}

		count++
		frequency := msg.GetFrequency()
		sum += frequency
		sumSq += frequency * frequency

		// Рассчитываем среднее и стандартное отклонение
		mean = sum / float64(count)
		variance := (sumSq / float64(count)) - (mean * mean) // Вычисление дисперсии
		if variance < 0 {
			variance = 0 // Гарантируем, что дисперсия не отрицательная
		}
		stdDev = math.Sqrt(variance)

		log.Printf("Received frequency: %f, Mean: %f, StdDev: %f", frequency, mean, stdDev)
		// Обнаружение аномалий
		if math.Abs(frequency-mean) > *k*stdDev {
			log.Printf("Anomaly detected! Frequency: %f, Mean: %f, StdDev: %f", frequency, mean, stdDev)
		}
	}
}
