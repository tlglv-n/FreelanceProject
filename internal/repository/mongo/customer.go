package mongo

import (
	"context"
	"errors"
	"exchanger/internal/domain/customer"
	"exchanger/pkg/market"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository struct {
	db *mongo.Collection
}

func NewCustomerRepository(db *mongo.Database) *CustomerRepository {
	return &CustomerRepository{
		db: db.Collection("customers"),
	}
}

func (r *CustomerRepository) List(ctx context.Context) (dest []customer.Entity, err error) {
	cur, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &dest); err != nil {
		return nil, err
	}

	return
}

func (r *CustomerRepository) Add(ctx context.Context, data customer.Entity) (id string, err error) {
	res, err := r.db.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).String(), nil
}

func (r *CustomerRepository) Get(ctx context.Context, id string) (dest customer.Entity, err error) {
	if err = r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&dest); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			err = market.ErrorNotFound
		}
	}

	return
}

func (r *CustomerRepository) Update(ctx context.Context, id string, data customer.Entity) (err error) {
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

func (r *CustomerRepository) prepareArgs(data customer.Entity) (args bson.M) {
	if data.FullName != nil {
		args["full_name"] = data.FullName
	}

	if data.Pseudonym != nil {
		args["pseudonym"] = data.Pseudonym
	}

	return
}

func (r *CustomerRepository) Delete(ctx context.Context, id string) (err error) {
	out, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if out.DeletedCount == 0 {
		return market.ErrorNotFound
	}

	return
}
