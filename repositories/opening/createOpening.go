package opening

import (
	"github.com/DamiaoCanndido/gopportunities/entities"
	"gorm.io/gorm"
)

type OpeningRepository interface {
	CreateOpening(opening entities.Opening) entities.Opening
}

type openingConnection struct {
	connection *gorm.DB
}

func NewOpeningRepository(db *gorm.DB) OpeningRepository {
	return &openingConnection{
		connection: db,
	}
}

func (db *openingConnection) CreateOpening(opening entities.Opening) entities.Opening {
	db.connection.Create(&opening)
	return opening
}
