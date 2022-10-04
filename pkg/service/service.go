package service

import (
	todo "github.com/bilmak/todo-app/pkg"
	"github.com/bilmak/todo-app/pkg/repository"
)



type Authorization interface{
	CreateUser(user todo.User) (int, error)

}

type TodoList interface{

}

type TodoItem interface{

}

type Service struct{
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service{
	return &Service{}
}