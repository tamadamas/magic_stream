package repositories

import (
	"context"
	"errors"

	"github.com/tamadamas/magic_stream/server/go/internal/app_errors"
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

	if err = cur.All(ctx, &movies); err != nil {
		return nil, errors.New("Failed to get movies")
	}

	return movies, nil
}

func (r *MoviesRepository) Find(ctx context.Context, id string) (*models.Movie, error) {
	var movie models.Movie

	err := r.col.FindOne(ctx, bson.M{"imdb_id": id}).Decode(&movie)

	if err != nil {
		return nil, app_errors.NewNotFoundError(nil, "Movie not found")
	}

	return &movie, nil
}
