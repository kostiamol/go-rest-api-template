package svc

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/kostiamol/go-rest-api-template/entities"
	"github.com/kostiamol/go-rest-api-template/storage"
	"github.com/palantir/stacktrace"
	"github.com/unrolled/render"
)

// Local refers to local environment
const (
	Local string = "LOCAL"
	Prod  string = "PROD"
)

// Storager defines all the database operations
type Storager interface {
	ListUsers() ([]entities.User, error)
	GetUser(i int) (entities.User, error)
	AddUser(u entities.User) (entities.User, error)
	UpdateUser(u entities.User) (entities.User, error)
	DeleteUser(i int) error
}

// Context holds application configuration data
type Context struct {
	Render  *render.Render
	Version string
	Env     string
	Port    string
	DB      Storager
}

// NewContext initialises an application context struct for testing purposes
func NewContext() Context {
	testVersion := "0.0.0"
	db := storage.NewMockDB()
	ctx := Context{
		Render:  render.New(),
		Version: testVersion,
		Env:     Local,
		Port:    "3001",
		DB:      db,
	}
	return ctx
}

// ParseVersionFile returns the version as a string, parsing and validating a file given the path
func ParseVersionFile(versionPath string) (string, error) {
	dat, err := ioutil.ReadFile(versionPath)
	if err != nil {
		return "", stacktrace.Propagate(err, "error reading version file")
	}
	version := string(dat)
	version = strings.Trim(strings.Trim(version, "\n"), " ")
	// regex pulled from official https://github.com/sindresorhus/semver-regex
	semverRegex := `^v?(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)\.(?:0|[1-9][0-9]*)(?:-[\da-z\-]+(?:\.[\da-z\-]+)*)?(?:\+[\da-z\-]+(?:\.[\da-z\-]+)*)?$`
	match, err := regexp.MatchString(semverRegex, version)
	if err != nil {
		return "", stacktrace.Propagate(err, "error executing regex match")
	}
	if !match {
		return "", stacktrace.NewError("string in VERSION is not a valid version number")
	}
	return version, nil
}
