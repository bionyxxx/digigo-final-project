package repositories

import (
	"database/sql"
	"gorm.io/gorm"
)

type Repo struct {
	db   *sql.DB
	gorm *gorm.DB
}

type RepoInterface interface {
	PhotoRepo
	CommentRepo
	SocialMediaRepo
}

func NewRepo(gorm *gorm.DB) *Repo {
	return &Repo{gorm: gorm}
}
