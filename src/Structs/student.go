package structs

import "gorm.io/gorm"

type Student struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `validate:"required"`
	Age            int32  `validate:"required,gte=0,lte=130"`
	CurrentPayment float64
	BalancePayment float64
	gorm.Model
}
