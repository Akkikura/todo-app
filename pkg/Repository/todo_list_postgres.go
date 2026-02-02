package Repository

import (
	"fmt"
	"strings"

	todo_app "github.com/akkikura/todo-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo_app.ToDoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createListQuerry := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListTable)
	row := tx.QueryRow(createListQuerry, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createUsersListQuerry := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) RETURNING user_id", usersListTable)
	_, err = tx.Exec(createUsersListQuerry, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *TodoListPostgres) GetListById(userId, listId int) (todo_app.ToDoList, error) {
	var list todo_app.ToDoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl "+
		"INNER JOIN %s ul on tl.id = il.list_id WHERE ul.user_id = $1 AND ul.list_id = $2", todoListTable, usersListTable)
	err := r.db.Get(&list, query, userId, listId)
	return list, err
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo_app.ToDoList, error) {
	var lists []todo_app.ToDoList

	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl "+
		"INNER JOIN %s ul on tl.id = il.list_id WHERE ul.user_id = $1", todoListTable, usersListTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", todoListTable, usersListTable)
	_, err := r.db.Exec(query, userId, listId)
	return err
}

func (r *TodoListPostgres) Update(userId, listId int, input todo_app.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id=ul.list_id AND ul.list_id=$%d AND "+
		"ul.user_id = $%d", todoListTable, setQuery, usersListTable, argId, argId+1)
	args = append(args, listId, userId)
	logrus.Debugf("query: %s", query)
	logrus.Debugf("query: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
