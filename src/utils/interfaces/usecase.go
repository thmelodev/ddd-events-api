package interfaces

import "context"

type IUsecase interface {
	Execute(ctx context.Context, dto any) (any, error)
}
