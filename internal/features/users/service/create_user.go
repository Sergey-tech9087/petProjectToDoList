package users_service

import (
	"context"
	"fmt"

	"github.com/Sergey-tech9087/petProjectToDoList/internal/core/domain"
)

func (s *UserService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("Validate user domain: %w", err)
	}

	user, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("Create user: %w", err)
	}

	return user, nil
}
