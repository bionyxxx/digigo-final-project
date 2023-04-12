package repositories

import (
	"Final_Project/models"
	"errors"
	"gorm.io/gorm"
)

type SocialMediaRepo interface {
	GetAllSocialMedia(in []models.SocialMedia) (res []models.SocialMedia, err error)
	CreateSocialMedia(in models.SocialMedia) (res models.SocialMedia, err error)
	GetSocialMedia(socialMediaId uint) (res models.SocialMedia, err error)
	UpdateSocialMedia(in models.SocialMedia) (res models.SocialMedia, err error)
	DeleteSocialMedia(in models.SocialMedia) (err error)
}

func (r Repo) DeleteSocialMedia(in models.SocialMedia) (err error) {
	err = r.gorm.Where("user_id = ?", in.UserID).Delete(&in).Error

	if err != nil {
		return err
	}

	return nil
}

func (r Repo) UpdateSocialMedia(in models.SocialMedia) (res models.SocialMedia, err error) {
	if err := r.gorm.Model(&in).Where("user_id = ?", in.UserID).Updates(&in).Scan(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetSocialMedia(socialMediaId uint) (res models.SocialMedia, err error) {
	err = r.gorm.Preload("User").First(&res, socialMediaId).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CreateSocialMedia(in models.SocialMedia) (res models.SocialMedia, err error) {
	if err := r.gorm.First(&in, "user_id = ?", in.UserID).Scan(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = r.gorm.Create(&in).Scan(&res).Error
			if err != nil {
				return res, err
			}
			return res, nil
		} else {
			if err != nil {
				return res, err
			}
		}
	}
	return res, nil
}

func (r Repo) GetAllSocialMedia(in []models.SocialMedia) (res []models.SocialMedia, err error) {
	err = r.gorm.Find(&in).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}
