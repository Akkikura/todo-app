package Service

import (
	todo_app "github.com/akkikura/todo-app"
	"github.com/akkikura/todo-app/pkg/Repository"
)

type TodoListService struct {
	repo Repository.TodoList
}

func NewTodoListService(repo Repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list todo_app.ToDoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]todo_app.ToDoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetListById(userId, listId int) (todo_app.ToDoList, error) {
	return s.repo.GetListById(userId, listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input todo_app.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)

}
