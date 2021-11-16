package model

type Customer interface{
	ID() int
	IsNotActive() bool
}