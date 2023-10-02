package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
)

type IUser interface {
	Create(ctx context.Context, data *entities.User) (int64, error)
	Read(ctx context.Context, id int64) (*entities.User, error)
	Update(ctx context.Context, data *entities.User) error
	Delete(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*entities.User, error)
}

type User struct {
	Model models.IUser
}

func NewUser() IUser {
	return &User{
		Model: models.User{},
	}
}

func (p *User) Create(ctx context.Context, data *entities.User) (int64, error) {
	data.UUID = uuid.NewString()
	return p.Model.Create(ctx, data)
}

func (p *User) Read(ctx context.Context, id int64) (*entities.User, error) {
	return p.Model.Read(ctx, id)
}

func (p *User) Update(ctx context.Context, data *entities.User) error {
	return p.Model.Update(ctx, data)
}

func (p *User) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *User) List(ctx context.Context) ([]*entities.User, error) {
	return p.Model.List(ctx)
}
