package bll

import "DDD/sample/sample_DAL"

type TestManager struct {
	modelService *dal.TestService
}

func NewTestManager() *TestManager {
	return &TestManager{
		modelService: dal.NewTestService(),
	}
}

func (m *TestManager) GetModelValue() int {
	return m.modelService.GetModelValue()
}

func (m *TestManager) SetModelValue(value int) {
	m.modelService.SetModelValue(value)
}
