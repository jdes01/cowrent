package postgres

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresWorkspace struct {
	ID          uint
	UUID        uuid.UUID
	CreatedAt   time.Time
	Name        string
	Seats       int
	CoworkingID uint
	Coworking   PostgresCoworking `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TableName   string            `gorm:"tableName:Workspaces"`
}

type CoworkingImage struct {
	gorm.Model
	ID          uint
	Name        string
	URL         string
	CoworkingID uint
	TableName   string `gorm:"tableName:CoworkingImages"`
}
type PostgresCoworking struct {
	gorm.Model
	ID         uint
	UUID       uuid.UUID
	CreatedAt  time.Time
	Name       string
	Images     []CoworkingImage    `gorm:"foreignKey:CoworkingID"`
	Workspaces []PostgresWorkspace `gorm:"foreignKey:CoworkingID"`
	TableName  string              `gorm:"tableName:Coworkings"`
}
