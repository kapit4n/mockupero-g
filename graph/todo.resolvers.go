package graph

import (
	"context"
	"fmt"
	"stats-mockupero/graph/common"
	"stats-mockupero/graph/model"
	models "stats-mockupero/graph/models"
	"strconv"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	//return &model.Todo{ID: "1", Text: input.Text}, nil
	//panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))

	context := common.GetContext(ctx)
	todo := &models.Todo{
		Text: input.Text,
		Done: false,
	}

	err := context.Database.Create(&todo).Error

	todoResult := &model.Todo{ID: string(todo.ID), Text: todo.Text}

	if err != nil {
		return nil, err
	}

	return todoResult, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	context := common.GetContext(ctx)
	var todos []*models.Todo
	var todosResult []*model.Todo

	err := context.Database.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	for _, todo := range todos {
		fmt.Println(todo.ID)

		xType := fmt.Sprintf("%T", todo.ID)
		fmt.Println(xType) // "[]int"

		todosResult = append(todosResult, &model.Todo{ID: strconv.FormatUint(uint64(todo.ID), 10), Text: todo.Text})
	}

	return todosResult, nil
}
