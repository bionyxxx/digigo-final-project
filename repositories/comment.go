package repositories

import (
	"Final_Project/models"
)

type CommentRepo interface {
	CreateComment(comment models.Comment, photoId uint) (res models.Comment, err error)
	GetAllComments(in []models.Comment) (res []models.Comment, err error)
	GetComment(commentId uint) (res models.Comment, err error)
	DeleteComment(commentId uint) (err error)
	UpdateComment(in models.Comment) (res models.Comment, err error)
}

func (r Repo) UpdateComment(in models.Comment) (res models.Comment, err error) {
	err = r.gorm.Model(&in).Updates(in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteComment(commentId uint) (err error) {
	err = r.gorm.Delete(&models.Comment{}, commentId).Error
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) GetComment(commentId uint) (res models.Comment, err error) {
	err = r.gorm.Preload("Photo").Preload("User").First(&res, commentId).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllComments(in []models.Comment) (res []models.Comment, err error) {
	err = r.gorm.Find(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CreateComment(comment models.Comment, photoId uint) (res models.Comment, err error) {
	err = r.gorm.Create(&comment).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return comment, nil
}
