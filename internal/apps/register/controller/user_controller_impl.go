package controller

import (
	"net/http"
	"rest_api/internal/apps/register/helper"
	"rest_api/internal/apps/register/model/web"
	"rest_api/internal/apps/register/service"

	"github.com/julienschmidt/httprouter"
)

// UserControllerImpl represents the implementation of the UserController interface
type UserControllerImpl struct {
	UserService service.UserService
}

// NewUserController creates a new instance of UserController
func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

// Create a new user
// swagger:operation POST /api/users users createUser
//
// ---
// summary: Create a new user
// description: Creates a new user in the system.
// parameters:
//   - name: X-API-Key
//     in: header
//     description: API key for authorization
//     required: true
//     type: string
//   - name: body
//     in: body
//     description: The user object to create.
//     required: true
//     schema:
//     "$ref": "#/definitions/UserCreateRequest"
//
// responses:
//
//	'200':
//	  description: Successfully created user.
//	  schema:
//	    "$ref": "#/responses/webResponse"
//	'400':
//	  description: Invalid request payload.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
//	'500':
//	  description: Internal server error.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Update an existing user
// swagger:operation PUT /api/users/{userEmail} users updateUser
//
// ---
// summary: Update an existing user
// description: Updates an existing user in the system.
// parameters:
//   - name: X-API-Key
//     in: header
//     description: API key for authorization
//     required: true
//     type: string
//   - name: userEmail
//     in: path
//     description: The email of the user to update
//     required: true
//     type: string
//   - name: body
//     in: body
//     description: The updated user object
//     required: true
//     schema:
//     "$ref": "#/definitions/UserUpdateRequest"
//
// responses:
//
//	'200':
//	  description: Successfully updated user.
//	  schema:
//	    "$ref": "#/responses/webResponse"
//	'400':
//	  description: Invalid request payload.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
//	'404':
//	  description: User not found.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
//	'500':
//	  description: Internal server error.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userEmail := params.ByName("userEmail")
	userUpdateRequest.Email = userEmail

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Delete an existing user
// swagger:operation DELETE /api/users/{userEmail} users deleteUser
//
// ---
// summary: Delete an existing user
// description: Deletes an existing user from the system.
// parameters:
//   - name: X-API-Key
//     in: header
//     description: API key for authorization
//     required: true
//     type: string
//   - name: userEmail
//     in: path
//     description: Email of the user to delete
//     required: true
//     type: string
//
// responses:
//
//	'200':
//	  description: Successfully deleted user.
//	  schema:
//	    "$ref": "#/responses/webResponse"
//	'404':
//	  description: User not found.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
//	'500':
//	  description: Internal server error.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userEmail := params.ByName("userEmail")

	controller.UserService.Delete(request.Context(), userEmail)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Find a user by Email
// swagger:operation GET /api/users/{userEmail} users getUser
//
// ---
// summary: Find a user by email
// description: Retrieves a user by its email.
// parameters:
//   - name: X-API-Key
//     in: header
//     description: API key for authorization
//     required: true
//     type: string
//   - name: userEmail
//     in: path
//     description: Email of the user to retrieve
//     required: true
//     type: string
//
// responses:
//
//	'200':
//	  description: Successfully retrieved user.
//	  schema:
//	    "$ref": "#/responses/webResponse"
//	'404':
//	  description: User not found.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
//	'500':
//	  description: Internal server error.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
func (controller *UserControllerImpl) FindByEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userEmail := params.ByName("userEmail")

	userResponse := controller.UserService.FindByEmail(request.Context(), userEmail)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// Find all users
// swagger:operation GET /api/users users listUsers
//
// ---
// summary: Find all users
// description: Retrieves all users.
// parameters:
//   - name: X-API-Key
//     in: header
//     description: API key for authorization
//     required: true
//     type: string
//
// responses:
//
//	'200':
//	  description: Successfully retrieved users.
//	  schema:
//	    "$ref": "#/responses/webResponse"
//	'500':
//	  description: Internal server error.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userResponses := controller.UserService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

// LoginHandler handles user login
// swagger:operation POST /api/login users loginUser
//
// ---
// summary: Authenticate a user
// description: Authenticates a user and returns user information.
// parameters:
//   - name: X-API-Key
//     in: header
//     description: API key for authorization
//     required: true
//     type: string
//   - name: body
//     in: body
//     description: User credentials for login.
//     required: true
//     schema:
//     "$ref": "#/definitions/UserUpdateRequest"
//
// responses:
//
//	'200':
//	  description: Successfully authenticated.
//	  schema:
//	    "$ref": "#/responses/webResponse"
//	'400':
//	  description: Invalid request payload.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
//	'401':
//	  description: Unauthorized access.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
//	'500':
//	  description: Internal server error.
//	  schema:
//	    "$ref": "#/responses/errorResponse"
func (controller *UserControllerImpl) LoginHandler(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	name := request.PostForm.Get("Name")
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindByName(request.Context(), name)
	if userResponse.Name == name {
		webResponse := web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   userResponse,
		}

		helper.WriteToResponseBody(writer, webResponse)
	}
}
