package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type userModel struct {
	UserId    string `json:"userId"`
	UserName  string `json:"userName"`
	UserEmail string `json:"userEmail"`
}

type UserService interface {
	GetUser(userId string) (userModel, error)
	UpdateUser(user *userModel) (int, error)
	GetAllUsers() []userModel
	AddUser(user *userModel) (int, error)
	RemoveUser(userId string) (int, error)
}

type DUserService struct{}

var dummyDBUserMap = map[string]userModel{"user_id": {"user_id", "name", "email"}}

var UserNotFoundErr = errors.New("user not found")

func (DUserService) GetUser(userId string) (userModel, error) {
	if user, ok := dummyDBUserMap[userId]; ok {
		return user, nil
	}
	return userModel{}, UserNotFoundErr
}

func (DUserService) AddUser(_ *userModel) (int, error) {
	return 1, nil
}

func (DUserService) UpdateUser(_ *userModel) (int, error) {
	return 1, nil
}

func (DUserService) GetAllUsers() []userModel {
	return []userModel{}
}

func (DUserService) RemoveUser(_ string) (int, error) {
	return 1, nil
}

type getUserRequest struct {
	UserId string `json:"userId"`
}

type getUserResponse struct {
	User userModel `json:"user"`
	Err  string    `json:"error"`
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
func makeGetUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserRequest)
		response, err := svc.GetUser(req.UserId)
		if err != nil {
			return nil, err
		}
		return getUserResponse{response, ""}, nil
	}
}
