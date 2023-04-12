package services

import "Final_Project/repositories"

type Service struct {
	repo repositories.RepoInterface
}

type ServiceInterface interface {
	PhotoService
	CommentService
	SocialMediaService
}

func NewService(repo repositories.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}
