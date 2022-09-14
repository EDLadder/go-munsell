package organization

import "context"

type Storage interface {
	Create(ctx context.Context, organization Organization) (string, error)
}
