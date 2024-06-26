package controller

import "context"

type Controller interface {
	Start(ctx context.Context) error
	ShutDown(ctx context.Context) error
}
