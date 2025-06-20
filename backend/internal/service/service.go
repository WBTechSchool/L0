package service

import "github.com/WBTechSchool/L0/internal/db"

type Service struct {
}

func NewService(db *db.Storage) *Service {
	return &Service{}
}
