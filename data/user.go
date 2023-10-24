package data

import (
	"time"

	up "github.com/upper/db/v4"
)

type User struct {
	ID        int       `db:"id,omitempty"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Active    int       `db:"user_active"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Token     Token     `db:"-"`
}

func (u *User) Table() string {
	return "users"
}

func (u *User) GetAll(condition up.Cond) ([]*User, error) {
	collection := upper.Collection(u.Table())

	var all []*User

	res := collection.Find(condition)
	err := res.All(&all)
	if err != nil {
		return nil, err
	}

	return all, nil
}

func (u *User) GetByEmail(email string) (*User, error) {
	collection := upper.Collection(u.Table())

	var user User

	res := collection.Find(up.Cond{"email =": email})
	err := res.One(&user)
	if err != nil {
		return nil, err
	}

	var token Token
	tokenCollection := upper.Collection(token.Table())
	tokenRes := tokenCollection.Find(up.Cond{"user_id =": user.ID, "expiry <": time.Now()}).OrderBy("created_at DESC")
	err = tokenRes.One(&token)
	if err != nil {
		if err != up.ErrNilRecord && err != up.ErrNoMoreRows {
			return nil, err
		}
	}

	user.Token = token

	return &user, nil
}
