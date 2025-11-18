package mongodb

import (
	"context"
	"fmt"

	"github.com/aaegamysta/card-trader-ex/internal/cards"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Repository struct {
	cfg Config
	mongodbClient *mongo.Client
}

func New(ctx context.Context, cfg Config) (*Repository, error) {
	connectionStr := fmt.Sprintf("mongodb+srv://%s:%s@%s.jjegak4.mongodb.net/?appName=%s", cfg.Username, cfg.Password, cfg.Project, cfg.Cluster)
	mongodbClient, err := mongo.Connect(options.Client().ApplyURI(connectionStr))
	if err != nil {
		return nil, err
	}

	r :=  &Repository{
		cfg:           cfg,
		mongodbClient: mongodbClient,
	}

	return r, nil
}