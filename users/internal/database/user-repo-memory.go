package database

import (
	"errors"

	"github.com/google/uuid"
	"github.com/melkdesousa/wottva/users/pkg/entities"
)

type InMemoryUserRepo struct {
	users map[string]*entities.User
}

// NewInMemoryUserRepo creates a new InMemoryUserRepo instance
func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]*entities.User),
	}
}

// Save saves a user entity in memory
func (repo *InMemoryUserRepo) Save(u *entities.User) error {
	if u.ID == "" {
		u.ID = uuid.New().String() // Generate a new ID if not provided
	}
	repo.users[u.ID] = u
	return nil
}

// Get retrieves a user entity from memory by its ID
func (repo *InMemoryUserRepo) List() []entities.User {
	var users []entities.User

	for _, user := range repo.users {
		users = append(users, entities.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}

	return users
}

func (repo *InMemoryUserRepo) Get(id string) (entities.User, error) {
	if user, ok := repo.users[id]; ok {
		return *user, nil
	}

	return entities.User{}, errors.New("user not found")
}
