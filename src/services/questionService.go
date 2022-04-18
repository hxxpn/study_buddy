package services

type Question struct {
	questionId      string
	questionScore   int
	questionContent string
}

type QuestionService interface {
	AddQuestion(question *Question) error
	AnswerQuestion(question *Question) error
	GetQuestion(questionId string) (Question, error)
	RemoveQuestion(questionId string) (int, error)
}

type questionService struct{}

func (questionService) AddQuestion(question *Question) error {
	return nil
}
func (questionService) AnswerQuestion(question *Question) error {
	return nil
}
func (questionService) RemoveQuestion(questionId string) (int, error) {
	return 1, nil
}
func (questionService) GetQuestion(questionId string) (Question, error) {
	return Question{}, nil
}
