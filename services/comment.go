package services

import "Final_Project/models"

type CommentService interface {
	CreateComment(comment models.Comment, photoId uint) (res models.Comment, err error)
	GetAllComments(in []models.Comment) (res []models.Comment, err error)
	GetComment(commentId uint) (res models.Comment, err error)
	DeleteComment(commentId uint) (err error)
	UpdateComment(in models.Comment) (res models.Comment, err error)
}

func (s *Service) UpdateComment(in models.Comment) (res models.Comment, err error) {
	res, err = s.repo.UpdateComment(in)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) DeleteComment(commentId uint) (err error) {
	err = s.repo.DeleteComment(commentId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetComment(commentId uint) (res models.Comment, err error) {
	res, err = s.repo.GetComment(commentId)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetAllComments(in []models.Comment) (res []models.Comment, err error) {
	res, err = s.repo.GetAllComments(in)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) CreateComment(comment models.Comment, photoId uint) (res models.Comment, err error) {
	res, err = s.repo.CreateComment(comment, photoId)
	if err != nil {
		return res, err
	}

	return res, nil
}
