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

// swagger:route POST /categories categories createCategory
//
// Create a new category
//
// Creates a new category in the system.
//
// Responses:
//   200: webResponse
//   400: errorResponse "Invalid request payload"
//   500: errorResponse "Internal server error"
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
// swagger:route PUT /categories/{categoryId} categories updateCategory
//
// Updates an existing category in the system.
//
// Responses:
//   200: webResponse
//   400: errorResponse "Invalid request payload"
//   404: errorResponse "Category not found"
//   500: errorResponse "Internal server error"
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

// Delete an existing category
// swagger:route DELETE /categories/{categoryId} categories deleteCategory
//
// Deletes an existing category in the system.
//
// Responses:
//   200: webResponse
//   404: errorResponse "Category not found"
//   500: errorResponse "Internal server error"
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

// Find a category by ID
// swagger:route GET /categories/{categoryId} categories getCategory
//
// Retrieves a category by its ID.
//
// Responses:
//   200: webResponse
//   404: errorResponse "Category not found"
//   500: errorResponse "Internal server error"
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
// swagger:route GET /categories categories listCategories
//
// Retrieves all categories.
//
// Responses:
//   200: webResponse
//   500: errorResponse "Internal server error"
func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    categoryResponses := controller.CategoryService.FindAll(request.Context())
    webResponse := web.WebResponse{
        Code:   200,
        Status: "OK",
        Data:   categoryResponses,
    }

    helper.WriteToResponseBody(writer, webResponse)
}

