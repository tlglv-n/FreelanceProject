package mongo

import (
	"context"
	"errors"
	"exchanger/internal/domain/hire"
	"exchanger/pkg/market"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HireRepository struct {
	db *mongo.Collection
}

func NewHireRepository(db *mongo.Database) *HireRepository {
	return &HireRepository{
		db: db.Collection("hires"),
	}
}

func (r *HireRepository) List(ctx context.Context) (dest []hire.Entity, err error) {
	cur, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &dest); err != nil {
		return nil, err
	}

	return
}

func (r *HireRepository) Add(ctx context.Context, data hire.Entity) (id string, err error) {
	res, err := r.db.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).String(), nil
}

func (r *HireRepository) Get(ctx context.Context, id string) (dest hire.Entity, err error) {
	if err = r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&dest); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = market.ErrorNotFound
		}
	}

	return
}

func (r *HireRepository) Update(ctx context.Context, id string, data hire.Entity) (err error) {
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

func (r *HireRepository) prepareArgs(data hire.Entity) (args bson.M) {
	if data.JobName != nil {
		args["job_name"] = data.JobName
	}

	if data.Amount != nil {
		args["amount"] = data.Amount
	}

	if data.Position != nil {
		args["position"] = data.Position
	}

	if data.Description != nil {
		args["description"] = data.Description
	}

	return
}

func (r *HireRepository) Delete(ctx context.Context, id string) (err error) {
	out, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if out.DeletedCount == 0 {
		return market.ErrorNotFound
	}

	return
}
