package repository

import (
	"errors"

	"gorm.io/gorm"
)

// BaseRepository provides common CRUD operations using GORM
type BaseRepository struct {
	db *gorm.DB
}

// NewBaseRepository creates a new BaseRepository
func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{
		db: db,
	}
}

// GetByID retrieves an entity by ID
func (r *BaseRepository) GetByID(id int64, dest interface{}) error {
	result := r.db.First(dest, id)
	return result.Error
}

// Create inserts a new entity
func (r *BaseRepository) Create(entity interface{}) error {
	result := r.db.Create(entity)
	return result.Error
}

// Update updates an existing entity
func (r *BaseRepository) Update(entity interface{}) error {
	result := r.db.Save(entity)
	return result.Error
}

// Delete removes an entity by ID
func (r *BaseRepository) Delete(model interface{}, id int64) error {
	result := r.db.Delete(model, id)
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return result.Error
}

// List retrieves entities with pagination
func (r *BaseRepository) List(dest interface{}, page, pageSize int, where ...interface{}) error {
	offset := (page - 1) * pageSize

	query := r.db.Offset(offset).Limit(pageSize)
	if len(where) > 0 {
		query = query.Where(where[0], where[1:]...)
	}

	result := query.Find(dest)
	return result.Error
}
