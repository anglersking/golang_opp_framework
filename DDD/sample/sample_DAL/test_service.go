package dal

import "DDD/sample/sample_model"

type TestService struct {
	testModel *model.TestModel
}

func NewTestService() *TestService {
	return &TestService{
		testModel: model.NewTestModel(),
	}
}

func (s *TestService) GetModelValue() int {
	return s.testModel.Model
}

func (s *TestService) SetModelValue(value int) {
	s.testModel.Model = value
}
