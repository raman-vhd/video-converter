package repository

import (
	"github.com/raman-vhd/video-converter/internal/lib"
	"go.mongodb.org/mongo-driver/mongo"
)

type ITemplateRepository interface{}

type templateRepository struct {
	db *mongo.Collection
}

func NewTemplate(
	db lib.Database,
) ITemplateRepository {
	dbColletion := db.GetCollection("template")
	return templateRepository{
		db: dbColletion,
	}
}
