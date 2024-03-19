package hiring

import (
	"context"
	"exchanger/internal/domain/worker"
	"exchanger/pkg/log"
	"exchanger/pkg/market"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (s *Service) ListWorkers(ctx context.Context) (res []worker.Response, err error) {
	logger := log.LoggerFromContext(ctx).Named("ListWorkers")

	data, err := s.workerRepository.List(ctx)
	if err != nil {
		logger.Error("failed to select", zap.Error(err))
		return
	}
	res = worker.ParseFromEntities(data)

	return
}

func (s *Service) AddWorker(ctx context.Context, req worker.Request) (res worker.Response, err error) {
	logger := log.LoggerFromContext(ctx).Named("AddWorker")

	data := worker.Entity{
		FullName:    &req.FullName,
		Pseudonym:   &req.Pseudonym,
		Description: &req.Description,
		Position:    &req.Position,
	}

	data.ID, err = s.workerRepository.Add(ctx, data)
	if err != nil {
		logger.Error("failed to add", zap.Error(err))
		return
	}
	res = worker.ParseFromEntity(data)

	return
}

func (s *Service) GetWorker(ctx context.Context, id string) (res worker.Response, err error) {
	logger := log.LoggerFromContext(ctx).Named("GetWorker").With(zap.String("id", id))

	data, err := s.workerRepository.Get(ctx, id)
	if err != nil && !errors.Is(err, market.ErrorNotFound) {
		logger.Error("failed to get", zap.Error(err))
		return
	}
	res = worker.ParseFromEntity(data)

	return
}

func (s *Service) UpdateWorker(ctx context.Context, id string, req worker.Request) (err error) {
	logger := log.LoggerFromContext(ctx).Named("UpdateWorker").With(zap.String("id", id))

	data := worker.Entity{
		FullName:    &req.FullName,
		Pseudonym:   &req.Pseudonym,
		Description: &req.Description,
		Position:    &req.Position,
	}

	err = s.workerRepository.Update(ctx, id, data)
	if err != nil && !errors.Is(err, market.ErrorNotFound) {
		logger.Error("failed to update by id", zap.Error(err))
		return
	}

	return
}

func (s *Service) DeleteWorker(ctx context.Context, id string) (err error) {
	logger := log.LoggerFromContext(ctx).Named("DeleteWorker").With(zap.String("id", id))

	err = s.workerRepository.Delete(ctx, id)
	if err != nil && !errors.Is(err, market.ErrorNotFound) {
		logger.Error("failed to delete", zap.Error(err))
		return
	}

	return
}
