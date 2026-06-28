package repository

import "gorm.io/gorm"

type Repository[T any] struct{}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) FindAll(db *gorm.DB) ([]*T, error) {
	var entities []*T
	err := db.Find(&entities).Error
	return entities, err
}

func (r *Repository[T]) FindById(db *gorm.DB, ID uint) (*T, error) {
	entity := new(T)
	err := db.First(entity, ID).Error
	return entity, err
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}
