package pkg

import (
	"context"

	mod "github.com/Filimonov-ua-d/test_task/models"
)

const CtxUserKey = "user"

type UseCase interface {
	Login(ctx context.Context, username, password string) (string, error)
	UploadPicture(ctx context.Context, u *mod.User, filename string) error
	GetImages(ctx context.Context, u *mod.User) ([]*mod.Image, error)
	ParseToken(ctx context.Context, accessToken string) (*mod.User, error)
}
