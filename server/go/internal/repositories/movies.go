package repositories

import (
	"context"

	"github.com/tamadamas/magic_stream/server/go/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MoviesRepository struct {
	db  *mongo.Database
	col *mongo.Collection
}

func NewMoviesRepository(db *mongo.Database) *MoviesRepository {
	return &MoviesRepository{
		db:  db,
		col: db.Collection("movies"),
	}
}

func (r *MoviesRepository) All(ctx context.Context) ([]models.Movie, error) {
	var movies []models.Movie

	cur, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var m models.Movie
		if err := cur.Decode(&m); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, cur.Err()
}

func (r *MoviesRepository) Find(ctx context.Context, id string) (*models.Movie, error) {
	var movie models.Movie

	err := r.col.FindOne(ctx, bson.M{"_id": id}).Decode(&movie)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}
