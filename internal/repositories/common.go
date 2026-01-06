package repositories

import (
	"boilerplate-api/internal/common/context"
	"boilerplate-api/internal/models"

	"gorm.io/gorm"
)

type CommonRepository struct {
	DB *gorm.DB
}

func NewCommonRepository(ctx *context.AppContext) *CommonRepository {
	return &CommonRepository{
		DB: ctx.DB,
	}
}

func (r *CommonRepository) FindCountryDialCodeByDialCode(dialCode string) (*models.CountryDialCode, error) {
	var countryDialCode models.CountryDialCode
	if err := r.DB.Where("dial_code = ?", dialCode).First(&countryDialCode).Error; err != nil {
		return nil, err
	}
	return &countryDialCode, nil
}
