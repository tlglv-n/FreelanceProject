package repository

import (
	"exchanger/internal/domain/customer"
	"exchanger/internal/domain/hire"
	"exchanger/internal/domain/worker"
	"exchanger/internal/repository/memory"
	"exchanger/internal/repository/mongo"
	"exchanger/internal/repository/postgres"
	"exchanger/pkg/market"
)

// Configuration is an alias for a function that will take in a pointer to a Repository and modify it
type Configuration func(r *Repository) error

// Repository is an implementation of the Repository
type Repository struct {
	mongo    market.Mongo
	postgres market.SQLX

	Customer customer.Repository
	Hire     hire.Repository
	Worker   worker.Repository
}

// New takes a variable amount of Configuration functions and returns a new Repository
// Each Configuration will be called in the order they are passed in
func New(configs ...Configuration) (s *Repository, err error) {
	// Create the repository
	s = &Repository{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the repository into the configuration function
		if err = cfg(s); err != nil {
			return
		}
	}

	return
}

// Close closes the repository and prevents new queries from starting.
// Close then waits for all queries that have started processing on the server to finish.
func (r *Repository) Close() {
	if r.postgres.Client != nil {
		r.postgres.Client.Close()
	}

	if r.mongo.Client != nil {
		r.mongo.Client.Disconnect(nil)
	}
}

// WithMemoryStore applies a memory store to the Repository
func WithMemoryStore() Configuration {
	return func(s *Repository) (err error) {
		// Create the memory store, if we needed parameters, such as connection strings they could be inputted here
		s.Customer = memory.NewCustomerRepository()
		s.Hire = memory.NewHireRepository()
		s.Worker = memory.NewWorkerRepository()

		return
	}
}

// WithMongoStore applies a mongo store to the Repository
func WithMongoStore(url, name string) Configuration {
	return func(s *Repository) (err error) {
		// Create the mongo store, if we needed parameters, such as connection strings they could be inputted here
		s.mongo, err = market.NewMongo(url)
		if err != nil {
			return
		}
		database := s.mongo.Client.Database(name)

		s.Customer = mongo.NewCustomerRepository(database)
		s.Hire = mongo.NewHireRepository(database)
		s.Worker = mongo.NewWorkerRepository(database)

		return
	}
}

// WithPostgresStore applies a postgres store to the Repository
func WithPostgresStore(dataSourceName string) Configuration {
	return func(s *Repository) (err error) {
		// Create the postgres store, if we needed parameters, such as connection strings they could be inputted here
		s.postgres, err = market.NewSQL(dataSourceName)
		if err != nil {
			return
		}

		if err = market.Migrate(dataSourceName); err != nil {
			return
		}

		s.Customer = postgres.NewCustomerRepository(s.postgres.Client)
		s.Hire = postgres.NewHireRepository(s.postgres.Client)
		s.Worker = postgres.NewWorkerRepository(s.postgres.Client)

		return
	}
}
