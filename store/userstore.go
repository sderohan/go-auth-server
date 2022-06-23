package store

import (
	"errors"
	"sync"
)

type Userstore interface {
	Save(*User) error
	FindUser(username string) (*User, error)
}

/*
	NOTE: As of now map us used to store the users, need to rework on this
*/
type InMemoryUserStore struct {
	mutex sync.RWMutex
	users map[string]*User
}

func (store *InMemoryUserStore) Finduser(username string) (*User, error) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	user, ok := store.users[username]
	if !ok {
		return nil, errors.New("user with given username doesn't exists, please consider registering first")
	}
	return user, nil
}

func (store *InMemoryUserStore) Save(user *User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.users[user.Username] != nil {
		return errors.New("User already exists")
	}

	store.users[user.Username] = user.Clone()
	return nil
}
