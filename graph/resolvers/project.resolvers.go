package resolvers

import (
	"context"
	"fmt"
	"stats-mockupero/graph/model"
)

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.GqlProject, error) {
	panic(fmt.Errorf("not implemented: CreateProject - createProject"))
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context) ([]*model.GqlProject, error) {
	panic(fmt.Errorf("not implemented: Projects - projects"))
}
