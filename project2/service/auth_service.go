package service

import (
	"context"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/model/web"
)

type AuthService interface {
	GenerateToken(ctx context.Context, request web.RequestAuth) web.ResponseToken
	Login(ctx context.Context, request web.RequestAuth) bool
}
