package mocks

import "github.com/stretchr/testify/mock"

type ReaderMock struct {
	mock.Mock
}

func NewReaderMock() *ReaderMock {
	return new(ReaderMock)
}

func (m *ReaderMock) Read() interface{} {
	args := m.Called()
	return args.Get(0)
}

func (m *ReaderMock) ExpectReturns(customerId int) *ReaderMock {
	m.On("Read").Return(customerId)
	return m
}