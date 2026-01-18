package repository

import "context"

type UserRepository interface {
	Find(ctx context.Context) error
}
