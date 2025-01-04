package services

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"bun-spreader/config"
	"bun-spreader/dto"
	"bun-spreader/models"
)

type UserService struct{}

var ctx = context.Background()

func (us *UserService) CreateUser(username dto.Customer) (models.Customer, error) {

	existingUser, err := us.GetUserByUserName(username.Name)
	if err == nil {
		return models.Customer{}, fmt.Errorf("username %s already exists", existingUser.Name)
	}

	customer := models.Customer{ID: uuid.New().String(), Name: username.Name}
	if _, err := config.DB.NewInsert().Model(&customer).Returning("c.created_at").Exec(ctx); err != nil {
		return models.Customer{}, fmt.Errorf("failed to create user: %w", err)
	}
	return customer, nil
}

func (us *UserService) GetAllUsers() ([]models.Customer, error) {
	customers := new([]models.Customer)
	if err := config.DB.NewSelect().Model(customers).Scan(ctx); err != nil {
		return []models.Customer{}, fmt.Errorf("failed to fetch users: %w", err)
	}

	return *customers, nil
}

func (us *UserService) GetUserByID(uuid string) (models.Customer, error) {
	customer := new(models.Customer)
	if err := config.DB.NewSelect().Model(customer).Where("\"id\" = ?", uuid).Scan(ctx); err != nil {
		return models.Customer{}, fmt.Errorf("failed to fetch user: %w", err)
	}

	return *customer, nil
}

func (us *UserService) GetUserByUserName(name string) (models.Customer, error) {
	customer := new(models.Customer)
	if err := config.DB.NewSelect().Model(customer).Where("\"name\" = ?", name).Scan(ctx); err != nil {
		return models.Customer{}, fmt.Errorf("failed to fetch user: %w", err)
	}

	return *customer, nil
}
