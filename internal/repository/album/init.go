package album

import "gorm.io/gorm"

func NewRepository(db *gorm.DB) Repository {
	return &albumRepository{
		DB: db,
	}
}
