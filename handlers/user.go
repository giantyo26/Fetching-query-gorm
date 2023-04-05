package handlers

import (
	dto "dumbmerch/dto/result"
	usersdto "dumbmerch/dto/users"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

// Declaring a struct type "handler" with a field "UserRepository" of type "repositories.UserRepository"
type handler struct {
	UserRepository repositories.UserRepository
}

// Defining a function named "HandlerUser" that returns a pointer to an instance of "handler" struct with "UserRepository" field
// initialized with the provided "repositories.UserRepository" parameter
func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

// Calling the "FindUsers" method on the "UserRepository" field of "handler" struct to get a list of users
func (h *handler) FindUsers(c echo.Context) error {
	users, err := h.UserRepository.FindUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// return a JSON response with a HTTP status code of 200 (OK) and the list of users in the response body
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: users})
}

func (h *handler) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id")) // Get the "id" parameter from the URL path, and store it in the "id" variable

	user, err := h.UserRepository.GetUser(id) // / Calling the "GetUser" method on the "UserRepository" field of "handler" struct to get a user with the provided "id
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// Return a JSON response with a HTTP status code of 200 (OK) and the user information converted to a custom response struct in the response body
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)})

}

func (h *handler) CreateUser(c echo.Context) error {
	// This is used to capture the request data sent by the client
	request := new(usersdto.CreateUserRequest) //creates a new pointer to an instance of the usersdto.CreateUserRequest
	// This line attempts to bind the request data to the request variable.
	// c.Bind(request) is used to populate the request struct with the data from the request body.
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	//This line validates the request struct against the validation rules defined in the usersdto.CreateUserRequest struct
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// data form pattern submit to pattern entity db user
	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	data, err := h.UserRepository.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	// This line returns a JSON response with the data returned by the convertResponse function,
	//  which converts the data to a response format defined in the dto.SuccessResult struct.
	// This response contains the newly created user data as the response body.

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func (h *handler) UpdateUser(c echo.Context) error {
	request := new(usersdto.UpdateUserRequest)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	data, err := h.UserRepository.UpdateUser(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

func (h *handler) DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	data, err := h.UserRepository.DeleteUser(user, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)})
}

// Define a function named "convertResponse" that takes a "models.User" parameter and returns a "usersdto.UserResponse" struct
func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID, // Set the "ID" field of the response struct with the "ID" field of the "models.User" parameter
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
