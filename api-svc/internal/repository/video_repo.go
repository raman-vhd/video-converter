package repository

import (
	"github.com/raman-vhd/video-converter/internal/lib"
	"github.com/raman-vhd/video-converter/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type IVideoRepository interface {
	CreateVideo(ctx context.Context, videoID string, ext string, size int) error
	AddVersion(ctx context.Context, videoID string, quality string) error
	GetVideo(ctx context.Context, videoID string) (model.Video, error)
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

func (r videoRepository) CreateVideo(ctx context.Context, videoID string, ext string, size int) error {
	_, err := r.db.InsertOne(ctx, model.Video{
		VideoID:  videoID,
		Ext:      ext,
		Size:     size,
		Versions: map[string]model.ConvertedVideo{},
	})
	if err != nil {
		return err
	}
	return nil
}

func (r videoRepository) AddVersion(ctx context.Context, videoID string, quality string) error {
	_, err := r.db.UpdateOne(ctx,
		bson.M{"videoid": videoID},
		bson.M{"$set": bson.M{
			"versions." + quality + ".state": "pending",
		}})
	return err
}

func (r videoRepository) GetVideo(ctx context.Context, videoID string) (model.Video, error) {
	var video model.Video
	err := r.db.FindOne(ctx, bson.M{"videoid": videoID}).Decode(&video)
	if err != nil {
		return model.Video{}, err
	}
	return video, nil
}
