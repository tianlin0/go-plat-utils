package lredis

import (
	"context"
)

type Lock interface {
	TryLock(ctx context.Context) (bool, error)
	Lock(ctx context.Context) (bool, error)
	UnLock(ctx context.Context) (bool, error)
}
