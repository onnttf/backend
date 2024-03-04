package dao

import (
	"backend/dal/model"
	"gorm.io/gorm"
)

type Entity interface {
	model.PetrolPrice
}

type DAO[T Entity] interface {
	Insert(newValue T) (*T, error)
	Update(newValue T, funcs ...func(*gorm.DB) *gorm.DB) error
	QueryOne(funcs ...func(*gorm.DB) *gorm.DB) (*T, error)
	QueryList(funcs ...func(*gorm.DB) *gorm.DB) ([]T, error)
}

type Dao[T Entity] struct {
	database *gorm.DB
}

func NewDao[T Entity](database *gorm.DB) *Dao[T] {
	return &Dao[T]{database: database}
}

func (dao *Dao[T]) Insert(newValue T) (*T, error) {
	if err := dao.database.Create(&newValue).Error; err != nil {
		return nil, err
	}
	return &newValue, nil
}

func (dao *Dao[T]) Update(newValue T, funcs ...func(*gorm.DB) *gorm.DB) error {
	return dao.database.Model(&newValue).Scopes(funcs...).Updates(&newValue).Error
}

func (dao *Dao[T]) QueryOne(funcs ...func(*gorm.DB) *gorm.DB) (*T, error) {
	var record T
	result := dao.database.Scopes(funcs...).Limit(1).Find(&record)
	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, nil
	}
	return &record, nil
}

func (dao *Dao[T]) QueryList(funcs ...func(*gorm.DB) *gorm.DB) ([]T, error) {
	var recordList []T
	result := dao.database.Scopes(funcs...).Find(&recordList)
	if result.Error != nil {
		return nil, result.Error
	}
	return recordList, nil
}

func Paginate(pageNo, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNo <= 0 {
			pageNo = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (pageNo - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
