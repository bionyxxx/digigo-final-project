package configs

import (
	"Final_Project/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Gorm struct {
	Username string
	Password string
	Port     string
	Address  string
	Database string
	DB       *gorm.DB
}

type GormDb struct {
	*Gorm
}

var (
	GORM *GormDb
	DB   *gorm.DB
	err  error
)

func GetDB() *gorm.DB {
	return DB
}

func InitDatabase() error {
	GORM = new(GormDb)

	GORM.Gorm = &Gorm{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Address:  os.Getenv("POSTGRES_ADDRESS"),
		Database: os.Getenv("POSTGRES_DB"),
	}

	// connect to database
	err := GORM.Gorm.OpenConnection()
	if err != nil {
		return err
	}

	return nil
}

func (p *Gorm) OpenConnection() error {
	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", p.Address, p.Port, p.Username, p.Database, p.Password)

	DB, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to db", err)
	}

	p.DB = DB

	err = p.DB.Debug().AutoMigrate(&models.User{}, &models.Photo{}, &models.Comment{}, &models.SocialMedia{})
	if err != nil {
		panic(err)
	}
	//admin := models.User{
	//	Username: "dutta",
	//	Email:    "dutta@mail.com",
	//	Password: helpers.BcryptHash("password"),
	//	Age:      21,
	//}
	//
	//if DB.Model(&admin).Where("email = ?", admin.Email).Updates(&admin).RowsAffected == 0 {
	//	err := DB.Create(&admin).Error
	//	if err != nil {
	//		panic(err)
	//	}
	//}

	return nil
}
