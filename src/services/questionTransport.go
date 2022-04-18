package services

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

func makeAddQuestionHandler(questionSvc QuestionService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var req = request.(addQuestionRequest)
		err := questionSvc.AddQuestion(&req.question)
		if err != nil {
			return nil, err
		}

		return addQuestionResponse{}, nil
	}
}
func makeAnswerQuestionHandler(questionSvc QuestionService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var req = request.(addQuestionRequest)
		err := questionSvc.AddQuestion(&req.question)
		if err != nil {
			return nil, err
		}

		return addQuestionResponse{}, nil
	}
}
func makeRemoveQuestionHandler(questionSvc QuestionService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var req = request.(removeQuestionRequest)
		status, err := questionSvc.RemoveQuestion(req.questionId)
		if err != nil {
			return nil, err
		}
		if status == 1 {
			return removeQuestionResponse{message: "Success"}, nil
		} else {
			return removeQuestionResponse{}, nil
		}
	}
}
func makeGetQuestionHandler(questionSvc QuestionService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var req = request.(getQuestionRequest)
		question, err := questionSvc.GetQuestion(req.questionId)
		if err != nil {
			return nil, err
		}

		return getQuestionResponse{question}, nil
	}
}

func decodeAddQuestionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request addQuestionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, err
	}
	return request, nil
}
func decodeAnswerQuestionRequest(_ context.Context, r *http.Request) (any, error)         {}
func decodeRemoveQuestionRequest(_ context.Context, r *http.Request) (interface{}, error) {}
func decodeGetQuestionRequest(_ context.Context, r *http.Request) (interface{}, error) {
}

func encodeQuestionResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type addQuestionRequest struct {
	question Question `json:"question"`
}

type addQuestionResponse struct {
	message string `json:"message"`
}

type removeQuestionRequest struct {
	questionId string `json:"questionId"`
}

type removeQuestionResponse struct {
	message string `json:"message"`
}

type getQuestionResponse struct {
	question Question `json:"question"`
}

type getQuestionRequest struct {
	questionId string `json:"questionId"`
}

type answerQuestionResponse struct {
	message string `json:"message"`
}

type answerQuestionRequest struct {
	question Question `json:"question"`
}
