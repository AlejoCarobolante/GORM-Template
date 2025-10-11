package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Entity1 struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Entity1Repository interface {
	Create(e context.Context, entity1 Entity1) error
	Fetch(e context.Context) ([]Entity1, error)
	FetchByID(e context.Context, id string) (Entity1, error)
	Update(e context.Context, updatedEntity1 Entity1) error
	Delete(e context.Context, id string) error
}
