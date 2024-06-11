package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/melkdesousa/wottva/users/pkg/entities"
)

type PgUserRepo struct {
	db *sql.DB
}

func NewPgUserRepo(db *sql.DB) *PgUserRepo {
	return &PgUserRepo{
		db: db,
	}
}

func (repo *PgUserRepo) Save(u *entities.User) error {
	_, err := repo.db.Exec("INSERT INTO users(id, name, created_at, updated_at) VALUES($1, $2, $3, $4);", u.ID, u.Name, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (repo *PgUserRepo) List() []entities.User {
	rows, err := repo.db.QueryContext(context.TODO(), "SELECT id, name FROM users;")

	if err != nil {
		log.Fatalf("Database query failed because %s", err)
	}

	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User

		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			log.Fatalf("Database scan failed because %s", err)
		}
		users = append(users, user)
	}

	return users
}

func (repo *PgUserRepo) Get(id string) (entities.User, error) {
	var user entities.User

	row := repo.db.QueryRowContext(context.TODO(), "SELECT id, name FROM users WHERE id = $1;", id)

	if err := row.Scan(&user.ID, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("user %s: no such album", id)
		}
		return user, fmt.Errorf("user %s: %v", id, err)
	}

	return user, nil
}
