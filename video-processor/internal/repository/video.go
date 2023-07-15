package repository

import (
	"context"

	"github.com/raman-vhd/video-converter/internal/lib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IVideoRepository interface {
	UpdateVideoState(videoID string, quality string, state string, size int) error
}

type videoRepository struct {
	db *mongo.Collection
}

func NewVideo(
	db lib.Database,
) IVideoRepository {
	dbColletion := db.GetCollection("video")
	return videoRepository{
		db: dbColletion,
	}
}

func (r videoRepository) UpdateVideoState(videoID string, quality string, state string, size int) error {
	ctx := context.Background()
	_, err := r.db.UpdateOne(ctx,
		bson.M{"videoid": videoID},
		bson.M{"$set": bson.M{
			"versions." + quality + ".state": state,
			"versions." + quality + ".size":  size,
		}})
	return err
}
