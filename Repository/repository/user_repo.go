package repository

import(
	models "RepositoryPattern/model"
)


type UserRepo interface {
	Select() ([]models.User, error)
	Insert(u models.User) (error)
}