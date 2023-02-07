package resolvers

import (
	"context"
	"stats-mockupero/graph/common"
	"stats-mockupero/graph/model"
	models "stats-mockupero/graph/models"
	"strconv"
)

// CreateProject is the resolver for the createProject field.
func (r *mutationResolver) CreateProject(ctx context.Context, input model.NewProject) (*model.GqlProject, error) {
	project := &models.Project{
		Name:        input.Name,
		Description: *input.Description,
	}

	context := common.GetContext(ctx)
	err := context.Database.Create(&project).Error

	if err != nil {
		return nil, err
	}

	projectResult := &model.GqlProject{
		ID:          strconv.FormatUint(uint64(project.ID), 10),
		Name:        project.Name,
		Description: &project.Description,
	}

	return projectResult, nil
}

// Projects is the resolver for the projects field.
func (r *queryResolver) Projects(ctx context.Context) ([]*model.GqlProject, error) {
	context := common.GetContext(ctx)

	var projects []*models.Project
	var projectsResult []*model.GqlProject

	err := context.Database.Find(&projects).Error

	if err != nil {
		return nil, err
	}

	for _, project := range projects {
		projectsResult = append(projectsResult, &model.GqlProject{
			ID:          strconv.FormatUint(uint64(project.ID), 10),
			Name:        project.Name,
			Description: &project.Description,
		})
	}

	return projectsResult, nil
}
