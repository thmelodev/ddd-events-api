package interfaces

import "context"

type IQuery interface {
	Execute(ctx context.Context, dto any) (any, error)
}
