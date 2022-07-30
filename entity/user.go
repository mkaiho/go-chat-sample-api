package entity

import (
	"errors"
	"fmt"
)

type UserID string

func ParseUserID(val string) (UserID, error) {
	id := UserID(val)
	if err := id.Validate(); err != nil {
		return "", fmt.Errorf("invalid user id: %w", err)
	}
	return id, nil
}

func (id UserID) Validate() error {
	if id.isEmpty() {
		return errors.New("empty")
	}
	return nil
}

func (id UserID) isEmpty() bool {
	return len(id) == 0
}

type User struct {
	id          UserID
	name        string
	displayName string
}

func NewUser(
	ID UserID,
	Name string,
	DisplayName string,
) *User {
	return &User{
		id:          ID,
		name:        Name,
		displayName: DisplayName,
	}
}

func (u *User) Validate() error {
	if err := u.id.Validate(); err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}
	if len(u.name) == 0 {
		return fmt.Errorf("invalid user name: %w", errors.New("empty"))
	}
	if len(u.displayName) == 0 {
		return fmt.Errorf("invalid user displayName: %w", errors.New("empty"))
	}
	return nil
}

type Users []*User
