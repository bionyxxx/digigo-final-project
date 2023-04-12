package services

import "Final_Project/models"

type PhotoService interface {
	CreatePhoto(in models.Photo) (res models.Photo, err error)
	GetAllPhotos(in []models.Photo) (res []models.Photo, err error)
	GetPhoto(photoId uint) (res models.Photo, err error)
	UpdatePhoto(in models.Photo) (res models.Photo, err error)
	DeletePhoto(photoId uint) (err error)
}

func (s *Service) DeletePhoto(photoId uint) (err error) {
	err = s.repo.DeletePhoto(photoId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdatePhoto(in models.Photo) (res models.Photo, err error) {
	res, err = s.repo.UpdatePhoto(in)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetPhoto(photoId uint) (res models.Photo, err error) {
	res, err = s.repo.GetPhoto(photoId)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) CreatePhoto(in models.Photo) (res models.Photo, err error) {
	res, err = s.repo.CreatePhoto(in)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetAllPhotos(in []models.Photo) (res []models.Photo, err error) {
	res, err = s.repo.GetAllPhotos(in)
	if err != nil {
		return res, err
	}

	return res, nil
}
