package policy

import (
	"context"
	"time"
)

type Repository interface {
	Get(ctx context.Context, id string) (Policy, error)
	List(ctx context.Context) ([]Policy, error)
	Create(ctx context.Context, pol Policy) (string, error)
	Update(ctx context.Context, pol Policy) (string, error)
}

type AuthzRepository interface {
	Add(ctx context.Context, policies []Policy) error
}

type Policy struct {
	ID          string
	RoleID      string
	NamespaceID string
	ActionID    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Filters struct {
	NamespaceID string
}
