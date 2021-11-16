package model

type Customer interface{
	IsNotActive() bool
}