package hiring

import (
	"exchanger/internal/domain/customer"
	"exchanger/internal/domain/hire"
	"exchanger/internal/domain/worker"
)

// Configuration is an alias for a function that will take in a pointer to a Service and modify it
type Configuration func(s *Service) error

// Service is an implementation of the Service
type Service struct {
	customerRepository customer.Repository
	workerRepository   worker.Repository
	hireRepository     hire.Repository
	customerCache      customer.Cache
	// TODO: workerCache
	// TODO: hireCache
}

// New takes a variable amount of Configuration functions and returns a new Service
// Each Configuration will be called in the order they are passed in
func New(configs ...Configuration) (s *Service, err error) {
	// Add the service
	s = &Service{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		if err = cfg(s); err != nil {
			return
		}
	}
	return
}

// WithCustomerRepository applies a given customer repository to the Service
func WithCustomerRepository(customerRepository customer.Repository) Configuration {
	// return a function that matches the Configuration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(s *Service) error {
		s.customerRepository = customerRepository
		return nil
	}
}

// WithWorkerRepository applies a given book repository to the Service
func WithWorkerRepository(workerRepository worker.Repository) Configuration {
	// Add the worker repository, if we needed parameters, such as connection strings they could be inputted here
	return func(s *Service) error {
		s.workerRepository = workerRepository
		return nil
	}
}

// WithHireRepository applies a given book repository to the Service
func WithHireRepository(hireRepository hire.Repository) Configuration {
	// Add the hire repository, if we needed parameters, such as connection strings they could be inputted here
	return func(s *Service) error {
		s.hireRepository = hireRepository
		return nil
	}
}

// WithCustomerCache applies a given author cache to the Service
func WithCustomerCache(customerCache customer.Cache) Configuration {
	// return a function that matches the Configuration alias,
	// You need to return this so that the parent function can take in all the needed parameters
	return func(s *Service) error {
		s.customerCache = customerCache
		return nil
	}
}

// TODO: WithWorkerCache

// TODO: WithHireCache
