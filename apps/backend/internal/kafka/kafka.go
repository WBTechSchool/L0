package kafka

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/WBTechSchool/L0/internal/models"
	"github.com/WBTechSchool/L0/internal/service"
	"github.com/segmentio/kafka-go"
)

func NewKafka(s *service.Service) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{os.Getenv("KAFKA_URL")},
		GroupID:   "group_id",
		Topic:     os.Getenv("KAFKA_TOPIC"),
		Partition: 0,
		MaxBytes:  10e6,
	})
	defer r.Close()

	slog.Info("Kafka starting")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	errChan := make(chan error, 1)

	go func() {
		for {
			msg, err := r.ReadMessage(context.Background())
			if err != nil {
				errChan <- err
				return
			}

			var order models.CreateOrderRequest
			if err := kafka.Unmarshal(msg.Value, &order); err != nil {
				errChan <- err
				return
			}

			slog.Info("[KAFKA][CONSUMER]", slog.Any("data", order))

			id, err := s.OrderService.CreateOrder(context.Background(), order)
			if err != nil {
				errChan <- err
				return
			}

			slog.Info("Order created", slog.Int64("id", id))

			if err := r.CommitMessages(context.Background(), msg); err != nil {
				errChan <- err
				return
			}
		}
	}()

	select {
	case <-sigChan:
		slog.Info("Kafka received shutdown signal")
	case err := <-errChan:
		slog.Error("Error", slog.Any("err", err))
	}
}
