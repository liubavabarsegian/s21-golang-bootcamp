package transmitter

import (
	"log"
	"log/slog"
	"math/rand/v2"
	"time"

	"github.com/google/uuid"

	// Сгенерированный код
	transmitter "Transmitter/pkg/transmitter"
)

type TransmitterServer struct {
	transmitter.UnimplementedTransmitterServer
	logger slog.Logger
	kTime  int
}

func NewTransmitterServer(logger slog.Logger, kTime int) *TransmitterServer {
	return &TransmitterServer{logger: logger, kTime: kTime}
}

func (s *TransmitterServer) Transmit(req *transmitter.TransmitRequest, stream transmitter.Transmitter_TransmitServer) error {
	s.logger.Info("Received request for streaming data")
	timer := time.NewTicker(time.Second / 5)

	for {
		select {
		// Exit on stream context done
		case <-stream.Context().Done():
			return nil
		case <-timer.C:
			response := &transmitter.TransmitResponse{
				SessionId: GetSessionId(),
				Frequency: GetRandomFrequency(),
				Timestamp: GetTimestamp(),
			}
			if err := stream.Send(response); err != nil {
				s.logger.Error("Failed to send response:", err)
				return err
			}
		}
	}
}

func GetSessionId() string {
	uuid := uuid.New().String()
	log.Println("Generated uuid: ", uuid)

	return uuid
}

func GetRandomFrequency() float64 {
	mean := rand.Float64()*20 - 10           // [-10, 10] range
	stdDev := rand.Float64()*(1.5-0.3) + 0.3 // [0.3, 1.5] range

	frequency := mean + stdDev*rand.NormFloat64()

	log.Println("Generated random mean: ", mean)
	log.Println("Generated random standard deviation: ", stdDev)
	log.Println("Got frequency: ", frequency)

	return frequency
}

func GetTimestamp() int64 {
	timestamp := time.Now().Unix()
	log.Println("Generated timestamp: ", timestamp)

	return timestamp
}
