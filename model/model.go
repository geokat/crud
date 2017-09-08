package model

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"sync"
)

type User struct {
	Id    string
	Email string
	Name  string
}

type Users struct {
	// Maps are not thread safe
	sync.RWMutex
	m map[string]User
}

var users Users

func init() {
	users = Users{m: make(map[string]User)}
}

func GetUsers() ([]byte, error) {
	uu := []User{}

	users.RLock()
	for _, u := range users.m {
		uu = append(uu, u)
	}
	users.RUnlock()

	return json.Marshal(uu)
}

func GetUser(id string) ([]byte, error) {
	users.RLock()
	u, ok := users.m[id]
	users.RUnlock()

	if !ok {
		return nil, fmt.Errorf("User ID doesn't exist: %s", id)
	}

	return json.Marshal(u)
}

func CreateUser(email, name string) error {
	// Create a unique ID by hashing user's email address
	h := md5.New()
	io.WriteString(h, email)
	id := fmt.Sprintf("%x", h.Sum(nil))

	// Lock for writing to avoid a race condition where another
	// request creates the same user before us
	users.Lock()
	defer users.Unlock()

	if _, ok := users.m[id]; ok {
		return fmt.Errorf("User email already registered: %s", email)
	}

	users.m[id] = User{
		Id:    id,
		Email: email,
		Name:  name,
	}

	return nil
}

func UpdateUser(id, email, name string) error {
	users.Lock()
	defer users.Unlock()

	if _, ok := users.m[id]; !ok {
		return fmt.Errorf("User ID does not exist: %s", id)
	}

	users.m[id] = User{
		Id:    id,
		Email: email,
		Name:  name,
	}

	return nil
}

func DeleteUser(id string) error {
	users.Lock()
	defer users.Unlock()

	if _, ok := users.m[id]; !ok {
		return fmt.Errorf("User ID does not exist: %s", id)
	}

	delete(users.m, id)

	return nil
}

/*
 * Local variables:
 * compile-command: "go build github.com/geokat/crud":
 * End:
 */
