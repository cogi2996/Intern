package services

import (
	"context"

	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
	"github.com/ideal-forward/assistant-public-api/pkg/e"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IRegister interface {
	Create(ctx context.Context, data *entities.Register) (int64, error)
	Read(ctx context.Context, id int64) (*entities.Register, error)
	Update(ctx context.Context, data *entities.Register) error
	Delete(ctx context.Context, id int64) error
	Authenticate(ctx context.Context, username, password string) (string, error)
}

type Register struct {
	Model models.IRegister
}

func NewRegister() IRegister {
	return &Register{
		Model: models.Register{},
	}
}

func (p *Register) Create(ctx context.Context, data *entities.Register) (int64, error) {
	return p.Model.Create(ctx, data)
}

func (p *Register) Read(ctx context.Context, id int64) (*entities.Register, error) {
	return p.Model.Read(ctx, id)
}

func (p *Register) Update(ctx context.Context, data *entities.Register) error {
	return p.Model.Update(ctx, data)
}

func (p *Register) Delete(ctx context.Context, id int64) error {
	return p.Model.Delete(ctx, id)
}

func (p *Register) Authenticate(ctx context.Context, username, password string) (string, error) {
	data, err := p.Model.ReadByUsername(ctx, username)
	if err != nil {
		return "", e.MakeDatabaseError(err, codes.Unauthenticated, "Username isnot found", "Không tìm thấy tên đăng nhập!")
	}

	if data.Password != password {
		st, _ := status.New(codes.Unauthenticated, "Password is incorrect").WithDetails(
			&errdetails.LocalizedMessage{
				Locale:  "vi",
				Message: "Mật khẩu không đúng!",
			},
		)

		return "", st.Err()
	}

	return TokenMaker.CreateToken(data.ID)
}
