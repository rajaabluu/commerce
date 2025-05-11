package repository

import "gorm.io/gorm"

type Repository[T any] struct{}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) FindAll(db *gorm.DB, entities *[]*T) error {
	return db.Find(&entities).Error
}

func (r *Repository[T]) FindById(db *gorm.DB, ID uint, entity *T) error {
	return db.First(entity, ID).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}
