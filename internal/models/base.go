package models

import (
	"time"

	"boilerplate-api/internal/common/errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	BaseModelTableName = "base_models"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

func (base *BaseModel) TableName() string {
	return BaseModelTableName
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return nil
}

func (base *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	base.UpdatedAt = time.Now()
	return nil
}

func (base *BaseModel) Delete(tx *gorm.DB) error {
	if base.DeletedAt != nil {
		return errors.ErrNotFound
	}

	now := time.Now()
	base.DeletedAt = &now

	return nil
}

func (base *BaseModel) IsDeleted() bool {
	return base.DeletedAt != nil
}

type BaseModelWithID struct {
	ID uuid.UUID `json:"id"`
}

func (base *BaseModelWithID) TableName() string {
	return BaseModelTableName
}
