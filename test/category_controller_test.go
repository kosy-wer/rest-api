package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"strconv"
	"fmt"
	"rest_api/api"
	"rest_api/internal/apps/database"
	"rest_api/internal/apps/register/middleware"
	"rest_api/internal/apps/register/controller"
	"rest_api/internal/apps/register/service"
	"rest_api/internal/apps/register/repository"
	"rest_api/internal/apps/register/model/domain"
	"testing"
)

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := api.NewRouter(userController)

	return middleware.NewAuthMiddleware(router)
}

func truncateUser(db *sql.DB) {
	db.Exec("TRUNCATE users")
}

func TestCreateCategorySuccess(t *testing.T) {

	db , err := database.GetConnection()                                  
	if err != nil {                                                       
	   panic(err)                                                                                                                                 
        }
	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "TV"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "TV", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db , err := database.GetConnection()
        if err != nil {                                                                  panic(err)                                                                                                                                               }

	truncateUser(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
    db, err := database.GetConnection()
    if err != nil {
        panic(err)
    }
    defer db.Close()

    truncateUser(db)

    tx, _ := db.Begin()
    userRepository := repository.NewUserRepository()
    user := userRepository.Save(context.Background(), tx, domain.User{
        Name: "Gadget",
    })
    tx.Commit()

    router := setupRouter(db)

    // Simpan ID kategori yang baru ditambahkan
    userID := user.Id

    // Buat permintaan PUT dengan menggunakan ID kategori yang baru ditambahkan
    requestBody := strings.NewReader(`{"name": "Updated Gadget"}`)
    request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/users/"+strconv.Itoa(userID), requestBody)
    request.Header.Add("Content-Type", "application/json")
    request.Header.Add("X-API-Key", "RAHASIA")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    response := recorder.Result()
    assert.Equal(t, 200, response.StatusCode)

    body, _ := io.ReadAll(response.Body)
    var responseBody map[string]interface{}
    json.Unmarshal(body, &responseBody)

    assert.Equal(t, 200, int(responseBody["code"].(float64)))
    assert.Equal(t, "OK", responseBody["status"])
    assert.Equal(t, userID, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
    assert.Equal(t, "Updated Gadget", responseBody["data"].(map[string]interface{})["name"])
}



func TestUpdateCategoryFailed(t *testing.T) {
	

	db , err := database.GetConnection()                                          
	if err != nil {
           panic(err)
        }
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Save(context.Background(), tx, domain.User{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/users/"+strconv.Itoa(user.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	

	db , err := database.GetConnection()
                                                                                    
	if err != nil {                                                               
		panic(err)                 
	}

	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Save(context.Background(), tx, domain.User{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users/"+strconv.Itoa(user.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, user.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, user.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	
	db , err := database.GetConnection()
                                                                                      if err != nil {                                                                  panic(err)                                                                 }
	truncateUser(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {	

	db , err := database.GetConnection()
                                                                                      if err != nil {                                                                  panic(err)                                                                 }
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Save(context.Background(), tx, domain.User{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/users/"+strconv.Itoa(user.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {	

	db , err := database.GetConnection()
                                                                                      if err != nil {                                                                  panic(err)                                                                 }
	truncateUser(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/users/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListCategoriesSuccess(t *testing.T) {
	db , err := database.GetConnection()
	if err != nil {                                                        panic(err)                                                  

        }
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user1 := userRepository.Save(context.Background(), tx, domain.User{
		Name: "Gadget",
	})
	user2 := userRepository.Save(context.Background(), tx, domain.User{
		Name: "Computer",
	})
	tx.Commit()

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	fmt.Println(responseBody)

	var users = responseBody["data"].([]interface{})

	userResponse1 := users[0].(map[string]interface{})
	userResponse2 := users[1].(map[string]interface{})

	assert.Equal(t, user1.Id, int(userResponse1["id"].(float64)))
	assert.Equal(t, user1.Name, userResponse1["name"])

	assert.Equal(t, user2.Id, int(userResponse2["id"].(float64)))
	assert.Equal(t, user2.Name, userResponse2["name"])
}

func TestLoginSuccess(t *testing.T) {
    db, err := database.GetConnection()
    if err != nil {
        panic(err)
    }

    truncateUser(db)

    tx, _ := db.Begin()
    userRepository := repository.NewUserRepository()
    user := userRepository.Save(context.Background(), tx, domain.User{
        Name: "Gadget",
    })
    tx.Commit()

    router := setupRouter(db)

    requestBody := strings.NewReader("id=" + strconv.Itoa(user.Id))
    request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", requestBody)
    request.Header.Add("X-API-Key", "RAHASIA")
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    recorder := httptest.NewRecorder()
    router.ServeHTTP(recorder, request)

    response := recorder.Result()
    assert.Equal(t, 200, response.StatusCode)

    body, _ := io.ReadAll(response.Body)
    var responseBody map[string]interface{}
    json.Unmarshal(body, &responseBody)

    assert.Equal(t, 200, int(responseBody["code"].(float64)))
    assert.Equal(t, "OK", responseBody["status"])
    assert.Equal(t, user.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
    assert.Equal(t, user.Name, responseBody["data"].(map[string]interface{})["name"])
}


func TestUnauthorized(t *testing.T) {


	db , err := database.GetConnection()
                                                                                    
	if err != nil {                                                             
		panic(err)                                                            
	}

	truncateUser(db)
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/categories", nil)
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}

