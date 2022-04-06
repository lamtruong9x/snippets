package models

import (
	"errors"
	"time"
)

var (
	ErrNoRecord = errors.New("models: no matching record found")
	// Add a new ErrInvalidCredentials error: a user
	//tries to login with an incorrect email address or password.
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// Add a new ErrDuplicateEmail error: a user
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)

type Snippet struct {
	ID      int
	Title   string //`validate:"required,max=10"`
	Content string //`validate:"required"`
	Created time.Time
	Expires time.Time //`validate:"required,oneof=1 7 365"`
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
	Active         bool
}

//func (s *Snippet) Validate() error {
//	v := validator.New()
//	return v.Struct(s)
//}
