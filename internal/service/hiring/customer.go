package hiring

import (
	"context"
	"exchanger/internal/domain/customer"
	"exchanger/pkg/log"
	"exchanger/pkg/market"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func (s *Service) ListCustomers(ctx context.Context) (res []customer.Response, err error) {
	logger := log.LoggerFromContext(ctx).Named("ListCustomers")

	data, err := s.customerRepository.List(ctx)
	if err != nil {
		logger.Error("failed to select", zap.Error(err))
		return
	}
	res = customer.ParseFromEntities(data)

	return
}

func (s *Service) AddCustomer(ctx context.Context, req customer.Request) (res customer.Response, err error) {
	logger := log.LoggerFromContext(ctx).Named("AddCustomer")

	data := customer.Entity{
		FullName:  &req.FullName,
		Pseudonym: &req.Pseudonym,
	}

	data.ID, err = s.customerRepository.Add(ctx, data)
	if err != nil {
		logger.Error("failed to add", zap.Error(err))
		return
	}
	res = customer.ParseFromEntity(data)

	return
}

func (s *Service) GetCustomer(ctx context.Context, id string) (res customer.Response, err error) {
	logger := log.LoggerFromContext(ctx).Named("GetCustomer").With(zap.String("id", id))

	data, err := s.customerRepository.Get(ctx, id)
	if err != nil && !errors.Is(err, market.ErrorNotFound) {
		logger.Error("failed to get", zap.Error(err))
		return
	}
	res = customer.ParseFromEntity(data)

	return
}

func (s *Service) UpdateCustomer(ctx context.Context, id string, req customer.Request) (err error) {
	logger := log.LoggerFromContext(ctx).Named("UpdateCustomer").With(zap.String("id", id))

	data := customer.Entity{
		FullName:  &req.FullName,
		Pseudonym: &req.Pseudonym,
	}

	err = s.customerRepository.Update(ctx, id, data)
	if err != nil && !errors.Is(err, market.ErrorNotFound) {
		logger.Error("failed to update by id", zap.Error(err))
		return
	}

	return
}

func (s *Service) DeleteCustomer(ctx context.Context, id string) (err error) {
	logger := log.LoggerFromContext(ctx).Named("DeleteCustomer").With(zap.String("id", id))

	err = s.customerRepository.Delete(ctx, id)
	if err != nil && !errors.Is(err, market.ErrorNotFound) {
		logger.Error("failed to delete", zap.Error(err))
		return
	}

	return
}
