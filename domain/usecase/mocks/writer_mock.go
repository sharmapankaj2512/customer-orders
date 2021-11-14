package mocks

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

type WriterMock struct {
	mock.Mock
}

func NewWriterMock() *WriterMock {
	return new(WriterMock)
}

func (m *WriterMock) Write(output interface{}) {
	m.Called(output)
}

func (m *WriterMock) ExpectReceivesError(message string) *WriterMock {
	m.On("Write", errors.New(message))
	return m
}
