package services

import "Final_Project/models"

type SocialMediaService interface {
	GetAllSocialMedia(in []models.SocialMedia) (res []models.SocialMedia, err error)
	CreateSocialMedia(in models.SocialMedia) (res models.SocialMedia, err error)
	GetSocialMedia(socialMediaId uint) (res models.SocialMedia, err error)
	UpdateSocialMedia(in models.SocialMedia) (res models.SocialMedia, err error)
	DeleteSocialMedia(in models.SocialMedia) (err error)
}

func (s *Service) DeleteSocialMedia(in models.SocialMedia) (err error) {
	err = s.repo.DeleteSocialMedia(in)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateSocialMedia(in models.SocialMedia) (res models.SocialMedia, err error) {
	res, err = s.repo.UpdateSocialMedia(in)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetSocialMedia(socialMediaId uint) (res models.SocialMedia, err error) {
	res, err = s.repo.GetSocialMedia(socialMediaId)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) CreateSocialMedia(in models.SocialMedia) (res models.SocialMedia, err error) {
	res, err = s.repo.CreateSocialMedia(in)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *Service) GetAllSocialMedia(in []models.SocialMedia) (res []models.SocialMedia, err error) {
	res, err = s.repo.GetAllSocialMedia(in)
	if err != nil {
		return res, err
	}

	return res, nil
}
