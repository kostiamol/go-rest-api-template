package svc

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kostiamol/go-rest-api-template/entities"
)

// HealthCheck will store information about its name and version
type HealthCheck struct {
	SvcName string `json:"svcName"`
	Version string `json:"version"`
}

// Status is a custom response object we pass around the system and send back to the customer
// 404: Not found
// 500: Internal Server Error
type Status struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HandlerFunc is a custom implementation of the http.HandlerFunc
type HandlerFunc func(http.ResponseWriter, *http.Request, Context)

// makeHandler allows us to pass an environment struct to our handlers, without resorting to global
// variables. It accepts an environment (Env) struct and our own handler function. It returns
// a function of the type http.HandlerFunc so can be passed on to the HandlerFunc in main.go.
func makeHandler(ctx Context, fn func(http.ResponseWriter, *http.Request, Context)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, ctx)
	}
}

// HealthCheckHandler returns useful info about the app
func HealthCheckHandler(w http.ResponseWriter, req *http.Request, ctx Context) {
	check := HealthCheck{
		SvcName: "go-rest-api-template",
		Version: ctx.Version,
	}
	ctx.Render.JSON(w, http.StatusOK, check)
}

// ListUsersHandler returns a list of users
func ListUsersHandler(w http.ResponseWriter, req *http.Request, ctx Context) {
	list, err := ctx.DB.ListUsers()
	if err != nil {
		response := Status{
			Status:  "404",
			Message: "can't find any users",
		}
		log.Println(err)
		ctx.Render.JSON(w, http.StatusNotFound, response)
		return
	}
	responseObject := make(map[string]interface{})
	responseObject["users"] = list
	responseObject["count"] = len(list)
	ctx.Render.JSON(w, http.StatusOK, responseObject)
}

// GetUserHandler returns a user object
func GetUserHandler(w http.ResponseWriter, req *http.Request, ctx Context) {
	vars := mux.Vars(req)
	uid, _ := strconv.Atoi(vars["uid"])
	user, err := ctx.DB.GetUser(uid)
	if err != nil {
		response := Status{
			Status:  "404",
			Message: "can't find user",
		}
		log.Println(err)
		ctx.Render.JSON(w, http.StatusNotFound, response)
		return
	}
	ctx.Render.JSON(w, http.StatusOK, user)
}

// CreateUserHandler adds a new user
func CreateUserHandler(w http.ResponseWriter, req *http.Request, ctx Context) {
	decoder := json.NewDecoder(req.Body)
	var u entities.User
	err := decoder.Decode(&u)
	if err != nil {
		response := Status{
			Status:  "400",
			Message: "malformed user object",
		}
		log.Println(err)
		ctx.Render.JSON(w, http.StatusBadRequest, response)
		return
	}
	user := entities.User{
		ID:              -1,
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		DateOfBirth:     u.DateOfBirth,
		LocationOfBirth: u.LocationOfBirth,
	}
	user, _ = ctx.DB.AddUser(user)
	ctx.Render.JSON(w, http.StatusCreated, user)
}

// UpdateUserHandler updates a user object
func UpdateUserHandler(w http.ResponseWriter, req *http.Request, ctx Context) {
	decoder := json.NewDecoder(req.Body)
	var u entities.User
	err := decoder.Decode(&u)
	if err != nil {
		response := Status{
			Status:  "400",
			Message: "malformed user object",
		}
		log.Println(err)
		ctx.Render.JSON(w, http.StatusBadRequest, response)
		return
	}
	user := entities.User{
		ID:              u.ID,
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		DateOfBirth:     u.DateOfBirth,
		LocationOfBirth: u.LocationOfBirth,
	}
	user, err = ctx.DB.UpdateUser(user)
	if err != nil {
		response := Status{
			Status:  "500",
			Message: "something went wrong",
		}
		log.Println(err)
		ctx.Render.JSON(w, http.StatusInternalServerError, response)
		return
	}
	ctx.Render.JSON(w, http.StatusOK, user)
}

// DeleteUserHandler deletes a user
func DeleteUserHandler(w http.ResponseWriter, req *http.Request, ctx Context) {
	vars := mux.Vars(req)
	uid, _ := strconv.Atoi(vars["uid"])
	err := ctx.DB.DeleteUser(uid)
	if err != nil {
		response := Status{
			Status:  "500",
			Message: "something went wrong",
		}
		log.Println(err)
		ctx.Render.JSON(w, http.StatusInternalServerError, response)
		return
	}
	ctx.Render.Text(w, http.StatusNoContent, "")
}

// PassportsHandler not implemented yet
func PassportsHandler(w http.ResponseWriter, req *http.Request, ctx Context) {
	log.Println("Handling Passports - Not implemented yet")
	ctx.Render.Text(w, http.StatusNotImplemented, "")
}
