package service

import (
	"github.com/ShadamHarizky/consumer-service/model"
	"github.com/ShadamHarizky/consumer-service/repository"
)

type ConsumerService struct {
	Repo repository.MessageRepository
}

func NewConsumerService(repo repository.MessageRepository) *ConsumerService {
	return &ConsumerService{Repo: repo}
}

func (s *ConsumerService) ProcessMessage(payload model.Message) error {
	message := &model.Message{Content: payload.Content, Source: payload.Source}
	return s.Repo.Save(message)
}
