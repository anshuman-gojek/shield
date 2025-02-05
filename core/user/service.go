package user

import (
	"context"
	"strings"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s Service) GetByID(ctx context.Context, id string) (User, error) {
	return s.repository.GetByID(ctx, id)
}

func (s Service) GetByIDs(ctx context.Context, userIDs []string) ([]User, error) {
	return s.repository.GetByIDs(ctx, userIDs)
}

func (s Service) GetByEmail(ctx context.Context, email string) (User, error) {
	return s.repository.GetByEmail(ctx, email)
}

func (s Service) Create(ctx context.Context, user User) (User, error) {
	newUser, err := s.repository.Create(ctx, User{
		Name:     user.Name,
		Email:    user.Email,
		Metadata: user.Metadata,
	})
	if err != nil {
		return User{}, err
	}

	return newUser, nil
}

func (s Service) CreateMetadataKey(ctx context.Context, key UserMetadataKey) (UserMetadataKey, error) {
	newUserMetadataKey, err := s.repository.CreateMetadataKey(ctx, UserMetadataKey{
		Key:         key.Key,
		Description: key.Description,
	})
	if err != nil {
		return UserMetadataKey{}, err
	}

	return newUserMetadataKey, nil
}

func (s Service) List(ctx context.Context, flt Filter) (PagedUsers, error) {
	users, err := s.repository.List(ctx, flt)
	if err != nil {
		return PagedUsers{}, err
	}
	//TODO might better to do this in handler level
	return PagedUsers{
		Count: int32(len(users)),
		Users: users,
	}, nil
}

func (s Service) UpdateByID(ctx context.Context, toUpdate User) (User, error) {
	return s.repository.UpdateByID(ctx, toUpdate)
}

func (s Service) UpdateByEmail(ctx context.Context, toUpdate User) (User, error) {
	return s.repository.UpdateByEmail(ctx, toUpdate)
}

func (s Service) FetchCurrentUser(ctx context.Context) (User, error) {
	email, ok := GetEmailFromContext(ctx)
	if !ok {
		return User{}, ErrMissingEmail
	}

	email = strings.TrimSpace(email)
	if email == "" {
		return User{}, ErrMissingEmail
	}

	fetchedUser, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}

	return fetchedUser, nil
}
