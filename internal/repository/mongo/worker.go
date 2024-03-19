package mongo

import (
	"context"
	"errors"
	"exchanger/internal/domain/worker"
	"exchanger/pkg/market"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WorkerRepository struct {
	db *mongo.Collection
}

func NewWorkerRepository(db *mongo.Database) *WorkerRepository {
	return &WorkerRepository{
		db: db.Collection("customers"),
	}
}

func (r *WorkerRepository) List(ctx context.Context) (dest []worker.Entity, err error) {
	cur, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &dest); err != nil {
		return nil, err
	}

	return
}

func (r *WorkerRepository) Add(ctx context.Context, data worker.Entity) (id string, err error) {
	res, err := r.db.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).String(), nil
}

func (r *WorkerRepository) Get(ctx context.Context, id string) (dest worker.Entity, err error) {
	if err = r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&dest); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = market.ErrorNotFound
		}
	}

	return
}

func (r *WorkerRepository) Update(ctx context.Context, id string, data worker.Entity) (err error) {
	args := r.prepareArgs(data)
	if len(args) > 0 {

		out, err := r.db.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": args})
		if err != nil {
			return err
		}

		if out.MatchedCount == 0 {
			return market.ErrorNotFound
		}
	}

	return
}

func (r *WorkerRepository) prepareArgs(data worker.Entity) (args bson.M) {
	if data.FullName != nil {
		args["full_name"] = data.FullName
	}

	if data.Pseudonym != nil {
		args["pseudonym"] = data.Pseudonym
	}

	if data.Description != nil {
		args["description"] = data.Description
	}

	if data.Position != nil {
		args["position"] = data.Pseudonym
	}

	return
}

func (r *WorkerRepository) Delete(ctx context.Context, id string) (err error) {
	out, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if out.DeletedCount == 0 {
		return market.ErrorNotFound
	}

	return
}
