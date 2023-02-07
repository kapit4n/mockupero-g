package resolvers

import (
	"context"
	"fmt"
	"stats-mockupero/graph/common"
	"stats-mockupero/graph/model"
	models "stats-mockupero/graph/models"
	"strconv"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.GqlTodo, error) {
	context := common.GetContext(ctx)
	todo := &models.Todo{
		Text: input.Text,
		Done: false,
	}

	err := context.Database.Create(&todo).Error

	todoResult := &model.GqlTodo{ID: string(todo.ID), Text: todo.Text}

	if err != nil {
		return nil, err
	}

	return todoResult, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.GqlTodo, error) {
	context := common.GetContext(ctx)
	var todos []*models.Todo
	var todosResult []*model.GqlTodo

	err := context.Database.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	for _, todo := range todos {
		fmt.Println(todo.ID)

		xType := fmt.Sprintf("%T", todo.ID)
		fmt.Println(xType) // "[]int"

		todosResult = append(todosResult, &model.GqlTodo{ID: strconv.FormatUint(uint64(todo.ID), 10), Text: todo.Text})
	}

	return todosResult, nil
}
