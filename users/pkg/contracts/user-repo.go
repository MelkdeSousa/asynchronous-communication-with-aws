package contracts

import "github.com/melkdesousa/wottva/users/pkg/entities"

type UserRepo interface {
	Save(u *entities.User) error
	List() []entities.User
	Get(id string) (entities.User, error)
}
