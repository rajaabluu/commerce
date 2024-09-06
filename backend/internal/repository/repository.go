package repository

import "gorm.io/gorm"

type Repository[T any] struct{}

func (repository *Repository[T]) Create(tx *gorm.DB, entity *T) error {
	return tx.Create(entity).Error
}

func (repository *Repository[T]) Find(tx *gorm.DB, entity *T, entities *[]*T) error {
	return tx.Where(entity).Find(entities).Error
}

func (repository *Repository[T]) FindOne(tx *gorm.DB, entity *T) error {
	return tx.Where(entity).First(entity).Error
}

func (repository *Repository[T]) FindById(tx *gorm.DB, ID uint, entity *T) error {
	return tx.Where("id = ?", ID).First(entity).Error
}

func (repository *Repository[T]) Update(tx *gorm.DB, entity *T) error {
	return tx.Save(entity).Error
}

func (repository *Repository[T]) Delete(tx *gorm.DB, entity *T) error {
	return tx.Delete(entity).Error
}
