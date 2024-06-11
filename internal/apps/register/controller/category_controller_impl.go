package controller

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "rest_api/internal/apps/register/helper"
    "rest_api/internal/apps/register/model/web"
    "rest_api/internal/apps/register/service"
    "strconv"
)

// CategoryControllerImpl represents the implementation of the CategoryController interface
type CategoryControllerImpl struct {
    CategoryService service.CategoryService
}

// NewCategoryController creates a new instance of CategoryController
func NewCategoryController(categoryService service.CategoryService) CategoryController {
    return &CategoryControllerImpl{
        CategoryService: categoryService,
    }
}

// Create a new category
// swagger:operation POST /api/categories categories createCategory
//
// ---
// summary: Create a new category
// description: Creates a new category in the system.
// parameters:
// - name: X-API-Key
//   in: header
//   description: API key for authorization
//   required: true
//   type: string
// - name: body
//   in: body
//   description: The category object to create.
//   required: true
//   schema:
//     "$ref": "#/definitions/CategoryCreateRequest"
// responses:
//   '200':
//     description: Successfully created category.
//     schema:
//       "$ref": "#/responses/webResponse"
//   '400':
//     description: Invalid request payload.
//     schema:
//       "$ref": "#/responses/errorResponse"
//   '500':
//     description: Internal server error.
//     schema:
//       "$ref": "#/responses/errorResponse"
func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    categoryCreateRequest := web.CategoryCreateRequest{}
    helper.ReadFromRequestBody(request, &categoryCreateRequest)

    categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
    webResponse := web.WebResponse{
        Code:   200,
        Status: "OK",
        Data:   categoryResponse,
    }

    helper.WriteToResponseBody(writer, webResponse)
}

// Update an existing category
// swagger:operation PUT /categories/{categoryId} categories updateCategory
//
// ---
// summary: Update an existing category
// description: Updates an existing category in the system.
// parameters:
// - name: categoryId
//   in: path
//   description: ID of the category to update
//   required: true
//   type: integer
// - name: body
//   in: body
//   description: The updated category object
//   required: true
//   schema:
//     "$ref": "#/definitions/CategoryUpdateRequest"
// responses:
//   '200':
//     description: Successfully updated category.
//     schema:
//       "$ref": "#/responses/webResponse"
//   '400':
//     description: Invalid request payload.
//     schema:
//       "$ref": "#/responses/errorResponse"
//   '404':
//     description: Category not found.
//     schema:
//       "$ref": "#/responses/errorResponse"
//   '500':
//     description: Internal server error.
//     schema:
//       "$ref": "#/responses/errorResponse"
func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    categoryUpdateRequest := web.CategoryUpdateRequest{}
    helper.ReadFromRequestBody(request, &categoryUpdateRequest)

    categoryId := params.ByName("categoryId")
    id, err := strconv.Atoi(categoryId)
    helper.PanicIfError(err)

    categoryUpdateRequest.Id = id

    categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
    webResponse := web.WebResponse{
        Code:   200,
        Status: "OK",
        Data:   categoryResponse,
    }

    helper.WriteToResponseBody(writer, webResponse)
}

// DeleteCategory deletes an existing category from the system.
//
// This endpoint allows deleting an existing category by providing its ID.
//
// swagger:operation DELETE /api/categories/{categoryId} categories deleteCategory
//
// ---
// summary: Delete an existing category
// description: Deletes an existing category in the system.
// parameters:
// - name: X-API-Key
//   in: header
//   description: API key for authorization
//   required: true
//   type: string
// - name: categoryId
//   in: path
//   description: ID of the category to delete
//   required: true
//   type: integer
// responses:
//   '200':
//     description: Successfully deleted category.
//     schema:
//       "$ref": "#/responses/webResponse"
//   '404':
//     description: Category not found.
//     schema:
//       "$ref": "#/responses/errorResponse"
//   '500':
//     description: Internal server error.
//     schema:
//       "$ref": "#/responses/errorResponse"
func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    categoryId := params.ByName("categoryId")
    id, err := strconv.Atoi(categoryId)
    helper.PanicIfError(err)

    controller.CategoryService.Delete(request.Context(), id)
    webResponse := web.WebResponse{
        Code:   200,
        Status: "OK",
    }

    helper.WriteToResponseBody(writer, webResponse)
}
// FindById retrieves a category by its ID.
//
// This endpoint allows retrieving a category by providing its ID.
//
// swagger:operation GET /api/categories/{categoryId} categories get
//
// ---
// summary: Find a category by ID
// description: Retrieves a category by its ID.
// parameters:
// - name: X-API-Key
//   in: header
//   description: API key for authorization
//   required: true
//   type: string
// - name: categoryId
//   in: path
//   description: ID of the category to retrieve
//   required: true
//   type: integer
// responses:
//   '200':
//     description: Successfully retrieved category.
//     schema:
//       "$ref": "#/responses/webResponse"
//   '404':
//     description: Category not found.
//     schema:
//       "$ref": "#/responses/errorResponse"
//   '500':
//     description: Internal server error.
//     schema:
//       "$ref": "#/responses/errorResponse"
func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    categoryId := params.ByName("categoryId")
    id, err := strconv.Atoi(categoryId)
    helper.PanicIfError(err)

    categoryResponse := controller.CategoryService.FindById(request.Context(), id)
    webResponse := web.WebResponse{
        Code:   200,
        Status: "OK",
        Data:   categoryResponse,
    }

    helper.WriteToResponseBody(writer, webResponse)
}

// Find all categories
// swagger:operation GET /categories categories listCategories
//
// ---
// summary: Find all categories
// description: Retrieves all categories.
// responses:
//   '200':
//     description: Successfully retrieved categories.
//     schema:
//       "$ref": "#/responses/webResponse"
//   '500':
//     description: Internal server error.
//     schema:
//       "$ref": "#/responses/errorResponse"
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    categoryResponses := controller.CategoryService.FindAll(request.Context())
    webResponse := web.WebResponse{
        Code:   200,
        Status: "OK",
        Data:   categoryResponses,
    }

    helper.WriteToResponseBody(writer, webResponse)
}

