package handlers

import (
	"FirstProject/internal/userService"
	"FirstProject/internal/web/users"
	"context"
)

type UserHandler struct {
	Service *userService.UserService
}

func (h *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		user := users.User{
			Id:       &usr.ID,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}
	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.Users{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}
	responce := users.PostUsers201JSONResponse{
		Id:       &createdUser.ID,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}
	return responce, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request users.DeleteUsersIdRequestObject) (users.DeleteUsersIdResponseObject, error) {
	id := request.Id
	err := h.Service.DeleteUserByID(uint(id))
	if err != nil {
		return nil, err
	}
	return users.DeleteUsersId204Response{}, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request users.PatchUsersIdRequestObject) (users.PatchUsersIdResponseObject, error) {
	id := request.Id
	userRequest := request.Body
	serviceUser := userService.Users{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}
	updatedUser, err := h.Service.UpdateUserByID(uint(id), serviceUser)
	if err != nil {
		return nil, err
	}
	responce := users.PatchUsersId200JSONResponse{
		Id:       &updatedUser.ID,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}
	return responce, nil
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{Service: service}
}
