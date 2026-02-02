package Repository

import (
	todo_app "github.com/akkikura/todo-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo_app.User) (int, error)
	GetUser(username, password string) (todo_app.User, error)
}

type TodoList interface {
	Create(userId int, list todo_app.ToDoList) (int, error)
	GetAll(userId int) ([]todo_app.ToDoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo_app.UpdateListInput) error
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
