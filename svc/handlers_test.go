package svc

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthHandler(t *testing.T) {
	ctx := NewContext()
	r, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	makeHandler(ctx, HealthHandler).ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code, "they should be equal")
	assert.Equal(t, "application/json; charset=UTF-8", w.HeaderMap["Content-Type"][0], "they should be equal")
	// parse json body
	var f interface{}
	json.Unmarshal(w.Body.Bytes(), &f)
	obj := f.(map[string]interface{})
	assert.Equal(t, "go-rest-api-template", obj["svcName"], "they should be equal")
	assert.Equal(t, ctx.Version, obj["version"], "they should be equal")
}

func TestListUsersHandler(t *testing.T) {
	ctx := NewContext()
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	makeHandler(ctx, ListUsersHandler).ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "they should be equal")
	assert.Equal(t, "application/json; charset=UTF-8", w.HeaderMap["Content-Type"][0], "they should be equal")
	//parse json body
	var f interface{}
	json.Unmarshal(w.Body.Bytes(), &f)
	m := f.(map[string]interface{})
	log.Println(m["users"])
}
