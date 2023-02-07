package resolvers

import (
	"context"
	"stats-mockupero/graph/common"
	"stats-mockupero/graph/model"
	models "stats-mockupero/graph/models"
	"strconv"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input model.ProjectInput) (*model.GqlProject, error) {
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

func (r *mutationResolver) UpdateProject(ctx context.Context, id string, input model.ProjectInput) (*model.GqlProject, error) {
	context := common.GetContext(ctx)

	project := &models.Project{
		Name:        input.Name,
		Description: *input.Description,
	}

	err := context.Database.Model(&project).Where("id = ?", id).Updates(project).Error

	if err != nil {
		return nil, err
	}

	projectResult := model.GqlProject{
		ID:          strconv.FormatUint(uint64(project.ID), 10),
		Name:        project.Name,
		Description: &project.Description,
	}

	return &projectResult, nil
}

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

func (r *mutationResolver) DeleteProject(ctx context.Context, id string) (bool, error) {
	context := common.GetContext(ctx)

	// TODO: review if the project exists

	err := context.Database.Delete(&models.Project{}, id).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) Project(ctx context.Context, id string) (*model.GqlProject, error) {
	context := common.GetContext(ctx)

	var projectResult *model.GqlProject
	var project *models.Project

	err := context.Database.Find(&project, id).Error

	if err != nil {
		return nil, err
	}

	projectResult = &model.GqlProject{
		ID:   strconv.FormatUint(uint64(project.ID), 10),
		Name: project.Name}

	return projectResult, nil
}
