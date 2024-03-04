package model

import (
	"time"
)

const TableNamePetrolPrice = "petrol_price"

type PetrolPrice struct {
	Id          int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT"`
	Region      string    `gorm:"column:region;type:varchar(255);NOT NULL"`
	Price0      string    `gorm:"column:price_0;type:varchar(5);default:0:00;NOT NULL"`
	Price92     string    `gorm:"column:price_92;type:varchar(5);default:0:00;NOT NULL"`
	Price95     string    `gorm:"column:price_95;type:varchar(5);default:0:00;NOT NULL"`
	Price98     string    `gorm:"column:price_98;type:varchar(5);default:0:00;NOT NULL"`
	ReleaseDate string    `gorm:"column:release_date;type:varchar(10);NOT NULL"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (*PetrolPrice) TableName() string {
	return TableNamePetrolPrice
}
