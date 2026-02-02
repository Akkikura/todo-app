package Service

import (
	todo_app "github.com/akkikura/todo-app"
	"github.com/akkikura/todo-app/pkg/Repository"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(tokenString string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo_app.ToDoList) (int, error)
	GetAll(userId int) ([]todo_app.ToDoList, error)
	GetListById(userId, listId int) (todo_app.ToDoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo_app.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *Repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
	}
}
