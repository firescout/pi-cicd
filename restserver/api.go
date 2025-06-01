package restserver

import (
	"context"
)

type DefaultApiServicer interface {
	OnPush(ctx context.Context, repo string) (ImplResponse, error)
	GetShutdown(context.Context) (ImplResponse, error)
}
