package model

type TestModel struct {
	Model int
}

func NewTestModel() *TestModel {
	return &TestModel{
		Model: 123,
	}
}
