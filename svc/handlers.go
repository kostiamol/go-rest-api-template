package svc

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kostiamol/go-rest-api-template/entities"
)

// healthcheck stores information about service' name and version
// swagger:response healthcheck
type healthcheck struct {
	// Service name
	SvcName string `json:"svcName"`
	// Version
	Version string `json:"version"`
}

// Status is used to produce different types of statuses with the same structure
// swagger:response status
type status struct {
	// HTTP status code
	Status string `json:"status"`
	// The status message
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
	// swagger:route GET /healthcheck service healthcheck
	//
	// Shows the service status.
	//
	// Checks whether the service is up and running.
	//
	//     Responses:
	//       200: healthcheck

	check := healthcheck{
		SvcName: "go-rest-api-template",
		Version: ctx.Version,
	}
	ctx.Render.JSON(w, http.StatusOK, check)
}

// users holds the map with the list of users and their quantity
// swagger:response users
type users map[string]interface{}

// ListUsersHandler returns a list of users
func ListUsersHandler(w http.ResponseWriter, req *http.Request, ctx Context) {
	// swagger:route GET /users users listUsers
	//
	// Lists users.
	//
	// This will show all available users.
	//
	//     Responses:
	//       200: users
	//       404: status

	list, err := ctx.DB.ListUsers()
	if err != nil {
		response := status{
			Status:  "404",
			Message: "can't find any users",
		}
		log.Println(err)
		ctx.Render.JSON(w, http.StatusNotFound, response)
		return
	}
	// responseObject := make(map[string]interface{})
	responseObject := users(make(map[string]interface{}))
	responseObject["users"] = list
	responseObject["count"] = len(list)
	ctx.Render.JSON(w, http.StatusOK, responseObject)
}

// GetUserHandler returns a user object
func GetUserHandler(w http.ResponseWriter, req *http.Request, ctx Context) {
	// swagger:route GET /users/{uid:[0-9]+} users listUsers
	//
	// Shows the user by uid.
	//
	// This will show the user with the specified uid.
	//
	//     Responses:
	//       200: user
	//       404: status

	vars := mux.Vars(req)
	uid, _ := strconv.Atoi(vars["uid"])
	user, err := ctx.DB.GetUser(uid)
	if err != nil {
		response := status{
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
	// swagger:route POST /users users createUser
	//
	// Creates the user.
	//
	// This will create the user.
	//
	//     Responses:
	//       201: user
	//       400: status

	decoder := json.NewDecoder(req.Body)
	var u entities.User
	err := decoder.Decode(&u)
	if err != nil {
		response := status{
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
	// swagger:route PUT /users users updateUser
	//
	// Updates the user.
	//
	// This will update the user.
	//
	//     Responses:
	//       200: user
	//       400: status
	//		 500: status

	decoder := json.NewDecoder(req.Body)
	var u entities.User
	err := decoder.Decode(&u)
	if err != nil {
		response := status{
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
		response := status{
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
	// swagger:route DELETE /users users deleteUser
	//
	// Deletes the user.
	//
	// This will delete the user.
	//
	//     Responses:
	//       204: status
	//		 500: status

	vars := mux.Vars(req)
	uid, _ := strconv.Atoi(vars["uid"])
	err := ctx.DB.DeleteUser(uid)
	if err != nil {
		response := status{
			Status:  "500",
			Message: "something went wrong",
		}
		log.Println(err)
		ctx.Render.JSON(w, http.StatusInternalServerError, response)
		return
	}
	ctx.Render.JSON(w, http.StatusNoContent, status{})
}

// PassportsHandler not implemented yet
func PassportsHandler(w http.ResponseWriter, req *http.Request, ctx Context) {
	log.Println("Handling Passports - Not implemented yet")
	ctx.Render.Text(w, http.StatusNotImplemented, "")
}
