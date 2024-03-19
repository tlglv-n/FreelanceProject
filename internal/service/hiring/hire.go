package hiring

import (
	"context"
	"exchanger/internal/domain/hire"
	"exchanger/pkg/log"
	"exchanger/pkg/market"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (s *Service) ListHires(ctx context.Context) (res []hire.Response, err error) {
	logger := log.LoggerFromContext(ctx).Named("ListHires")

	data, err := s.hireRepository.List(ctx)
	if err != nil {
		logger.Error("failed to select", zap.Error(err))
		return
	}
	res = hire.ParseFromEntities(data)

	return
}

func (s *Service) AddHire(ctx context.Context, req hire.Request) (res hire.Response, err error) {
	logger := log.LoggerFromContext(ctx).Named("AddHire")

	data := hire.Entity{
		JobName:     &req.JobName,
		Amount:      &req.Amount,
		Description: &req.Description,
		Position:    &req.Position,
		CustomerID:  req.CustomerID,
	}

	data.ID, err = s.hireRepository.Add(ctx, data)
	if err != nil {
		logger.Error("failed to add", zap.Error(err))
		return
	}
	res = hire.ParseFromEntity(data)

	return
}

func (s *Service) GetHire(ctx context.Context, id string) (res hire.Response, err error) {
	logger := log.LoggerFromContext(ctx).Named("GetHire").With(zap.String("id", id))

	data, err := s.hireRepository.Get(ctx, id)
	if err != nil && !errors.Is(err, market.ErrorNotFound) {
		logger.Error("failed to get", zap.Error(err))
		return
	}
	res = hire.ParseFromEntity(data)

	return
}

func (s *Service) UpdateHire(ctx context.Context, id string, req hire.Request) (err error) {
	logger := log.LoggerFromContext(ctx).Named("UpdateHire").With(zap.String("id", id))

	data := hire.Entity{
		JobName:     &req.JobName,
		Amount:      &req.Amount,
		Description: &req.Description,
		Position:    &req.Position,
		CustomerID:  req.CustomerID,
	}

	err = s.hireRepository.Update(ctx, id, data)
	if err != nil && !errors.Is(err, market.ErrorNotFound) {
		logger.Error("failed to update by id", zap.Error(err))
		return
	}

	return
}

func (s *Service) DeleteHire(ctx context.Context, id string) (err error) {
	logger := log.LoggerFromContext(ctx).Named("DeleteHire").With(zap.String("id", id))

	err = s.hireRepository.Delete(ctx, id)
	if err != nil && !errors.Is(err, market.ErrorNotFound) {
		logger.Error("failed to delete", zap.Error(err))
		return
	}

	return
}
