package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/google/uuid"
	"github.com/melkdesousa/wottva/users/api/dtos"
	"github.com/melkdesousa/wottva/users/pkg/contracts"
	"github.com/melkdesousa/wottva/users/pkg/entities"
)

type UserController struct {
	userRepo contracts.UserRepo
	broker   contracts.Broker
}

var (
	UserWithIDRe = regexp.MustCompile(`^/users/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
	UsersRe      = regexp.MustCompile(`^/users$`)
)

func NewUserController(r contracts.UserRepo, b contracts.Broker) *UserController {
	return &UserController{
		userRepo: r,
		broker:   b,
	}
}

func (c *UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodPost && UsersRe.MatchString(r.URL.Path):
		c.CreateUser(w, r)
		return

	case r.Method == http.MethodGet && UsersRe.MatchString(r.URL.Path):
		c.ListUsers(w, r)
		return

	case r.Method == http.MethodGet && UserWithIDRe.MatchString(r.URL.Path):
		c.GetUser(w, r)
		return

	default:
		return
	}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dtos.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.userRepo.Save(&entities.User{
		ID:   uuid.New().String(),
		Name: user.Name,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userCreatedJson, _ := json.Marshal(user)

	err = c.broker.Publish(contracts.UserCreated, userCreatedJson)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	users := c.userRepo.List()

	response, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	matches := UserWithIDRe.FindStringSubmatch(r.URL.Path)

	if len(matches) != 2 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := c.userRepo.Get(matches[1])
	if err != nil {
		http.Error(w, fmt.Sprintf("User with ID \"%s\" not found", matches[1]), http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
