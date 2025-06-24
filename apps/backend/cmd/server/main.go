package main

import (
	"context"
	"log/slog"

	"github.com/WBTechSchool/L0/internal/db"
	"github.com/WBTechSchool/L0/internal/http/router"
	"github.com/WBTechSchool/L0/internal/http/server"
	"github.com/WBTechSchool/L0/internal/kafka"
	"github.com/WBTechSchool/L0/internal/logger"
	"github.com/WBTechSchool/L0/internal/service"
)

func main() {
	logger.New()

	db, err := db.New(context.Background())
	if err != nil {
		slog.Error("Database init", slog.Any("err", err))
		return
	}

	s := service.NewService(db)
	r := router.NewRouter(s)

	go kafka.NewKafka(s)
	server.NewServer(r)
}
