package storage

import (
	"github.com/kostiamol/go-rest-api-template/entities"
	"github.com/palantir/stacktrace"
)

// MockDB will hold the connection and key db info
type MockDB struct {
	UserList  map[int]entities.User
	MaxUserID int
}

// ListUsers returns a list of JSON documents
func (db *MockDB) ListUsers() ([]entities.User, error) {
	var list []entities.User
	for _, v := range db.UserList {
		list = append(list, v)
	}
	return list, nil
}

// GetUser returns a single JSON document
func (db *MockDB) GetUser(i int) (entities.User, error) {
	user, ok := db.UserList[i]
	if !ok {
		return entities.User{}, stacktrace.NewError("Failure trying to retrieve user")
	}
	return user, nil
}

// AddUser adds a User JSON document, returns the JSON document with the generated id
func (db *MockDB) AddUser(u entities.User) (entities.User, error) {
	db.MaxUserID = db.MaxUserID + 1
	u.ID = db.MaxUserID
	db.UserList[db.MaxUserID] = u
	return u, nil
}

// UpdateUser updates an existing user
func (db *MockDB) UpdateUser(u entities.User) (entities.User, error) {
	id := u.ID
	_, ok := db.UserList[id]
	if !ok {
		return u, stacktrace.NewError("Failure trying to update user")
	}
	db.UserList[id] = u
	return db.UserList[id], nil
}

// DeleteUser deletes a user
func (db *MockDB) DeleteUser(i int) error {
	_, ok := db.UserList[i]
	if !ok {
		return stacktrace.NewError("Failure trying to delete user")
	}
	delete(db.UserList, i)
	return nil
}
