package interfaces

import (
	"github.com/IlyaZayats/managus/internal/entity"
)

type UserRepository interface {
	Login(user entity.User) (int, error)
}
