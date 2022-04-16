package orm

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Model primary key is auto increment
type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `gorm:"index" json:"created_at"`
	UpdatedAt time.Time      `gorm:"index" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// UModel primary key is uuid model
type UModel struct {
	ID        string         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `gorm:"index" json:"created_at"`
	UpdatedAt time.Time      `gorm:"index" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (m *UModel) BeforeCreate(tx *gorm.DB) error {
	m.ID = uuid.NewString()
	return nil
}

type Operator struct {
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
	DeletedBy string `json:"deleted_by"`
}
