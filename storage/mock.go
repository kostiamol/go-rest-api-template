package storage

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/kostiamol/go-rest-api-template/entities"
	"github.com/palantir/stacktrace"
)

// MockDB will hold the connection and key db info
type MockDB struct {
	UserList  map[int]entities.User
	MaxUserID int
}

// NewMockDB initialises a database for test purposes
func NewMockDB() *MockDB {
	list := make(map[int]entities.User)
	dt, _ := time.Parse(time.RFC3339, "1985-12-31T00:00:00Z")
	list[0] = entities.User{
		ID:              0,
		FirstName:       "John",
		LastName:        "Doe",
		DateOfBirth:     dt,
		LocationOfBirth: "London",
	}
	dt, _ = time.Parse(time.RFC3339, "1992-01-01T00:00:00Z")
	list[1] = entities.User{
		ID:              1,
		FirstName:       "Jane",
		LastName:        "Doe",
		DateOfBirth:     dt,
		LocationOfBirth: "Milton Keynes",
	}
	return &MockDB{
		UserList:  list,
		MaxUserID: 1,
	}
}

// LoadFixturesIntoMockDB loads data from fixtures file into MockDB
func LoadFixturesIntoMockDB(fixturesFile string) (*MockDB, error) {
	var jsonObject map[string][]entities.User
	file, err := ioutil.ReadFile(fixturesFile)
	if err != nil {
		return nil, stacktrace.Propagate(err, "error reading fixtures file")
	}
	err = json.Unmarshal(file, &jsonObject)
	if err != nil {
		return nil, stacktrace.Propagate(err, "error parsing fixtures file")
	}
	list := make(map[int]entities.User)
	list[0] = jsonObject["users"][0]
	list[1] = jsonObject["users"][1]
	return &MockDB{
		UserList:  list,
		MaxUserID: 1,
	}, nil
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
