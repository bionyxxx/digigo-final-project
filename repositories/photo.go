package repositories

import "Final_Project/models"

type PhotoRepo interface {
	CreatePhoto(in models.Photo) (res models.Photo, err error)
	GetAllPhotos(in []models.Photo) (res []models.Photo, err error)
	GetPhoto(photoId uint) (res models.Photo, err error)
	UpdatePhoto(in models.Photo) (res models.Photo, err error)
	DeletePhoto(photoId uint) (err error)
}

func (r Repo) DeletePhoto(photoId uint) (err error) {
	err = r.gorm.Delete(&models.Photo{}, photoId).Error

	if err != nil {
		return err
	}

	return nil
}

func (r Repo) UpdatePhoto(in models.Photo) (res models.Photo, err error) {
	err = r.gorm.Model(&in).Updates(&in).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetPhoto(photoId uint) (res models.Photo, err error) {
	err = r.gorm.Preload("User").Preload("Comments").First(&res, photoId).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) CreatePhoto(in models.Photo) (res models.Photo, err error) {
	err = r.gorm.Create(&in).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetAllPhotos(in []models.Photo) (res []models.Photo, err error) {
	err = r.gorm.Find(&in).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}
