package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/WBTechSchool/L0/internal/models"
	"github.com/segmentio/kafka-go"
)

func mockOrderData() models.CreateOrderRequest {
	order := models.CreateOrderRequest{
		Order: models.OrderRequest{
			TrackNumber:       "WBILMTESTTRACK",
			Entry:             "WBIL",
			Locale:            "en",
			InternalSignature: "",
			CustomerID:        1,
			DeliveryService:   "meest",
			ShardKey:          "9",
			SmID:              99,
			OffShard:          "1",
		},
		Delivery: models.Delivery{
			Name:    "Test Testov",
			Phone:   "+9720000000",
			Zip:     "2639809",
			City:    "Kiryat Mozkin",
			Address: "Ploshad Mira 15",
			Region:  "Kraiot",
			Email:   "test@gmail.com",
		},
		Payment: models.Payment{
			Transaction:  "b563feb7b2b84b6test",
			RequestID:    "",
			Currency:     "USD",
			Provider:     "wbpay",
			Amount:       1817,
			PaymentDt:    1637907727,
			Bank:         "aplha",
			DeliveryCost: 1500,
			GoodsTotal:   317,
			CustomFee:    0,
		},
		OrderItems: []models.OrderItem{
			{
				ChrtID:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				Rid:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        30,
				Size:        "0",
				TotalPrice:  317,
				NmID:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
		},
	}

	return order
}

func main() {
	topic := os.Getenv("KAFKA_TOPIC")
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", os.Getenv("KAFKA_EXTERN_URL"), topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	value, err := kafka.Marshal(mockOrderData())
	if err != nil {
		log.Fatal("Failed marshaling data")
	}

	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(value)},
	)

	if err != nil {
		log.Fatal("Failed to write messages:", err)
	} else {
		fmt.Println("Message writed")
	}

	if err := conn.Close(); err != nil {
		log.Fatal("Failed to close writer:", err)
	}
}
