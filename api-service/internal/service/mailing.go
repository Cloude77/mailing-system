package service

import (
	"context"
	"github.com/Cloude77/mailing-system/api-service/internal/model"
	"github.com/Cloude77/mailing-system/api-service/internal/repository"
)

type MailingService struct {
	repo *repository.Postgres
}

func NewMailingService(repo *repository.Postgres) *MailingService {
	return &MailingService{repo: repo}
}

func (s *MailingService) CreateMailing(ctx context.Context, req *model.MailingRequest) (*model.Mailing, error) {
	mailing := &model.Mailing{
		Message:     req.Message,
		GroupName:   req.GroupName,
		SendTime:    req.SendTime,
		Periodicity: req.Periodicity,
	}
	if err := s.repo.CreateMailing(ctx, mailing); err != nil {
		return nil, err
	}
	return mailing, nil
}
