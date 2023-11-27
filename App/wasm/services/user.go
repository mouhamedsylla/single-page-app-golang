package services

import "fmt"

type UserRepository interface {
	GetUserByID(userID int) string
}

type DatabaseUserRepository struct {
	dbConnection string
}

func (repo *DatabaseUserRepository) GetUserByID(userID int) string {
	return fmt.Sprintf("Utilisateur avec l'ID %d", userID)
}

type UserService struct {
	userRepository UserRepository
}

func (service *UserService) GetUserByID(userID int) string {
	user := service.userRepository.GetUserByID(userID)
	return user
}
